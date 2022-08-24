package connectors

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/lightningnetwork/lnd/lnwallet/chainfee"
)

const ()

const (
	getNodesEndpoint             = "/getnodes"
	getPairsEndpoint             = "/getpairs"
	createSwapEndpoint           = "/createswap"
	routingHintsEndpoint         = "/routinghints"
	swapStatusEndpoint           = "/swapstatus"
	broadcastTransactionEndpoint = "/broadcasttransaction"
	claimWitnessInputSize        = 1 + 1 + 8 + 73 + 1 + 32 + 1 + 100
)

var ErrSwapNotFound = errors.New("transaction not in mempool or settled/canceled")

type BoltzConnector interface {
	GetPair() (PairResponse, error)
	GetReverseSwapInfo() (*ReverseSwapInfo, error)
	NewReverseSwap(pairId string, orderSide string, amt btcutil.Amount, feesHash string, routingNode []byte) (*ReverseSwap, error)
	CheckTransaction(transactionHex, lockupAddress string, amt int64) (string, error)
	GetTransaction(id, lockupAddress string, amt int64) (status, txid, tx string, eta int, err error)
	ClaimFee(claimAddress string, feePerKw int64) (int64, error)
	ClaimTransaction(redeemScript, transactionHex string, claimAddress string, preimage, key string, fees int64) (string, error)
	GetNodePubkey() (string, error)
	GetRoutingHints(routingNode []byte) ([]RoutingHint, error)
}

type Boltz struct {
	rest         *RestClient
	apiURL       string
	chain        *chaincfg.Params
	claimAddress string
}

func NewBoltz(apiURL string, chain *chaincfg.Params, claimAddress string) (*Boltz, error) {
	rest, _ := NewRestClient(apiURL)
	return &Boltz{
		rest,
		apiURL,
		chain,
		claimAddress,
	}, nil
}

func (boltz *Boltz) GetPair() (response PairResponse, err error) {
	err = boltz.rest.Get(getPairsEndpoint, &response)
	return
}

func (boltz *Boltz) GetReverseSwapInfo() (*ReverseSwapInfo, error) {

	pairs, err := boltz.GetPair()
	if err != nil {
		return nil, err
	}

	for _, w := range pairs.Warnings {
		if w == "reverse.swaps.disabled" {
			return nil, fmt.Errorf("reverse.swaps.disabled")
		}
	}
	btcPair, ok := pairs.Pairs["BTC/rBTC"]
	if !ok {
		return nil, fmt.Errorf("no BTC/rBTC pair")
	}
	return &ReverseSwapInfo{
		FeesHash: btcPair.Hash,
		Max:      btcPair.Limits.Maximal,
		Min:      btcPair.Limits.Minimal,
		Fees: Fees{
			Percentage: btcPair.Fees.Percentage,
			Lockup:     btcPair.Fees.MinerFees.BaseAsset.Reverse.Lockup,
			Claim:      btcPair.Fees.MinerFees.BaseAsset.Reverse.Claim,
		},
	}, nil
}

// NewReverseSwap begins the reverse submarine process.
func (boltz *Boltz) NewReverseSwap(pairId string, orderSide string, amt btcutil.Amount, feesHash string, routingNode []byte) (*ReverseSwap, error) {
	preimage := boltz.getPreimage()
	preimageHash := sha256.Sum256(preimage)
	key, err := boltz.getPrivate()
	if err != nil {
		return nil, fmt.Errorf("getPrivate: %w", err)
	}

	rs, err := boltz.createReverseSwap(pairId, orderSide, int64(amt), feesHash, preimage, preimageHash, boltz.claimAddress, routingNode)
	if err != nil {
		return nil, fmt.Errorf("createReverseSwap amt:%v, preimage:%x, key:%x; %w", amt, preimage, key, err)
	}

	return &ReverseSwap{*rs, hex.EncodeToString(preimage), hex.EncodeToString(preimageHash[:]), hex.EncodeToString(key.Serialize())}, nil
}

