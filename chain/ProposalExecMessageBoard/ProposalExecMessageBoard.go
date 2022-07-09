// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ProposalExecMessageBoard

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

// ProposalExecMessageBoardMetaData contains all meta data concerning the ProposalExecMessageBoard contract.
var ProposalExecMessageBoardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"execProposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nowProposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"EMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addrRegistry\",\"outputs\":[{\"internalType\":\"contractIAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"execProposalID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"execMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrReg\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProposalExecMessageBoardABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalExecMessageBoardMetaData.ABI instead.
var ProposalExecMessageBoardABI = ProposalExecMessageBoardMetaData.ABI

// ProposalExecMessageBoard is an auto generated Go binding around an Ethereum contract.
type ProposalExecMessageBoard struct {
	ProposalExecMessageBoardCaller     // Read-only binding to the contract
	ProposalExecMessageBoardTransactor // Write-only binding to the contract
	ProposalExecMessageBoardFilterer   // Log filterer for contract events
}

// ProposalExecMessageBoardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposalExecMessageBoardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalExecMessageBoardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposalExecMessageBoardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalExecMessageBoardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposalExecMessageBoardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalExecMessageBoardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposalExecMessageBoardSession struct {
	Contract     *ProposalExecMessageBoard // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ProposalExecMessageBoardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposalExecMessageBoardCallerSession struct {
	Contract *ProposalExecMessageBoardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ProposalExecMessageBoardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposalExecMessageBoardTransactorSession struct {
	Contract     *ProposalExecMessageBoardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ProposalExecMessageBoardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposalExecMessageBoardRaw struct {
	Contract *ProposalExecMessageBoard // Generic contract binding to access the raw methods on
}

// ProposalExecMessageBoardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposalExecMessageBoardCallerRaw struct {
	Contract *ProposalExecMessageBoardCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalExecMessageBoardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposalExecMessageBoardTransactorRaw struct {
	Contract *ProposalExecMessageBoardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposalExecMessageBoard creates a new instance of ProposalExecMessageBoard, bound to a specific deployed contract.
func NewProposalExecMessageBoard(address common.Address, backend bind.ContractBackend) (*ProposalExecMessageBoard, error) {
	contract, err := bindProposalExecMessageBoard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProposalExecMessageBoard{ProposalExecMessageBoardCaller: ProposalExecMessageBoardCaller{contract: contract}, ProposalExecMessageBoardTransactor: ProposalExecMessageBoardTransactor{contract: contract}, ProposalExecMessageBoardFilterer: ProposalExecMessageBoardFilterer{contract: contract}}, nil
}

// NewProposalExecMessageBoardCaller creates a new read-only instance of ProposalExecMessageBoard, bound to a specific deployed contract.
func NewProposalExecMessageBoardCaller(address common.Address, caller bind.ContractCaller) (*ProposalExecMessageBoardCaller, error) {
	contract, err := bindProposalExecMessageBoard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalExecMessageBoardCaller{contract: contract}, nil
}

// NewProposalExecMessageBoardTransactor creates a new write-only instance of ProposalExecMessageBoard, bound to a specific deployed contract.
func NewProposalExecMessageBoardTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalExecMessageBoardTransactor, error) {
	contract, err := bindProposalExecMessageBoard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalExecMessageBoardTransactor{contract: contract}, nil
}

// NewProposalExecMessageBoardFilterer creates a new log filterer instance of ProposalExecMessageBoard, bound to a specific deployed contract.
func NewProposalExecMessageBoardFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalExecMessageBoardFilterer, error) {
	contract, err := bindProposalExecMessageBoard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalExecMessageBoardFilterer{contract: contract}, nil
}

// bindProposalExecMessageBoard binds a generic wrapper to an already deployed contract.
func bindProposalExecMessageBoard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProposalExecMessageBoardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalExecMessageBoard *ProposalExecMessageBoardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalExecMessageBoard.Contract.ProposalExecMessageBoardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalExecMessageBoard *ProposalExecMessageBoardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalExecMessageBoard.Contract.ProposalExecMessageBoardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalExecMessageBoard *ProposalExecMessageBoardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalExecMessageBoard.Contract.ProposalExecMessageBoardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalExecMessageBoard *ProposalExecMessageBoardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalExecMessageBoard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalExecMessageBoard *ProposalExecMessageBoardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalExecMessageBoard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalExecMessageBoard *ProposalExecMessageBoardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalExecMessageBoard.Contract.contract.Transact(opts, method, params...)
}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_ProposalExecMessageBoard *ProposalExecMessageBoardCaller) AddrRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposalExecMessageBoard.contract.Call(opts, &out, "addrRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_ProposalExecMessageBoard *ProposalExecMessageBoardSession) AddrRegistry() (common.Address, error) {
	return _ProposalExecMessageBoard.Contract.AddrRegistry(&_ProposalExecMessageBoard.CallOpts)
}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_ProposalExecMessageBoard *ProposalExecMessageBoardCallerSession) AddrRegistry() (common.Address, error) {
	return _ProposalExecMessageBoard.Contract.AddrRegistry(&_ProposalExecMessageBoard.CallOpts)
}

// ExecMessage is a paid mutator transaction binding the contract method 0x5b47f056.
//
// Solidity: function execMessage(bytes32 execProposalID, bytes message) returns()
func (_ProposalExecMessageBoard *ProposalExecMessageBoardTransactor) ExecMessage(opts *bind.TransactOpts, execProposalID [32]byte, message []byte) (*types.Transaction, error) {
	return _ProposalExecMessageBoard.contract.Transact(opts, "execMessage", execProposalID, message)
}

