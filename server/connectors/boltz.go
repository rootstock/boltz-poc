package connectors

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/patogallaiov/boltz-poc/abi"
	"github.com/patogallaiov/boltz-poc/abi/erc20"
	"github.com/patogallaiov/boltz-poc/config"
	"github.com/patogallaiov/boltz-poc/storage"
	log "github.com/sirupsen/logrus"
)

const (
	getNodesEndpoint             = "/getnodes"
	getPairsEndpoint             = "/getpairs"
	getContractsEndpoint         = "/getcontracts"
	createSwapEndpoint           = "/createswap"
	routingHintsEndpoint         = "/routinghints"
	swapStatusEndpoint           = "/swapstatus"
	broadcastTransactionEndpoint = "/broadcasttransaction"
	claimWitnessInputSize        = 1 + 1 + 8 + 73 + 1 + 32 + 1 + 100
	lockupEventHex               = "0x15b4b8206809535e547317cd5cedc86cff6e7d203551f93701786ddaf14fd9f9"
	lockupEventErc20Hex          = "0xa98eaa2bd8230d87a1a4c356f5c1d41cb85ff88131122ec8b1931cb9d31ae145"
)

var ErrSwapNotFound = errors.New("transaction not in mempool or settled/canceled")

type BoltzConnector interface {
	Initialize() error
	GetPair() (PairResponse, error)
	GetContracts() (ContractsResponse, error)
	GetReverseSwapInfo() (*ReverseSwapInfo, error)
	NewReverseSwap(pairId string, orderSide string, amt btcutil.Amount, feesHash string, routingNode []byte) (*ReverseSwap, error)
	GetNodePubkey() (string, error)
	GetRoutingHints(routingNode []byte) ([]RoutingHint, error)
}

type Boltz struct {
	rest             *RestClient
	apiURL           string
	chain            *chaincfg.Params
	claimAddress     string
	rsk              RSKConnector
	db               storage.DBConnector
	etherSwapAddress common.Address
	erc20SwapAddress common.Address
}

func NewBoltz(appCfg config.Config, chain *chaincfg.Params, claimAddress string, rsk RSKConnector, db storage.DBConnector) (*Boltz, error) {
	rest, _ := NewRestClient(appCfg.Boltz.Endpoint)
	return &Boltz{
		rest:         rest,
		apiURL:       appCfg.Boltz.Endpoint,
		chain:        chain,
		claimAddress: claimAddress,
		rsk:          rsk,
		db:           db,
	}, nil
}
func (boltz *Boltz) Initialize() error {
	// Check endpoint
	if _, err := boltz.GetReverseSwapInfo(); err != nil {
		return err
	}

	// Subscribe to RSK events
	contracts, err := boltz.GetContracts()
	if err != nil {
		return err
	}
	boltz.etherSwapAddress = common.HexToAddress(contracts.Ethereum.SwapContracts.EtherSwap)
	boltz.erc20SwapAddress = common.HexToAddress(contracts.Ethereum.SwapContracts.ERC20Swap)

	boltz.rsk.Subscribe([]common.Address{boltz.etherSwapAddress, boltz.erc20SwapAddress}, boltz.receiveEvent)
	return nil
}

func (boltz *Boltz) receiveEvent(elog gethTypes.Log) {
	log.Debugf("[BOLTZ-client] - Event received -> %v", elog)
	eventHash := elog.Topics[0].Hex()
	switch eventHash {
	case lockupEventHex:
		boltz.handleLockup(elog)
	case lockupEventErc20Hex:
		boltz.handleLockupERC20(elog)
	default:
		log.Infof("Event received, not mapped: %s", eventHash)
	}

}

func (boltz *Boltz) GetPair() (response PairResponse, err error) {
	err = boltz.rest.Get(getPairsEndpoint, &response)
	return
}