// CheckTransaction checks that the transaction corresponds to the adresss and amount
func (boltz *Boltz) CheckTransaction(transactionHex, lockupAddress string, amt int64) (string, error) {
	txSerialized, err := hex.DecodeString(transactionHex)
	if err != nil {
		return "", fmt.Errorf("hex.DecodeString(%v): %w", transactionHex, err)
	}
	tx, err := btcutil.NewTxFromBytes(txSerialized)
	if err != nil {
		return "", fmt.Errorf("btcutil.NewTxFromBytes(%x): %w", txSerialized, err)
	}
	var out *wire.OutPoint
	for i, txout := range tx.MsgTx().TxOut {
		class, addresses, requiredsigs, err := txscript.ExtractPkScriptAddrs(txout.PkScript, boltz.chain)
		if err != nil {
			return "", fmt.Errorf("txscript.ExtractPkScriptAddrs(%x) %w", txout.PkScript, err)
		}
		if class == txscript.WitnessV0ScriptHashTy && len(addresses) == 1 && addresses[0].EncodeAddress() == lockupAddress && requiredsigs == 1 {
			out = wire.NewOutPoint(tx.Hash(), uint32(i))
			if int64(amt) != txout.Value {
				return "", fmt.Errorf("bad amount: %v != %v", int64(amt), txout.Value)
			}
		}
	}
	if out == nil {
		return "", fmt.Errorf("lockupAddress: %v not found in the transaction: %v", lockupAddress, transactionHex)
	}
	return tx.Hash().String(), nil
}

// GetTransaction return the transaction after paying the ln invoice
func (boltz *Boltz) GetTransaction(id, lockupAddress string, amt int64) (status, txid, tx string, eta int, err error) {
	request := transactionRequest{ID: id}
	var response transactionStatus
	err = boltz.rest.Post(swapStatusEndpoint, request, &response)

	if err != nil {
		err = fmt.Errorf("json decode (status ok): %w", err)
		return
	}
	if response.Status != "transaction.mempool" && response.Status != "transaction.confirmed" {
		err = ErrSwapNotFound
		return
	}

	if lockupAddress != "" {
		var calculatedTxid string
		calculatedTxid, err = boltz.CheckTransaction(response.Transaction.Hex, lockupAddress, amt)
		if err != nil {
			err = fmt.Errorf("CheckTransaction(%v, %v, %v): %w)", response.Transaction.Hex, lockupAddress, amt, err)
			return
		}
		if calculatedTxid != response.Transaction.ID {
			err = fmt.Errorf("bad txid: %v != %v", response.Transaction.ID, calculatedTxid)
			return
		}
	}
	status = response.Status
	tx = response.Transaction.Hex
	txid = response.Transaction.ID
	eta = response.Transaction.ETA
	return
}

//ClaimFees return the fees needed for the claimed transaction for a feePerKw
func (boltz *Boltz) ClaimFee(claimAddress string, feePerKw int64) (int64, error) {
	addr, err := btcutil.DecodeAddress(claimAddress, boltz.chain)
	if err != nil {
		return 0, fmt.Errorf("btcutil.DecodeAddress(%v) %w", addr, err)
	}
	claimScript, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return 0, fmt.Errorf("txscript.PayToAddrScript(%v): %w", addr.String(), err)
	}
	claimTx := wire.NewMsgTx(1)
	txIn := wire.NewTxIn(&wire.OutPoint{}, nil, nil)
	txIn.Sequence = 0
	claimTx.AddTxIn(txIn)
	txOut := wire.TxOut{PkScript: claimScript}
	claimTx.AddTxOut(&txOut)

	// Calcluate the weight and the fee
	weight := 4*claimTx.SerializeSizeStripped() + claimWitnessInputSize*len(claimTx.TxIn)
	fee := chainfee.SatPerKWeight(feePerKw).FeeForWeight(int64(weight))
	return int64(fee), nil
}