// ExecMessage is a paid mutator transaction binding the contract method 0x5b47f056.
//
// Solidity: function execMessage(bytes32 execProposalID, bytes message) returns()
func (_ProposalExecMessageBoard *ProposalExecMessageBoardSession) ExecMessage(execProposalID [32]byte, message []byte) (*types.Transaction, error) {
	return _ProposalExecMessageBoard.Contract.ExecMessage(&_ProposalExecMessageBoard.TransactOpts, execProposalID, message)
}

// ExecMessage is a paid mutator transaction binding the contract method 0x5b47f056.
//
// Solidity: function execMessage(bytes32 execProposalID, bytes message) returns()
func (_ProposalExecMessageBoard *ProposalExecMessageBoardTransactorSession) ExecMessage(execProposalID [32]byte, message []byte) (*types.Transaction, error) {
	return _ProposalExecMessageBoard.Contract.ExecMessage(&_ProposalExecMessageBoard.TransactOpts, execProposalID, message)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_ProposalExecMessageBoard *ProposalExecMessageBoardTransactor) Init(opts *bind.TransactOpts, addrReg common.Address) (*types.Transaction, error) {
	return _ProposalExecMessageBoard.contract.Transact(opts, "init", addrReg)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_ProposalExecMessageBoard *ProposalExecMessageBoardSession) Init(addrReg common.Address) (*types.Transaction, error) {
	return _ProposalExecMessageBoard.Contract.Init(&_ProposalExecMessageBoard.TransactOpts, addrReg)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_ProposalExecMessageBoard *ProposalExecMessageBoardTransactorSession) Init(addrReg common.Address) (*types.Transaction, error) {
	return _ProposalExecMessageBoard.Contract.Init(&_ProposalExecMessageBoard.TransactOpts, addrReg)
}

// ProposalExecMessageBoardEMessageIterator is returned from FilterEMessage and is used to iterate over the raw logs and unpacked data for EMessage events raised by the ProposalExecMessageBoard contract.
type ProposalExecMessageBoardEMessageIterator struct {
	Event *ProposalExecMessageBoardEMessage // Event containing the contract specifics and raw log

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
func (it *ProposalExecMessageBoardEMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalExecMessageBoardEMessage)
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
		it.Event = new(ProposalExecMessageBoardEMessage)
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
func (it *ProposalExecMessageBoardEMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalExecMessageBoardEMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalExecMessageBoardEMessage represents a EMessage event raised by the ProposalExecMessageBoard contract.
type ProposalExecMessageBoardEMessage struct {
	TopicID        [32]byte
	Addr           common.Address
	ExecProposalID [32]byte
	NowProposalID  [32]byte
	Message        []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEMessage is a free log retrieval operation binding the contract event 0x6f3f3ec434c9727383cf666515d89caeb650069a19b7ae96e2624b7fbd0b1f0e.
//
// Solidity: event EMessage(bytes32 indexed topicID, address indexed addr, bytes32 indexed execProposalID, bytes32 nowProposalID, bytes message)
func (_ProposalExecMessageBoard *ProposalExecMessageBoardFilterer) FilterEMessage(opts *bind.FilterOpts, topicID [][32]byte, addr []common.Address, execProposalID [][32]byte) (*ProposalExecMessageBoardEMessageIterator, error) {

	var topicIDRule []interface{}
	for _, topicIDItem := range topicID {
		topicIDRule = append(topicIDRule, topicIDItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var execProposalIDRule []interface{}
	for _, execProposalIDItem := range execProposalID {
		execProposalIDRule = append(execProposalIDRule, execProposalIDItem)
	}

	logs, sub, err := _ProposalExecMessageBoard.contract.FilterLogs(opts, "EMessage", topicIDRule, addrRule, execProposalIDRule)
	if err != nil {
		return nil, err
	}
	return &ProposalExecMessageBoardEMessageIterator{contract: _ProposalExecMessageBoard.contract, event: "EMessage", logs: logs, sub: sub}, nil
}

// WatchEMessage is a free log subscription operation binding the contract event 0x6f3f3ec434c9727383cf666515d89caeb650069a19b7ae96e2624b7fbd0b1f0e.
//
// Solidity: event EMessage(bytes32 indexed topicID, address indexed addr, bytes32 indexed execProposalID, bytes32 nowProposalID, bytes message)
func (_ProposalExecMessageBoard *ProposalExecMessageBoardFilterer) WatchEMessage(opts *bind.WatchOpts, sink chan<- *ProposalExecMessageBoardEMessage, topicID [][32]byte, addr []common.Address, execProposalID [][32]byte) (event.Subscription, error) {

	var topicIDRule []interface{}
	for _, topicIDItem := range topicID {
		topicIDRule = append(topicIDRule, topicIDItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var execProposalIDRule []interface{}
	for _, execProposalIDItem := range execProposalID {
		execProposalIDRule = append(execProposalIDRule, execProposalIDItem)
	}

	logs, sub, err := _ProposalExecMessageBoard.contract.WatchLogs(opts, "EMessage", topicIDRule, addrRule, execProposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalExecMessageBoardEMessage)
				if err := _ProposalExecMessageBoard.contract.UnpackLog(event, "EMessage", log); err != nil {
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

// ParseEMessage is a log parse operation binding the contract event 0x6f3f3ec434c9727383cf666515d89caeb650069a19b7ae96e2624b7fbd0b1f0e.
//
// Solidity: event EMessage(bytes32 indexed topicID, address indexed addr, bytes32 indexed execProposalID, bytes32 nowProposalID, bytes message)
func (_ProposalExecMessageBoard *ProposalExecMessageBoardFilterer) ParseEMessage(log types.Log) (*ProposalExecMessageBoardEMessage, error) {
	event := new(ProposalExecMessageBoardEMessage)
	if err := _ProposalExecMessageBoard.contract.UnpackLog(event, "EMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