func (boltz *Boltz) GetContracts() (response ContractsResponse, err error) {
	err = boltz.rest.Get(getContractsEndpoint, &response)
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

	rs, err := boltz.createReverseSwap(pairId, orderSide, int64(amt), feesHash, preimage, preimageHash, boltz.rsk.GetAddress(), routingNode)
	if err != nil {
		return nil, fmt.Errorf("createReverseSwap amt:%v, preimage:%x, key:%x; %w", amt, preimage, key, err)
	}

	return &ReverseSwap{*rs, hex.EncodeToString(preimage), hex.EncodeToString(preimageHash[:]), hex.EncodeToString(key.Serialize())}, nil
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
 [
      { name: 'pairId', type: 'string' },
      { name: 'orderSide', type: 'string' },
      { name: 'preimageHash', type: 'string', hex: true },
      { name: 'pairHash', type: 'string', optional: true },
      { name: 'referralId', type: 'string', optional: true },
      { name: 'routingNode', type: 'string', optional: true },
      { name: 'claimAddress', type: 'string', optional: true, },
      { name: 'invoiceAmount', type: 'number', optional: true },
      { name: 'onchainAmount', type: 'number', optional: true },
      { name: 'prepayMinerFee', type: 'boolean', optional: true },
      { name: 'claimPublicKey', type: 'string', hex: true, optional: true },
    ]
*/
func (boltz *Boltz) createReverseSwap(pairId string, orderSide string, amt int64, feesHash string, preimage []byte, preimageHash [32]byte, claimAddress common.Address, routingNode []byte) (*boltzReverseSwap, error) {
	var request = struct {
		Type          string `json:"type"`
		PairID        string `json:"pairId"`
		OrderSide     string `json:"orderSide"`
		InvoiceAmount int64  `json:"invoiceAmount"`
		//OnchainAmount int64  `json:"onchainAmount"`
		PreimageHash string `json:"preimageHash"`
		PairHash     string `json:"pairHash,omitempty"`
		ClaimAddress string `json:"claimAddress"` //ClaimPublicKey string `json:"claimPublicKey"`
		RoutingNode  string `json:"routingNode,omitempty"`
	}{
		Type:          "reversesubmarine",
		PairID:        pairId,
		OrderSide:     orderSide,
		InvoiceAmount: amt,
		PreimageHash:  hex.EncodeToString(preimageHash[:]),
		PairHash:      feesHash,
		ClaimAddress:  claimAddress.Hex(), //ClaimPublicKey: hex.EncodeToString(key.PubKey().SerializeCompressed()), -> key *btcec.PrivateKey
		RoutingNode:   hex.EncodeToString(routingNode),
	}

	log.Debugf("Creating reverse swap: %v", request)

	var response boltzReverseSwap

	err := boltz.rest.Post(createSwapEndpoint, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
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

func (boltz *Boltz) handleLockup(elog gethTypes.Log) ([]byte, error) {
	log.Debugf("Lockup event: %v", elog)
	// Get Ether swap contract, and send Claim.
	swapContract, err := abi.NewAbi(boltz.etherSwapAddress, boltz.rsk.GetClient())
	if err != nil {
		log.Fatalf("Error getting contract abi: %v", err)
		return nil, err
	}
	auth, err := boltz.rsk.GetTransactor(bind.NewKeyedTransactorWithChainID)
	if err != nil {
		log.Fatalf("Error building transactor: %v", err)
		return nil, err
	}
	// Decode event parameters
	event, err := swapContract.AbiFilterer.ParseLockup(elog)
	if err != nil {
		return nil, err
	}
	log.Debugf("Lockup event decoded.")

	//Decode preimageHash from event
	preimageHash := hexutil.Encode(event.PreimageHash[:])
	preimageHash = strings.TrimPrefix(preimageHash, "0x")
	log.Debugf("Lockup received preimageHash: %s", preimageHash)

	// Get payment by preimageHash from DB.
	payment, err := boltz.db.GetPayment(preimageHash)
	if err != nil {
		log.Fatalf("Payment (%s) NOT found: %v", preimageHash, err)
		return nil, err
	}
	var preimage [32]byte = common.HexToHash(payment.Preimage)
	// todo(shree): this is where we diverge .. start using claimviaDocmint
	// where is the gaslimit set (only in ../rsk.go?)
	//auth.NoSend = true
	//result, err := swapContract.AbiTransactor.Claim(auth, preimage, event.Amount, event.RefundAddress, event.Timelock)
	log.Debugf("Lockup preimage: %s", payment.Preimage)
	log.Debugf("Lockup event.Amount: %s", event.Amount)
	log.Debugf("Lockup event.RefundAddress: %s", event.RefundAddress)
	log.Debugf("Lockup event.Timelock: %s", event.Timelock)
	log.Debugf("Lockup event.ClaimAddress: %s", event.ClaimAddress)
	result, err := swapContract.AbiTransactor.ClaimDoCViaMint(auth, preimage, event.Amount, event.RefundAddress, event.Timelock, new(big.Int).Div(event.Amount, big.NewInt(2)), event.ClaimAddress, event.ClaimAddress)

	if err != nil {
		log.Fatalf("Error executing claim tx: %v", err)
		payment.Status = "Error"
		boltz.db.SavePayment(&payment)
		return nil, err
	}
	log.Infof("Claim transaction sent data: %v", common.Bytes2Hex(result.Data()))
	payment.Status = "Completed"
	payment.Tx = result.Hash().Hex()
	boltz.db.SavePayment(&payment)
	return result.Data(), nil
}

func (boltz *Boltz) handleLockupERC20(elog gethTypes.Log) ([]byte, error) {
	log.Debugf("Lockup ERC20 event: %v", elog)
	// Get ERC20 swap contract, and send Claim.
	swapContract, err := erc20.NewAbi(boltz.erc20SwapAddress, boltz.rsk.GetClient())
	if err != nil {
		log.Fatalf("Error getting contract abi: %v", err)
		return nil, err
	}
	auth, err := boltz.rsk.GetTransactor(bind.NewKeyedTransactorWithChainID)
	if err != nil {
		log.Fatalf("Error building transactor: %v", err)
		return nil, err
	}

	event, err := swapContract.AbiFilterer.ParseLockup(elog)
	if err != nil {
		log.Fatalf("Error parsing lockup: %v", err)
		return nil, err
	}

	//Decode preimageHash from event
	preimageHash := hexutil.Encode(event.PreimageHash[:])
	preimageHash = strings.TrimPrefix(preimageHash, "0x")
	log.Debugf("Lockup ERC20 received preimageHash: %s", preimageHash)

	// Get payment by preimageHash from DB.
	payment, err := boltz.db.GetPayment(preimageHash)
	if err != nil {
		log.Fatalf("Payment (%s) NOT found: %v", preimageHash, err)
		return nil, err
	}
	preimage := common.HexToHash(payment.Preimage)

	log.Debugf("Claiming auth.From: %s", auth.From)
	//auth.NoSend = true
	result, err := swapContract.AbiTransactor.Claim(auth, preimage, event.Amount, event.TokenAddress, event.RefundAddress, event.Timelock)
	if err != nil {
		log.Fatalf("Error executing claim tx: %v", err)
		return nil, err
	}
	log.Infof("Claim transaction sent data: %v", common.Bytes2Hex(result.Data()))
	return result.Data(), nil
}
