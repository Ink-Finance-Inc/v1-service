// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AuditReporter

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

// AuditReporterMetaData contains all meta data concerning the AuditReporter contract.
var AuditReporterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"IncomeReport\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addrRegistry\",\"outputs\":[{\"internalType\":\"contractIAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"commitReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrReg\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AuditReporterABI is the input ABI used to generate the binding from.
// Deprecated: Use AuditReporterMetaData.ABI instead.
var AuditReporterABI = AuditReporterMetaData.ABI

// AuditReporter is an auto generated Go binding around an Ethereum contract.
type AuditReporter struct {
	AuditReporterCaller     // Read-only binding to the contract
	AuditReporterTransactor // Write-only binding to the contract
	AuditReporterFilterer   // Log filterer for contract events
}

// AuditReporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuditReporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuditReporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuditReporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuditReporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuditReporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuditReporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuditReporterSession struct {
	Contract     *AuditReporter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuditReporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuditReporterCallerSession struct {
	Contract *AuditReporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AuditReporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuditReporterTransactorSession struct {
	Contract     *AuditReporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AuditReporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuditReporterRaw struct {
	Contract *AuditReporter // Generic contract binding to access the raw methods on
}

// AuditReporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuditReporterCallerRaw struct {
	Contract *AuditReporterCaller // Generic read-only contract binding to access the raw methods on
}

// AuditReporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuditReporterTransactorRaw struct {
	Contract *AuditReporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuditReporter creates a new instance of AuditReporter, bound to a specific deployed contract.
func NewAuditReporter(address common.Address, backend bind.ContractBackend) (*AuditReporter, error) {
	contract, err := bindAuditReporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuditReporter{AuditReporterCaller: AuditReporterCaller{contract: contract}, AuditReporterTransactor: AuditReporterTransactor{contract: contract}, AuditReporterFilterer: AuditReporterFilterer{contract: contract}}, nil
}

// NewAuditReporterCaller creates a new read-only instance of AuditReporter, bound to a specific deployed contract.
func NewAuditReporterCaller(address common.Address, caller bind.ContractCaller) (*AuditReporterCaller, error) {
	contract, err := bindAuditReporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuditReporterCaller{contract: contract}, nil
}

// NewAuditReporterTransactor creates a new write-only instance of AuditReporter, bound to a specific deployed contract.
func NewAuditReporterTransactor(address common.Address, transactor bind.ContractTransactor) (*AuditReporterTransactor, error) {
	contract, err := bindAuditReporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuditReporterTransactor{contract: contract}, nil
}

// NewAuditReporterFilterer creates a new log filterer instance of AuditReporter, bound to a specific deployed contract.
func NewAuditReporterFilterer(address common.Address, filterer bind.ContractFilterer) (*AuditReporterFilterer, error) {
	contract, err := bindAuditReporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuditReporterFilterer{contract: contract}, nil
}

// bindAuditReporter binds a generic wrapper to an already deployed contract.
func bindAuditReporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuditReporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuditReporter *AuditReporterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuditReporter.Contract.AuditReporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuditReporter *AuditReporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuditReporter.Contract.AuditReporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuditReporter *AuditReporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuditReporter.Contract.AuditReporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuditReporter *AuditReporterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuditReporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuditReporter *AuditReporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuditReporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuditReporter *AuditReporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuditReporter.Contract.contract.Transact(opts, method, params...)
}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_AuditReporter *AuditReporterCaller) AddrRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuditReporter.contract.Call(opts, &out, "addrRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_AuditReporter *AuditReporterSession) AddrRegistry() (common.Address, error) {
	return _AuditReporter.Contract.AddrRegistry(&_AuditReporter.CallOpts)
}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_AuditReporter *AuditReporterCallerSession) AddrRegistry() (common.Address, error) {
	return _AuditReporter.Contract.AddrRegistry(&_AuditReporter.CallOpts)
}

// CommitReport is a paid mutator transaction binding the contract method 0xdbad9d2a.
//
// Solidity: function commitReport(address dao, bytes data) returns()
func (_AuditReporter *AuditReporterTransactor) CommitReport(opts *bind.TransactOpts, dao common.Address, data []byte) (*types.Transaction, error) {
	return _AuditReporter.contract.Transact(opts, "commitReport", dao, data)
}

