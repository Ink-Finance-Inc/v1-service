// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ProposalRegistry

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

// IProposalOperatorProposalApplyInfo is an auto generated low-level Go binding around an user-defined struct.
type IProposalOperatorProposalApplyInfo struct {
	Items   [][]byte
	Headers [][]byte
}

// IProposalRegistryInfoProposal is an auto generated low-level Go binding around an user-defined struct.
type IProposalRegistryInfoProposal struct {
	Status     uint8
	ProposalID [32]byte
	TopicID    [32]byte
	Dao        common.Address
}

// IProposalRegistryInfoTopic is an auto generated low-level Go binding around an user-defined struct.
type IProposalRegistryInfoTopic struct {
	TopicID     [32]byte
	ProposalIDs [][32]byte
	Dao         common.Address
}

// ProposalRegistryMetaData contains all meta data concerning the ProposalRegistry contract.
var ProposalRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"metadata\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"kvData\",\"type\":\"bytes[]\"}],\"name\":\"ENewProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"kvData\",\"type\":\"bytes[]\"}],\"name\":\"EProposalAppend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"agree\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"}],\"name\":\"EProposalDecide\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"ETopicCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"ETopicFix\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addrRegistry\",\"outputs\":[{\"internalType\":\"contractIAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"kvData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"changeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"agree\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"decideProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"startKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getProposalAllMetadata\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"kvData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getProposalKvData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"typeID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"startKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getProposalKvDataKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getProposalMetadata\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"typeID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"getProposalSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIProposalRegistryInfo.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"}],\"internalType\":\"structIProposalRegistryInfo.Proposal\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"startKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getTopicAllMetadata\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"kvData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"}],\"name\":\"getTopicInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proposalIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"}],\"internalType\":\"structIProposalRegistryInfo.Topic\",\"name\":\"topic\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getTopicKeyProposal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getTopicMetadata\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"typeID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"topicID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"startKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getTopicMetadataKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"typeID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"}],\"name\":\"helpEncodeMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrReg\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"items\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"headers\",\"type\":\"bytes[]\"}],\"internalType\":\"structIProposalOperator.ProposalApplyInfo\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"newProposal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ProposalRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalRegistryMetaData.ABI instead.
var ProposalRegistryABI = ProposalRegistryMetaData.ABI

// ProposalRegistry is an auto generated Go binding around an Ethereum contract.
type ProposalRegistry struct {
	ProposalRegistryCaller     // Read-only binding to the contract
	ProposalRegistryTransactor // Write-only binding to the contract
	ProposalRegistryFilterer   // Log filterer for contract events
}

// ProposalRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposalRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposalRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposalRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposalRegistrySession struct {
	Contract     *ProposalRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposalRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposalRegistryCallerSession struct {
	Contract *ProposalRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProposalRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposalRegistryTransactorSession struct {
	Contract     *ProposalRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProposalRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposalRegistryRaw struct {
	Contract *ProposalRegistry // Generic contract binding to access the raw methods on
}

// ProposalRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposalRegistryCallerRaw struct {
	Contract *ProposalRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposalRegistryTransactorRaw struct {
	Contract *ProposalRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposalRegistry creates a new instance of ProposalRegistry, bound to a specific deployed contract.
func NewProposalRegistry(address common.Address, backend bind.ContractBackend) (*ProposalRegistry, error) {
	contract, err := bindProposalRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistry{ProposalRegistryCaller: ProposalRegistryCaller{contract: contract}, ProposalRegistryTransactor: ProposalRegistryTransactor{contract: contract}, ProposalRegistryFilterer: ProposalRegistryFilterer{contract: contract}}, nil
}

// NewProposalRegistryCaller creates a new read-only instance of ProposalRegistry, bound to a specific deployed contract.
func NewProposalRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProposalRegistryCaller, error) {
	contract, err := bindProposalRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryCaller{contract: contract}, nil
}

// NewProposalRegistryTransactor creates a new write-only instance of ProposalRegistry, bound to a specific deployed contract.
func NewProposalRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalRegistryTransactor, error) {
	contract, err := bindProposalRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryTransactor{contract: contract}, nil
}

// NewProposalRegistryFilterer creates a new log filterer instance of ProposalRegistry, bound to a specific deployed contract.
func NewProposalRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalRegistryFilterer, error) {
	contract, err := bindProposalRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryFilterer{contract: contract}, nil
}

// bindProposalRegistry binds a generic wrapper to an already deployed contract.
func bindProposalRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProposalRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalRegistry *ProposalRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalRegistry.Contract.ProposalRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalRegistry *ProposalRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.ProposalRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalRegistry *ProposalRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.ProposalRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalRegistry *ProposalRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalRegistry *ProposalRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalRegistry *ProposalRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.contract.Transact(opts, method, params...)
}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_ProposalRegistry *ProposalRegistryCaller) AddrRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "addrRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_ProposalRegistry *ProposalRegistrySession) AddrRegistry() (common.Address, error) {
	return _ProposalRegistry.Contract.AddrRegistry(&_ProposalRegistry.CallOpts)
}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_ProposalRegistry *ProposalRegistryCallerSession) AddrRegistry() (common.Address, error) {
	return _ProposalRegistry.Contract.AddrRegistry(&_ProposalRegistry.CallOpts)
}

