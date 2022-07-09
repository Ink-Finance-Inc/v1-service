// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MasterDAO

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

// DAOBaseInitData is an auto generated low-level Go binding around an user-defined struct.
type DAOBaseInitData struct {
	Name                      string
	Describe                  string
	Mds                       [][]byte
	GovTokenAddr              common.Address
	GovTokenAmountRequirement *big.Int
	StakingAddr               common.Address
	Flows                     []IDAOFlowInfo
	BadgeName                 string
	BadgeTotal                *big.Int
}

// IDAOFlowInfo is an auto generated low-level Go binding around an user-defined struct.
type IDAOFlowInfo struct {
	FlowID     [32]byte
	Committees []IDAOHandleProposalProposalCommitteeInfo
}

// IDAOHandleProposalDAOProcessInfo is an auto generated low-level Go binding around an user-defined struct.
type IDAOHandleProposalDAOProcessInfo struct {
	ProposalID             [32]byte
	ProcessCategory        [32]byte
	NextCommittee          IDAOHandleProposalProposalCommitteeInfo
	LastOperationTimestamp *big.Int
	Committees             []IDAOHandleProposalProposalCommitteeInfo
}

// IDAOHandleProposalProposalCommitteeInfo is an auto generated low-level Go binding around an user-defined struct.
type IDAOHandleProposalProposalCommitteeInfo struct {
	Step      [32]byte
	Committee common.Address
	Sensitive *big.Int
	Name      string
}

// IDAOUserFlowInfo is an auto generated low-level Go binding around an user-defined struct.
type IDAOUserFlowInfo struct {
	FlowID     [32]byte
	Committees []common.Address
}

// IProposalOperatorProposalApplyInfo is an auto generated low-level Go binding around an user-defined struct.
type IProposalOperatorProposalApplyInfo struct {
	Items   [][]byte
	Headers [][]byte
}

// MasterDAOMetaData contains all meta data concerning the MasterDAO contract.
var MasterDAOMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ucv\",\"type\":\"address\"}],\"name\":\"EAddUCV\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"EBadge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governanceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ECreateDAO\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"agree\",\"type\":\"bool\"}],\"name\":\"EDecideProposal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_STEP_NUM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NORMAL_CATEGORY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addrRegistry\",\"outputs\":[{\"internalType\":\"contractIAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"describe\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"mds\",\"type\":\"bytes[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"govTokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"govTokenAmountRequirement\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakingAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"flowID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"step\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sensitive\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structIDAOHandleProposal.ProposalCommitteeInfo[]\",\"name\":\"committees\",\"type\":\"tuple[]\"}],\"internalType\":\"structIDAO.FlowInfo[]\",\"name\":\"flows\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"badgeName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"badgeTotal\",\"type\":\"uint256\"}],\"internalType\":\"structDAO.BaseInitData\",\"name\":\"initData\",\"type\":\"tuple\"}],\"name\":\"buildInitData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"changeProposal\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"checkMemberPosition\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"committee\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"agree\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"decideProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callbackEvent\",\"type\":\"bytes\"}],\"name\":\"delegateCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"describe\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllCommittees\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"flowID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"step\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sensitive\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structIDAOHandleProposal.ProposalCommitteeInfo[]\",\"name\":\"committees\",\"type\":\"tuple[]\"}],\"internalType\":\"structIDAO.FlowInfo[]\",\"name\":\"flowInfos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBadge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"badge\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"getDAOProcessInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"processCategory\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"step\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sensitive\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structIDAOHandleProposal.ProposalCommitteeInfo\",\"name\":\"nextCommittee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastOperationTimestamp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"step\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sensitive\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structIDAOHandleProposal.ProposalCommitteeInfo[]\",\"name\":\"committees\",\"type\":\"tuple[]\"}],\"internalType\":\"structIDAOHandleProposal.DAOProcessInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"getLastOperationTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processCategory\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMemberCommittees\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"committees\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getMetadata\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"typeID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"getNextCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"step\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sensitive\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structIDAOHandleProposal.ProposalCommitteeInfo\",\"name\":\"nextInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"stepIdx\",\"type\":\"uint256\"}],\"name\":\"getProposalStep\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"step\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sensitive\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structIDAOHandleProposal.ProposalCommitteeInfo\",\"name\":\"stepInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"staking\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUCVs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"ucvs\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserCommitteeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"flowID\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"committees\",\"type\":\"address[]\"}],\"internalType\":\"structIDAO.UserFlowInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govTokenAmountRequirement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrReg\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addrRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"callbackEvent\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"items\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"headers\",\"type\":\"bytes[]\"}],\"internalType\":\"structIProposalOperator.ProposalApplyInfo\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"newProposal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processCategory\",\"type\":\"bytes32\"}],\"name\":\"processCategoryFlow\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"step\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sensitive\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structIDAOHandleProposal.ProposalCommitteeInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"proposalCanExecute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reason\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"flowID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"step\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sensitive\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structIDAOHandleProposal.ProposalCommitteeInfo[]\",\"name\":\"committees\",\"type\":\"tuple[]\"}],\"internalType\":\"structIDAO.FlowInfo\",\"name\":\"flow\",\"type\":\"tuple\"}],\"name\":\"setFlowStep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ucv\",\"type\":\"address\"}],\"name\":\"setUCV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processCategory\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"step\",\"type\":\"bytes32\"}],\"name\":\"stepCommittee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MasterDAOABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterDAOMetaData.ABI instead.
var MasterDAOABI = MasterDAOMetaData.ABI

// MasterDAO is an auto generated Go binding around an Ethereum contract.
type MasterDAO struct {
	MasterDAOCaller     // Read-only binding to the contract
	MasterDAOTransactor // Write-only binding to the contract
	MasterDAOFilterer   // Log filterer for contract events
}

