// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package InkPayManage

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

// InkPayManageScheduleMember is an auto generated low-level Go binding around an user-defined struct.
type InkPayManageScheduleMember struct {
	Coin     common.Address
	OncePay  *big.Int
	Active   *big.Int
	HasTimes *big.Int
	MaxTimes *big.Int
}

// InkPayManageSetPayMember is an auto generated low-level Go binding around an user-defined struct.
type InkPayManageSetPayMember struct {
	Addr    common.Address
	Coin    common.Address
	OncePay *big.Int
	Desc    []byte
}

// InkPayManageMetaData contains all meta data concerning the InkPayManage contract.
var InkPayManageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oncePay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"desc\",\"type\":\"bytes\"}],\"name\":\"EAddMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ERemoveMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"ESetPaySchedule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timeID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Sign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"hasTimeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oncePay\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"desc\",\"type\":\"bytes\"}],\"internalType\":\"structInkPayManage.SetPayMember[]\",\"name\":\"members\",\"type\":\"tuple[]\"}],\"name\":\"addPayMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"members\",\"type\":\"address[]\"}],\"name\":\"delPayMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getPayMemberInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oncePay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"active\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hasTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTimes\",\"type\":\"uint256\"}],\"internalType\":\"structInkPayManage.ScheduleMember\",\"name\":\"info\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"nowMaxID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"avai\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getScheduleInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ensureSIgnTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getSign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nowID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oncePay\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"desc\",\"type\":\"bytes\"}],\"internalType\":\"structInkPayManage.SetPayMember[]\",\"name\":\"members\",\"type\":\"tuple[]\"}],\"name\":\"setPaySchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeID\",\"type\":\"uint256\"}],\"name\":\"sign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InkPayManageABI is the input ABI used to generate the binding from.
// Deprecated: Use InkPayManageMetaData.ABI instead.
var InkPayManageABI = InkPayManageMetaData.ABI

// InkPayManage is an auto generated Go binding around an Ethereum contract.
type InkPayManage struct {
	InkPayManageCaller     // Read-only binding to the contract
	InkPayManageTransactor // Write-only binding to the contract
	InkPayManageFilterer   // Log filterer for contract events
}

