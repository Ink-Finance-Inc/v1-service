// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FactoryManager

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

// IFactoryManagerFactoryInfo is an auto generated low-level Go binding around an user-defined struct.
type IFactoryManagerFactoryInfo struct {
	Disable     *big.Int
	FactoryType [32]byte
	FactoryID   [32]byte
	BeaconAddr  common.Address
	CodeAddr    common.Address
}

// FactoryManagerMetaData contains all meta data concerning the FactoryManager contract.
var FactoryManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sandBoxAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"codeAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"factoryType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"factoryID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beaconAddr\",\"type\":\"address\"}],\"name\":\"EAddNewFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deployAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"factoryType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"factoryID\",\"type\":\"bytes32\"}],\"name\":\"ECreateContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"factoryID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disableFlag\",\"type\":\"bool\"}],\"name\":\"EFactoryDisable\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"factoryType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"factoryID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"codeAddr\",\"type\":\"address\"}],\"name\":\"addNewFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addrRegistry\",\"outputs\":[{\"internalType\":\"contractIAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"factoryID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reqID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"onlyOne\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"deployContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"factoryID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"disable\",\"type\":\"bool\"}],\"name\":\"disableFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generatesNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"factoryID\",\"type\":\"bytes32\"}],\"name\":\"getFactoryInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"disable\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"factoryType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"factoryID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beaconAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"codeAddr\",\"type\":\"address\"}],\"internalType\":\"structIFactoryManager.FactoryInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrReg\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inkProxy\",\"outputs\":[{\"internalType\":\"contractInkProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sandBox\",\"outputs\":[{\"internalType\":\"contractSandBox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tagFactoryIDs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FactoryManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryManagerMetaData.ABI instead.
var FactoryManagerABI = FactoryManagerMetaData.ABI

// FactoryManager is an auto generated Go binding around an Ethereum contract.
type FactoryManager struct {
	FactoryManagerCaller     // Read-only binding to the contract
	FactoryManagerTransactor // Write-only binding to the contract
	FactoryManagerFilterer   // Log filterer for contract events
}

// FactoryManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactoryManagerSession struct {
	Contract     *FactoryManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryManagerCallerSession struct {
	Contract *FactoryManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FactoryManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryManagerTransactorSession struct {
	Contract     *FactoryManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FactoryManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryManagerRaw struct {
	Contract *FactoryManager // Generic contract binding to access the raw methods on
}

// FactoryManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryManagerCallerRaw struct {
	Contract *FactoryManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryManagerTransactorRaw struct {
	Contract *FactoryManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactoryManager creates a new instance of FactoryManager, bound to a specific deployed contract.
func NewFactoryManager(address common.Address, backend bind.ContractBackend) (*FactoryManager, error) {
	contract, err := bindFactoryManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FactoryManager{FactoryManagerCaller: FactoryManagerCaller{contract: contract}, FactoryManagerTransactor: FactoryManagerTransactor{contract: contract}, FactoryManagerFilterer: FactoryManagerFilterer{contract: contract}}, nil
}

// NewFactoryManagerCaller creates a new read-only instance of FactoryManager, bound to a specific deployed contract.
func NewFactoryManagerCaller(address common.Address, caller bind.ContractCaller) (*FactoryManagerCaller, error) {
	contract, err := bindFactoryManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryManagerCaller{contract: contract}, nil
}

// NewFactoryManagerTransactor creates a new write-only instance of FactoryManager, bound to a specific deployed contract.
func NewFactoryManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryManagerTransactor, error) {
	contract, err := bindFactoryManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryManagerTransactor{contract: contract}, nil
}

// NewFactoryManagerFilterer creates a new log filterer instance of FactoryManager, bound to a specific deployed contract.
func NewFactoryManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryManagerFilterer, error) {
	contract, err := bindFactoryManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryManagerFilterer{contract: contract}, nil
}

// bindFactoryManager binds a generic wrapper to an already deployed contract.
func bindFactoryManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FactoryManager *FactoryManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactoryManager.Contract.FactoryManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FactoryManager *FactoryManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryManager.Contract.FactoryManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FactoryManager *FactoryManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactoryManager.Contract.FactoryManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FactoryManager *FactoryManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactoryManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FactoryManager *FactoryManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FactoryManager *FactoryManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactoryManager.Contract.contract.Transact(opts, method, params...)
}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_FactoryManager *FactoryManagerCaller) AddrRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryManager.contract.Call(opts, &out, "addrRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_FactoryManager *FactoryManagerSession) AddrRegistry() (common.Address, error) {
	return _FactoryManager.Contract.AddrRegistry(&_FactoryManager.CallOpts)
}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_FactoryManager *FactoryManagerCallerSession) AddrRegistry() (common.Address, error) {
	return _FactoryManager.Contract.AddrRegistry(&_FactoryManager.CallOpts)
}