// MasterDAOCaller is an auto generated read-only Go binding around an Ethereum contract.
type MasterDAOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterDAOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterDAOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterDAOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterDAOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterDAOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterDAOSession struct {
	Contract     *MasterDAO        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterDAOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterDAOCallerSession struct {
	Contract *MasterDAOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MasterDAOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterDAOTransactorSession struct {
	Contract     *MasterDAOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MasterDAORaw is an auto generated low-level Go binding around an Ethereum contract.
type MasterDAORaw struct {
	Contract *MasterDAO // Generic contract binding to access the raw methods on
}

// MasterDAOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterDAOCallerRaw struct {
	Contract *MasterDAOCaller // Generic read-only contract binding to access the raw methods on
}

// MasterDAOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterDAOTransactorRaw struct {
	Contract *MasterDAOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterDAO creates a new instance of MasterDAO, bound to a specific deployed contract.
func NewMasterDAO(address common.Address, backend bind.ContractBackend) (*MasterDAO, error) {
	contract, err := bindMasterDAO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterDAO{MasterDAOCaller: MasterDAOCaller{contract: contract}, MasterDAOTransactor: MasterDAOTransactor{contract: contract}, MasterDAOFilterer: MasterDAOFilterer{contract: contract}}, nil
}

// NewMasterDAOCaller creates a new read-only instance of MasterDAO, bound to a specific deployed contract.
func NewMasterDAOCaller(address common.Address, caller bind.ContractCaller) (*MasterDAOCaller, error) {
	contract, err := bindMasterDAO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterDAOCaller{contract: contract}, nil
}

// NewMasterDAOTransactor creates a new write-only instance of MasterDAO, bound to a specific deployed contract.
func NewMasterDAOTransactor(address common.Address, transactor bind.ContractTransactor) (*MasterDAOTransactor, error) {
	contract, err := bindMasterDAO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterDAOTransactor{contract: contract}, nil
}

// NewMasterDAOFilterer creates a new log filterer instance of MasterDAO, bound to a specific deployed contract.
func NewMasterDAOFilterer(address common.Address, filterer bind.ContractFilterer) (*MasterDAOFilterer, error) {
	contract, err := bindMasterDAO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterDAOFilterer{contract: contract}, nil
}

// bindMasterDAO binds a generic wrapper to an already deployed contract.
func bindMasterDAO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasterDAOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterDAO *MasterDAORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterDAO.Contract.MasterDAOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterDAO *MasterDAORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterDAO.Contract.MasterDAOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterDAO *MasterDAORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterDAO.Contract.MasterDAOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterDAO *MasterDAOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterDAO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterDAO *MasterDAOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterDAO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterDAO *MasterDAOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterDAO.Contract.contract.Transact(opts, method, params...)
}

// MAXSTEPNUM is a free data retrieval call binding the contract method 0xa3f1a470.
//
// Solidity: function MAX_STEP_NUM() view returns(uint256)
func (_MasterDAO *MasterDAOCaller) MAXSTEPNUM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "MAX_STEP_NUM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTEPNUM is a free data retrieval call binding the contract method 0xa3f1a470.
//
// Solidity: function MAX_STEP_NUM() view returns(uint256)
func (_MasterDAO *MasterDAOSession) MAXSTEPNUM() (*big.Int, error) {
	return _MasterDAO.Contract.MAXSTEPNUM(&_MasterDAO.CallOpts)
}

// MAXSTEPNUM is a free data retrieval call binding the contract method 0xa3f1a470.
//
// Solidity: function MAX_STEP_NUM() view returns(uint256)
func (_MasterDAO *MasterDAOCallerSession) MAXSTEPNUM() (*big.Int, error) {
	return _MasterDAO.Contract.MAXSTEPNUM(&_MasterDAO.CallOpts)
}

// NORMALCATEGORY is a free data retrieval call binding the contract method 0x3ab1e93b.
//
// Solidity: function NORMAL_CATEGORY() view returns(bytes32)
func (_MasterDAO *MasterDAOCaller) NORMALCATEGORY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "NORMAL_CATEGORY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NORMALCATEGORY is a free data retrieval call binding the contract method 0x3ab1e93b.
//
// Solidity: function NORMAL_CATEGORY() view returns(bytes32)
func (_MasterDAO *MasterDAOSession) NORMALCATEGORY() ([32]byte, error) {
	return _MasterDAO.Contract.NORMALCATEGORY(&_MasterDAO.CallOpts)
}

// NORMALCATEGORY is a free data retrieval call binding the contract method 0x3ab1e93b.
//
// Solidity: function NORMAL_CATEGORY() view returns(bytes32)
func (_MasterDAO *MasterDAOCallerSession) NORMALCATEGORY() ([32]byte, error) {
	return _MasterDAO.Contract.NORMALCATEGORY(&_MasterDAO.CallOpts)
}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_MasterDAO *MasterDAOCaller) AddrRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "addrRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_MasterDAO *MasterDAOSession) AddrRegistry() (common.Address, error) {
	return _MasterDAO.Contract.AddrRegistry(&_MasterDAO.CallOpts)
}

// AddrRegistry is a free data retrieval call binding the contract method 0x82d7c176.
//
// Solidity: function addrRegistry() view returns(address)
func (_MasterDAO *MasterDAOCallerSession) AddrRegistry() (common.Address, error) {
	return _MasterDAO.Contract.AddrRegistry(&_MasterDAO.CallOpts)
}

