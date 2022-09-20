// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_mocAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_docAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ChangeRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"preimage\",\"type\":\"bytes32\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"preimage\",\"type\":\"bytes32\"}],\"name\":\"ClaimDocViaMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"}],\"name\":\"Lockup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Minting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferredDoc\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"preimage\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"preimage\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"btcToMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"docReceiverAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"leftoverRbtcAddr\",\"type\":\"address\"}],\"name\":\"claimDoCViaMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"claimAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"}],\"name\":\"hashValues\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"claimAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"claimAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prepayAmount\",\"type\":\"uint256\"}],\"name\":\"lockPrepayMinerfee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"claimAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"swaps\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// HashValues is a free data retrieval call binding the contract method 0x8b2f8f82.
//
// Solidity: function hashValues(bytes32 preimageHash, uint256 amount, address claimAddress, address refundAddress, uint256 timelock) pure returns(bytes32)
func (_Abi *AbiCaller) HashValues(opts *bind.CallOpts, preimageHash [32]byte, amount *big.Int, claimAddress common.Address, refundAddress common.Address, timelock *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "hashValues", preimageHash, amount, claimAddress, refundAddress, timelock)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashValues is a free data retrieval call binding the contract method 0x8b2f8f82.
//
// Solidity: function hashValues(bytes32 preimageHash, uint256 amount, address claimAddress, address refundAddress, uint256 timelock) pure returns(bytes32)
func (_Abi *AbiSession) HashValues(preimageHash [32]byte, amount *big.Int, claimAddress common.Address, refundAddress common.Address, timelock *big.Int) ([32]byte, error) {
	return _Abi.Contract.HashValues(&_Abi.CallOpts, preimageHash, amount, claimAddress, refundAddress, timelock)
}

// HashValues is a free data retrieval call binding the contract method 0x8b2f8f82.
//
// Solidity: function hashValues(bytes32 preimageHash, uint256 amount, address claimAddress, address refundAddress, uint256 timelock) pure returns(bytes32)
func (_Abi *AbiCallerSession) HashValues(preimageHash [32]byte, amount *big.Int, claimAddress common.Address, refundAddress common.Address, timelock *big.Int) ([32]byte, error) {
	return _Abi.Contract.HashValues(&_Abi.CallOpts, preimageHash, amount, claimAddress, refundAddress, timelock)
}