// ClaimTransaction returns the claim transaction to broadcast after sending it
// also to boltz
func (boltz *Boltz) ClaimTransaction(
	redeemScript, transactionHex string,
	claimAddress string,
	preimage, key string,
	fees int64,
) (string, error) {
	txSerialized, err := hex.DecodeString(transactionHex)
	if err != nil {
		return "", fmt.Errorf("hex.DecodeString(%v): %w", transactionHex, err)
	}
	tx, err := btcutil.NewTxFromBytes(txSerialized)
	if err != nil {
		return "", fmt.Errorf("btcutil.NewTxFromBytes(%x): %w", txSerialized, err)
	}

	script, err := hex.DecodeString(redeemScript)
	if err != nil {
		return "", fmt.Errorf("hex.DecodeString(%v): %w", redeemScript, err)
	}
	lockupAddress, err := boltz.addressWitnessScriptHash(script, boltz.chain)
	if err != nil {
		return "", fmt.Errorf("addressWitnessScriptHash %v: %w", script, err)
	}
	var out *wire.OutPoint
	var amt btcutil.Amount
	for i, txout := range tx.MsgTx().TxOut {
		class, addresses, requiredsigs, err := txscript.ExtractPkScriptAddrs(txout.PkScript, boltz.chain)
		if err != nil {
			return "", fmt.Errorf("txscript.ExtractPkScriptAddrs(%x) %w", txout.PkScript, err)
		}
		if class == txscript.WitnessV0ScriptHashTy && requiredsigs == 1 &&
			len(addresses) == 1 && addresses[0].EncodeAddress() == lockupAddress.EncodeAddress() {
			out = wire.NewOutPoint(tx.Hash(), uint32(i))
			amt = btcutil.Amount(txout.Value)
		}
	}

	addr, err := btcutil.DecodeAddress(claimAddress, boltz.chain)
	if err != nil {
		return "", fmt.Errorf("btcutil.DecodeAddress(%v) %w", claimAddress, err)
	}

	preim, err := hex.DecodeString(preimage)
	if err != nil {
		return "", fmt.Errorf("hex.DecodeString(%v): %w", preimage, err)
	}
	privateKey, err := hex.DecodeString(key)
	if err != nil {
		return "", fmt.Errorf("hex.DecodeString(%v): %w", key, err)
	}

	ctx, err := boltz.claimTransaction(script, amt, out, addr, preim, privateKey, btcutil.Amount(fees))
	if err != nil {
		return "", fmt.Errorf("claimTransaction: %w", err)
	}
	ctxHex := hex.EncodeToString(ctx)
	//Ignore the result of broadcasting the transaction via boltz
	_, _ = boltz.broadcastTransaction(ctxHex)
	return ctxHex, nil
}

func (boltz *Boltz) GetNodePubkey() (string, error) {
	var nodes struct {
		Nodes map[string]struct {
			URIS    []string `json:"uris"`
			NodeKey string   `json:"nodeKey"`
		} `json:"nodes"`
	}
	err := boltz.rest.Get(getNodesEndpoint, nodes)
	if err != nil {
		return "", err
	}

	if b, ok := nodes.Nodes["BTC"]; ok {
		return b.NodeKey, nil
	}
	return "", fmt.Errorf("pubkey not found")
}
func (boltz *Boltz) GetRoutingHints(routingNode []byte) ([]RoutingHint, error) {
	var request = struct {
		Symbol      string `json:"symbol"`
		RoutingNode string `json:"routingNode"`
	}{
		Symbol:      "BTC",
		RoutingNode: hex.EncodeToString(routingNode),
	}

	var response struct {
		RoutingHints []RoutingHint `json:"routingHints"`
	}

	err := boltz.rest.Post(routingHintsEndpoint, request, &response)
	if err != nil {
		return nil, err
	}

	return response.RoutingHints, nil
}