// CommitReport is a paid mutator transaction binding the contract method 0xdbad9d2a.
//
// Solidity: function commitReport(address dao, bytes data) returns()
func (_AuditReporter *AuditReporterSession) CommitReport(dao common.Address, data []byte) (*types.Transaction, error) {
	return _AuditReporter.Contract.CommitReport(&_AuditReporter.TransactOpts, dao, data)
}

// CommitReport is a paid mutator transaction binding the contract method 0xdbad9d2a.
//
// Solidity: function commitReport(address dao, bytes data) returns()
func (_AuditReporter *AuditReporterTransactorSession) CommitReport(dao common.Address, data []byte) (*types.Transaction, error) {
	return _AuditReporter.Contract.CommitReport(&_AuditReporter.TransactOpts, dao, data)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_AuditReporter *AuditReporterTransactor) Init(opts *bind.TransactOpts, addrReg common.Address) (*types.Transaction, error) {
	return _AuditReporter.contract.Transact(opts, "init", addrReg)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_AuditReporter *AuditReporterSession) Init(addrReg common.Address) (*types.Transaction, error) {
	return _AuditReporter.Contract.Init(&_AuditReporter.TransactOpts, addrReg)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_AuditReporter *AuditReporterTransactorSession) Init(addrReg common.Address) (*types.Transaction, error) {
	return _AuditReporter.Contract.Init(&_AuditReporter.TransactOpts, addrReg)
}

// AuditReporterIncomeReportIterator is returned from FilterIncomeReport and is used to iterate over the raw logs and unpacked data for IncomeReport events raised by the AuditReporter contract.
type AuditReporterIncomeReportIterator struct {
	Event *AuditReporterIncomeReport // Event containing the contract specifics and raw log

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
func (it *AuditReporterIncomeReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuditReporterIncomeReport)
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
		it.Event = new(AuditReporterIncomeReport)
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
func (it *AuditReporterIncomeReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuditReporterIncomeReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuditReporterIncomeReport represents a IncomeReport event raised by the AuditReporter contract.
type AuditReporterIncomeReport struct {
	Dao      common.Address
	Operator common.Address
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIncomeReport is a free log retrieval operation binding the contract event 0x4a4cbeb1671a7df1da02b962b3d22164ea9ab2bb3c60ba63db735839e0a52267.
//
// Solidity: event IncomeReport(address indexed dao, address indexed operator, bytes data)
func (_AuditReporter *AuditReporterFilterer) FilterIncomeReport(opts *bind.FilterOpts, dao []common.Address, operator []common.Address) (*AuditReporterIncomeReportIterator, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AuditReporter.contract.FilterLogs(opts, "IncomeReport", daoRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AuditReporterIncomeReportIterator{contract: _AuditReporter.contract, event: "IncomeReport", logs: logs, sub: sub}, nil
}

// WatchIncomeReport is a free log subscription operation binding the contract event 0x4a4cbeb1671a7df1da02b962b3d22164ea9ab2bb3c60ba63db735839e0a52267.
//
// Solidity: event IncomeReport(address indexed dao, address indexed operator, bytes data)
func (_AuditReporter *AuditReporterFilterer) WatchIncomeReport(opts *bind.WatchOpts, sink chan<- *AuditReporterIncomeReport, dao []common.Address, operator []common.Address) (event.Subscription, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AuditReporter.contract.WatchLogs(opts, "IncomeReport", daoRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuditReporterIncomeReport)
				if err := _AuditReporter.contract.UnpackLog(event, "IncomeReport", log); err != nil {
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

// ParseIncomeReport is a log parse operation binding the contract event 0x4a4cbeb1671a7df1da02b962b3d22164ea9ab2bb3c60ba63db735839e0a52267.
//
// Solidity: event IncomeReport(address indexed dao, address indexed operator, bytes data)
func (_AuditReporter *AuditReporterFilterer) ParseIncomeReport(log types.Log) (*AuditReporterIncomeReport, error) {
	event := new(AuditReporterIncomeReport)
	if err := _AuditReporter.contract.UnpackLog(event, "IncomeReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