// Swaps is a free data retrieval call binding the contract method 0xeb84e7f2.
//
// Solidity: function swaps(bytes32 ) view returns(bool)
func (_Abi *AbiCaller) Swaps(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "swaps", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Swaps is a free data retrieval call binding the contract method 0xeb84e7f2.
//
// Solidity: function swaps(bytes32 ) view returns(bool)
func (_Abi *AbiSession) Swaps(arg0 [32]byte) (bool, error) {
	return _Abi.Contract.Swaps(&_Abi.CallOpts, arg0)
}

// Swaps is a free data retrieval call binding the contract method 0xeb84e7f2.
//
// Solidity: function swaps(bytes32 ) view returns(bool)
func (_Abi *AbiCallerSession) Swaps(arg0 [32]byte) (bool, error) {
	return _Abi.Contract.Swaps(&_Abi.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_Abi *AbiCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_Abi *AbiSession) Version() (uint8, error) {
	return _Abi.Contract.Version(&_Abi.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_Abi *AbiCallerSession) Version() (uint8, error) {
	return _Abi.Contract.Version(&_Abi.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xc3c37fbc.
//
// Solidity: function claim(bytes32 preimage, uint256 amount, address refundAddress, uint256 timelock) returns()
func (_Abi *AbiTransactor) Claim(opts *bind.TransactOpts, preimage [32]byte, amount *big.Int, refundAddress common.Address, timelock *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "claim", preimage, amount, refundAddress, timelock)
}

// Claim is a paid mutator transaction binding the contract method 0xc3c37fbc.
//
// Solidity: function claim(bytes32 preimage, uint256 amount, address refundAddress, uint256 timelock) returns()
func (_Abi *AbiSession) Claim(preimage [32]byte, amount *big.Int, refundAddress common.Address, timelock *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Claim(&_Abi.TransactOpts, preimage, amount, refundAddress, timelock)
}

// Claim is a paid mutator transaction binding the contract method 0xc3c37fbc.
//
// Solidity: function claim(bytes32 preimage, uint256 amount, address refundAddress, uint256 timelock) returns()
func (_Abi *AbiTransactorSession) Claim(preimage [32]byte, amount *big.Int, refundAddress common.Address, timelock *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Claim(&_Abi.TransactOpts, preimage, amount, refundAddress, timelock)
}

// ClaimDoCViaMint is a paid mutator transaction binding the contract method 0xa179c70f.
//
// Solidity: function claimDoCViaMint(bytes32 preimage, uint256 amount, address refundAddress, uint256 timelock, uint256 btcToMint, address docReceiverAddress, address leftoverRbtcAddr) returns()
func (_Abi *AbiTransactor) ClaimDoCViaMint(opts *bind.TransactOpts, preimage [32]byte, amount *big.Int, refundAddress common.Address, timelock *big.Int, btcToMint *big.Int, docReceiverAddress common.Address, leftoverRbtcAddr common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "claimDoCViaMint", preimage, amount, refundAddress, timelock, btcToMint, docReceiverAddress, leftoverRbtcAddr)
}

// ClaimDoCViaMint is a paid mutator transaction binding the contract method 0xa179c70f.
//
// Solidity: function claimDoCViaMint(bytes32 preimage, uint256 amount, address refundAddress, uint256 timelock, uint256 btcToMint, address docReceiverAddress, address leftoverRbtcAddr) returns()
func (_Abi *AbiSession) ClaimDoCViaMint(preimage [32]byte, amount *big.Int, refundAddress common.Address, timelock *big.Int, btcToMint *big.Int, docReceiverAddress common.Address, leftoverRbtcAddr common.Address) (*types.Transaction, error) {
	return _Abi.Contract.ClaimDoCViaMint(&_Abi.TransactOpts, preimage, amount, refundAddress, timelock, btcToMint, docReceiverAddress, leftoverRbtcAddr)
}

// ClaimDoCViaMint is a paid mutator transaction binding the contract method 0xa179c70f.
//
// Solidity: function claimDoCViaMint(bytes32 preimage, uint256 amount, address refundAddress, uint256 timelock, uint256 btcToMint, address docReceiverAddress, address leftoverRbtcAddr) returns()
func (_Abi *AbiTransactorSession) ClaimDoCViaMint(preimage [32]byte, amount *big.Int, refundAddress common.Address, timelock *big.Int, btcToMint *big.Int, docReceiverAddress common.Address, leftoverRbtcAddr common.Address) (*types.Transaction, error) {
	return _Abi.Contract.ClaimDoCViaMint(&_Abi.TransactOpts, preimage, amount, refundAddress, timelock, btcToMint, docReceiverAddress, leftoverRbtcAddr)
}

// Lock is a paid mutator transaction binding the contract method 0x0899146b.
//
// Solidity: function lock(bytes32 preimageHash, address claimAddress, uint256 timelock) payable returns()
func (_Abi *AbiTransactor) Lock(opts *bind.TransactOpts, preimageHash [32]byte, claimAddress common.Address, timelock *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "lock", preimageHash, claimAddress, timelock)
}

// Lock is a paid mutator transaction binding the contract method 0x0899146b.
//
// Solidity: function lock(bytes32 preimageHash, address claimAddress, uint256 timelock) payable returns()
func (_Abi *AbiSession) Lock(preimageHash [32]byte, claimAddress common.Address, timelock *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Lock(&_Abi.TransactOpts, preimageHash, claimAddress, timelock)
}

// Lock is a paid mutator transaction binding the contract method 0x0899146b.
//
// Solidity: function lock(bytes32 preimageHash, address claimAddress, uint256 timelock) payable returns()
func (_Abi *AbiTransactorSession) Lock(preimageHash [32]byte, claimAddress common.Address, timelock *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Lock(&_Abi.TransactOpts, preimageHash, claimAddress, timelock)
}

// LockPrepayMinerfee is a paid mutator transaction binding the contract method 0x6fa4ae60.
//
// Solidity: function lockPrepayMinerfee(bytes32 preimageHash, address claimAddress, uint256 timelock, uint256 prepayAmount) payable returns()
func (_Abi *AbiTransactor) LockPrepayMinerfee(opts *bind.TransactOpts, preimageHash [32]byte, claimAddress common.Address, timelock *big.Int, prepayAmount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "lockPrepayMinerfee", preimageHash, claimAddress, timelock, prepayAmount)
}

// LockPrepayMinerfee is a paid mutator transaction binding the contract method 0x6fa4ae60.
//
// Solidity: function lockPrepayMinerfee(bytes32 preimageHash, address claimAddress, uint256 timelock, uint256 prepayAmount) payable returns()
func (_Abi *AbiSession) LockPrepayMinerfee(preimageHash [32]byte, claimAddress common.Address, timelock *big.Int, prepayAmount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.LockPrepayMinerfee(&_Abi.TransactOpts, preimageHash, claimAddress, timelock, prepayAmount)
}

// LockPrepayMinerfee is a paid mutator transaction binding the contract method 0x6fa4ae60.
//
// Solidity: function lockPrepayMinerfee(bytes32 preimageHash, address claimAddress, uint256 timelock, uint256 prepayAmount) payable returns()
func (_Abi *AbiTransactorSession) LockPrepayMinerfee(preimageHash [32]byte, claimAddress common.Address, timelock *big.Int, prepayAmount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.LockPrepayMinerfee(&_Abi.TransactOpts, preimageHash, claimAddress, timelock, prepayAmount)
}

// Refund is a paid mutator transaction binding the contract method 0x35cd4ccb.
//
// Solidity: function refund(bytes32 preimageHash, uint256 amount, address claimAddress, uint256 timelock) returns()
func (_Abi *AbiTransactor) Refund(opts *bind.TransactOpts, preimageHash [32]byte, amount *big.Int, claimAddress common.Address, timelock *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "refund", preimageHash, amount, claimAddress, timelock)
}