/**
pairId = "BTC/BTC"
orderSide = "buy"
*/
func (boltz *Boltz) createReverseSwap(pairId string, orderSide string, amt int64, feesHash string, preimage []byte, preimageHash [32]byte, claimAddress string, routingNode []byte) (*boltzReverseSwap, error) {
	var request = struct {
		Type          string `json:"type"`
		PairID        string `json:"pairId"`
		OrderSide     string `json:"orderSide"`
		InvoiceAmount int64  `json:"invoiceAmount"`
		PreimageHash  string `json:"preimageHash"`
		PairHash      string `json:"pairHash,omitempty"`
		ClaimAddress  string `json:"claimAddress"` //ClaimPublicKey string `json:"claimPublicKey"`
		RoutingNode   string `json:"routingNode,omitempty"`
	}{
		Type:          "reversesubmarine",
		PairID:        pairId,
		OrderSide:     orderSide,
		InvoiceAmount: amt,
		PreimageHash:  hex.EncodeToString(preimageHash[:]),
		PairHash:      feesHash,
		ClaimAddress:  claimAddress, //ClaimPublicKey: hex.EncodeToString(key.PubKey().SerializeCompressed()), -> key *btcec.PrivateKey
		RoutingNode:   hex.EncodeToString(routingNode),
	}

	var response boltzReverseSwap

	err := boltz.rest.Post(createSwapEndpoint, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (boltz *Boltz) checkHeight(h int64, hs string) string {
	b1, err := hex.DecodeString(hs)
	if err != nil {
		return ""
	}
	b := make([]byte, 8)
	copy(b, b1)
	if binary.LittleEndian.Uint64(b) == uint64(h) {
		return hs
	}
	return ""
}

func (boltz *Boltz) addressWitnessScriptHash(script []byte, net *chaincfg.Params) (*btcutil.AddressWitnessScriptHash, error) {
	witnessProg := sha256.Sum256(script)
	return btcutil.NewAddressWitnessScriptHash(witnessProg[:], net)
}

func (boltz *Boltz) getPreimage() []byte {
	preimage := make([]byte, 32)
	rand.Read(preimage)
	return preimage
}

func (boltz *Boltz) getPrivate() (*btcec.PrivateKey, error) {
	k, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, fmt.Errorf("btcec.NewPrivateKey: %w", err)
	}
	return k, nil
}

func (boltz *Boltz) claimTransaction(
	script []byte,
	amt btcutil.Amount,
	txout *wire.OutPoint,
	claimAddress btcutil.Address,
	preimage []byte,
	privateKey []byte,
	fees btcutil.Amount,
) ([]byte, error) {
	claimTx := wire.NewMsgTx(1)
	txIn := wire.NewTxIn(txout, nil, nil)
	txIn.Sequence = 0
	claimTx.AddTxIn(txIn)

	claimScript, err := txscript.PayToAddrScript(claimAddress)
	if err != nil {
		return nil, fmt.Errorf("txscript.PayToAddrScript(%v): %w", claimAddress.String(), err)
	}
	txOut := wire.TxOut{PkScript: claimScript}
	claimTx.AddTxOut(&txOut)

	// Adjust the amount in the txout
	claimTx.TxOut[0].Value = int64(amt - fees)

	sigHashes := txscript.NewTxSigHashes(claimTx)
	key, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKey)
	scriptSig, err := txscript.RawTxInWitnessSignature(claimTx, sigHashes, 0, int64(amt), script, txscript.SigHashAll, key)
	if err != nil {
		return nil, fmt.Errorf("txscript.RawTxInWitnessSignature: %w", err)
	}
	claimTx.TxIn[0].Witness = [][]byte{scriptSig, preimage, script}

	var rawTx bytes.Buffer
	err = claimTx.Serialize(&rawTx)
	if err != nil {
		return nil, fmt.Errorf("claimTx.Serialize %#v: %w", claimTx, err)
	}
	return rawTx.Bytes(), nil
}

func (boltz *Boltz) broadcastTransaction(tx string) (string, error) {
	var request = struct {
		Currency       string `json:"currency"`
		TransactionHex string `json:"transactionHex"`
	}{"BTC", tx}

	var response struct {
		TransactionID string `json:"transactionId"`
	}

	err := boltz.rest.Post(broadcastTransactionEndpoint, request, &response)
	if err != nil {
		return "", err
	}

	return response.TransactionID, nil
}