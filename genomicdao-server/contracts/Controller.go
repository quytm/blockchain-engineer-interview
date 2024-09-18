// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// ControllerDataDoc is an auto generated low-level Go binding around an user-defined struct.
type ControllerDataDoc struct {
	Id          string
	HashContent string
}

// ControllerUploadSession is an auto generated low-level Go binding around an user-defined struct.
type ControllerUploadSession struct {
	Id        *big.Int
	User      common.Address
	Proof     string
	Confirmed bool
}

// ControllerMetaData contains all meta data concerning the Controller contract.
var ControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pcspAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"grantedTo\",\"type\":\"address\"}],\"name\":\"GrantedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"MintedGNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"revokedFrom\",\"type\":\"address\"}],\"name\":\"RevokedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"riskScore\",\"type\":\"uint256\"}],\"name\":\"RewardedPCSP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"UploadData\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contentHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"riskScore\",\"type\":\"uint256\"}],\"name\":\"confirm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"geneNFT\",\"outputs\":[{\"internalType\":\"contractGeneNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"}],\"name\":\"getDoc\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hashContent\",\"type\":\"string\"}],\"internalType\":\"structController.DataDoc\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"getSession\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"confirmed\",\"type\":\"bool\"}],\"internalType\":\"structController.UploadSession\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pcspToken\",\"outputs\":[{\"internalType\":\"contractPostCovidStrokePrevention\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"requestAccess\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hashContent\",\"type\":\"string\"}],\"internalType\":\"structController.DataDoc\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokeAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"}],\"name\":\"uploadData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerMetaData.ABI instead.
var ControllerABI = ControllerMetaData.ABI

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// GeneNFT is a free data retrieval call binding the contract method 0x5231f627.
//
// Solidity: function geneNFT() view returns(address)
func (_Controller *ControllerCaller) GeneNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "geneNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneNFT is a free data retrieval call binding the contract method 0x5231f627.
//
// Solidity: function geneNFT() view returns(address)
func (_Controller *ControllerSession) GeneNFT() (common.Address, error) {
	return _Controller.Contract.GeneNFT(&_Controller.CallOpts)
}

// GeneNFT is a free data retrieval call binding the contract method 0x5231f627.
//
// Solidity: function geneNFT() view returns(address)
func (_Controller *ControllerCallerSession) GeneNFT() (common.Address, error) {
	return _Controller.Contract.GeneNFT(&_Controller.CallOpts)
}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string docId) view returns((string,string))
func (_Controller *ControllerCaller) GetDoc(opts *bind.CallOpts, docId string) (ControllerDataDoc, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getDoc", docId)

	if err != nil {
		return *new(ControllerDataDoc), err
	}

	out0 := *abi.ConvertType(out[0], new(ControllerDataDoc)).(*ControllerDataDoc)

	return out0, err

}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string docId) view returns((string,string))
func (_Controller *ControllerSession) GetDoc(docId string) (ControllerDataDoc, error) {
	return _Controller.Contract.GetDoc(&_Controller.CallOpts, docId)
}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string docId) view returns((string,string))
func (_Controller *ControllerCallerSession) GetDoc(docId string) (ControllerDataDoc, error) {
	return _Controller.Contract.GetDoc(&_Controller.CallOpts, docId)
}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((uint256,address,string,bool))
func (_Controller *ControllerCaller) GetSession(opts *bind.CallOpts, sessionId *big.Int) (ControllerUploadSession, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getSession", sessionId)

	if err != nil {
		return *new(ControllerUploadSession), err
	}

	out0 := *abi.ConvertType(out[0], new(ControllerUploadSession)).(*ControllerUploadSession)

	return out0, err

}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((uint256,address,string,bool))
func (_Controller *ControllerSession) GetSession(sessionId *big.Int) (ControllerUploadSession, error) {
	return _Controller.Contract.GetSession(&_Controller.CallOpts, sessionId)
}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((uint256,address,string,bool))
func (_Controller *ControllerCallerSession) GetSession(sessionId *big.Int) (ControllerUploadSession, error) {
	return _Controller.Contract.GetSession(&_Controller.CallOpts, sessionId)
}