// Refund is a paid mutator transaction binding the contract method 0x35cd4ccb.
//
// Solidity: function refund(bytes32 preimageHash, uint256 amount, address claimAddress, uint256 timelock) returns()
func (_Abi *AbiSession) Refund(preimageHash [32]byte, amount *big.Int, claimAddress common.Address, timelock *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Refund(&_Abi.TransactOpts, preimageHash, amount, claimAddress, timelock)
}

// Refund is a paid mutator transaction binding the contract method 0x35cd4ccb.
//
// Solidity: function refund(bytes32 preimageHash, uint256 amount, address claimAddress, uint256 timelock) returns()
func (_Abi *AbiTransactorSession) Refund(preimageHash [32]byte, amount *big.Int, claimAddress common.Address, timelock *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Refund(&_Abi.TransactOpts, preimageHash, amount, claimAddress, timelock)
}

// AbiChangeRefundIterator is returned from FilterChangeRefund and is used to iterate over the raw logs and unpacked data for ChangeRefund events raised by the Abi contract.
type AbiChangeRefundIterator struct {
	Event *AbiChangeRefund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiChangeRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiChangeRefund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiChangeRefund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiChangeRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiChangeRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiChangeRefund represents a ChangeRefund event raised by the Abi contract.
type AbiChangeRefund struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangeRefund is a free log retrieval operation binding the contract event 0x5b089f4746873d9e61e20908bcec0ab9b916b18894f34e506847aa1a18f98354.
//
// Solidity: event ChangeRefund(uint256 value)
func (_Abi *AbiFilterer) FilterChangeRefund(opts *bind.FilterOpts) (*AbiChangeRefundIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "ChangeRefund")
	if err != nil {
		return nil, err
	}
	return &AbiChangeRefundIterator{contract: _Abi.contract, event: "ChangeRefund", logs: logs, sub: sub}, nil
}