// BuildInitData is a free data retrieval call binding the contract method 0xfbc47fa4.
//
// Solidity: function buildInitData((string,string,bytes[],address,uint256,address,(bytes32,(bytes32,address,uint256,string)[])[],string,uint256) initData) pure returns(bytes data)
func (_MasterDAO *MasterDAOCaller) BuildInitData(opts *bind.CallOpts, initData DAOBaseInitData) ([]byte, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "buildInitData", initData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BuildInitData is a free data retrieval call binding the contract method 0xfbc47fa4.
//
// Solidity: function buildInitData((string,string,bytes[],address,uint256,address,(bytes32,(bytes32,address,uint256,string)[])[],string,uint256) initData) pure returns(bytes data)
func (_MasterDAO *MasterDAOSession) BuildInitData(initData DAOBaseInitData) ([]byte, error) {
	return _MasterDAO.Contract.BuildInitData(&_MasterDAO.CallOpts, initData)
}

// BuildInitData is a free data retrieval call binding the contract method 0xfbc47fa4.
//
// Solidity: function buildInitData((string,string,bytes[],address,uint256,address,(bytes32,(bytes32,address,uint256,string)[])[],string,uint256) initData) pure returns(bytes data)
func (_MasterDAO *MasterDAOCallerSession) BuildInitData(initData DAOBaseInitData) ([]byte, error) {
	return _MasterDAO.Contract.BuildInitData(&_MasterDAO.CallOpts, initData)
}

// ChangeProposal is a free data retrieval call binding the contract method 0xf406f799.
//
// Solidity: function changeProposal(bytes32 , bytes[] , bytes ) view returns()
func (_MasterDAO *MasterDAOCaller) ChangeProposal(opts *bind.CallOpts, arg0 [32]byte, arg1 [][]byte, arg2 []byte) error {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "changeProposal", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// ChangeProposal is a free data retrieval call binding the contract method 0xf406f799.
//
// Solidity: function changeProposal(bytes32 , bytes[] , bytes ) view returns()
func (_MasterDAO *MasterDAOSession) ChangeProposal(arg0 [32]byte, arg1 [][]byte, arg2 []byte) error {
	return _MasterDAO.Contract.ChangeProposal(&_MasterDAO.CallOpts, arg0, arg1, arg2)
}

// ChangeProposal is a free data retrieval call binding the contract method 0xf406f799.
//
// Solidity: function changeProposal(bytes32 , bytes[] , bytes ) view returns()
func (_MasterDAO *MasterDAOCallerSession) ChangeProposal(arg0 [32]byte, arg1 [][]byte, arg2 []byte) error {
	return _MasterDAO.Contract.ChangeProposal(&_MasterDAO.CallOpts, arg0, arg1, arg2)
}

// CheckMemberPosition is a free data retrieval call binding the contract method 0xdf4901a9.
//
// Solidity: function checkMemberPosition(bytes32 proposalID, address user) view returns(address committee)
func (_MasterDAO *MasterDAOCaller) CheckMemberPosition(opts *bind.CallOpts, proposalID [32]byte, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "checkMemberPosition", proposalID, user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckMemberPosition is a free data retrieval call binding the contract method 0xdf4901a9.
//
// Solidity: function checkMemberPosition(bytes32 proposalID, address user) view returns(address committee)
func (_MasterDAO *MasterDAOSession) CheckMemberPosition(proposalID [32]byte, user common.Address) (common.Address, error) {
	return _MasterDAO.Contract.CheckMemberPosition(&_MasterDAO.CallOpts, proposalID, user)
}

// CheckMemberPosition is a free data retrieval call binding the contract method 0xdf4901a9.
//
// Solidity: function checkMemberPosition(bytes32 proposalID, address user) view returns(address committee)
func (_MasterDAO *MasterDAOCallerSession) CheckMemberPosition(proposalID [32]byte, user common.Address) (common.Address, error) {
	return _MasterDAO.Contract.CheckMemberPosition(&_MasterDAO.CallOpts, proposalID, user)
}

// Describe is a free data retrieval call binding the contract method 0x1ad184ff.
//
// Solidity: function describe() view returns(string)
func (_MasterDAO *MasterDAOCaller) Describe(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "describe")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Describe is a free data retrieval call binding the contract method 0x1ad184ff.
//
// Solidity: function describe() view returns(string)
func (_MasterDAO *MasterDAOSession) Describe() (string, error) {
	return _MasterDAO.Contract.Describe(&_MasterDAO.CallOpts)
}

// Describe is a free data retrieval call binding the contract method 0x1ad184ff.
//
// Solidity: function describe() view returns(string)
func (_MasterDAO *MasterDAOCallerSession) Describe() (string, error) {
	return _MasterDAO.Contract.Describe(&_MasterDAO.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address admin)
func (_MasterDAO *MasterDAOCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address admin)
func (_MasterDAO *MasterDAOSession) GetAdmin() (common.Address, error) {
	return _MasterDAO.Contract.GetAdmin(&_MasterDAO.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address admin)
func (_MasterDAO *MasterDAOCallerSession) GetAdmin() (common.Address, error) {
	return _MasterDAO.Contract.GetAdmin(&_MasterDAO.CallOpts)
}

// GetAllCommittees is a free data retrieval call binding the contract method 0x7bff2ca6.
//
// Solidity: function getAllCommittees() view returns((bytes32,(bytes32,address,uint256,string)[])[] flowInfos)
func (_MasterDAO *MasterDAOCaller) GetAllCommittees(opts *bind.CallOpts) ([]IDAOFlowInfo, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getAllCommittees")

	if err != nil {
		return *new([]IDAOFlowInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDAOFlowInfo)).(*[]IDAOFlowInfo)

	return out0, err

}

// GetAllCommittees is a free data retrieval call binding the contract method 0x7bff2ca6.
//
// Solidity: function getAllCommittees() view returns((bytes32,(bytes32,address,uint256,string)[])[] flowInfos)
func (_MasterDAO *MasterDAOSession) GetAllCommittees() ([]IDAOFlowInfo, error) {
	return _MasterDAO.Contract.GetAllCommittees(&_MasterDAO.CallOpts)
}

// GetAllCommittees is a free data retrieval call binding the contract method 0x7bff2ca6.
//
// Solidity: function getAllCommittees() view returns((bytes32,(bytes32,address,uint256,string)[])[] flowInfos)
func (_MasterDAO *MasterDAOCallerSession) GetAllCommittees() ([]IDAOFlowInfo, error) {
	return _MasterDAO.Contract.GetAllCommittees(&_MasterDAO.CallOpts)
}

// GetBadge is a free data retrieval call binding the contract method 0xfb05cffb.
//
// Solidity: function getBadge() view returns(address badge)
func (_MasterDAO *MasterDAOCaller) GetBadge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getBadge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBadge is a free data retrieval call binding the contract method 0xfb05cffb.
//
// Solidity: function getBadge() view returns(address badge)
func (_MasterDAO *MasterDAOSession) GetBadge() (common.Address, error) {
	return _MasterDAO.Contract.GetBadge(&_MasterDAO.CallOpts)
}

// GetBadge is a free data retrieval call binding the contract method 0xfb05cffb.
//
// Solidity: function getBadge() view returns(address badge)
func (_MasterDAO *MasterDAOCallerSession) GetBadge() (common.Address, error) {
	return _MasterDAO.Contract.GetBadge(&_MasterDAO.CallOpts)
}

// GetDAOProcessInfo is a free data retrieval call binding the contract method 0x85536fd2.
//
// Solidity: function getDAOProcessInfo(bytes32 proposalID) view returns((bytes32,bytes32,(bytes32,address,uint256,string),uint256,(bytes32,address,uint256,string)[]) info)
func (_MasterDAO *MasterDAOCaller) GetDAOProcessInfo(opts *bind.CallOpts, proposalID [32]byte) (IDAOHandleProposalDAOProcessInfo, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getDAOProcessInfo", proposalID)

	if err != nil {
		return *new(IDAOHandleProposalDAOProcessInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IDAOHandleProposalDAOProcessInfo)).(*IDAOHandleProposalDAOProcessInfo)

	return out0, err

}

// GetDAOProcessInfo is a free data retrieval call binding the contract method 0x85536fd2.
//
// Solidity: function getDAOProcessInfo(bytes32 proposalID) view returns((bytes32,bytes32,(bytes32,address,uint256,string),uint256,(bytes32,address,uint256,string)[]) info)
func (_MasterDAO *MasterDAOSession) GetDAOProcessInfo(proposalID [32]byte) (IDAOHandleProposalDAOProcessInfo, error) {
	return _MasterDAO.Contract.GetDAOProcessInfo(&_MasterDAO.CallOpts, proposalID)
}

// GetDAOProcessInfo is a free data retrieval call binding the contract method 0x85536fd2.
//
// Solidity: function getDAOProcessInfo(bytes32 proposalID) view returns((bytes32,bytes32,(bytes32,address,uint256,string),uint256,(bytes32,address,uint256,string)[]) info)
func (_MasterDAO *MasterDAOCallerSession) GetDAOProcessInfo(proposalID [32]byte) (IDAOHandleProposalDAOProcessInfo, error) {
	return _MasterDAO.Contract.GetDAOProcessInfo(&_MasterDAO.CallOpts, proposalID)
}

// GetLastOperationTimestamp is a free data retrieval call binding the contract method 0xb7b77f7a.
//
// Solidity: function getLastOperationTimestamp(bytes32 proposalID) view returns(uint256 timestamp)
func (_MasterDAO *MasterDAOCaller) GetLastOperationTimestamp(opts *bind.CallOpts, proposalID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getLastOperationTimestamp", proposalID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastOperationTimestamp is a free data retrieval call binding the contract method 0xb7b77f7a.
//
// Solidity: function getLastOperationTimestamp(bytes32 proposalID) view returns(uint256 timestamp)
func (_MasterDAO *MasterDAOSession) GetLastOperationTimestamp(proposalID [32]byte) (*big.Int, error) {
	return _MasterDAO.Contract.GetLastOperationTimestamp(&_MasterDAO.CallOpts, proposalID)
}

// GetLastOperationTimestamp is a free data retrieval call binding the contract method 0xb7b77f7a.
//
// Solidity: function getLastOperationTimestamp(bytes32 proposalID) view returns(uint256 timestamp)
func (_MasterDAO *MasterDAOCallerSession) GetLastOperationTimestamp(proposalID [32]byte) (*big.Int, error) {
	return _MasterDAO.Contract.GetLastOperationTimestamp(&_MasterDAO.CallOpts, proposalID)
}

// GetMemberCommittees is a free data retrieval call binding the contract method 0x6c64d506.
//
// Solidity: function getMemberCommittees(bytes32 processCategory, address addr) view returns(address[] committees)
func (_MasterDAO *MasterDAOCaller) GetMemberCommittees(opts *bind.CallOpts, processCategory [32]byte, addr common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getMemberCommittees", processCategory, addr)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMemberCommittees is a free data retrieval call binding the contract method 0x6c64d506.
//
// Solidity: function getMemberCommittees(bytes32 processCategory, address addr) view returns(address[] committees)
func (_MasterDAO *MasterDAOSession) GetMemberCommittees(processCategory [32]byte, addr common.Address) ([]common.Address, error) {
	return _MasterDAO.Contract.GetMemberCommittees(&_MasterDAO.CallOpts, processCategory, addr)
}

// GetMemberCommittees is a free data retrieval call binding the contract method 0x6c64d506.
//
// Solidity: function getMemberCommittees(bytes32 processCategory, address addr) view returns(address[] committees)
func (_MasterDAO *MasterDAOCallerSession) GetMemberCommittees(processCategory [32]byte, addr common.Address) ([]common.Address, error) {
	return _MasterDAO.Contract.GetMemberCommittees(&_MasterDAO.CallOpts, processCategory, addr)
}

// GetMetadata is a free data retrieval call binding the contract method 0xd11a1f07.
//
// Solidity: function getMetadata(string key) view returns(bytes32 typeID, bytes data)
func (_MasterDAO *MasterDAOCaller) GetMetadata(opts *bind.CallOpts, key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getMetadata", key)

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

// GetMetadata is a free data retrieval call binding the contract method 0xd11a1f07.
//
// Solidity: function getMetadata(string key) view returns(bytes32 typeID, bytes data)
func (_MasterDAO *MasterDAOSession) GetMetadata(key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	return _MasterDAO.Contract.GetMetadata(&_MasterDAO.CallOpts, key)
}

// GetMetadata is a free data retrieval call binding the contract method 0xd11a1f07.
//
// Solidity: function getMetadata(string key) view returns(bytes32 typeID, bytes data)
func (_MasterDAO *MasterDAOCallerSession) GetMetadata(key string) (struct {
	TypeID [32]byte
	Data   []byte
}, error) {
	return _MasterDAO.Contract.GetMetadata(&_MasterDAO.CallOpts, key)
}

// GetNextCommittee is a free data retrieval call binding the contract method 0xb098e341.
//
// Solidity: function getNextCommittee(bytes32 proposalID) view returns((bytes32,address,uint256,string) nextInfo)
func (_MasterDAO *MasterDAOCaller) GetNextCommittee(opts *bind.CallOpts, proposalID [32]byte) (IDAOHandleProposalProposalCommitteeInfo, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getNextCommittee", proposalID)

	if err != nil {
		return *new(IDAOHandleProposalProposalCommitteeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IDAOHandleProposalProposalCommitteeInfo)).(*IDAOHandleProposalProposalCommitteeInfo)

	return out0, err

}

// GetNextCommittee is a free data retrieval call binding the contract method 0xb098e341.
//
// Solidity: function getNextCommittee(bytes32 proposalID) view returns((bytes32,address,uint256,string) nextInfo)
func (_MasterDAO *MasterDAOSession) GetNextCommittee(proposalID [32]byte) (IDAOHandleProposalProposalCommitteeInfo, error) {
	return _MasterDAO.Contract.GetNextCommittee(&_MasterDAO.CallOpts, proposalID)
}

// GetNextCommittee is a free data retrieval call binding the contract method 0xb098e341.
//
// Solidity: function getNextCommittee(bytes32 proposalID) view returns((bytes32,address,uint256,string) nextInfo)
func (_MasterDAO *MasterDAOCallerSession) GetNextCommittee(proposalID [32]byte) (IDAOHandleProposalProposalCommitteeInfo, error) {
	return _MasterDAO.Contract.GetNextCommittee(&_MasterDAO.CallOpts, proposalID)
}

// GetProposalStep is a free data retrieval call binding the contract method 0x96adb52d.
//
// Solidity: function getProposalStep(bytes32 proposalID, uint256 stepIdx) view returns((bytes32,address,uint256,string) stepInfo)
func (_MasterDAO *MasterDAOCaller) GetProposalStep(opts *bind.CallOpts, proposalID [32]byte, stepIdx *big.Int) (IDAOHandleProposalProposalCommitteeInfo, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getProposalStep", proposalID, stepIdx)

	if err != nil {
		return *new(IDAOHandleProposalProposalCommitteeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IDAOHandleProposalProposalCommitteeInfo)).(*IDAOHandleProposalProposalCommitteeInfo)

	return out0, err

}

// GetProposalStep is a free data retrieval call binding the contract method 0x96adb52d.
//
// Solidity: function getProposalStep(bytes32 proposalID, uint256 stepIdx) view returns((bytes32,address,uint256,string) stepInfo)
func (_MasterDAO *MasterDAOSession) GetProposalStep(proposalID [32]byte, stepIdx *big.Int) (IDAOHandleProposalProposalCommitteeInfo, error) {
	return _MasterDAO.Contract.GetProposalStep(&_MasterDAO.CallOpts, proposalID, stepIdx)
}

// GetProposalStep is a free data retrieval call binding the contract method 0x96adb52d.
//
// Solidity: function getProposalStep(bytes32 proposalID, uint256 stepIdx) view returns((bytes32,address,uint256,string) stepInfo)
func (_MasterDAO *MasterDAOCallerSession) GetProposalStep(proposalID [32]byte, stepIdx *big.Int) (IDAOHandleProposalProposalCommitteeInfo, error) {
	return _MasterDAO.Contract.GetProposalStep(&_MasterDAO.CallOpts, proposalID, stepIdx)
}

// GetStakingAddr is a free data retrieval call binding the contract method 0x37b1d6bc.
//
// Solidity: function getStakingAddr() view returns(address staking)
func (_MasterDAO *MasterDAOCaller) GetStakingAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getStakingAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingAddr is a free data retrieval call binding the contract method 0x37b1d6bc.
//
// Solidity: function getStakingAddr() view returns(address staking)
func (_MasterDAO *MasterDAOSession) GetStakingAddr() (common.Address, error) {
	return _MasterDAO.Contract.GetStakingAddr(&_MasterDAO.CallOpts)
}

// GetStakingAddr is a free data retrieval call binding the contract method 0x37b1d6bc.
//
// Solidity: function getStakingAddr() view returns(address staking)
func (_MasterDAO *MasterDAOCallerSession) GetStakingAddr() (common.Address, error) {
	return _MasterDAO.Contract.GetStakingAddr(&_MasterDAO.CallOpts)
}

// GetUCVs is a free data retrieval call binding the contract method 0x4e2782f3.
//
// Solidity: function getUCVs() view returns(address[] ucvs)
func (_MasterDAO *MasterDAOCaller) GetUCVs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getUCVs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUCVs is a free data retrieval call binding the contract method 0x4e2782f3.
//
// Solidity: function getUCVs() view returns(address[] ucvs)
func (_MasterDAO *MasterDAOSession) GetUCVs() ([]common.Address, error) {
	return _MasterDAO.Contract.GetUCVs(&_MasterDAO.CallOpts)
}

// GetUCVs is a free data retrieval call binding the contract method 0x4e2782f3.
//
// Solidity: function getUCVs() view returns(address[] ucvs)
func (_MasterDAO *MasterDAOCallerSession) GetUCVs() ([]common.Address, error) {
	return _MasterDAO.Contract.GetUCVs(&_MasterDAO.CallOpts)
}

// GetUserCommitteeInfo is a free data retrieval call binding the contract method 0xca12b11b.
//
// Solidity: function getUserCommitteeInfo(address user) view returns((bytes32,address[])[] infos)
func (_MasterDAO *MasterDAOCaller) GetUserCommitteeInfo(opts *bind.CallOpts, user common.Address) ([]IDAOUserFlowInfo, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "getUserCommitteeInfo", user)

	if err != nil {
		return *new([]IDAOUserFlowInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDAOUserFlowInfo)).(*[]IDAOUserFlowInfo)

	return out0, err

}

// GetUserCommitteeInfo is a free data retrieval call binding the contract method 0xca12b11b.
//
// Solidity: function getUserCommitteeInfo(address user) view returns((bytes32,address[])[] infos)
func (_MasterDAO *MasterDAOSession) GetUserCommitteeInfo(user common.Address) ([]IDAOUserFlowInfo, error) {
	return _MasterDAO.Contract.GetUserCommitteeInfo(&_MasterDAO.CallOpts, user)
}

// GetUserCommitteeInfo is a free data retrieval call binding the contract method 0xca12b11b.
//
// Solidity: function getUserCommitteeInfo(address user) view returns((bytes32,address[])[] infos)
func (_MasterDAO *MasterDAOCallerSession) GetUserCommitteeInfo(user common.Address) ([]IDAOUserFlowInfo, error) {
	return _MasterDAO.Contract.GetUserCommitteeInfo(&_MasterDAO.CallOpts, user)
}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_MasterDAO *MasterDAOCaller) GovToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "govToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_MasterDAO *MasterDAOSession) GovToken() (common.Address, error) {
	return _MasterDAO.Contract.GovToken(&_MasterDAO.CallOpts)
}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_MasterDAO *MasterDAOCallerSession) GovToken() (common.Address, error) {
	return _MasterDAO.Contract.GovToken(&_MasterDAO.CallOpts)
}

// GovTokenAmountRequirement is a free data retrieval call binding the contract method 0x7dcaecab.
//
// Solidity: function govTokenAmountRequirement() view returns(uint256)
func (_MasterDAO *MasterDAOCaller) GovTokenAmountRequirement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "govTokenAmountRequirement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GovTokenAmountRequirement is a free data retrieval call binding the contract method 0x7dcaecab.
//
// Solidity: function govTokenAmountRequirement() view returns(uint256)
func (_MasterDAO *MasterDAOSession) GovTokenAmountRequirement() (*big.Int, error) {
	return _MasterDAO.Contract.GovTokenAmountRequirement(&_MasterDAO.CallOpts)
}

// GovTokenAmountRequirement is a free data retrieval call binding the contract method 0x7dcaecab.
//
// Solidity: function govTokenAmountRequirement() view returns(uint256)
func (_MasterDAO *MasterDAOCallerSession) GovTokenAmountRequirement() (*big.Int, error) {
	return _MasterDAO.Contract.GovTokenAmountRequirement(&_MasterDAO.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MasterDAO *MasterDAOCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MasterDAO *MasterDAOSession) Name() (string, error) {
	return _MasterDAO.Contract.Name(&_MasterDAO.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MasterDAO *MasterDAOCallerSession) Name() (string, error) {
	return _MasterDAO.Contract.Name(&_MasterDAO.CallOpts)
}

// ProcessCategoryFlow is a free data retrieval call binding the contract method 0x26e193c4.
//
// Solidity: function processCategoryFlow(bytes32 processCategory) view returns((bytes32,address,uint256,string)[] infos)
func (_MasterDAO *MasterDAOCaller) ProcessCategoryFlow(opts *bind.CallOpts, processCategory [32]byte) ([]IDAOHandleProposalProposalCommitteeInfo, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "processCategoryFlow", processCategory)

	if err != nil {
		return *new([]IDAOHandleProposalProposalCommitteeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDAOHandleProposalProposalCommitteeInfo)).(*[]IDAOHandleProposalProposalCommitteeInfo)

	return out0, err

}

// ProcessCategoryFlow is a free data retrieval call binding the contract method 0x26e193c4.
//
// Solidity: function processCategoryFlow(bytes32 processCategory) view returns((bytes32,address,uint256,string)[] infos)
func (_MasterDAO *MasterDAOSession) ProcessCategoryFlow(processCategory [32]byte) ([]IDAOHandleProposalProposalCommitteeInfo, error) {
	return _MasterDAO.Contract.ProcessCategoryFlow(&_MasterDAO.CallOpts, processCategory)
}

// ProcessCategoryFlow is a free data retrieval call binding the contract method 0x26e193c4.
//
// Solidity: function processCategoryFlow(bytes32 processCategory) view returns((bytes32,address,uint256,string)[] infos)
func (_MasterDAO *MasterDAOCallerSession) ProcessCategoryFlow(processCategory [32]byte) ([]IDAOHandleProposalProposalCommitteeInfo, error) {
	return _MasterDAO.Contract.ProcessCategoryFlow(&_MasterDAO.CallOpts, processCategory)
}

// ProposalCanExecute is a free data retrieval call binding the contract method 0xca292f82.
//
// Solidity: function proposalCanExecute(bytes32 proposalID) view returns(uint256 reason)
func (_MasterDAO *MasterDAOCaller) ProposalCanExecute(opts *bind.CallOpts, proposalID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "proposalCanExecute", proposalID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCanExecute is a free data retrieval call binding the contract method 0xca292f82.
//
// Solidity: function proposalCanExecute(bytes32 proposalID) view returns(uint256 reason)
func (_MasterDAO *MasterDAOSession) ProposalCanExecute(proposalID [32]byte) (*big.Int, error) {
	return _MasterDAO.Contract.ProposalCanExecute(&_MasterDAO.CallOpts, proposalID)
}

// ProposalCanExecute is a free data retrieval call binding the contract method 0xca292f82.
//
// Solidity: function proposalCanExecute(bytes32 proposalID) view returns(uint256 reason)
func (_MasterDAO *MasterDAOCallerSession) ProposalCanExecute(proposalID [32]byte) (*big.Int, error) {
	return _MasterDAO.Contract.ProposalCanExecute(&_MasterDAO.CallOpts, proposalID)
}

// StepCommittee is a free data retrieval call binding the contract method 0xd55e8b04.
//
// Solidity: function stepCommittee(bytes32 processCategory, bytes32 step) view returns(address)
func (_MasterDAO *MasterDAOCaller) StepCommittee(opts *bind.CallOpts, processCategory [32]byte, step [32]byte) (common.Address, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "stepCommittee", processCategory, step)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StepCommittee is a free data retrieval call binding the contract method 0xd55e8b04.
//
// Solidity: function stepCommittee(bytes32 processCategory, bytes32 step) view returns(address)
func (_MasterDAO *MasterDAOSession) StepCommittee(processCategory [32]byte, step [32]byte) (common.Address, error) {
	return _MasterDAO.Contract.StepCommittee(&_MasterDAO.CallOpts, processCategory, step)
}

// StepCommittee is a free data retrieval call binding the contract method 0xd55e8b04.
//
// Solidity: function stepCommittee(bytes32 processCategory, bytes32 step) view returns(address)
func (_MasterDAO *MasterDAOCallerSession) StepCommittee(processCategory [32]byte, step [32]byte) (common.Address, error) {
	return _MasterDAO.Contract.StepCommittee(&_MasterDAO.CallOpts, processCategory, step)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MasterDAO *MasterDAOCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MasterDAO.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MasterDAO *MasterDAOSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MasterDAO.Contract.SupportsInterface(&_MasterDAO.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MasterDAO *MasterDAOCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MasterDAO.Contract.SupportsInterface(&_MasterDAO.CallOpts, interfaceId)
}

// DecideProposal is a paid mutator transaction binding the contract method 0x67024e77.
//
// Solidity: function decideProposal(bytes32 proposalID, bool agree, bytes ) returns()
func (_MasterDAO *MasterDAOTransactor) DecideProposal(opts *bind.TransactOpts, proposalID [32]byte, agree bool, arg2 []byte) (*types.Transaction, error) {
	return _MasterDAO.contract.Transact(opts, "decideProposal", proposalID, agree, arg2)
}

// DecideProposal is a paid mutator transaction binding the contract method 0x67024e77.
//
// Solidity: function decideProposal(bytes32 proposalID, bool agree, bytes ) returns()
func (_MasterDAO *MasterDAOSession) DecideProposal(proposalID [32]byte, agree bool, arg2 []byte) (*types.Transaction, error) {
	return _MasterDAO.Contract.DecideProposal(&_MasterDAO.TransactOpts, proposalID, agree, arg2)
}

// DecideProposal is a paid mutator transaction binding the contract method 0x67024e77.
//
// Solidity: function decideProposal(bytes32 proposalID, bool agree, bytes ) returns()
func (_MasterDAO *MasterDAOTransactorSession) DecideProposal(proposalID [32]byte, agree bool, arg2 []byte) (*types.Transaction, error) {
	return _MasterDAO.Contract.DecideProposal(&_MasterDAO.TransactOpts, proposalID, agree, arg2)
}

// DelegateCallback is a paid mutator transaction binding the contract method 0xf96baba1.
//
// Solidity: function delegateCallback(bytes callbackEvent) returns()
func (_MasterDAO *MasterDAOTransactor) DelegateCallback(opts *bind.TransactOpts, callbackEvent []byte) (*types.Transaction, error) {
	return _MasterDAO.contract.Transact(opts, "delegateCallback", callbackEvent)
}

// DelegateCallback is a paid mutator transaction binding the contract method 0xf96baba1.
//
// Solidity: function delegateCallback(bytes callbackEvent) returns()
func (_MasterDAO *MasterDAOSession) DelegateCallback(callbackEvent []byte) (*types.Transaction, error) {
	return _MasterDAO.Contract.DelegateCallback(&_MasterDAO.TransactOpts, callbackEvent)
}

// DelegateCallback is a paid mutator transaction binding the contract method 0xf96baba1.
//
// Solidity: function delegateCallback(bytes callbackEvent) returns()
func (_MasterDAO *MasterDAOTransactorSession) DelegateCallback(callbackEvent []byte) (*types.Transaction, error) {
	return _MasterDAO.Contract.DelegateCallback(&_MasterDAO.TransactOpts, callbackEvent)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x980ff6c6.
//
// Solidity: function executeProposal(bytes32 proposalID) returns()
func (_MasterDAO *MasterDAOTransactor) ExecuteProposal(opts *bind.TransactOpts, proposalID [32]byte) (*types.Transaction, error) {
	return _MasterDAO.contract.Transact(opts, "executeProposal", proposalID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x980ff6c6.
//
// Solidity: function executeProposal(bytes32 proposalID) returns()
func (_MasterDAO *MasterDAOSession) ExecuteProposal(proposalID [32]byte) (*types.Transaction, error) {
	return _MasterDAO.Contract.ExecuteProposal(&_MasterDAO.TransactOpts, proposalID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x980ff6c6.
//
// Solidity: function executeProposal(bytes32 proposalID) returns()
func (_MasterDAO *MasterDAOTransactorSession) ExecuteProposal(proposalID [32]byte) (*types.Transaction, error) {
	return _MasterDAO.Contract.ExecuteProposal(&_MasterDAO.TransactOpts, proposalID)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_MasterDAO *MasterDAOTransactor) Init(opts *bind.TransactOpts, addrReg common.Address) (*types.Transaction, error) {
	return _MasterDAO.contract.Transact(opts, "init", addrReg)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_MasterDAO *MasterDAOSession) Init(addrReg common.Address) (*types.Transaction, error) {
	return _MasterDAO.Contract.Init(&_MasterDAO.TransactOpts, addrReg)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address addrReg) returns()
func (_MasterDAO *MasterDAOTransactorSession) Init(addrReg common.Address) (*types.Transaction, error) {
	return _MasterDAO.Contract.Init(&_MasterDAO.TransactOpts, addrReg)
}

// Init0 is a paid mutator transaction binding the contract method 0x378dfd8e.
//
// Solidity: function init(address admin, address addrRegistry, bytes data) returns(bytes callbackEvent)
func (_MasterDAO *MasterDAOTransactor) Init0(opts *bind.TransactOpts, admin common.Address, addrRegistry common.Address, data []byte) (*types.Transaction, error) {
	return _MasterDAO.contract.Transact(opts, "init0", admin, addrRegistry, data)
}

// Init0 is a paid mutator transaction binding the contract method 0x378dfd8e.
//
// Solidity: function init(address admin, address addrRegistry, bytes data) returns(bytes callbackEvent)
func (_MasterDAO *MasterDAOSession) Init0(admin common.Address, addrRegistry common.Address, data []byte) (*types.Transaction, error) {
	return _MasterDAO.Contract.Init0(&_MasterDAO.TransactOpts, admin, addrRegistry, data)
}

// Init0 is a paid mutator transaction binding the contract method 0x378dfd8e.
//
// Solidity: function init(address admin, address addrRegistry, bytes data) returns(bytes callbackEvent)
func (_MasterDAO *MasterDAOTransactorSession) Init0(admin common.Address, addrRegistry common.Address, data []byte) (*types.Transaction, error) {
	return _MasterDAO.Contract.Init0(&_MasterDAO.TransactOpts, admin, addrRegistry, data)
}

// NewProposal is a paid mutator transaction binding the contract method 0x8d2a2869.
//
// Solidity: function newProposal((bytes[],bytes[]) proposal, bytes data) returns(bytes32 proposalID)
func (_MasterDAO *MasterDAOTransactor) NewProposal(opts *bind.TransactOpts, proposal IProposalOperatorProposalApplyInfo, data []byte) (*types.Transaction, error) {
	return _MasterDAO.contract.Transact(opts, "newProposal", proposal, data)
}

// NewProposal is a paid mutator transaction binding the contract method 0x8d2a2869.
//
// Solidity: function newProposal((bytes[],bytes[]) proposal, bytes data) returns(bytes32 proposalID)
func (_MasterDAO *MasterDAOSession) NewProposal(proposal IProposalOperatorProposalApplyInfo, data []byte) (*types.Transaction, error) {
	return _MasterDAO.Contract.NewProposal(&_MasterDAO.TransactOpts, proposal, data)
}

// NewProposal is a paid mutator transaction binding the contract method 0x8d2a2869.
//
// Solidity: function newProposal((bytes[],bytes[]) proposal, bytes data) returns(bytes32 proposalID)
func (_MasterDAO *MasterDAOTransactorSession) NewProposal(proposal IProposalOperatorProposalApplyInfo, data []byte) (*types.Transaction, error) {
	return _MasterDAO.Contract.NewProposal(&_MasterDAO.TransactOpts, proposal, data)
}

// SetFlowStep is a paid mutator transaction binding the contract method 0x1fdb32cb.
//
// Solidity: function setFlowStep((bytes32,(bytes32,address,uint256,string)[]) flow) returns()
func (_MasterDAO *MasterDAOTransactor) SetFlowStep(opts *bind.TransactOpts, flow IDAOFlowInfo) (*types.Transaction, error) {
	return _MasterDAO.contract.Transact(opts, "setFlowStep", flow)
}

// SetFlowStep is a paid mutator transaction binding the contract method 0x1fdb32cb.
//
// Solidity: function setFlowStep((bytes32,(bytes32,address,uint256,string)[]) flow) returns()
func (_MasterDAO *MasterDAOSession) SetFlowStep(flow IDAOFlowInfo) (*types.Transaction, error) {
	return _MasterDAO.Contract.SetFlowStep(&_MasterDAO.TransactOpts, flow)
}

// SetFlowStep is a paid mutator transaction binding the contract method 0x1fdb32cb.
//
// Solidity: function setFlowStep((bytes32,(bytes32,address,uint256,string)[]) flow) returns()
func (_MasterDAO *MasterDAOTransactorSession) SetFlowStep(flow IDAOFlowInfo) (*types.Transaction, error) {
	return _MasterDAO.Contract.SetFlowStep(&_MasterDAO.TransactOpts, flow)
}

// SetUCV is a paid mutator transaction binding the contract method 0x75a3e954.
//
// Solidity: function setUCV(address ucv) returns()
func (_MasterDAO *MasterDAOTransactor) SetUCV(opts *bind.TransactOpts, ucv common.Address) (*types.Transaction, error) {
	return _MasterDAO.contract.Transact(opts, "setUCV", ucv)
}

// SetUCV is a paid mutator transaction binding the contract method 0x75a3e954.
//
// Solidity: function setUCV(address ucv) returns()
func (_MasterDAO *MasterDAOSession) SetUCV(ucv common.Address) (*types.Transaction, error) {
	return _MasterDAO.Contract.SetUCV(&_MasterDAO.TransactOpts, ucv)
}

// SetUCV is a paid mutator transaction binding the contract method 0x75a3e954.
//
// Solidity: function setUCV(address ucv) returns()
func (_MasterDAO *MasterDAOTransactorSession) SetUCV(ucv common.Address) (*types.Transaction, error) {
	return _MasterDAO.Contract.SetUCV(&_MasterDAO.TransactOpts, ucv)
}

// MasterDAOEAddUCVIterator is returned from FilterEAddUCV and is used to iterate over the raw logs and unpacked data for EAddUCV events raised by the MasterDAO contract.
type MasterDAOEAddUCVIterator struct {
	Event *MasterDAOEAddUCV // Event containing the contract specifics and raw log

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
func (it *MasterDAOEAddUCVIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterDAOEAddUCV)
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
		it.Event = new(MasterDAOEAddUCV)
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
func (it *MasterDAOEAddUCVIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterDAOEAddUCVIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterDAOEAddUCV represents a EAddUCV event raised by the MasterDAO contract.
type MasterDAOEAddUCV struct {
	Ucv common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEAddUCV is a free log retrieval operation binding the contract event 0x81905b57cc2663b6c8fa014c1471f293aab627ae83ea96f2b9bd5e92b7a7c154.
//
// Solidity: event EAddUCV(address indexed ucv)
func (_MasterDAO *MasterDAOFilterer) FilterEAddUCV(opts *bind.FilterOpts, ucv []common.Address) (*MasterDAOEAddUCVIterator, error) {

	var ucvRule []interface{}
	for _, ucvItem := range ucv {
		ucvRule = append(ucvRule, ucvItem)
	}

	logs, sub, err := _MasterDAO.contract.FilterLogs(opts, "EAddUCV", ucvRule)
	if err != nil {
		return nil, err
	}
	return &MasterDAOEAddUCVIterator{contract: _MasterDAO.contract, event: "EAddUCV", logs: logs, sub: sub}, nil
}

// WatchEAddUCV is a free log subscription operation binding the contract event 0x81905b57cc2663b6c8fa014c1471f293aab627ae83ea96f2b9bd5e92b7a7c154.
//
// Solidity: event EAddUCV(address indexed ucv)
func (_MasterDAO *MasterDAOFilterer) WatchEAddUCV(opts *bind.WatchOpts, sink chan<- *MasterDAOEAddUCV, ucv []common.Address) (event.Subscription, error) {

	var ucvRule []interface{}
	for _, ucvItem := range ucv {
		ucvRule = append(ucvRule, ucvItem)
	}

	logs, sub, err := _MasterDAO.contract.WatchLogs(opts, "EAddUCV", ucvRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterDAOEAddUCV)
				if err := _MasterDAO.contract.UnpackLog(event, "EAddUCV", log); err != nil {
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

// ParseEAddUCV is a log parse operation binding the contract event 0x81905b57cc2663b6c8fa014c1471f293aab627ae83ea96f2b9bd5e92b7a7c154.
//
// Solidity: event EAddUCV(address indexed ucv)
func (_MasterDAO *MasterDAOFilterer) ParseEAddUCV(log types.Log) (*MasterDAOEAddUCV, error) {
	event := new(MasterDAOEAddUCV)
	if err := _MasterDAO.contract.UnpackLog(event, "EAddUCV", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterDAOEBadgeIterator is returned from FilterEBadge and is used to iterate over the raw logs and unpacked data for EBadge events raised by the MasterDAO contract.
type MasterDAOEBadgeIterator struct {
	Event *MasterDAOEBadge // Event containing the contract specifics and raw log

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
func (it *MasterDAOEBadgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterDAOEBadge)
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
		it.Event = new(MasterDAOEBadge)
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
func (it *MasterDAOEBadgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterDAOEBadgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterDAOEBadge represents a EBadge event raised by the MasterDAO contract.
type MasterDAOEBadge struct {
	Token common.Address
	Name  string
	Total *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEBadge is a free log retrieval operation binding the contract event 0xfd4c7fa6aa14112ccf22002a905c4a1b770855e4422f84264bf02fc8a9a03978.
//
// Solidity: event EBadge(address token, string name, uint256 total)
func (_MasterDAO *MasterDAOFilterer) FilterEBadge(opts *bind.FilterOpts) (*MasterDAOEBadgeIterator, error) {

	logs, sub, err := _MasterDAO.contract.FilterLogs(opts, "EBadge")
	if err != nil {
		return nil, err
	}
	return &MasterDAOEBadgeIterator{contract: _MasterDAO.contract, event: "EBadge", logs: logs, sub: sub}, nil
}

// WatchEBadge is a free log subscription operation binding the contract event 0xfd4c7fa6aa14112ccf22002a905c4a1b770855e4422f84264bf02fc8a9a03978.
//
// Solidity: event EBadge(address token, string name, uint256 total)
func (_MasterDAO *MasterDAOFilterer) WatchEBadge(opts *bind.WatchOpts, sink chan<- *MasterDAOEBadge) (event.Subscription, error) {

	logs, sub, err := _MasterDAO.contract.WatchLogs(opts, "EBadge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterDAOEBadge)
				if err := _MasterDAO.contract.UnpackLog(event, "EBadge", log); err != nil {
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

// ParseEBadge is a log parse operation binding the contract event 0xfd4c7fa6aa14112ccf22002a905c4a1b770855e4422f84264bf02fc8a9a03978.
//
// Solidity: event EBadge(address token, string name, uint256 total)
func (_MasterDAO *MasterDAOFilterer) ParseEBadge(log types.Log) (*MasterDAOEBadge, error) {
	event := new(MasterDAOEBadge)
	if err := _MasterDAO.contract.UnpackLog(event, "EBadge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterDAOECreateDAOIterator is returned from FilterECreateDAO and is used to iterate over the raw logs and unpacked data for ECreateDAO events raised by the MasterDAO contract.
type MasterDAOECreateDAOIterator struct {
	Event *MasterDAOECreateDAO // Event containing the contract specifics and raw log

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
func (it *MasterDAOECreateDAOIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterDAOECreateDAO)
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
		it.Event = new(MasterDAOECreateDAO)
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
func (it *MasterDAOECreateDAOIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterDAOECreateDAOIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterDAOECreateDAO represents a ECreateDAO event raised by the MasterDAO contract.
type MasterDAOECreateDAO struct {
	Addr            common.Address
	Admin           common.Address
	GovernanceToken common.Address
	Name            string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterECreateDAO is a free log retrieval operation binding the contract event 0xc06ef9efcb36a8aa8456d39009491feb0f86daa0ca66d81c4e20426f74658139.
//
// Solidity: event ECreateDAO(address indexed addr, address indexed admin, address indexed governanceToken, string name)
func (_MasterDAO *MasterDAOFilterer) FilterECreateDAO(opts *bind.FilterOpts, addr []common.Address, admin []common.Address, governanceToken []common.Address) (*MasterDAOECreateDAOIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var governanceTokenRule []interface{}
	for _, governanceTokenItem := range governanceToken {
		governanceTokenRule = append(governanceTokenRule, governanceTokenItem)
	}

	logs, sub, err := _MasterDAO.contract.FilterLogs(opts, "ECreateDAO", addrRule, adminRule, governanceTokenRule)
	if err != nil {
		return nil, err
	}
	return &MasterDAOECreateDAOIterator{contract: _MasterDAO.contract, event: "ECreateDAO", logs: logs, sub: sub}, nil
}

// WatchECreateDAO is a free log subscription operation binding the contract event 0xc06ef9efcb36a8aa8456d39009491feb0f86daa0ca66d81c4e20426f74658139.
//
// Solidity: event ECreateDAO(address indexed addr, address indexed admin, address indexed governanceToken, string name)
func (_MasterDAO *MasterDAOFilterer) WatchECreateDAO(opts *bind.WatchOpts, sink chan<- *MasterDAOECreateDAO, addr []common.Address, admin []common.Address, governanceToken []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var governanceTokenRule []interface{}
	for _, governanceTokenItem := range governanceToken {
		governanceTokenRule = append(governanceTokenRule, governanceTokenItem)
	}

	logs, sub, err := _MasterDAO.contract.WatchLogs(opts, "ECreateDAO", addrRule, adminRule, governanceTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterDAOECreateDAO)
				if err := _MasterDAO.contract.UnpackLog(event, "ECreateDAO", log); err != nil {
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

// ParseECreateDAO is a log parse operation binding the contract event 0xc06ef9efcb36a8aa8456d39009491feb0f86daa0ca66d81c4e20426f74658139.
//
// Solidity: event ECreateDAO(address indexed addr, address indexed admin, address indexed governanceToken, string name)
func (_MasterDAO *MasterDAOFilterer) ParseECreateDAO(log types.Log) (*MasterDAOECreateDAO, error) {
	event := new(MasterDAOECreateDAO)
	if err := _MasterDAO.contract.UnpackLog(event, "ECreateDAO", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterDAOEDecideProposalIterator is returned from FilterEDecideProposal and is used to iterate over the raw logs and unpacked data for EDecideProposal events raised by the MasterDAO contract.
type MasterDAOEDecideProposalIterator struct {
	Event *MasterDAOEDecideProposal // Event containing the contract specifics and raw log

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
func (it *MasterDAOEDecideProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterDAOEDecideProposal)
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
		it.Event = new(MasterDAOEDecideProposal)
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
func (it *MasterDAOEDecideProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterDAOEDecideProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterDAOEDecideProposal represents a EDecideProposal event raised by the MasterDAO contract.
type MasterDAOEDecideProposal struct {
	ProposalID [32]byte
	Agree      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEDecideProposal is a free log retrieval operation binding the contract event 0xe1763b5acf2e710d7669c0f3044f674f971a95c2817c35b06806381012406045.
//
// Solidity: event EDecideProposal(bytes32 indexed proposalID, bool agree)
func (_MasterDAO *MasterDAOFilterer) FilterEDecideProposal(opts *bind.FilterOpts, proposalID [][32]byte) (*MasterDAOEDecideProposalIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _MasterDAO.contract.FilterLogs(opts, "EDecideProposal", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &MasterDAOEDecideProposalIterator{contract: _MasterDAO.contract, event: "EDecideProposal", logs: logs, sub: sub}, nil
}

// WatchEDecideProposal is a free log subscription operation binding the contract event 0xe1763b5acf2e710d7669c0f3044f674f971a95c2817c35b06806381012406045.
//
// Solidity: event EDecideProposal(bytes32 indexed proposalID, bool agree)
func (_MasterDAO *MasterDAOFilterer) WatchEDecideProposal(opts *bind.WatchOpts, sink chan<- *MasterDAOEDecideProposal, proposalID [][32]byte) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _MasterDAO.contract.WatchLogs(opts, "EDecideProposal", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterDAOEDecideProposal)
				if err := _MasterDAO.contract.UnpackLog(event, "EDecideProposal", log); err != nil {
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

// ParseEDecideProposal is a log parse operation binding the contract event 0xe1763b5acf2e710d7669c0f3044f674f971a95c2817c35b06806381012406045.
//
// Solidity: event EDecideProposal(bytes32 indexed proposalID, bool agree)
func (_MasterDAO *MasterDAOFilterer) ParseEDecideProposal(log types.Log) (*MasterDAOEDecideProposal, error) {
	event := new(MasterDAOEDecideProposal)
	if err := _MasterDAO.contract.UnpackLog(event, "EDecideProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