// PcspToken is a free data retrieval call binding the contract method 0xdab3761e.
//
// Solidity: function pcspToken() view returns(address)
func (_Controller *ControllerCaller) PcspToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "pcspToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PcspToken is a free data retrieval call binding the contract method 0xdab3761e.
//
// Solidity: function pcspToken() view returns(address)
func (_Controller *ControllerSession) PcspToken() (common.Address, error) {
	return _Controller.Contract.PcspToken(&_Controller.CallOpts)
}

// PcspToken is a free data retrieval call binding the contract method 0xdab3761e.
//
// Solidity: function pcspToken() view returns(address)
func (_Controller *ControllerCallerSession) PcspToken() (common.Address, error) {
	return _Controller.Contract.PcspToken(&_Controller.CallOpts)
}

// RequestAccess is a free data retrieval call binding the contract method 0x7fdc2cdc.
//
// Solidity: function requestAccess(uint256 nftId) view returns((string,string))
func (_Controller *ControllerCaller) RequestAccess(opts *bind.CallOpts, nftId *big.Int) (ControllerDataDoc, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "requestAccess", nftId)

	if err != nil {
		return *new(ControllerDataDoc), err
	}

	out0 := *abi.ConvertType(out[0], new(ControllerDataDoc)).(*ControllerDataDoc)

	return out0, err

}

// RequestAccess is a free data retrieval call binding the contract method 0x7fdc2cdc.
//
// Solidity: function requestAccess(uint256 nftId) view returns((string,string))
func (_Controller *ControllerSession) RequestAccess(nftId *big.Int) (ControllerDataDoc, error) {
	return _Controller.Contract.RequestAccess(&_Controller.CallOpts, nftId)
}

// RequestAccess is a free data retrieval call binding the contract method 0x7fdc2cdc.
//
// Solidity: function requestAccess(uint256 nftId) view returns((string,string))
func (_Controller *ControllerCallerSession) RequestAccess(nftId *big.Int) (ControllerDataDoc, error) {
	return _Controller.Contract.RequestAccess(&_Controller.CallOpts, nftId)
}

// Confirm is a paid mutator transaction binding the contract method 0xb62fdfce.
//
// Solidity: function confirm(string docId, string contentHash, string proof, uint256 sessionId, uint256 riskScore) returns(uint256)
func (_Controller *ControllerTransactor) Confirm(opts *bind.TransactOpts, docId string, contentHash string, proof string, sessionId *big.Int, riskScore *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "confirm", docId, contentHash, proof, sessionId, riskScore)
}

// Confirm is a paid mutator transaction binding the contract method 0xb62fdfce.
//
// Solidity: function confirm(string docId, string contentHash, string proof, uint256 sessionId, uint256 riskScore) returns(uint256)
func (_Controller *ControllerSession) Confirm(docId string, contentHash string, proof string, sessionId *big.Int, riskScore *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Confirm(&_Controller.TransactOpts, docId, contentHash, proof, sessionId, riskScore)
}