// WatchChangeRefund is a free log subscription operation binding the contract event 0x5b089f4746873d9e61e20908bcec0ab9b916b18894f34e506847aa1a18f98354.
//
// Solidity: event ChangeRefund(uint256 value)
func (_Abi *AbiFilterer) WatchChangeRefund(opts *bind.WatchOpts, sink chan<- *AbiChangeRefund) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "ChangeRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiChangeRefund)
				if err := _Abi.contract.UnpackLog(event, "ChangeRefund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangeRefund is a log parse operation binding the contract event 0x5b089f4746873d9e61e20908bcec0ab9b916b18894f34e506847aa1a18f98354.
//
// Solidity: event ChangeRefund(uint256 value)
func (_Abi *AbiFilterer) ParseChangeRefund(log types.Log) (*AbiChangeRefund, error) {
	event := new(AbiChangeRefund)
	if err := _Abi.contract.UnpackLog(event, "ChangeRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Abi contract.
type AbiClaimIterator struct {
	Event *AbiClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiClaim represents a Claim event raised by the Abi contract.
type AbiClaim struct {
	PreimageHash [32]byte
	Preimage     [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x5664142af3dcfc3dc3de45a43f75c746bd1d8c11170a5037fdf98bdb35775137.
//
// Solidity: event Claim(bytes32 indexed preimageHash, bytes32 preimage)
func (_Abi *AbiFilterer) FilterClaim(opts *bind.FilterOpts, preimageHash [][32]byte) (*AbiClaimIterator, error) {

	var preimageHashRule []interface{}
	for _, preimageHashItem := range preimageHash {
		preimageHashRule = append(preimageHashRule, preimageHashItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Claim", preimageHashRule)
	if err != nil {
		return nil, err
	}
	return &AbiClaimIterator{contract: _Abi.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x5664142af3dcfc3dc3de45a43f75c746bd1d8c11170a5037fdf98bdb35775137.
//
// Solidity: event Claim(bytes32 indexed preimageHash, bytes32 preimage)
func (_Abi *AbiFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *AbiClaim, preimageHash [][32]byte) (event.Subscription, error) {

	var preimageHashRule []interface{}
	for _, preimageHashItem := range preimageHash {
		preimageHashRule = append(preimageHashRule, preimageHashItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Claim", preimageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiClaim)
				if err := _Abi.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x5664142af3dcfc3dc3de45a43f75c746bd1d8c11170a5037fdf98bdb35775137.
//
// Solidity: event Claim(bytes32 indexed preimageHash, bytes32 preimage)
func (_Abi *AbiFilterer) ParseClaim(log types.Log) (*AbiClaim, error) {
	event := new(AbiClaim)
	if err := _Abi.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiClaimDocViaMintIterator is returned from FilterClaimDocViaMint and is used to iterate over the raw logs and unpacked data for ClaimDocViaMint events raised by the Abi contract.
type AbiClaimDocViaMintIterator struct {
	Event *AbiClaimDocViaMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiClaimDocViaMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiClaimDocViaMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiClaimDocViaMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiClaimDocViaMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiClaimDocViaMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiClaimDocViaMint represents a ClaimDocViaMint event raised by the Abi contract.
type AbiClaimDocViaMint struct {
	PreimageHash [32]byte
	Preimage     [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaimDocViaMint is a free log retrieval operation binding the contract event 0xe84c7e87cfb5d2a2c9dd4263e0c8942861dd96a78ec9b260a10c3588903e0832.
//
// Solidity: event ClaimDocViaMint(bytes32 indexed preimageHash, bytes32 preimage)
func (_Abi *AbiFilterer) FilterClaimDocViaMint(opts *bind.FilterOpts, preimageHash [][32]byte) (*AbiClaimDocViaMintIterator, error) {

	var preimageHashRule []interface{}
	for _, preimageHashItem := range preimageHash {
		preimageHashRule = append(preimageHashRule, preimageHashItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "ClaimDocViaMint", preimageHashRule)
	if err != nil {
		return nil, err
	}
	return &AbiClaimDocViaMintIterator{contract: _Abi.contract, event: "ClaimDocViaMint", logs: logs, sub: sub}, nil
}

// WatchClaimDocViaMint is a free log subscription operation binding the contract event 0xe84c7e87cfb5d2a2c9dd4263e0c8942861dd96a78ec9b260a10c3588903e0832.
//
// Solidity: event ClaimDocViaMint(bytes32 indexed preimageHash, bytes32 preimage)
func (_Abi *AbiFilterer) WatchClaimDocViaMint(opts *bind.WatchOpts, sink chan<- *AbiClaimDocViaMint, preimageHash [][32]byte) (event.Subscription, error) {

	var preimageHashRule []interface{}
	for _, preimageHashItem := range preimageHash {
		preimageHashRule = append(preimageHashRule, preimageHashItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "ClaimDocViaMint", preimageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiClaimDocViaMint)
				if err := _Abi.contract.UnpackLog(event, "ClaimDocViaMint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimDocViaMint is a log parse operation binding the contract event 0xe84c7e87cfb5d2a2c9dd4263e0c8942861dd96a78ec9b260a10c3588903e0832.
//
// Solidity: event ClaimDocViaMint(bytes32 indexed preimageHash, bytes32 preimage)
func (_Abi *AbiFilterer) ParseClaimDocViaMint(log types.Log) (*AbiClaimDocViaMint, error) {
	event := new(AbiClaimDocViaMint)
	if err := _Abi.contract.UnpackLog(event, "ClaimDocViaMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiLockupIterator is returned from FilterLockup and is used to iterate over the raw logs and unpacked data for Lockup events raised by the Abi contract.
type AbiLockupIterator struct {
	Event *AbiLockup // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiLockupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiLockup)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiLockup)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiLockupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiLockupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiLockup represents a Lockup event raised by the Abi contract.
type AbiLockup struct {
	PreimageHash  [32]byte
	Amount        *big.Int
	ClaimAddress  common.Address
	RefundAddress common.Address
	Timelock      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLockup is a free log retrieval operation binding the contract event 0x15b4b8206809535e547317cd5cedc86cff6e7d203551f93701786ddaf14fd9f9.
//
// Solidity: event Lockup(bytes32 indexed preimageHash, uint256 amount, address claimAddress, address indexed refundAddress, uint256 timelock)
func (_Abi *AbiFilterer) FilterLockup(opts *bind.FilterOpts, preimageHash [][32]byte, refundAddress []common.Address) (*AbiLockupIterator, error) {

	var preimageHashRule []interface{}
	for _, preimageHashItem := range preimageHash {
		preimageHashRule = append(preimageHashRule, preimageHashItem)
	}

	var refundAddressRule []interface{}
	for _, refundAddressItem := range refundAddress {
		refundAddressRule = append(refundAddressRule, refundAddressItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Lockup", preimageHashRule, refundAddressRule)
	if err != nil {
		return nil, err
	}
	return &AbiLockupIterator{contract: _Abi.contract, event: "Lockup", logs: logs, sub: sub}, nil
}

// WatchLockup is a free log subscription operation binding the contract event 0x15b4b8206809535e547317cd5cedc86cff6e7d203551f93701786ddaf14fd9f9.
//
// Solidity: event Lockup(bytes32 indexed preimageHash, uint256 amount, address claimAddress, address indexed refundAddress, uint256 timelock)
func (_Abi *AbiFilterer) WatchLockup(opts *bind.WatchOpts, sink chan<- *AbiLockup, preimageHash [][32]byte, refundAddress []common.Address) (event.Subscription, error) {

	var preimageHashRule []interface{}
	for _, preimageHashItem := range preimageHash {
		preimageHashRule = append(preimageHashRule, preimageHashItem)
	}

	var refundAddressRule []interface{}
	for _, refundAddressItem := range refundAddress {
		refundAddressRule = append(refundAddressRule, refundAddressItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Lockup", preimageHashRule, refundAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiLockup)
				if err := _Abi.contract.UnpackLog(event, "Lockup", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockup is a log parse operation binding the contract event 0x15b4b8206809535e547317cd5cedc86cff6e7d203551f93701786ddaf14fd9f9.
//
// Solidity: event Lockup(bytes32 indexed preimageHash, uint256 amount, address claimAddress, address indexed refundAddress, uint256 timelock)
func (_Abi *AbiFilterer) ParseLockup(log types.Log) (*AbiLockup, error) {
	event := new(AbiLockup)
	if err := _Abi.contract.UnpackLog(event, "Lockup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Abi contract.
type AbiMintedIterator struct {
	Event *AbiMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiMinted represents a Minted event raised by the Abi contract.
type AbiMinted struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x176b02bb2d12439ff7a20b59f402cca16c76f50508b13ef3166a600eb719354a.
//
// Solidity: event Minted(uint256 value)
func (_Abi *AbiFilterer) FilterMinted(opts *bind.FilterOpts) (*AbiMintedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &AbiMintedIterator{contract: _Abi.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x176b02bb2d12439ff7a20b59f402cca16c76f50508b13ef3166a600eb719354a.
//
// Solidity: event Minted(uint256 value)
func (_Abi *AbiFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *AbiMinted) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiMinted)
				if err := _Abi.contract.UnpackLog(event, "Minted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinted is a log parse operation binding the contract event 0x176b02bb2d12439ff7a20b59f402cca16c76f50508b13ef3166a600eb719354a.
//
// Solidity: event Minted(uint256 value)
func (_Abi *AbiFilterer) ParseMinted(log types.Log) (*AbiMinted, error) {
	event := new(AbiMinted)
	if err := _Abi.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiMintingIterator is returned from FilterMinting and is used to iterate over the raw logs and unpacked data for Minting events raised by the Abi contract.
type AbiMintingIterator struct {
	Event *AbiMinting // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiMintingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiMinting)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiMinting)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiMintingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiMintingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiMinting represents a Minting event raised by the Abi contract.
type AbiMinting struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinting is a free log retrieval operation binding the contract event 0x867eb97d3472487867f1142a56e0853d70ef8258f4c73f7d515a4210045dda24.
//
// Solidity: event Minting(uint256 value)
func (_Abi *AbiFilterer) FilterMinting(opts *bind.FilterOpts) (*AbiMintingIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Minting")
	if err != nil {
		return nil, err
	}
	return &AbiMintingIterator{contract: _Abi.contract, event: "Minting", logs: logs, sub: sub}, nil
}

// WatchMinting is a free log subscription operation binding the contract event 0x867eb97d3472487867f1142a56e0853d70ef8258f4c73f7d515a4210045dda24.
//
// Solidity: event Minting(uint256 value)
func (_Abi *AbiFilterer) WatchMinting(opts *bind.WatchOpts, sink chan<- *AbiMinting) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Minting")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiMinting)
				if err := _Abi.contract.UnpackLog(event, "Minting", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinting is a log parse operation binding the contract event 0x867eb97d3472487867f1142a56e0853d70ef8258f4c73f7d515a4210045dda24.
//
// Solidity: event Minting(uint256 value)
func (_Abi *AbiFilterer) ParseMinting(log types.Log) (*AbiMinting, error) {
	event := new(AbiMinting)
	if err := _Abi.contract.UnpackLog(event, "Minting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the Abi contract.
type AbiRefundIterator struct {
	Event *AbiRefund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRefund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiRefund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRefund represents a Refund event raised by the Abi contract.
type AbiRefund struct {
	PreimageHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x3fbd469ec3a5ce074f975f76ce27e727ba21c99176917b97ae2e713695582a12.
//
// Solidity: event Refund(bytes32 indexed preimageHash)
func (_Abi *AbiFilterer) FilterRefund(opts *bind.FilterOpts, preimageHash [][32]byte) (*AbiRefundIterator, error) {

	var preimageHashRule []interface{}
	for _, preimageHashItem := range preimageHash {
		preimageHashRule = append(preimageHashRule, preimageHashItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Refund", preimageHashRule)
	if err != nil {
		return nil, err
	}
	return &AbiRefundIterator{contract: _Abi.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x3fbd469ec3a5ce074f975f76ce27e727ba21c99176917b97ae2e713695582a12.
//
// Solidity: event Refund(bytes32 indexed preimageHash)
func (_Abi *AbiFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *AbiRefund, preimageHash [][32]byte) (event.Subscription, error) {

	var preimageHashRule []interface{}
	for _, preimageHashItem := range preimageHash {
		preimageHashRule = append(preimageHashRule, preimageHashItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Refund", preimageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRefund)
				if err := _Abi.contract.UnpackLog(event, "Refund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefund is a log parse operation binding the contract event 0x3fbd469ec3a5ce074f975f76ce27e727ba21c99176917b97ae2e713695582a12.
//
// Solidity: event Refund(bytes32 indexed preimageHash)
func (_Abi *AbiFilterer) ParseRefund(log types.Log) (*AbiRefund, error) {
	event := new(AbiRefund)
	if err := _Abi.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiTransferredDocIterator is returned from FilterTransferredDoc and is used to iterate over the raw logs and unpacked data for TransferredDoc events raised by the Abi contract.
type AbiTransferredDocIterator struct {
	Event *AbiTransferredDoc // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbiTransferredDocIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTransferredDoc)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbiTransferredDoc)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbiTransferredDocIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTransferredDocIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTransferredDoc represents a TransferredDoc event raised by the Abi contract.
type AbiTransferredDoc struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferredDoc is a free log retrieval operation binding the contract event 0xade78c61bd38ba5e6104aa4db36526d71980db1f2d860531f78031507cc90b82.
//
// Solidity: event TransferredDoc(uint256 value)
func (_Abi *AbiFilterer) FilterTransferredDoc(opts *bind.FilterOpts) (*AbiTransferredDocIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "TransferredDoc")
	if err != nil {
		return nil, err
	}
	return &AbiTransferredDocIterator{contract: _Abi.contract, event: "TransferredDoc", logs: logs, sub: sub}, nil
}

// WatchTransferredDoc is a free log subscription operation binding the contract event 0xade78c61bd38ba5e6104aa4db36526d71980db1f2d860531f78031507cc90b82.
//
// Solidity: event TransferredDoc(uint256 value)
func (_Abi *AbiFilterer) WatchTransferredDoc(opts *bind.WatchOpts, sink chan<- *AbiTransferredDoc) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "TransferredDoc")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTransferredDoc)
				if err := _Abi.contract.UnpackLog(event, "TransferredDoc", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferredDoc is a log parse operation binding the contract event 0xade78c61bd38ba5e6104aa4db36526d71980db1f2d860531f78031507cc90b82.
//
// Solidity: event TransferredDoc(uint256 value)
func (_Abi *AbiFilterer) ParseTransferredDoc(log types.Log) (*AbiTransferredDoc, error) {
	event := new(AbiTransferredDoc)
	if err := _Abi.contract.UnpackLog(event, "TransferredDoc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