// InkPayManageCaller is an auto generated read-only Go binding around an Ethereum contract.
type InkPayManageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InkPayManageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InkPayManageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InkPayManageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InkPayManageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InkPayManageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InkPayManageSession struct {
	Contract     *InkPayManage     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InkPayManageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InkPayManageCallerSession struct {
	Contract *InkPayManageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// InkPayManageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InkPayManageTransactorSession struct {
	Contract     *InkPayManageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InkPayManageRaw is an auto generated low-level Go binding around an Ethereum contract.
type InkPayManageRaw struct {
	Contract *InkPayManage // Generic contract binding to access the raw methods on
}

// InkPayManageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InkPayManageCallerRaw struct {
	Contract *InkPayManageCaller // Generic read-only contract binding to access the raw methods on
}

// InkPayManageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InkPayManageTransactorRaw struct {
	Contract *InkPayManageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInkPayManage creates a new instance of InkPayManage, bound to a specific deployed contract.
func NewInkPayManage(address common.Address, backend bind.ContractBackend) (*InkPayManage, error) {
	contract, err := bindInkPayManage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InkPayManage{InkPayManageCaller: InkPayManageCaller{contract: contract}, InkPayManageTransactor: InkPayManageTransactor{contract: contract}, InkPayManageFilterer: InkPayManageFilterer{contract: contract}}, nil
}

// NewInkPayManageCaller creates a new read-only instance of InkPayManage, bound to a specific deployed contract.
func NewInkPayManageCaller(address common.Address, caller bind.ContractCaller) (*InkPayManageCaller, error) {
	contract, err := bindInkPayManage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InkPayManageCaller{contract: contract}, nil
}

// NewInkPayManageTransactor creates a new write-only instance of InkPayManage, bound to a specific deployed contract.
func NewInkPayManageTransactor(address common.Address, transactor bind.ContractTransactor) (*InkPayManageTransactor, error) {
	contract, err := bindInkPayManage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InkPayManageTransactor{contract: contract}, nil
}

// NewInkPayManageFilterer creates a new log filterer instance of InkPayManage, bound to a specific deployed contract.
func NewInkPayManageFilterer(address common.Address, filterer bind.ContractFilterer) (*InkPayManageFilterer, error) {
	contract, err := bindInkPayManage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InkPayManageFilterer{contract: contract}, nil
}

// bindInkPayManage binds a generic wrapper to an already deployed contract.
func bindInkPayManage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InkPayManageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InkPayManage *InkPayManageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InkPayManage.Contract.InkPayManageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InkPayManage *InkPayManageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InkPayManage.Contract.InkPayManageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InkPayManage *InkPayManageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InkPayManage.Contract.InkPayManageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InkPayManage *InkPayManageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InkPayManage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InkPayManage *InkPayManageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InkPayManage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InkPayManage *InkPayManageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InkPayManage.Contract.contract.Transact(opts, method, params...)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_InkPayManage *InkPayManageCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InkPayManage.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_InkPayManage *InkPayManageSession) Dao() (common.Address, error) {
	return _InkPayManage.Contract.Dao(&_InkPayManage.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_InkPayManage *InkPayManageCallerSession) Dao() (common.Address, error) {
	return _InkPayManage.Contract.Dao(&_InkPayManage.CallOpts)
}

// GetPayMemberInfo is a free data retrieval call binding the contract method 0xa408f46b.
//
// Solidity: function getPayMemberInfo(uint256 id, address addr) view returns((address,uint256,uint256,uint256,uint256) info, uint256 nowMaxID, uint256 avai)
func (_InkPayManage *InkPayManageCaller) GetPayMemberInfo(opts *bind.CallOpts, id *big.Int, addr common.Address) (struct {
	Info     InkPayManageScheduleMember
	NowMaxID *big.Int
	Avai     *big.Int
}, error) {
	var out []interface{}
	err := _InkPayManage.contract.Call(opts, &out, "getPayMemberInfo", id, addr)

	outstruct := new(struct {
		Info     InkPayManageScheduleMember
		NowMaxID *big.Int
		Avai     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Info = *abi.ConvertType(out[0], new(InkPayManageScheduleMember)).(*InkPayManageScheduleMember)
	outstruct.NowMaxID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Avai = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPayMemberInfo is a free data retrieval call binding the contract method 0xa408f46b.
//
// Solidity: function getPayMemberInfo(uint256 id, address addr) view returns((address,uint256,uint256,uint256,uint256) info, uint256 nowMaxID, uint256 avai)
func (_InkPayManage *InkPayManageSession) GetPayMemberInfo(id *big.Int, addr common.Address) (struct {
	Info     InkPayManageScheduleMember
	NowMaxID *big.Int
	Avai     *big.Int
}, error) {
	return _InkPayManage.Contract.GetPayMemberInfo(&_InkPayManage.CallOpts, id, addr)
}

// GetPayMemberInfo is a free data retrieval call binding the contract method 0xa408f46b.
//
// Solidity: function getPayMemberInfo(uint256 id, address addr) view returns((address,uint256,uint256,uint256,uint256) info, uint256 nowMaxID, uint256 avai)
func (_InkPayManage *InkPayManageCallerSession) GetPayMemberInfo(id *big.Int, addr common.Address) (struct {
	Info     InkPayManageScheduleMember
	NowMaxID *big.Int
	Avai     *big.Int
}, error) {
	return _InkPayManage.Contract.GetPayMemberInfo(&_InkPayManage.CallOpts, id, addr)
}

// GetScheduleInfo is a free data retrieval call binding the contract method 0x033d5e4d.
//
// Solidity: function getScheduleInfo(uint256 id) view returns(uint256 ensureSIgnTime, uint256 duration, uint256 times, uint256 startTime)
func (_InkPayManage *InkPayManageCaller) GetScheduleInfo(opts *bind.CallOpts, id *big.Int) (struct {
	EnsureSIgnTime *big.Int
	Duration       *big.Int
	Times          *big.Int
	StartTime      *big.Int
}, error) {
	var out []interface{}
	err := _InkPayManage.contract.Call(opts, &out, "getScheduleInfo", id)

	outstruct := new(struct {
		EnsureSIgnTime *big.Int
		Duration       *big.Int
		Times          *big.Int
		StartTime      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EnsureSIgnTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Times = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetScheduleInfo is a free data retrieval call binding the contract method 0x033d5e4d.
//
// Solidity: function getScheduleInfo(uint256 id) view returns(uint256 ensureSIgnTime, uint256 duration, uint256 times, uint256 startTime)
func (_InkPayManage *InkPayManageSession) GetScheduleInfo(id *big.Int) (struct {
	EnsureSIgnTime *big.Int
	Duration       *big.Int
	Times          *big.Int
	StartTime      *big.Int
}, error) {
	return _InkPayManage.Contract.GetScheduleInfo(&_InkPayManage.CallOpts, id)
}

// GetScheduleInfo is a free data retrieval call binding the contract method 0x033d5e4d.
//
// Solidity: function getScheduleInfo(uint256 id) view returns(uint256 ensureSIgnTime, uint256 duration, uint256 times, uint256 startTime)
func (_InkPayManage *InkPayManageCallerSession) GetScheduleInfo(id *big.Int) (struct {
	EnsureSIgnTime *big.Int
	Duration       *big.Int
	Times          *big.Int
	StartTime      *big.Int
}, error) {
	return _InkPayManage.Contract.GetScheduleInfo(&_InkPayManage.CallOpts, id)
}

// GetSign is a free data retrieval call binding the contract method 0x564b762f.
//
// Solidity: function getSign(uint256 id, uint256 timeID, address addr) view returns(uint256)
func (_InkPayManage *InkPayManageCaller) GetSign(opts *bind.CallOpts, id *big.Int, timeID *big.Int, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InkPayManage.contract.Call(opts, &out, "getSign", id, timeID, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSign is a free data retrieval call binding the contract method 0x564b762f.
//
// Solidity: function getSign(uint256 id, uint256 timeID, address addr) view returns(uint256)
func (_InkPayManage *InkPayManageSession) GetSign(id *big.Int, timeID *big.Int, addr common.Address) (*big.Int, error) {
	return _InkPayManage.Contract.GetSign(&_InkPayManage.CallOpts, id, timeID, addr)
}

// GetSign is a free data retrieval call binding the contract method 0x564b762f.
//
// Solidity: function getSign(uint256 id, uint256 timeID, address addr) view returns(uint256)
func (_InkPayManage *InkPayManageCallerSession) GetSign(id *big.Int, timeID *big.Int, addr common.Address) (*big.Int, error) {
	return _InkPayManage.Contract.GetSign(&_InkPayManage.CallOpts, id, timeID, addr)
}

// NowID is a free data retrieval call binding the contract method 0xc279577e.
//
// Solidity: function nowID() view returns(uint256)
func (_InkPayManage *InkPayManageCaller) NowID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InkPayManage.contract.Call(opts, &out, "nowID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NowID is a free data retrieval call binding the contract method 0xc279577e.
//
// Solidity: function nowID() view returns(uint256)
func (_InkPayManage *InkPayManageSession) NowID() (*big.Int, error) {
	return _InkPayManage.Contract.NowID(&_InkPayManage.CallOpts)
}

// NowID is a free data retrieval call binding the contract method 0xc279577e.
//
// Solidity: function nowID() view returns(uint256)
func (_InkPayManage *InkPayManageCallerSession) NowID() (*big.Int, error) {
	return _InkPayManage.Contract.NowID(&_InkPayManage.CallOpts)
}

// AddPayMember is a paid mutator transaction binding the contract method 0x875ee74a.
//
// Solidity: function addPayMember(uint256 id, (address,address,uint256,bytes)[] members) returns()
func (_InkPayManage *InkPayManageTransactor) AddPayMember(opts *bind.TransactOpts, id *big.Int, members []InkPayManageSetPayMember) (*types.Transaction, error) {
	return _InkPayManage.contract.Transact(opts, "addPayMember", id, members)
}

// AddPayMember is a paid mutator transaction binding the contract method 0x875ee74a.
//
// Solidity: function addPayMember(uint256 id, (address,address,uint256,bytes)[] members) returns()
func (_InkPayManage *InkPayManageSession) AddPayMember(id *big.Int, members []InkPayManageSetPayMember) (*types.Transaction, error) {
	return _InkPayManage.Contract.AddPayMember(&_InkPayManage.TransactOpts, id, members)
}

// AddPayMember is a paid mutator transaction binding the contract method 0x875ee74a.
//
// Solidity: function addPayMember(uint256 id, (address,address,uint256,bytes)[] members) returns()
func (_InkPayManage *InkPayManageTransactorSession) AddPayMember(id *big.Int, members []InkPayManageSetPayMember) (*types.Transaction, error) {
	return _InkPayManage.Contract.AddPayMember(&_InkPayManage.TransactOpts, id, members)
}

// DelPayMember is a paid mutator transaction binding the contract method 0x5a250334.
//
// Solidity: function delPayMember(uint256 id, address[] members) returns()
func (_InkPayManage *InkPayManageTransactor) DelPayMember(opts *bind.TransactOpts, id *big.Int, members []common.Address) (*types.Transaction, error) {
	return _InkPayManage.contract.Transact(opts, "delPayMember", id, members)
}

// DelPayMember is a paid mutator transaction binding the contract method 0x5a250334.
//
// Solidity: function delPayMember(uint256 id, address[] members) returns()
func (_InkPayManage *InkPayManageSession) DelPayMember(id *big.Int, members []common.Address) (*types.Transaction, error) {
	return _InkPayManage.Contract.DelPayMember(&_InkPayManage.TransactOpts, id, members)
}

// DelPayMember is a paid mutator transaction binding the contract method 0x5a250334.
//
// Solidity: function delPayMember(uint256 id, address[] members) returns()
func (_InkPayManage *InkPayManageTransactorSession) DelPayMember(id *big.Int, members []common.Address) (*types.Transaction, error) {
	return _InkPayManage.Contract.DelPayMember(&_InkPayManage.TransactOpts, id, members)
}

// SetPaySchedule is a paid mutator transaction binding the contract method 0xc793f99b.
//
// Solidity: function setPaySchedule(uint256 duration, uint256 times, uint256 startTime, (address,address,uint256,bytes)[] members) returns()
func (_InkPayManage *InkPayManageTransactor) SetPaySchedule(opts *bind.TransactOpts, duration *big.Int, times *big.Int, startTime *big.Int, members []InkPayManageSetPayMember) (*types.Transaction, error) {
	return _InkPayManage.contract.Transact(opts, "setPaySchedule", duration, times, startTime, members)
}

// SetPaySchedule is a paid mutator transaction binding the contract method 0xc793f99b.
//
// Solidity: function setPaySchedule(uint256 duration, uint256 times, uint256 startTime, (address,address,uint256,bytes)[] members) returns()
func (_InkPayManage *InkPayManageSession) SetPaySchedule(duration *big.Int, times *big.Int, startTime *big.Int, members []InkPayManageSetPayMember) (*types.Transaction, error) {
	return _InkPayManage.Contract.SetPaySchedule(&_InkPayManage.TransactOpts, duration, times, startTime, members)
}

// SetPaySchedule is a paid mutator transaction binding the contract method 0xc793f99b.
//
// Solidity: function setPaySchedule(uint256 duration, uint256 times, uint256 startTime, (address,address,uint256,bytes)[] members) returns()
func (_InkPayManage *InkPayManageTransactorSession) SetPaySchedule(duration *big.Int, times *big.Int, startTime *big.Int, members []InkPayManageSetPayMember) (*types.Transaction, error) {
	return _InkPayManage.Contract.SetPaySchedule(&_InkPayManage.TransactOpts, duration, times, startTime, members)
}

// Sign is a paid mutator transaction binding the contract method 0xb53003ca.
//
// Solidity: function sign(uint256 id, uint256 timeID) returns()
func (_InkPayManage *InkPayManageTransactor) Sign(opts *bind.TransactOpts, id *big.Int, timeID *big.Int) (*types.Transaction, error) {
	return _InkPayManage.contract.Transact(opts, "sign", id, timeID)
}

// Sign is a paid mutator transaction binding the contract method 0xb53003ca.
//
// Solidity: function sign(uint256 id, uint256 timeID) returns()
func (_InkPayManage *InkPayManageSession) Sign(id *big.Int, timeID *big.Int) (*types.Transaction, error) {
	return _InkPayManage.Contract.Sign(&_InkPayManage.TransactOpts, id, timeID)
}

// Sign is a paid mutator transaction binding the contract method 0xb53003ca.
//
// Solidity: function sign(uint256 id, uint256 timeID) returns()
func (_InkPayManage *InkPayManageTransactorSession) Sign(id *big.Int, timeID *big.Int) (*types.Transaction, error) {
	return _InkPayManage.Contract.Sign(&_InkPayManage.TransactOpts, id, timeID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 id) returns()
func (_InkPayManage *InkPayManageTransactor) Withdraw(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _InkPayManage.contract.Transact(opts, "withdraw", id)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 id) returns()
func (_InkPayManage *InkPayManageSession) Withdraw(id *big.Int) (*types.Transaction, error) {
	return _InkPayManage.Contract.Withdraw(&_InkPayManage.TransactOpts, id)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 id) returns()
func (_InkPayManage *InkPayManageTransactorSession) Withdraw(id *big.Int) (*types.Transaction, error) {
	return _InkPayManage.Contract.Withdraw(&_InkPayManage.TransactOpts, id)
}

// InkPayManageEAddMemberIterator is returned from FilterEAddMember and is used to iterate over the raw logs and unpacked data for EAddMember events raised by the InkPayManage contract.
type InkPayManageEAddMemberIterator struct {
	Event *InkPayManageEAddMember // Event containing the contract specifics and raw log

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
func (it *InkPayManageEAddMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InkPayManageEAddMember)
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
		it.Event = new(InkPayManageEAddMember)
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
func (it *InkPayManageEAddMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InkPayManageEAddMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InkPayManageEAddMember represents a EAddMember event raised by the InkPayManage contract.
type InkPayManageEAddMember struct {
	Id      *big.Int
	Addr    common.Address
	Coin    common.Address
	OncePay *big.Int
	Desc    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEAddMember is a free log retrieval operation binding the contract event 0x3ac2a5194b22ad7ce4c32695ca2370742c19e6d46bafad90e325eb21572900d8.
//
// Solidity: event EAddMember(uint256 indexed id, address indexed addr, address coin, uint256 oncePay, bytes desc)
func (_InkPayManage *InkPayManageFilterer) FilterEAddMember(opts *bind.FilterOpts, id []*big.Int, addr []common.Address) (*InkPayManageEAddMemberIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _InkPayManage.contract.FilterLogs(opts, "EAddMember", idRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &InkPayManageEAddMemberIterator{contract: _InkPayManage.contract, event: "EAddMember", logs: logs, sub: sub}, nil
}

// WatchEAddMember is a free log subscription operation binding the contract event 0x3ac2a5194b22ad7ce4c32695ca2370742c19e6d46bafad90e325eb21572900d8.
//
// Solidity: event EAddMember(uint256 indexed id, address indexed addr, address coin, uint256 oncePay, bytes desc)
func (_InkPayManage *InkPayManageFilterer) WatchEAddMember(opts *bind.WatchOpts, sink chan<- *InkPayManageEAddMember, id []*big.Int, addr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _InkPayManage.contract.WatchLogs(opts, "EAddMember", idRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InkPayManageEAddMember)
				if err := _InkPayManage.contract.UnpackLog(event, "EAddMember", log); err != nil {
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

// ParseEAddMember is a log parse operation binding the contract event 0x3ac2a5194b22ad7ce4c32695ca2370742c19e6d46bafad90e325eb21572900d8.
//
// Solidity: event EAddMember(uint256 indexed id, address indexed addr, address coin, uint256 oncePay, bytes desc)
func (_InkPayManage *InkPayManageFilterer) ParseEAddMember(log types.Log) (*InkPayManageEAddMember, error) {
	event := new(InkPayManageEAddMember)
	if err := _InkPayManage.contract.UnpackLog(event, "EAddMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InkPayManageERemoveMemberIterator is returned from FilterERemoveMember and is used to iterate over the raw logs and unpacked data for ERemoveMember events raised by the InkPayManage contract.
type InkPayManageERemoveMemberIterator struct {
	Event *InkPayManageERemoveMember // Event containing the contract specifics and raw log

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
func (it *InkPayManageERemoveMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InkPayManageERemoveMember)
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
		it.Event = new(InkPayManageERemoveMember)
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
func (it *InkPayManageERemoveMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InkPayManageERemoveMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InkPayManageERemoveMember represents a ERemoveMember event raised by the InkPayManage contract.
type InkPayManageERemoveMember struct {
	Id   *big.Int
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterERemoveMember is a free log retrieval operation binding the contract event 0xb90f923cec7d37f73da72e0aa0b6f3a58dbecdf654470db39981a0f1fab2e769.
//
// Solidity: event ERemoveMember(uint256 indexed id, address indexed addr)
func (_InkPayManage *InkPayManageFilterer) FilterERemoveMember(opts *bind.FilterOpts, id []*big.Int, addr []common.Address) (*InkPayManageERemoveMemberIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _InkPayManage.contract.FilterLogs(opts, "ERemoveMember", idRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &InkPayManageERemoveMemberIterator{contract: _InkPayManage.contract, event: "ERemoveMember", logs: logs, sub: sub}, nil
}

// WatchERemoveMember is a free log subscription operation binding the contract event 0xb90f923cec7d37f73da72e0aa0b6f3a58dbecdf654470db39981a0f1fab2e769.
//
// Solidity: event ERemoveMember(uint256 indexed id, address indexed addr)
func (_InkPayManage *InkPayManageFilterer) WatchERemoveMember(opts *bind.WatchOpts, sink chan<- *InkPayManageERemoveMember, id []*big.Int, addr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _InkPayManage.contract.WatchLogs(opts, "ERemoveMember", idRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InkPayManageERemoveMember)
				if err := _InkPayManage.contract.UnpackLog(event, "ERemoveMember", log); err != nil {
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

// ParseERemoveMember is a log parse operation binding the contract event 0xb90f923cec7d37f73da72e0aa0b6f3a58dbecdf654470db39981a0f1fab2e769.
//
// Solidity: event ERemoveMember(uint256 indexed id, address indexed addr)
func (_InkPayManage *InkPayManageFilterer) ParseERemoveMember(log types.Log) (*InkPayManageERemoveMember, error) {
	event := new(InkPayManageERemoveMember)
	if err := _InkPayManage.contract.UnpackLog(event, "ERemoveMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InkPayManageESetPayScheduleIterator is returned from FilterESetPaySchedule and is used to iterate over the raw logs and unpacked data for ESetPaySchedule events raised by the InkPayManage contract.
type InkPayManageESetPayScheduleIterator struct {
	Event *InkPayManageESetPaySchedule // Event containing the contract specifics and raw log

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
func (it *InkPayManageESetPayScheduleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InkPayManageESetPaySchedule)
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
		it.Event = new(InkPayManageESetPaySchedule)
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
func (it *InkPayManageESetPayScheduleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InkPayManageESetPayScheduleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InkPayManageESetPaySchedule represents a ESetPaySchedule event raised by the InkPayManage contract.
type InkPayManageESetPaySchedule struct {
	Id        *big.Int
	Duration  *big.Int
	Times     *big.Int
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterESetPaySchedule is a free log retrieval operation binding the contract event 0x166beaec75778cf5da131a86b5bbe95a1a5060bc049f9064ebaa3e13c814eedc.
//
// Solidity: event ESetPaySchedule(uint256 indexed id, uint256 duration, uint256 times, uint256 startTime)
func (_InkPayManage *InkPayManageFilterer) FilterESetPaySchedule(opts *bind.FilterOpts, id []*big.Int) (*InkPayManageESetPayScheduleIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _InkPayManage.contract.FilterLogs(opts, "ESetPaySchedule", idRule)
	if err != nil {
		return nil, err
	}
	return &InkPayManageESetPayScheduleIterator{contract: _InkPayManage.contract, event: "ESetPaySchedule", logs: logs, sub: sub}, nil
}

// WatchESetPaySchedule is a free log subscription operation binding the contract event 0x166beaec75778cf5da131a86b5bbe95a1a5060bc049f9064ebaa3e13c814eedc.
//
// Solidity: event ESetPaySchedule(uint256 indexed id, uint256 duration, uint256 times, uint256 startTime)
func (_InkPayManage *InkPayManageFilterer) WatchESetPaySchedule(opts *bind.WatchOpts, sink chan<- *InkPayManageESetPaySchedule, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _InkPayManage.contract.WatchLogs(opts, "ESetPaySchedule", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InkPayManageESetPaySchedule)
				if err := _InkPayManage.contract.UnpackLog(event, "ESetPaySchedule", log); err != nil {
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

// ParseESetPaySchedule is a log parse operation binding the contract event 0x166beaec75778cf5da131a86b5bbe95a1a5060bc049f9064ebaa3e13c814eedc.
//
// Solidity: event ESetPaySchedule(uint256 indexed id, uint256 duration, uint256 times, uint256 startTime)
func (_InkPayManage *InkPayManageFilterer) ParseESetPaySchedule(log types.Log) (*InkPayManageESetPaySchedule, error) {
	event := new(InkPayManageESetPaySchedule)
	if err := _InkPayManage.contract.UnpackLog(event, "ESetPaySchedule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InkPayManageSignIterator is returned from FilterSign and is used to iterate over the raw logs and unpacked data for Sign events raised by the InkPayManage contract.
type InkPayManageSignIterator struct {
	Event *InkPayManageSign // Event containing the contract specifics and raw log

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
func (it *InkPayManageSignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InkPayManageSign)
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
		it.Event = new(InkPayManageSign)
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
func (it *InkPayManageSignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InkPayManageSignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InkPayManageSign represents a Sign event raised by the InkPayManage contract.
type InkPayManageSign struct {
	Id     *big.Int
	TimeID *big.Int
	Addr   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSign is a free log retrieval operation binding the contract event 0xe4d0f4b9a5e9ea15391ac52534f8bda735157ffb6b50cb9e12d39aec3beb6977.
//
// Solidity: event Sign(uint256 indexed id, uint256 indexed timeID, address indexed addr)
func (_InkPayManage *InkPayManageFilterer) FilterSign(opts *bind.FilterOpts, id []*big.Int, timeID []*big.Int, addr []common.Address) (*InkPayManageSignIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var timeIDRule []interface{}
	for _, timeIDItem := range timeID {
		timeIDRule = append(timeIDRule, timeIDItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _InkPayManage.contract.FilterLogs(opts, "Sign", idRule, timeIDRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &InkPayManageSignIterator{contract: _InkPayManage.contract, event: "Sign", logs: logs, sub: sub}, nil
}

// WatchSign is a free log subscription operation binding the contract event 0xe4d0f4b9a5e9ea15391ac52534f8bda735157ffb6b50cb9e12d39aec3beb6977.
//
// Solidity: event Sign(uint256 indexed id, uint256 indexed timeID, address indexed addr)
func (_InkPayManage *InkPayManageFilterer) WatchSign(opts *bind.WatchOpts, sink chan<- *InkPayManageSign, id []*big.Int, timeID []*big.Int, addr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var timeIDRule []interface{}
	for _, timeIDItem := range timeID {
		timeIDRule = append(timeIDRule, timeIDItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _InkPayManage.contract.WatchLogs(opts, "Sign", idRule, timeIDRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InkPayManageSign)
				if err := _InkPayManage.contract.UnpackLog(event, "Sign", log); err != nil {
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

// ParseSign is a log parse operation binding the contract event 0xe4d0f4b9a5e9ea15391ac52534f8bda735157ffb6b50cb9e12d39aec3beb6977.
//
// Solidity: event Sign(uint256 indexed id, uint256 indexed timeID, address indexed addr)
func (_InkPayManage *InkPayManageFilterer) ParseSign(log types.Log) (*InkPayManageSign, error) {
	event := new(InkPayManageSign)
	if err := _InkPayManage.contract.UnpackLog(event, "Sign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InkPayManageWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the InkPayManage contract.
type InkPayManageWithdrawIterator struct {
	Event *InkPayManageWithdraw // Event containing the contract specifics and raw log

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
func (it *InkPayManageWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InkPayManageWithdraw)
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
		it.Event = new(InkPayManageWithdraw)
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
func (it *InkPayManageWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InkPayManageWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InkPayManageWithdraw represents a Withdraw event raised by the InkPayManage contract.
type InkPayManageWithdraw struct {
	Id        *big.Int
	Addr      common.Address
	HasTimeID *big.Int
	Total     *big.Int
	Coin      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x6df5b29a0ae372e6e7a943a78c8e1fb0054174551a9d1e6a0f169e1209797e8d.
//
// Solidity: event Withdraw(uint256 indexed id, address indexed addr, uint256 indexed hasTimeID, uint256 total, address coin)
func (_InkPayManage *InkPayManageFilterer) FilterWithdraw(opts *bind.FilterOpts, id []*big.Int, addr []common.Address, hasTimeID []*big.Int) (*InkPayManageWithdrawIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var hasTimeIDRule []interface{}
	for _, hasTimeIDItem := range hasTimeID {
		hasTimeIDRule = append(hasTimeIDRule, hasTimeIDItem)
	}

	logs, sub, err := _InkPayManage.contract.FilterLogs(opts, "Withdraw", idRule, addrRule, hasTimeIDRule)
	if err != nil {
		return nil, err
	}
	return &InkPayManageWithdrawIterator{contract: _InkPayManage.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x6df5b29a0ae372e6e7a943a78c8e1fb0054174551a9d1e6a0f169e1209797e8d.
//
// Solidity: event Withdraw(uint256 indexed id, address indexed addr, uint256 indexed hasTimeID, uint256 total, address coin)
func (_InkPayManage *InkPayManageFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *InkPayManageWithdraw, id []*big.Int, addr []common.Address, hasTimeID []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var hasTimeIDRule []interface{}
	for _, hasTimeIDItem := range hasTimeID {
		hasTimeIDRule = append(hasTimeIDRule, hasTimeIDItem)
	}

	logs, sub, err := _InkPayManage.contract.WatchLogs(opts, "Withdraw", idRule, addrRule, hasTimeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InkPayManageWithdraw)
				if err := _InkPayManage.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x6df5b29a0ae372e6e7a943a78c8e1fb0054174551a9d1e6a0f169e1209797e8d.
//
// Solidity: event Withdraw(uint256 indexed id, address indexed addr, uint256 indexed hasTimeID, uint256 total, address coin)
func (_InkPayManage *InkPayManageFilterer) ParseWithdraw(log types.Log) (*InkPayManageWithdraw, error) {
	event := new(InkPayManageWithdraw)
	if err := _InkPayManage.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
