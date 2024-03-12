// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vote

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

// VotingMember is an auto generated low-level Go binding around an user-defined struct.
type VotingMember struct {
	Age    uint16
	Name   string
	Person common.Address
}

// VoteMetaData contains all meta data concerning the Vote contract.
var VoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vote\",\"type\":\"address\"}],\"name\":\"GiveVote\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCandidate\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMembers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"age\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"}],\"internalType\":\"structvoting.Member[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"_age\",\"type\":\"uint16\"}],\"name\":\"setMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterVote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VoteABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteMetaData.ABI instead.
var VoteABI = VoteMetaData.ABI

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// GetAllVote is a free data retrieval call binding the contract method 0x8a95baaa.
//
// Solidity: function getAllVote() view returns(uint256, uint256, uint256)
func (_Vote *VoteCaller) GetAllVote(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getAllVote")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAllVote is a free data retrieval call binding the contract method 0x8a95baaa.
//
// Solidity: function getAllVote() view returns(uint256, uint256, uint256)
func (_Vote *VoteSession) GetAllVote() (*big.Int, *big.Int, *big.Int, error) {
	return _Vote.Contract.GetAllVote(&_Vote.CallOpts)
}

// GetAllVote is a free data retrieval call binding the contract method 0x8a95baaa.
//
// Solidity: function getAllVote() view returns(uint256, uint256, uint256)
func (_Vote *VoteCallerSession) GetAllVote() (*big.Int, *big.Int, *big.Int, error) {
	return _Vote.Contract.GetAllVote(&_Vote.CallOpts)
}

// GetCandidate is a free data retrieval call binding the contract method 0x70218b85.
//
// Solidity: function getCandidate() view returns(uint16)
func (_Vote *VoteCaller) GetCandidate(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getCandidate")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetCandidate is a free data retrieval call binding the contract method 0x70218b85.
//
// Solidity: function getCandidate() view returns(uint16)
func (_Vote *VoteSession) GetCandidate() (uint16, error) {
	return _Vote.Contract.GetCandidate(&_Vote.CallOpts)
}

// GetCandidate is a free data retrieval call binding the contract method 0x70218b85.
//
// Solidity: function getCandidate() view returns(uint16)
func (_Vote *VoteCallerSession) GetCandidate() (uint16, error) {
	return _Vote.Contract.GetCandidate(&_Vote.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns((uint16,string,address)[])
func (_Vote *VoteCaller) GetMembers(opts *bind.CallOpts) ([]VotingMember, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getMembers")

	if err != nil {
		return *new([]VotingMember), err
	}

	out0 := *abi.ConvertType(out[0], new([]VotingMember)).(*[]VotingMember)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns((uint16,string,address)[])
func (_Vote *VoteSession) GetMembers() ([]VotingMember, error) {
	return _Vote.Contract.GetMembers(&_Vote.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns((uint16,string,address)[])
func (_Vote *VoteCallerSession) GetMembers() ([]VotingMember, error) {
	return _Vote.Contract.GetMembers(&_Vote.CallOpts)
}

// VoterVote is a free data retrieval call binding the contract method 0x9e4ecbce.
//
// Solidity: function voterVote() view returns(address)
func (_Vote *VoteCaller) VoterVote(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "voterVote")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterVote is a free data retrieval call binding the contract method 0x9e4ecbce.
//
// Solidity: function voterVote() view returns(address)
func (_Vote *VoteSession) VoterVote() (common.Address, error) {
	return _Vote.Contract.VoterVote(&_Vote.CallOpts)
}

// VoterVote is a free data retrieval call binding the contract method 0x9e4ecbce.
//
// Solidity: function voterVote() view returns(address)
func (_Vote *VoteCallerSession) VoterVote() (common.Address, error) {
	return _Vote.Contract.VoterVote(&_Vote.CallOpts)
}

// GiveVote is a paid mutator transaction binding the contract method 0x44f40d01.
//
// Solidity: function GiveVote(address _vote) returns(uint8)
func (_Vote *VoteTransactor) GiveVote(opts *bind.TransactOpts, _vote common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "GiveVote", _vote)
}

// GiveVote is a paid mutator transaction binding the contract method 0x44f40d01.
//
// Solidity: function GiveVote(address _vote) returns(uint8)
func (_Vote *VoteSession) GiveVote(_vote common.Address) (*types.Transaction, error) {
	return _Vote.Contract.GiveVote(&_Vote.TransactOpts, _vote)
}

// GiveVote is a paid mutator transaction binding the contract method 0x44f40d01.
//
// Solidity: function GiveVote(address _vote) returns(uint8)
func (_Vote *VoteTransactorSession) GiveVote(_vote common.Address) (*types.Transaction, error) {
	return _Vote.Contract.GiveVote(&_Vote.TransactOpts, _vote)
}

// SetMember is a paid mutator transaction binding the contract method 0xec95372f.
//
// Solidity: function setMember(string _name, uint16 _age) returns()
func (_Vote *VoteTransactor) SetMember(opts *bind.TransactOpts, _name string, _age uint16) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setMember", _name, _age)
}

// SetMember is a paid mutator transaction binding the contract method 0xec95372f.
//
// Solidity: function setMember(string _name, uint16 _age) returns()
func (_Vote *VoteSession) SetMember(_name string, _age uint16) (*types.Transaction, error) {
	return _Vote.Contract.SetMember(&_Vote.TransactOpts, _name, _age)
}

// SetMember is a paid mutator transaction binding the contract method 0xec95372f.
//
// Solidity: function setMember(string _name, uint16 _age) returns()
func (_Vote *VoteTransactorSession) SetMember(_name string, _age uint16) (*types.Transaction, error) {
	return _Vote.Contract.SetMember(&_Vote.TransactOpts, _name, _age)
}