// GeneratesNum is a free data retrieval call binding the contract method 0x80548569.
//
// Solidity: function generatesNum() view returns(uint256)
func (_FactoryManager *FactoryManagerCaller) GeneratesNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FactoryManager.contract.Call(opts, &out, "generatesNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneratesNum is a free data retrieval call binding the contract method 0x80548569.
//
// Solidity: function generatesNum() view returns(uint256)
func (_FactoryManager *FactoryManagerSession) GeneratesNum() (*big.Int, error) {
	return _FactoryManager.Contract.GeneratesNum(&_FactoryManager.CallOpts)
}

// GeneratesNum is a free data retrieval call binding the contract method 0x80548569.
//
// Solidity: function generatesNum() view returns(uint256)
func (_FactoryManager *FactoryManagerCallerSession) GeneratesNum() (*big.Int, error) {
	return _FactoryManager.Contract.GeneratesNum(&_FactoryManager.CallOpts)
}

// GetFactoryInfo is a free data retrieval call binding the contract method 0x9ad1d589.
//
// Solidity: function getFactoryInfo(bytes32 factoryID) view returns((uint256,bytes32,bytes32,address,address) info)
func (_FactoryManager *FactoryManagerCaller) GetFactoryInfo(opts *bind.CallOpts, factoryID [32]byte) (IFactoryManagerFactoryInfo, error) {
	var out []interface{}
	err := _FactoryManager.contract.Call(opts, &out, "getFactoryInfo", factoryID)

	if err != nil {
		return *new(IFactoryManagerFactoryInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IFactoryManagerFactoryInfo)).(*IFactoryManagerFactoryInfo)

	return out0, err

}

// GetFactoryInfo is a free data retrieval call binding the contract method 0x9ad1d589.
//
// Solidity: function getFactoryInfo(bytes32 factoryID) view returns((uint256,bytes32,bytes32,address,address) info)
func (_FactoryManager *FactoryManagerSession) GetFactoryInfo(factoryID [32]byte) (IFactoryManagerFactoryInfo, error) {
	return _FactoryManager.Contract.GetFactoryInfo(&_FactoryManager.CallOpts, factoryID)
}

// GetFactoryInfo is a free data retrieval call binding the contract method 0x9ad1d589.
//
// Solidity: function getFactoryInfo(bytes32 factoryID) view returns((uint256,bytes32,bytes32,address,address) info)
func (_FactoryManager *FactoryManagerCallerSession) GetFactoryInfo(factoryID [32]byte) (IFactoryManagerFactoryInfo, error) {
	return _FactoryManager.Contract.GetFactoryInfo(&_FactoryManager.CallOpts, factoryID)
}

// InkProxy is a free data retrieval call binding the contract method 0xcc79a84b.
//
// Solidity: function inkProxy() view returns(address)
func (_FactoryManager *FactoryManagerCaller) InkProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryManager.contract.Call(opts, &out, "inkProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InkProxy is a free data retrieval call binding the contract method 0xcc79a84b.
//
// Solidity: function inkProxy() view returns(address)
func (_FactoryManager *FactoryManagerSession) InkProxy() (common.Address, error) {
	return _FactoryManager.Contract.InkProxy(&_FactoryManager.CallOpts)
}

// InkProxy is a free data retrieval call binding the contract method 0xcc79a84b.
//
// Solidity: function inkProxy() view returns(address)
func (_FactoryManager *FactoryManagerCallerSession) InkProxy() (common.Address, error) {
	return _FactoryManager.Contract.InkProxy(&_FactoryManager.CallOpts)
}

// SandBox is a free data retrieval call binding the contract method 0xf4cc4127.
//
// Solidity: function sandBox() view returns(address)
func (_FactoryManager *FactoryManagerCaller) SandBox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryManager.contract.Call(opts, &out, "sandBox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SandBox is a free data retrieval call binding the contract method 0xf4cc4127.
//
// Solidity: function sandBox() view returns(address)
func (_FactoryManager *FactoryManagerSession) SandBox() (common.Address, error) {
	return _FactoryManager.Contract.SandBox(&_FactoryManager.CallOpts)
}

// SandBox is a free data retrieval call binding the contract method 0xf4cc4127.
//
// Solidity: function sandBox() view returns(address)
func (_FactoryManager *FactoryManagerCallerSession) SandBox() (common.Address, error) {
	return _FactoryManager.Contract.SandBox(&_FactoryManager.CallOpts)
}

// TagFactoryIDs is a free data retrieval call binding the contract method 0x36adaaf2.
//
// Solidity: function tagFactoryIDs(bytes32 , uint256 ) view returns(bytes32)
func (_FactoryManager *FactoryManagerCaller) TagFactoryIDs(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _FactoryManager.contract.Call(opts, &out, "tagFactoryIDs", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TagFactoryIDs is a free data retrieval call binding the contract method 0x36adaaf2.
//
// Solidity: function tagFactoryIDs(bytes32 , uint256 ) view returns(bytes32)
func (_FactoryManager *FactoryManagerSession) TagFactoryIDs(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _FactoryManager.Contract.TagFactoryIDs(&_FactoryManager.CallOpts, arg0, arg1)
}

// TagFactoryIDs is a free data retrieval call binding the contract method 0x36adaaf2.
//
// Solidity: function tagFactoryIDs(bytes32 , uint256 ) view returns(bytes32)
func (_FactoryManager *FactoryManagerCallerSession) TagFactoryIDs(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _FactoryManager.Contract.TagFactoryIDs(&_FactoryManager.CallOpts, arg0, arg1)
}

// AddNewFactory is a paid mutator transaction binding the contract method 0x7bd3b183.
//
// Solidity: function addNewFactory(bytes32 factoryType, bytes32 factoryID, address codeAddr) returns()
func (_FactoryManager *FactoryManagerTransactor) AddNewFactory(opts *bind.TransactOpts, factoryType [32]byte, factoryID [32]byte, codeAddr common.Address) (*types.Transaction, error) {
	return _FactoryManager.contract.Transact(opts, "addNewFactory", factoryType, factoryID, codeAddr)
}

// AddNewFactory is a paid mutator transaction binding the contract method 0x7bd3b183.
//
// Solidity: function addNewFactory(bytes32 factoryType, bytes32 factoryID, address codeAddr) returns()
func (_FactoryManager *FactoryManagerSession) AddNewFactory(factoryType [32]byte, factoryID [32]byte, codeAddr common.Address) (*types.Transaction, error) {
	return _FactoryManager.Contract.AddNewFactory(&_FactoryManager.TransactOpts, factoryType, factoryID, codeAddr)
}

// AddNewFactory is a paid mutator transaction binding the contract method 0x7bd3b183.
//
// Solidity: function addNewFactory(bytes32 factoryType, bytes32 factoryID, address codeAddr) returns()
func (_FactoryManager *FactoryManagerTransactorSession) AddNewFactory(factoryType [32]byte, factoryID [32]byte, codeAddr common.Address) (*types.Transaction, error) {
	return _FactoryManager.Contract.AddNewFactory(&_FactoryManager.TransactOpts, factoryType, factoryID, codeAddr)
}

// DeployContract is a paid mutator transaction binding the contract method 0xa87d7d58.
//
// Solidity: function deployContract(address admin, bytes32 factoryID, bytes32 reqID, bool onlyOne, bytes initData) returns()
func (_FactoryManager *FactoryManagerTransactor) DeployContract(opts *bind.TransactOpts, admin common.Address, factoryID [32]byte, reqID [32]byte, onlyOne bool, initData []byte) (*types.Transaction, error) {
	return _FactoryManager.contract.Transact(opts, "deployContract", admin, factoryID, reqID, onlyOne, initData)
}

// DeployContract is a paid mutator transaction binding the contract method 0xa87d7d58.
//
// Solidity: function deployContract(address admin, bytes32 factoryID, bytes32 reqID, bool onlyOne, bytes initData) returns()
func (_FactoryManager *FactoryManagerSession) DeployContract(admin common.Address, factoryID [32]byte, reqID [32]byte, onlyOne bool, initData []byte) (*types.Transaction, error) {
	return _FactoryManager.Contract.DeployContract(&_FactoryManager.TransactOpts, admin, factoryID, reqID, onlyOne, initData)
}

// DeployContract is a paid mutator transaction binding the contract method 0xa87d7d58.
//
// Solidity: function deployContract(address admin, bytes32 factoryID, bytes32 reqID, bool onlyOne, bytes initData) returns()
func (_FactoryManager *FactoryManagerTransactorSession) DeployContract(admin common.Address, factoryID [32]byte, reqID [32]byte, onlyOne bool, initData []byte) (*types.Transaction, error) {
	return _FactoryManager.Contract.DeployContract(&_FactoryManager.TransactOpts, admin, factoryID, reqID, onlyOne, initData)
}

// DisableFactory is a paid mutator transaction binding the contract method 0x7e95ac01.
//
// Solidity: function disableFactory(bytes32 factoryID, bool disable) returns()
func (_FactoryManager *FactoryManagerTransactor) DisableFactory(opts *bind.TransactOpts, factoryID [32]byte, disable bool) (*types.Transaction, error) {
	return _FactoryManager.contract.Transact(opts, "disableFactory", factoryID, disable)
}

// DisableFactory is a paid mutator transaction binding the contract method 0x7e95ac01.
//
// Solidity: function disableFactory(bytes32 factoryID, bool disable) returns()
func (_FactoryManager *FactoryManagerSession) DisableFactory(factoryID [32]byte, disable bool) (*types.Transaction, error) {
	return _FactoryManager.Contract.DisableFactory(&_FactoryManager.TransactOpts, factoryID, disable)
}

// DisableFactory is a paid mutator transaction binding the contract method 0x7e95ac01.
//
// Solidity: function disableFactory(bytes32 factoryID, bool disable) returns()
func (_FactoryManager *FactoryManagerTransactorSession) DisableFactory(factoryID [32]byte, disable bool) (*types.Transaction, error) {
	return _FactoryManager.Contract.DisableFactory(&_FactoryManager.TransactOpts, factoryID, disable)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_FactoryManager *FactoryManagerTransactor) Init(opts *bind.TransactOpts, addrReg common.Address) (*types.Transaction, error) {
	return _FactoryManager.contract.Transact(opts, "init", addrReg)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_FactoryManager *FactoryManagerSession) Init(addrReg common.Address) (*types.Transaction, error) {
	return _FactoryManager.Contract.Init(&_FactoryManager.TransactOpts, addrReg)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_FactoryManager *FactoryManagerTransactorSession) Init(addrReg common.Address) (*types.Transaction, error) {
	return _FactoryManager.Contract.Init(&_FactoryManager.TransactOpts, addrReg)
}

// FactoryManagerEAddNewFactoryIterator is returned from FilterEAddNewFactory and is used to iterate over the raw logs and unpacked data for EAddNewFactory events raised by the FactoryManager contract.
type FactoryManagerEAddNewFactoryIterator struct {
	Event *FactoryManagerEAddNewFactory // Event containing the contract specifics and raw log

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
func (it *FactoryManagerEAddNewFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryManagerEAddNewFactory)
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
		it.Event = new(FactoryManagerEAddNewFactory)
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
func (it *FactoryManagerEAddNewFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryManagerEAddNewFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryManagerEAddNewFactory represents a EAddNewFactory event raised by the FactoryManager contract.
type FactoryManagerEAddNewFactory struct {
	CodeAddr    common.Address
	FactoryType [32]byte
	FactoryID   [32]byte
	BeaconAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEAddNewFactory is a free log retrieval operation binding the contract event 0xfaf2fcfca30e4eb145c3ee6b3be4ccdc9e56c68bad4a1cd8da85dbb364c9b8ad.
//
// Solidity: event EAddNewFactory(address indexed codeAddr, bytes32 indexed factoryType, bytes32 indexed factoryID, address beaconAddr)
func (_FactoryManager *FactoryManagerFilterer) FilterEAddNewFactory(opts *bind.FilterOpts, codeAddr []common.Address, factoryType [][32]byte, factoryID [][32]byte) (*FactoryManagerEAddNewFactoryIterator, error) {

	var codeAddrRule []interface{}
	for _, codeAddrItem := range codeAddr {
		codeAddrRule = append(codeAddrRule, codeAddrItem)
	}
	var factoryTypeRule []interface{}
	for _, factoryTypeItem := range factoryType {
		factoryTypeRule = append(factoryTypeRule, factoryTypeItem)
	}
	var factoryIDRule []interface{}
	for _, factoryIDItem := range factoryID {
		factoryIDRule = append(factoryIDRule, factoryIDItem)
	}

	logs, sub, err := _FactoryManager.contract.FilterLogs(opts, "EAddNewFactory", codeAddrRule, factoryTypeRule, factoryIDRule)
	if err != nil {
		return nil, err
	}
	return &FactoryManagerEAddNewFactoryIterator{contract: _FactoryManager.contract, event: "EAddNewFactory", logs: logs, sub: sub}, nil
}

// WatchEAddNewFactory is a free log subscription operation binding the contract event 0xfaf2fcfca30e4eb145c3ee6b3be4ccdc9e56c68bad4a1cd8da85dbb364c9b8ad.
//
// Solidity: event EAddNewFactory(address indexed codeAddr, bytes32 indexed factoryType, bytes32 indexed factoryID, address beaconAddr)
func (_FactoryManager *FactoryManagerFilterer) WatchEAddNewFactory(opts *bind.WatchOpts, sink chan<- *FactoryManagerEAddNewFactory, codeAddr []common.Address, factoryType [][32]byte, factoryID [][32]byte) (event.Subscription, error) {

	var codeAddrRule []interface{}
	for _, codeAddrItem := range codeAddr {
		codeAddrRule = append(codeAddrRule, codeAddrItem)
	}
	var factoryTypeRule []interface{}
	for _, factoryTypeItem := range factoryType {
		factoryTypeRule = append(factoryTypeRule, factoryTypeItem)
	}
	var factoryIDRule []interface{}
	for _, factoryIDItem := range factoryID {
		factoryIDRule = append(factoryIDRule, factoryIDItem)
	}

	logs, sub, err := _FactoryManager.contract.WatchLogs(opts, "EAddNewFactory", codeAddrRule, factoryTypeRule, factoryIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryManagerEAddNewFactory)
				if err := _FactoryManager.contract.UnpackLog(event, "EAddNewFactory", log); err != nil {
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

// ParseEAddNewFactory is a log parse operation binding the contract event 0xfaf2fcfca30e4eb145c3ee6b3be4ccdc9e56c68bad4a1cd8da85dbb364c9b8ad.
//
// Solidity: event EAddNewFactory(address indexed codeAddr, bytes32 indexed factoryType, bytes32 indexed factoryID, address beaconAddr)
func (_FactoryManager *FactoryManagerFilterer) ParseEAddNewFactory(log types.Log) (*FactoryManagerEAddNewFactory, error) {
	event := new(FactoryManagerEAddNewFactory)
	if err := _FactoryManager.contract.UnpackLog(event, "EAddNewFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryManagerECreateContractIterator is returned from FilterECreateContract and is used to iterate over the raw logs and unpacked data for ECreateContract events raised by the FactoryManager contract.
type FactoryManagerECreateContractIterator struct {
	Event *FactoryManagerECreateContract // Event containing the contract specifics and raw log

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
func (it *FactoryManagerECreateContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryManagerECreateContract)
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
		it.Event = new(FactoryManagerECreateContract)
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
func (it *FactoryManagerECreateContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryManagerECreateContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryManagerECreateContract represents a ECreateContract event raised by the FactoryManager contract.
type FactoryManagerECreateContract struct {
	Admin       common.Address
	ReqID       [32]byte
	DeployAddr  common.Address
	FactoryType [32]byte
	FactoryID   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterECreateContract is a free log retrieval operation binding the contract event 0x6964bb94ec599265662c19aeef03c5816b5799164f587e6f3d29a376ae8b8925.
//
// Solidity: event ECreateContract(address indexed admin, bytes32 indexed reqID, address indexed deployAddr, bytes32 factoryType, bytes32 factoryID)
func (_FactoryManager *FactoryManagerFilterer) FilterECreateContract(opts *bind.FilterOpts, admin []common.Address, reqID [][32]byte, deployAddr []common.Address) (*FactoryManagerECreateContractIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var reqIDRule []interface{}
	for _, reqIDItem := range reqID {
		reqIDRule = append(reqIDRule, reqIDItem)
	}
	var deployAddrRule []interface{}
	for _, deployAddrItem := range deployAddr {
		deployAddrRule = append(deployAddrRule, deployAddrItem)
	}

	logs, sub, err := _FactoryManager.contract.FilterLogs(opts, "ECreateContract", adminRule, reqIDRule, deployAddrRule)
	if err != nil {
		return nil, err
	}
	return &FactoryManagerECreateContractIterator{contract: _FactoryManager.contract, event: "ECreateContract", logs: logs, sub: sub}, nil
}

// WatchECreateContract is a free log subscription operation binding the contract event 0x6964bb94ec599265662c19aeef03c5816b5799164f587e6f3d29a376ae8b8925.
//
// Solidity: event ECreateContract(address indexed admin, bytes32 indexed reqID, address indexed deployAddr, bytes32 factoryType, bytes32 factoryID)
func (_FactoryManager *FactoryManagerFilterer) WatchECreateContract(opts *bind.WatchOpts, sink chan<- *FactoryManagerECreateContract, admin []common.Address, reqID [][32]byte, deployAddr []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var reqIDRule []interface{}
	for _, reqIDItem := range reqID {
		reqIDRule = append(reqIDRule, reqIDItem)
	}
	var deployAddrRule []interface{}
	for _, deployAddrItem := range deployAddr {
		deployAddrRule = append(deployAddrRule, deployAddrItem)
	}

	logs, sub, err := _FactoryManager.contract.WatchLogs(opts, "ECreateContract", adminRule, reqIDRule, deployAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryManagerECreateContract)
				if err := _FactoryManager.contract.UnpackLog(event, "ECreateContract", log); err != nil {
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

// ParseECreateContract is a log parse operation binding the contract event 0x6964bb94ec599265662c19aeef03c5816b5799164f587e6f3d29a376ae8b8925.
//
// Solidity: event ECreateContract(address indexed admin, bytes32 indexed reqID, address indexed deployAddr, bytes32 factoryType, bytes32 factoryID)
func (_FactoryManager *FactoryManagerFilterer) ParseECreateContract(log types.Log) (*FactoryManagerECreateContract, error) {
	event := new(FactoryManagerECreateContract)
	if err := _FactoryManager.contract.UnpackLog(event, "ECreateContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryManagerEFactoryDisableIterator is returned from FilterEFactoryDisable and is used to iterate over the raw logs and unpacked data for EFactoryDisable events raised by the FactoryManager contract.
type FactoryManagerEFactoryDisableIterator struct {
	Event *FactoryManagerEFactoryDisable // Event containing the contract specifics and raw log

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
func (it *FactoryManagerEFactoryDisableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryManagerEFactoryDisable)
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
		it.Event = new(FactoryManagerEFactoryDisable)
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
func (it *FactoryManagerEFactoryDisableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryManagerEFactoryDisableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryManagerEFactoryDisable represents a EFactoryDisable event raised by the FactoryManager contract.
type FactoryManagerEFactoryDisable struct {
	FactoryID   [32]byte
	DisableFlag bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEFactoryDisable is a free log retrieval operation binding the contract event 0x46e800e93c8d9180fa0fe6ae8519a6fe6490a30a7af9db8e627808d85c518ef3.
//
// Solidity: event EFactoryDisable(bytes32 indexed factoryID, bool disableFlag)
func (_FactoryManager *FactoryManagerFilterer) FilterEFactoryDisable(opts *bind.FilterOpts, factoryID [][32]byte) (*FactoryManagerEFactoryDisableIterator, error) {

	var factoryIDRule []interface{}
	for _, factoryIDItem := range factoryID {
		factoryIDRule = append(factoryIDRule, factoryIDItem)
	}

	logs, sub, err := _FactoryManager.contract.FilterLogs(opts, "EFactoryDisable", factoryIDRule)
	if err != nil {
		return nil, err
	}
	return &FactoryManagerEFactoryDisableIterator{contract: _FactoryManager.contract, event: "EFactoryDisable", logs: logs, sub: sub}, nil
}

// WatchEFactoryDisable is a free log subscription operation binding the contract event 0x46e800e93c8d9180fa0fe6ae8519a6fe6490a30a7af9db8e627808d85c518ef3.
//
// Solidity: event EFactoryDisable(bytes32 indexed factoryID, bool disableFlag)
func (_FactoryManager *FactoryManagerFilterer) WatchEFactoryDisable(opts *bind.WatchOpts, sink chan<- *FactoryManagerEFactoryDisable, factoryID [][32]byte) (event.Subscription, error) {

	var factoryIDRule []interface{}
	for _, factoryIDItem := range factoryID {
		factoryIDRule = append(factoryIDRule, factoryIDItem)
	}

	logs, sub, err := _FactoryManager.contract.WatchLogs(opts, "EFactoryDisable", factoryIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryManagerEFactoryDisable)
				if err := _FactoryManager.contract.UnpackLog(event, "EFactoryDisable", log); err != nil {
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

// ParseEFactoryDisable is a log parse operation binding the contract event 0x46e800e93c8d9180fa0fe6ae8519a6fe6490a30a7af9db8e627808d85c518ef3.
//
// Solidity: event EFactoryDisable(bytes32 indexed factoryID, bool disableFlag)
func (_FactoryManager *FactoryManagerFilterer) ParseEFactoryDisable(log types.Log) (*FactoryManagerEFactoryDisable, error) {
	event := new(FactoryManagerEFactoryDisable)
	if err := _FactoryManager.contract.UnpackLog(event, "EFactoryDisable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