// GetProposalAllMetadata is a free data retrieval call binding the contract method 0x2c1d49f7.
//
// Solidity: function getProposalAllMetadata(bytes32 proposalID, string startKey, uint256 pageSize) view returns(bytes[] kvData)
func (_ProposalRegistry *ProposalRegistryCaller) GetProposalAllMetadata(opts *bind.CallOpts, proposalID [32]byte, startKey string, pageSize *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getProposalAllMetadata", proposalID, startKey, pageSize)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetProposalAllMetadata is a free data retrieval call binding the contract method 0x2c1d49f7.
//
// Solidity: function getProposalAllMetadata(bytes32 proposalID, string startKey, uint256 pageSize) view returns(bytes[] kvData)
func (_ProposalRegistry *ProposalRegistrySession) GetProposalAllMetadata(proposalID [32]byte, startKey string, pageSize *big.Int) ([][]byte, error) {
	return _ProposalRegistry.Contract.GetProposalAllMetadata(&_ProposalRegistry.CallOpts, proposalID, startKey, pageSize)
}

// GetProposalAllMetadata is a free data retrieval call binding the contract method 0x2c1d49f7.
//
// Solidity: function getProposalAllMetadata(bytes32 proposalID, string startKey, uint256 pageSize) view returns(bytes[] kvData)
func (_ProposalRegistry *ProposalRegistryCallerSession) GetProposalAllMetadata(proposalID [32]byte, startKey string, pageSize *big.Int) ([][]byte, error) {
	return _ProposalRegistry.Contract.GetProposalAllMetadata(&_ProposalRegistry.CallOpts, proposalID, startKey, pageSize)
}

// GetProposalKvData is a free data retrieval call binding the contract method 0xac8e1bb6.
//
// Solidity: function getProposalKvData(bytes32 proposalID, string key) view returns(bytes32 typeID, bytes data)
func (_ProposalRegistry *ProposalRegistryCaller) GetProposalKvData(opts *bind.CallOpts, proposalID [32]byte, key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getProposalKvData", proposalID, key)

	outstruct := new(struct {
		TypeID [32]byte
		Data   []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TypeID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Data = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetProposalKvData is a free data retrieval call binding the contract method 0xac8e1bb6.
//
// Solidity: function getProposalKvData(bytes32 proposalID, string key) view returns(bytes32 typeID, bytes data)
func (_ProposalRegistry *ProposalRegistrySession) GetProposalKvData(proposalID [32]byte, key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	return _ProposalRegistry.Contract.GetProposalKvData(&_ProposalRegistry.CallOpts, proposalID, key)
}

// GetProposalKvData is a free data retrieval call binding the contract method 0xac8e1bb6.
//
// Solidity: function getProposalKvData(bytes32 proposalID, string key) view returns(bytes32 typeID, bytes data)
func (_ProposalRegistry *ProposalRegistryCallerSession) GetProposalKvData(proposalID [32]byte, key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	return _ProposalRegistry.Contract.GetProposalKvData(&_ProposalRegistry.CallOpts, proposalID, key)
}

// GetProposalKvDataKeys is a free data retrieval call binding the contract method 0xf0413cd9.
//
// Solidity: function getProposalKvDataKeys(bytes32 proposalID, string startKey, uint256 pageSize) view returns(string[] keys)
func (_ProposalRegistry *ProposalRegistryCaller) GetProposalKvDataKeys(opts *bind.CallOpts, proposalID [32]byte, startKey string, pageSize *big.Int) ([]string, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getProposalKvDataKeys", proposalID, startKey, pageSize)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetProposalKvDataKeys is a free data retrieval call binding the contract method 0xf0413cd9.
//
// Solidity: function getProposalKvDataKeys(bytes32 proposalID, string startKey, uint256 pageSize) view returns(string[] keys)
func (_ProposalRegistry *ProposalRegistrySession) GetProposalKvDataKeys(proposalID [32]byte, startKey string, pageSize *big.Int) ([]string, error) {
	return _ProposalRegistry.Contract.GetProposalKvDataKeys(&_ProposalRegistry.CallOpts, proposalID, startKey, pageSize)
}

// GetProposalKvDataKeys is a free data retrieval call binding the contract method 0xf0413cd9.
//
// Solidity: function getProposalKvDataKeys(bytes32 proposalID, string startKey, uint256 pageSize) view returns(string[] keys)
func (_ProposalRegistry *ProposalRegistryCallerSession) GetProposalKvDataKeys(proposalID [32]byte, startKey string, pageSize *big.Int) ([]string, error) {
	return _ProposalRegistry.Contract.GetProposalKvDataKeys(&_ProposalRegistry.CallOpts, proposalID, startKey, pageSize)
}

// GetProposalMetadata is a free data retrieval call binding the contract method 0x692c2b63.
//
// Solidity: function getProposalMetadata(bytes32 proposalID, string key) view returns(bytes32 typeID, bytes data)
func (_ProposalRegistry *ProposalRegistryCaller) GetProposalMetadata(opts *bind.CallOpts, proposalID [32]byte, key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getProposalMetadata", proposalID, key)

	outstruct := new(struct {
		TypeID [32]byte
		Data   []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TypeID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Data = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetProposalMetadata is a free data retrieval call binding the contract method 0x692c2b63.
//
// Solidity: function getProposalMetadata(bytes32 proposalID, string key) view returns(bytes32 typeID, bytes data)
func (_ProposalRegistry *ProposalRegistrySession) GetProposalMetadata(proposalID [32]byte, key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	return _ProposalRegistry.Contract.GetProposalMetadata(&_ProposalRegistry.CallOpts, proposalID, key)
}

// GetProposalMetadata is a free data retrieval call binding the contract method 0x692c2b63.
//
// Solidity: function getProposalMetadata(bytes32 proposalID, string key) view returns(bytes32 typeID, bytes data)
func (_ProposalRegistry *ProposalRegistryCallerSession) GetProposalMetadata(proposalID [32]byte, key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	return _ProposalRegistry.Contract.GetProposalMetadata(&_ProposalRegistry.CallOpts, proposalID, key)
}

// GetProposalSummary is a free data retrieval call binding the contract method 0xfb1a8136.
//
// Solidity: function getProposalSummary(bytes32 proposalID) view returns((uint8,bytes32,bytes32,address) proposal)
func (_ProposalRegistry *ProposalRegistryCaller) GetProposalSummary(opts *bind.CallOpts, proposalID [32]byte) (IProposalRegistryInfoProposal, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getProposalSummary", proposalID)

	if err != nil {
		return *new(IProposalRegistryInfoProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(IProposalRegistryInfoProposal)).(*IProposalRegistryInfoProposal)

	return out0, err

}

// GetProposalSummary is a free data retrieval call binding the contract method 0xfb1a8136.
//
// Solidity: function getProposalSummary(bytes32 proposalID) view returns((uint8,bytes32,bytes32,address) proposal)
func (_ProposalRegistry *ProposalRegistrySession) GetProposalSummary(proposalID [32]byte) (IProposalRegistryInfoProposal, error) {
	return _ProposalRegistry.Contract.GetProposalSummary(&_ProposalRegistry.CallOpts, proposalID)
}

// GetProposalSummary is a free data retrieval call binding the contract method 0xfb1a8136.
//
// Solidity: function getProposalSummary(bytes32 proposalID) view returns((uint8,bytes32,bytes32,address) proposal)
func (_ProposalRegistry *ProposalRegistryCallerSession) GetProposalSummary(proposalID [32]byte) (IProposalRegistryInfoProposal, error) {
	return _ProposalRegistry.Contract.GetProposalSummary(&_ProposalRegistry.CallOpts, proposalID)
}

// GetTopicAllMetadata is a free data retrieval call binding the contract method 0x56c4ffa7.
//
// Solidity: function getTopicAllMetadata(bytes32 topicID, string startKey, uint256 pageSize) view returns(bytes[] kvData)
func (_ProposalRegistry *ProposalRegistryCaller) GetTopicAllMetadata(opts *bind.CallOpts, topicID [32]byte, startKey string, pageSize *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getTopicAllMetadata", topicID, startKey, pageSize)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetTopicAllMetadata is a free data retrieval call binding the contract method 0x56c4ffa7.
//
// Solidity: function getTopicAllMetadata(bytes32 topicID, string startKey, uint256 pageSize) view returns(bytes[] kvData)
func (_ProposalRegistry *ProposalRegistrySession) GetTopicAllMetadata(topicID [32]byte, startKey string, pageSize *big.Int) ([][]byte, error) {
	return _ProposalRegistry.Contract.GetTopicAllMetadata(&_ProposalRegistry.CallOpts, topicID, startKey, pageSize)
}

// GetTopicAllMetadata is a free data retrieval call binding the contract method 0x56c4ffa7.
//
// Solidity: function getTopicAllMetadata(bytes32 topicID, string startKey, uint256 pageSize) view returns(bytes[] kvData)
func (_ProposalRegistry *ProposalRegistryCallerSession) GetTopicAllMetadata(topicID [32]byte, startKey string, pageSize *big.Int) ([][]byte, error) {
	return _ProposalRegistry.Contract.GetTopicAllMetadata(&_ProposalRegistry.CallOpts, topicID, startKey, pageSize)
}

// GetTopicInfo is a free data retrieval call binding the contract method 0xa10d4f2f.
//
// Solidity: function getTopicInfo(bytes32 topicID) view returns((bytes32,bytes32[],address) topic)
func (_ProposalRegistry *ProposalRegistryCaller) GetTopicInfo(opts *bind.CallOpts, topicID [32]byte) (IProposalRegistryInfoTopic, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getTopicInfo", topicID)

	if err != nil {
		return *new(IProposalRegistryInfoTopic), err
	}

	out0 := *abi.ConvertType(out[0], new(IProposalRegistryInfoTopic)).(*IProposalRegistryInfoTopic)

	return out0, err

}

// GetTopicInfo is a free data retrieval call binding the contract method 0xa10d4f2f.
//
// Solidity: function getTopicInfo(bytes32 topicID) view returns((bytes32,bytes32[],address) topic)
func (_ProposalRegistry *ProposalRegistrySession) GetTopicInfo(topicID [32]byte) (IProposalRegistryInfoTopic, error) {
	return _ProposalRegistry.Contract.GetTopicInfo(&_ProposalRegistry.CallOpts, topicID)
}

// GetTopicInfo is a free data retrieval call binding the contract method 0xa10d4f2f.
//
// Solidity: function getTopicInfo(bytes32 topicID) view returns((bytes32,bytes32[],address) topic)
func (_ProposalRegistry *ProposalRegistryCallerSession) GetTopicInfo(topicID [32]byte) (IProposalRegistryInfoTopic, error) {
	return _ProposalRegistry.Contract.GetTopicInfo(&_ProposalRegistry.CallOpts, topicID)
}

// GetTopicKeyProposal is a free data retrieval call binding the contract method 0x712653a8.
//
// Solidity: function getTopicKeyProposal(bytes32 topicID, string key) view returns(bytes32 proposalID)
func (_ProposalRegistry *ProposalRegistryCaller) GetTopicKeyProposal(opts *bind.CallOpts, topicID [32]byte, key string) ([32]byte, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getTopicKeyProposal", topicID, key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTopicKeyProposal is a free data retrieval call binding the contract method 0x712653a8.
//
// Solidity: function getTopicKeyProposal(bytes32 topicID, string key) view returns(bytes32 proposalID)
func (_ProposalRegistry *ProposalRegistrySession) GetTopicKeyProposal(topicID [32]byte, key string) ([32]byte, error) {
	return _ProposalRegistry.Contract.GetTopicKeyProposal(&_ProposalRegistry.CallOpts, topicID, key)
}

// GetTopicKeyProposal is a free data retrieval call binding the contract method 0x712653a8.
//
// Solidity: function getTopicKeyProposal(bytes32 topicID, string key) view returns(bytes32 proposalID)
func (_ProposalRegistry *ProposalRegistryCallerSession) GetTopicKeyProposal(topicID [32]byte, key string) ([32]byte, error) {
	return _ProposalRegistry.Contract.GetTopicKeyProposal(&_ProposalRegistry.CallOpts, topicID, key)
}

// GetTopicMetadata is a free data retrieval call binding the contract method 0x15813c18.
//
// Solidity: function getTopicMetadata(bytes32 topicID, string key) view returns(bytes32 typeID, bytes data)
func (_ProposalRegistry *ProposalRegistryCaller) GetTopicMetadata(opts *bind.CallOpts, topicID [32]byte, key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getTopicMetadata", topicID, key)

	outstruct := new(struct {
		TypeID [32]byte
		Data   []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TypeID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Data = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetTopicMetadata is a free data retrieval call binding the contract method 0x15813c18.
//
// Solidity: function getTopicMetadata(bytes32 topicID, string key) view returns(bytes32 typeID, bytes data)
func (_ProposalRegistry *ProposalRegistrySession) GetTopicMetadata(topicID [32]byte, key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	return _ProposalRegistry.Contract.GetTopicMetadata(&_ProposalRegistry.CallOpts, topicID, key)
}

// GetTopicMetadata is a free data retrieval call binding the contract method 0x15813c18.
//
// Solidity: function getTopicMetadata(bytes32 topicID, string key) view returns(bytes32 typeID, bytes data)
func (_ProposalRegistry *ProposalRegistryCallerSession) GetTopicMetadata(topicID [32]byte, key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	return _ProposalRegistry.Contract.GetTopicMetadata(&_ProposalRegistry.CallOpts, topicID, key)
}

// GetTopicMetadataKeys is a free data retrieval call binding the contract method 0x34388ecc.
//
// Solidity: function getTopicMetadataKeys(bytes32 topicID, string startKey, uint256 pageSize) view returns(string[] keys)
func (_ProposalRegistry *ProposalRegistryCaller) GetTopicMetadataKeys(opts *bind.CallOpts, topicID [32]byte, startKey string, pageSize *big.Int) ([]string, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "getTopicMetadataKeys", topicID, startKey, pageSize)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetTopicMetadataKeys is a free data retrieval call binding the contract method 0x34388ecc.
//
// Solidity: function getTopicMetadataKeys(bytes32 topicID, string startKey, uint256 pageSize) view returns(string[] keys)
func (_ProposalRegistry *ProposalRegistrySession) GetTopicMetadataKeys(topicID [32]byte, startKey string, pageSize *big.Int) ([]string, error) {
	return _ProposalRegistry.Contract.GetTopicMetadataKeys(&_ProposalRegistry.CallOpts, topicID, startKey, pageSize)
}

// GetTopicMetadataKeys is a free data retrieval call binding the contract method 0x34388ecc.
//
// Solidity: function getTopicMetadataKeys(bytes32 topicID, string startKey, uint256 pageSize) view returns(string[] keys)
func (_ProposalRegistry *ProposalRegistryCallerSession) GetTopicMetadataKeys(topicID [32]byte, startKey string, pageSize *big.Int) ([]string, error) {
	return _ProposalRegistry.Contract.GetTopicMetadataKeys(&_ProposalRegistry.CallOpts, topicID, startKey, pageSize)
}

// HelpEncodeMetadata is a free data retrieval call binding the contract method 0x6bb80798.
//
// Solidity: function helpEncodeMetadata(string key, bytes32 typeID, bytes data, string desc) pure returns(bytes)
func (_ProposalRegistry *ProposalRegistryCaller) HelpEncodeMetadata(opts *bind.CallOpts, key string, typeID [32]byte, data []byte, desc string) ([]byte, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "helpEncodeMetadata", key, typeID, data, desc)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// HelpEncodeMetadata is a free data retrieval call binding the contract method 0x6bb80798.
//
// Solidity: function helpEncodeMetadata(string key, bytes32 typeID, bytes data, string desc) pure returns(bytes)
func (_ProposalRegistry *ProposalRegistrySession) HelpEncodeMetadata(key string, typeID [32]byte, data []byte, desc string) ([]byte, error) {
	return _ProposalRegistry.Contract.HelpEncodeMetadata(&_ProposalRegistry.CallOpts, key, typeID, data, desc)
}

// HelpEncodeMetadata is a free data retrieval call binding the contract method 0x6bb80798.
//
// Solidity: function helpEncodeMetadata(string key, bytes32 typeID, bytes data, string desc) pure returns(bytes)
func (_ProposalRegistry *ProposalRegistryCallerSession) HelpEncodeMetadata(key string, typeID [32]byte, data []byte, desc string) ([]byte, error) {
	return _ProposalRegistry.Contract.HelpEncodeMetadata(&_ProposalRegistry.CallOpts, key, typeID, data, desc)
}

// ProposalNum is a free data retrieval call binding the contract method 0x9f2f1942.
//
// Solidity: function proposalNum() view returns(uint256)
func (_ProposalRegistry *ProposalRegistryCaller) ProposalNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "proposalNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalNum is a free data retrieval call binding the contract method 0x9f2f1942.
//
// Solidity: function proposalNum() view returns(uint256)
func (_ProposalRegistry *ProposalRegistrySession) ProposalNum() (*big.Int, error) {
	return _ProposalRegistry.Contract.ProposalNum(&_ProposalRegistry.CallOpts)
}

// ProposalNum is a free data retrieval call binding the contract method 0x9f2f1942.
//
// Solidity: function proposalNum() view returns(uint256)
func (_ProposalRegistry *ProposalRegistryCallerSession) ProposalNum() (*big.Int, error) {
	return _ProposalRegistry.Contract.ProposalNum(&_ProposalRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProposalRegistry *ProposalRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ProposalRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProposalRegistry *ProposalRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProposalRegistry.Contract.SupportsInterface(&_ProposalRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProposalRegistry *ProposalRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProposalRegistry.Contract.SupportsInterface(&_ProposalRegistry.CallOpts, interfaceId)
}

// ChangeProposal is a paid mutator transaction binding the contract method 0xf406f799.
//
// Solidity: function changeProposal(bytes32 proposalID, bytes[] kvData, bytes ) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) ChangeProposal(opts *bind.TransactOpts, proposalID [32]byte, kvData [][]byte, arg2 []byte) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "changeProposal", proposalID, kvData, arg2)
}

// ChangeProposal is a paid mutator transaction binding the contract method 0xf406f799.
//
// Solidity: function changeProposal(bytes32 proposalID, bytes[] kvData, bytes ) returns()
func (_ProposalRegistry *ProposalRegistrySession) ChangeProposal(proposalID [32]byte, kvData [][]byte, arg2 []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.ChangeProposal(&_ProposalRegistry.TransactOpts, proposalID, kvData, arg2)
}

// ChangeProposal is a paid mutator transaction binding the contract method 0xf406f799.
//
// Solidity: function changeProposal(bytes32 proposalID, bytes[] kvData, bytes ) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) ChangeProposal(proposalID [32]byte, kvData [][]byte, arg2 []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.ChangeProposal(&_ProposalRegistry.TransactOpts, proposalID, kvData, arg2)
}

// DecideProposal is a paid mutator transaction binding the contract method 0x67024e77.
//
// Solidity: function decideProposal(bytes32 proposalID, bool agree, bytes ) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) DecideProposal(opts *bind.TransactOpts, proposalID [32]byte, agree bool, arg2 []byte) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "decideProposal", proposalID, agree, arg2)
}

// DecideProposal is a paid mutator transaction binding the contract method 0x67024e77.
//
// Solidity: function decideProposal(bytes32 proposalID, bool agree, bytes ) returns()
func (_ProposalRegistry *ProposalRegistrySession) DecideProposal(proposalID [32]byte, agree bool, arg2 []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.DecideProposal(&_ProposalRegistry.TransactOpts, proposalID, agree, arg2)
}

// DecideProposal is a paid mutator transaction binding the contract method 0x67024e77.
//
// Solidity: function decideProposal(bytes32 proposalID, bool agree, bytes ) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) DecideProposal(proposalID [32]byte, agree bool, arg2 []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.DecideProposal(&_ProposalRegistry.TransactOpts, proposalID, agree, arg2)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_ProposalRegistry *ProposalRegistryTransactor) Init(opts *bind.TransactOpts, addrReg common.Address) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "init", addrReg)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_ProposalRegistry *ProposalRegistrySession) Init(addrReg common.Address) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.Init(&_ProposalRegistry.TransactOpts, addrReg)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_ProposalRegistry *ProposalRegistryTransactorSession) Init(addrReg common.Address) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.Init(&_ProposalRegistry.TransactOpts, addrReg)
}

// NewProposal is a paid mutator transaction binding the contract method 0x8d2a2869.
//
// Solidity: function newProposal((bytes[],bytes[]) proposal, bytes ) returns(bytes32 proposalID)
func (_ProposalRegistry *ProposalRegistryTransactor) NewProposal(opts *bind.TransactOpts, proposal IProposalOperatorProposalApplyInfo, arg1 []byte) (*types.Transaction, error) {
	return _ProposalRegistry.contract.Transact(opts, "newProposal", proposal, arg1)
}

// NewProposal is a paid mutator transaction binding the contract method 0x8d2a2869.
//
// Solidity: function newProposal((bytes[],bytes[]) proposal, bytes ) returns(bytes32 proposalID)
func (_ProposalRegistry *ProposalRegistrySession) NewProposal(proposal IProposalOperatorProposalApplyInfo, arg1 []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.NewProposal(&_ProposalRegistry.TransactOpts, proposal, arg1)
}

// NewProposal is a paid mutator transaction binding the contract method 0x8d2a2869.
//
// Solidity: function newProposal((bytes[],bytes[]) proposal, bytes ) returns(bytes32 proposalID)
func (_ProposalRegistry *ProposalRegistryTransactorSession) NewProposal(proposal IProposalOperatorProposalApplyInfo, arg1 []byte) (*types.Transaction, error) {
	return _ProposalRegistry.Contract.NewProposal(&_ProposalRegistry.TransactOpts, proposal, arg1)
}

// ProposalRegistryENewProposalIterator is returned from FilterENewProposal and is used to iterate over the raw logs and unpacked data for ENewProposal events raised by the ProposalRegistry contract.
type ProposalRegistryENewProposalIterator struct {
	Event *ProposalRegistryENewProposal // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryENewProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryENewProposal)
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
		it.Event = new(ProposalRegistryENewProposal)
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
func (it *ProposalRegistryENewProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryENewProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryENewProposal represents a ENewProposal event raised by the ProposalRegistry contract.
type ProposalRegistryENewProposal struct {
	Dao        common.Address
	ProposalID [32]byte
	Metadata   [][]byte
	KvData     [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterENewProposal is a free log retrieval operation binding the contract event 0x418d36d9e4d07417178f14808c2225270a5e5996c3c97053e33d4f5cee1c1176.
//
// Solidity: event ENewProposal(address indexed dao, bytes32 indexed proposalID, bytes[] metadata, bytes[] kvData)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterENewProposal(opts *bind.FilterOpts, dao []common.Address, proposalID [][32]byte) (*ProposalRegistryENewProposalIterator, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "ENewProposal", daoRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryENewProposalIterator{contract: _ProposalRegistry.contract, event: "ENewProposal", logs: logs, sub: sub}, nil
}

// WatchENewProposal is a free log subscription operation binding the contract event 0x418d36d9e4d07417178f14808c2225270a5e5996c3c97053e33d4f5cee1c1176.
//
// Solidity: event ENewProposal(address indexed dao, bytes32 indexed proposalID, bytes[] metadata, bytes[] kvData)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchENewProposal(opts *bind.WatchOpts, sink chan<- *ProposalRegistryENewProposal, dao []common.Address, proposalID [][32]byte) (event.Subscription, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "ENewProposal", daoRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryENewProposal)
				if err := _ProposalRegistry.contract.UnpackLog(event, "ENewProposal", log); err != nil {
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

// ParseENewProposal is a log parse operation binding the contract event 0x418d36d9e4d07417178f14808c2225270a5e5996c3c97053e33d4f5cee1c1176.
//
// Solidity: event ENewProposal(address indexed dao, bytes32 indexed proposalID, bytes[] metadata, bytes[] kvData)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseENewProposal(log types.Log) (*ProposalRegistryENewProposal, error) {
	event := new(ProposalRegistryENewProposal)
	if err := _ProposalRegistry.contract.UnpackLog(event, "ENewProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryEProposalAppendIterator is returned from FilterEProposalAppend and is used to iterate over the raw logs and unpacked data for EProposalAppend events raised by the ProposalRegistry contract.
type ProposalRegistryEProposalAppendIterator struct {
	Event *ProposalRegistryEProposalAppend // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryEProposalAppendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryEProposalAppend)
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
		it.Event = new(ProposalRegistryEProposalAppend)
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
func (it *ProposalRegistryEProposalAppendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryEProposalAppendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryEProposalAppend represents a EProposalAppend event raised by the ProposalRegistry contract.
type ProposalRegistryEProposalAppend struct {
	Dao        common.Address
	ProposalID [32]byte
	KvData     [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEProposalAppend is a free log retrieval operation binding the contract event 0x36363313a6f56ef7a6adb6b774e101360859b98961d3a98680a3a2440c1dcbcf.
//
// Solidity: event EProposalAppend(address indexed dao, bytes32 indexed proposalID, bytes[] kvData)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterEProposalAppend(opts *bind.FilterOpts, dao []common.Address, proposalID [][32]byte) (*ProposalRegistryEProposalAppendIterator, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "EProposalAppend", daoRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryEProposalAppendIterator{contract: _ProposalRegistry.contract, event: "EProposalAppend", logs: logs, sub: sub}, nil
}

// WatchEProposalAppend is a free log subscription operation binding the contract event 0x36363313a6f56ef7a6adb6b774e101360859b98961d3a98680a3a2440c1dcbcf.
//
// Solidity: event EProposalAppend(address indexed dao, bytes32 indexed proposalID, bytes[] kvData)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchEProposalAppend(opts *bind.WatchOpts, sink chan<- *ProposalRegistryEProposalAppend, dao []common.Address, proposalID [][32]byte) (event.Subscription, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "EProposalAppend", daoRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryEProposalAppend)
				if err := _ProposalRegistry.contract.UnpackLog(event, "EProposalAppend", log); err != nil {
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

// ParseEProposalAppend is a log parse operation binding the contract event 0x36363313a6f56ef7a6adb6b774e101360859b98961d3a98680a3a2440c1dcbcf.
//
// Solidity: event EProposalAppend(address indexed dao, bytes32 indexed proposalID, bytes[] kvData)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseEProposalAppend(log types.Log) (*ProposalRegistryEProposalAppend, error) {
	event := new(ProposalRegistryEProposalAppend)
	if err := _ProposalRegistry.contract.UnpackLog(event, "EProposalAppend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryEProposalDecideIterator is returned from FilterEProposalDecide and is used to iterate over the raw logs and unpacked data for EProposalDecide events raised by the ProposalRegistry contract.
type ProposalRegistryEProposalDecideIterator struct {
	Event *ProposalRegistryEProposalDecide // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryEProposalDecideIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryEProposalDecide)
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
		it.Event = new(ProposalRegistryEProposalDecide)
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
func (it *ProposalRegistryEProposalDecideIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryEProposalDecideIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryEProposalDecide represents a EProposalDecide event raised by the ProposalRegistry contract.
type ProposalRegistryEProposalDecide struct {
	Dao        common.Address
	ProposalID [32]byte
	Agree      bool
	TopicID    [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEProposalDecide is a free log retrieval operation binding the contract event 0xf2a3129651ad2fdf3330fb1d1db190764bc759088b4c6eacf1fa9c32c4405016.
//
// Solidity: event EProposalDecide(address indexed dao, bytes32 indexed proposalID, bool indexed agree, bytes32 topicID)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterEProposalDecide(opts *bind.FilterOpts, dao []common.Address, proposalID [][32]byte, agree []bool) (*ProposalRegistryEProposalDecideIterator, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var agreeRule []interface{}
	for _, agreeItem := range agree {
		agreeRule = append(agreeRule, agreeItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "EProposalDecide", daoRule, proposalIDRule, agreeRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryEProposalDecideIterator{contract: _ProposalRegistry.contract, event: "EProposalDecide", logs: logs, sub: sub}, nil
}

// WatchEProposalDecide is a free log subscription operation binding the contract event 0xf2a3129651ad2fdf3330fb1d1db190764bc759088b4c6eacf1fa9c32c4405016.
//
// Solidity: event EProposalDecide(address indexed dao, bytes32 indexed proposalID, bool indexed agree, bytes32 topicID)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchEProposalDecide(opts *bind.WatchOpts, sink chan<- *ProposalRegistryEProposalDecide, dao []common.Address, proposalID [][32]byte, agree []bool) (event.Subscription, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var agreeRule []interface{}
	for _, agreeItem := range agree {
		agreeRule = append(agreeRule, agreeItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "EProposalDecide", daoRule, proposalIDRule, agreeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryEProposalDecide)
				if err := _ProposalRegistry.contract.UnpackLog(event, "EProposalDecide", log); err != nil {
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

// ParseEProposalDecide is a log parse operation binding the contract event 0xf2a3129651ad2fdf3330fb1d1db190764bc759088b4c6eacf1fa9c32c4405016.
//
// Solidity: event EProposalDecide(address indexed dao, bytes32 indexed proposalID, bool indexed agree, bytes32 topicID)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseEProposalDecide(log types.Log) (*ProposalRegistryEProposalDecide, error) {
	event := new(ProposalRegistryEProposalDecide)
	if err := _ProposalRegistry.contract.UnpackLog(event, "EProposalDecide", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryETopicCreateIterator is returned from FilterETopicCreate and is used to iterate over the raw logs and unpacked data for ETopicCreate events raised by the ProposalRegistry contract.
type ProposalRegistryETopicCreateIterator struct {
	Event *ProposalRegistryETopicCreate // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryETopicCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryETopicCreate)
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
		it.Event = new(ProposalRegistryETopicCreate)
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
func (it *ProposalRegistryETopicCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryETopicCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryETopicCreate represents a ETopicCreate event raised by the ProposalRegistry contract.
type ProposalRegistryETopicCreate struct {
	Dao        common.Address
	TopicID    [32]byte
	ProposalID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterETopicCreate is a free log retrieval operation binding the contract event 0xf7644cde7ea87265ebadfd31e56ab7041dd6ae5e34d5a39d943a8d255c31f944.
//
// Solidity: event ETopicCreate(address indexed dao, bytes32 indexed topicID, bytes32 indexed proposalID)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterETopicCreate(opts *bind.FilterOpts, dao []common.Address, topicID [][32]byte, proposalID [][32]byte) (*ProposalRegistryETopicCreateIterator, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var topicIDRule []interface{}
	for _, topicIDItem := range topicID {
		topicIDRule = append(topicIDRule, topicIDItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "ETopicCreate", daoRule, topicIDRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryETopicCreateIterator{contract: _ProposalRegistry.contract, event: "ETopicCreate", logs: logs, sub: sub}, nil
}

// WatchETopicCreate is a free log subscription operation binding the contract event 0xf7644cde7ea87265ebadfd31e56ab7041dd6ae5e34d5a39d943a8d255c31f944.
//
// Solidity: event ETopicCreate(address indexed dao, bytes32 indexed topicID, bytes32 indexed proposalID)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchETopicCreate(opts *bind.WatchOpts, sink chan<- *ProposalRegistryETopicCreate, dao []common.Address, topicID [][32]byte, proposalID [][32]byte) (event.Subscription, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var topicIDRule []interface{}
	for _, topicIDItem := range topicID {
		topicIDRule = append(topicIDRule, topicIDItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "ETopicCreate", daoRule, topicIDRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryETopicCreate)
				if err := _ProposalRegistry.contract.UnpackLog(event, "ETopicCreate", log); err != nil {
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

// ParseETopicCreate is a log parse operation binding the contract event 0xf7644cde7ea87265ebadfd31e56ab7041dd6ae5e34d5a39d943a8d255c31f944.
//
// Solidity: event ETopicCreate(address indexed dao, bytes32 indexed topicID, bytes32 indexed proposalID)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseETopicCreate(log types.Log) (*ProposalRegistryETopicCreate, error) {
	event := new(ProposalRegistryETopicCreate)
	if err := _ProposalRegistry.contract.UnpackLog(event, "ETopicCreate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalRegistryETopicFixIterator is returned from FilterETopicFix and is used to iterate over the raw logs and unpacked data for ETopicFix events raised by the ProposalRegistry contract.
type ProposalRegistryETopicFixIterator struct {
	Event *ProposalRegistryETopicFix // Event containing the contract specifics and raw log

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
func (it *ProposalRegistryETopicFixIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalRegistryETopicFix)
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
		it.Event = new(ProposalRegistryETopicFix)
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
func (it *ProposalRegistryETopicFixIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalRegistryETopicFixIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalRegistryETopicFix represents a ETopicFix event raised by the ProposalRegistry contract.
type ProposalRegistryETopicFix struct {
	Dao        common.Address
	TopicID    [32]byte
	ProposalID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterETopicFix is a free log retrieval operation binding the contract event 0xd34cc36cbd4fae71068412905d78e6f9e28c954a2ec0b9d8520c0fefca645404.
//
// Solidity: event ETopicFix(address indexed dao, bytes32 indexed topicID, bytes32 indexed proposalID)
func (_ProposalRegistry *ProposalRegistryFilterer) FilterETopicFix(opts *bind.FilterOpts, dao []common.Address, topicID [][32]byte, proposalID [][32]byte) (*ProposalRegistryETopicFixIterator, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var topicIDRule []interface{}
	for _, topicIDItem := range topicID {
		topicIDRule = append(topicIDRule, topicIDItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.FilterLogs(opts, "ETopicFix", daoRule, topicIDRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &ProposalRegistryETopicFixIterator{contract: _ProposalRegistry.contract, event: "ETopicFix", logs: logs, sub: sub}, nil
}

// WatchETopicFix is a free log subscription operation binding the contract event 0xd34cc36cbd4fae71068412905d78e6f9e28c954a2ec0b9d8520c0fefca645404.
//
// Solidity: event ETopicFix(address indexed dao, bytes32 indexed topicID, bytes32 indexed proposalID)
func (_ProposalRegistry *ProposalRegistryFilterer) WatchETopicFix(opts *bind.WatchOpts, sink chan<- *ProposalRegistryETopicFix, dao []common.Address, topicID [][32]byte, proposalID [][32]byte) (event.Subscription, error) {

	var daoRule []interface{}
	for _, daoItem := range dao {
		daoRule = append(daoRule, daoItem)
	}
	var topicIDRule []interface{}
	for _, topicIDItem := range topicID {
		topicIDRule = append(topicIDRule, topicIDItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _ProposalRegistry.contract.WatchLogs(opts, "ETopicFix", daoRule, topicIDRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalRegistryETopicFix)
				if err := _ProposalRegistry.contract.UnpackLog(event, "ETopicFix", log); err != nil {
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

// ParseETopicFix is a log parse operation binding the contract event 0xd34cc36cbd4fae71068412905d78e6f9e28c954a2ec0b9d8520c0fefca645404.
//
// Solidity: event ETopicFix(address indexed dao, bytes32 indexed topicID, bytes32 indexed proposalID)
func (_ProposalRegistry *ProposalRegistryFilterer) ParseETopicFix(log types.Log) (*ProposalRegistryETopicFix, error) {
	event := new(ProposalRegistryETopicFix)
	if err := _ProposalRegistry.contract.UnpackLog(event, "ETopicFix", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