// Confirm is a paid mutator transaction binding the contract method 0xb62fdfce.
//
// Solidity: function confirm(string docId, string contentHash, string proof, uint256 sessionId, uint256 riskScore) returns(uint256)
func (_Controller *ControllerTransactorSession) Confirm(docId string, contentHash string, proof string, sessionId *big.Int, riskScore *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.Confirm(&_Controller.TransactOpts, docId, contentHash, proof, sessionId, riskScore)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x65dd152c.
//
// Solidity: function grantAccess(uint256 nftId, address user) returns()
func (_Controller *ControllerTransactor) GrantAccess(opts *bind.TransactOpts, nftId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "grantAccess", nftId, user)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x65dd152c.
//
// Solidity: function grantAccess(uint256 nftId, address user) returns()
func (_Controller *ControllerSession) GrantAccess(nftId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Controller.Contract.GrantAccess(&_Controller.TransactOpts, nftId, user)
}

// GrantAccess is a paid mutator transaction binding the contract method 0x65dd152c.
//
// Solidity: function grantAccess(uint256 nftId, address user) returns()
func (_Controller *ControllerTransactorSession) GrantAccess(nftId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Controller.Contract.GrantAccess(&_Controller.TransactOpts, nftId, user)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0xd9a5ec1b.
//
// Solidity: function revokeAccess(uint256 nftId, address user) returns()
func (_Controller *ControllerTransactor) RevokeAccess(opts *bind.TransactOpts, nftId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "revokeAccess", nftId, user)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0xd9a5ec1b.
//
// Solidity: function revokeAccess(uint256 nftId, address user) returns()
func (_Controller *ControllerSession) RevokeAccess(nftId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RevokeAccess(&_Controller.TransactOpts, nftId, user)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0xd9a5ec1b.
//
// Solidity: function revokeAccess(uint256 nftId, address user) returns()
func (_Controller *ControllerTransactorSession) RevokeAccess(nftId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Controller.Contract.RevokeAccess(&_Controller.TransactOpts, nftId, user)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(string docId) returns(uint256)
func (_Controller *ControllerTransactor) UploadData(opts *bind.TransactOpts, docId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "uploadData", docId)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(string docId) returns(uint256)
func (_Controller *ControllerSession) UploadData(docId string) (*types.Transaction, error) {
	return _Controller.Contract.UploadData(&_Controller.TransactOpts, docId)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(string docId) returns(uint256)
func (_Controller *ControllerTransactorSession) UploadData(docId string) (*types.Transaction, error) {
	return _Controller.Contract.UploadData(&_Controller.TransactOpts, docId)
}

// ControllerGrantedAccessIterator is returned from FilterGrantedAccess and is used to iterate over the raw logs and unpacked data for GrantedAccess events raised by the Controller contract.
type ControllerGrantedAccessIterator struct {
	Event *ControllerGrantedAccess // Event containing the contract specifics and raw log

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
func (it *ControllerGrantedAccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerGrantedAccess)
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
		it.Event = new(ControllerGrantedAccess)
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
func (it *ControllerGrantedAccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerGrantedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerGrantedAccess represents a GrantedAccess event raised by the Controller contract.
type ControllerGrantedAccess struct {
	NftId     *big.Int
	GrantedTo common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGrantedAccess is a free log retrieval operation binding the contract event 0x6a583816c9b620557c8e39c28c88cd423493e16c503abe9d25b84c75009ee65d.
//
// Solidity: event GrantedAccess(uint256 nftId, address grantedTo)
func (_Controller *ControllerFilterer) FilterGrantedAccess(opts *bind.FilterOpts) (*ControllerGrantedAccessIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "GrantedAccess")
	if err != nil {
		return nil, err
	}
	return &ControllerGrantedAccessIterator{contract: _Controller.contract, event: "GrantedAccess", logs: logs, sub: sub}, nil
}

// WatchGrantedAccess is a free log subscription operation binding the contract event 0x6a583816c9b620557c8e39c28c88cd423493e16c503abe9d25b84c75009ee65d.
//
// Solidity: event GrantedAccess(uint256 nftId, address grantedTo)
func (_Controller *ControllerFilterer) WatchGrantedAccess(opts *bind.WatchOpts, sink chan<- *ControllerGrantedAccess) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "GrantedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerGrantedAccess)
				if err := _Controller.contract.UnpackLog(event, "GrantedAccess", log); err != nil {
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

// ParseGrantedAccess is a log parse operation binding the contract event 0x6a583816c9b620557c8e39c28c88cd423493e16c503abe9d25b84c75009ee65d.
//
// Solidity: event GrantedAccess(uint256 nftId, address grantedTo)
func (_Controller *ControllerFilterer) ParseGrantedAccess(log types.Log) (*ControllerGrantedAccess, error) {
	event := new(ControllerGrantedAccess)
	if err := _Controller.contract.UnpackLog(event, "GrantedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerMintedGNFTIterator is returned from FilterMintedGNFT and is used to iterate over the raw logs and unpacked data for MintedGNFT events raised by the Controller contract.
type ControllerMintedGNFTIterator struct {
	Event *ControllerMintedGNFT // Event containing the contract specifics and raw log

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
func (it *ControllerMintedGNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerMintedGNFT)
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
		it.Event = new(ControllerMintedGNFT)
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
func (it *ControllerMintedGNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerMintedGNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerMintedGNFT represents a MintedGNFT event raised by the Controller contract.
type ControllerMintedGNFT struct {
	DocId     string
	SessionId *big.Int
	TokenId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintedGNFT is a free log retrieval operation binding the contract event 0x319e534953ee19933fcfb7956c223ddebce772a5baa82cd5f4656d285dc8971e.
//
// Solidity: event MintedGNFT(string docId, uint256 sessionId, uint256 tokenId)
func (_Controller *ControllerFilterer) FilterMintedGNFT(opts *bind.FilterOpts) (*ControllerMintedGNFTIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "MintedGNFT")
	if err != nil {
		return nil, err
	}
	return &ControllerMintedGNFTIterator{contract: _Controller.contract, event: "MintedGNFT", logs: logs, sub: sub}, nil
}

// WatchMintedGNFT is a free log subscription operation binding the contract event 0x319e534953ee19933fcfb7956c223ddebce772a5baa82cd5f4656d285dc8971e.
//
// Solidity: event MintedGNFT(string docId, uint256 sessionId, uint256 tokenId)
func (_Controller *ControllerFilterer) WatchMintedGNFT(opts *bind.WatchOpts, sink chan<- *ControllerMintedGNFT) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "MintedGNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerMintedGNFT)
				if err := _Controller.contract.UnpackLog(event, "MintedGNFT", log); err != nil {
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

// ParseMintedGNFT is a log parse operation binding the contract event 0x319e534953ee19933fcfb7956c223ddebce772a5baa82cd5f4656d285dc8971e.
//
// Solidity: event MintedGNFT(string docId, uint256 sessionId, uint256 tokenId)
func (_Controller *ControllerFilterer) ParseMintedGNFT(log types.Log) (*ControllerMintedGNFT, error) {
	event := new(ControllerMintedGNFT)
	if err := _Controller.contract.UnpackLog(event, "MintedGNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerRevokedAccessIterator is returned from FilterRevokedAccess and is used to iterate over the raw logs and unpacked data for RevokedAccess events raised by the Controller contract.
type ControllerRevokedAccessIterator struct {
	Event *ControllerRevokedAccess // Event containing the contract specifics and raw log

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
func (it *ControllerRevokedAccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRevokedAccess)
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
		it.Event = new(ControllerRevokedAccess)
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
func (it *ControllerRevokedAccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRevokedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRevokedAccess represents a RevokedAccess event raised by the Controller contract.
type ControllerRevokedAccess struct {
	NftId       *big.Int
	RevokedFrom common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevokedAccess is a free log retrieval operation binding the contract event 0xb9c3e2ec3ed558de01f11e512d264d5da714b3dc925c852002f6c9d25dd81bce.
//
// Solidity: event RevokedAccess(uint256 nftId, address revokedFrom)
func (_Controller *ControllerFilterer) FilterRevokedAccess(opts *bind.FilterOpts) (*ControllerRevokedAccessIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "RevokedAccess")
	if err != nil {
		return nil, err
	}
	return &ControllerRevokedAccessIterator{contract: _Controller.contract, event: "RevokedAccess", logs: logs, sub: sub}, nil
}

// WatchRevokedAccess is a free log subscription operation binding the contract event 0xb9c3e2ec3ed558de01f11e512d264d5da714b3dc925c852002f6c9d25dd81bce.
//
// Solidity: event RevokedAccess(uint256 nftId, address revokedFrom)
func (_Controller *ControllerFilterer) WatchRevokedAccess(opts *bind.WatchOpts, sink chan<- *ControllerRevokedAccess) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "RevokedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRevokedAccess)
				if err := _Controller.contract.UnpackLog(event, "RevokedAccess", log); err != nil {
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

// ParseRevokedAccess is a log parse operation binding the contract event 0xb9c3e2ec3ed558de01f11e512d264d5da714b3dc925c852002f6c9d25dd81bce.
//
// Solidity: event RevokedAccess(uint256 nftId, address revokedFrom)
func (_Controller *ControllerFilterer) ParseRevokedAccess(log types.Log) (*ControllerRevokedAccess, error) {
	event := new(ControllerRevokedAccess)
	if err := _Controller.contract.UnpackLog(event, "RevokedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerRewardedPCSPIterator is returned from FilterRewardedPCSP and is used to iterate over the raw logs and unpacked data for RewardedPCSP events raised by the Controller contract.
type ControllerRewardedPCSPIterator struct {
	Event *ControllerRewardedPCSP // Event containing the contract specifics and raw log

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
func (it *ControllerRewardedPCSPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRewardedPCSP)
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
		it.Event = new(ControllerRewardedPCSP)
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
func (it *ControllerRewardedPCSPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRewardedPCSPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRewardedPCSP represents a RewardedPCSP event raised by the Controller contract.
type ControllerRewardedPCSP struct {
	DocId     string
	SessionId *big.Int
	RiskScore *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardedPCSP is a free log retrieval operation binding the contract event 0x8881aa49bda3c31624847247d51afc28aa797fa81466dbd9df6d05d0b34d0d12.
//
// Solidity: event RewardedPCSP(string docId, uint256 sessionId, uint256 riskScore)
func (_Controller *ControllerFilterer) FilterRewardedPCSP(opts *bind.FilterOpts) (*ControllerRewardedPCSPIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "RewardedPCSP")
	if err != nil {
		return nil, err
	}
	return &ControllerRewardedPCSPIterator{contract: _Controller.contract, event: "RewardedPCSP", logs: logs, sub: sub}, nil
}

// WatchRewardedPCSP is a free log subscription operation binding the contract event 0x8881aa49bda3c31624847247d51afc28aa797fa81466dbd9df6d05d0b34d0d12.
//
// Solidity: event RewardedPCSP(string docId, uint256 sessionId, uint256 riskScore)
func (_Controller *ControllerFilterer) WatchRewardedPCSP(opts *bind.WatchOpts, sink chan<- *ControllerRewardedPCSP) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "RewardedPCSP")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRewardedPCSP)
				if err := _Controller.contract.UnpackLog(event, "RewardedPCSP", log); err != nil {
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

// ParseRewardedPCSP is a log parse operation binding the contract event 0x8881aa49bda3c31624847247d51afc28aa797fa81466dbd9df6d05d0b34d0d12.
//
// Solidity: event RewardedPCSP(string docId, uint256 sessionId, uint256 riskScore)
func (_Controller *ControllerFilterer) ParseRewardedPCSP(log types.Log) (*ControllerRewardedPCSP, error) {
	event := new(ControllerRewardedPCSP)
	if err := _Controller.contract.UnpackLog(event, "RewardedPCSP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerUploadDataIterator is returned from FilterUploadData and is used to iterate over the raw logs and unpacked data for UploadData events raised by the Controller contract.
type ControllerUploadDataIterator struct {
	Event *ControllerUploadData // Event containing the contract specifics and raw log

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
func (it *ControllerUploadDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerUploadData)
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
		it.Event = new(ControllerUploadData)
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
func (it *ControllerUploadDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerUploadDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerUploadData represents a UploadData event raised by the Controller contract.
type ControllerUploadData struct {
	DocId     string
	SessionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUploadData is a free log retrieval operation binding the contract event 0x698b35ede3baa51dbaa3b9a040c287690e40d0101d312c80eb364c7b17c458bc.
//
// Solidity: event UploadData(string docId, uint256 sessionId)
func (_Controller *ControllerFilterer) FilterUploadData(opts *bind.FilterOpts) (*ControllerUploadDataIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "UploadData")
	if err != nil {
		return nil, err
	}
	return &ControllerUploadDataIterator{contract: _Controller.contract, event: "UploadData", logs: logs, sub: sub}, nil
}

// WatchUploadData is a free log subscription operation binding the contract event 0x698b35ede3baa51dbaa3b9a040c287690e40d0101d312c80eb364c7b17c458bc.
//
// Solidity: event UploadData(string docId, uint256 sessionId)
func (_Controller *ControllerFilterer) WatchUploadData(opts *bind.WatchOpts, sink chan<- *ControllerUploadData) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "UploadData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerUploadData)
				if err := _Controller.contract.UnpackLog(event, "UploadData", log); err != nil {
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

// ParseUploadData is a log parse operation binding the contract event 0x698b35ede3baa51dbaa3b9a040c287690e40d0101d312c80eb364c7b17c458bc.
//
// Solidity: event UploadData(string docId, uint256 sessionId)
func (_Controller *ControllerFilterer) ParseUploadData(log types.Log) (*ControllerUploadData, error) {
	event := new(ControllerUploadData)
	if err := _Controller.contract.UnpackLog(event, "UploadData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
