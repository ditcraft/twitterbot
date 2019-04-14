// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ditCoordinator

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// KNWTokenContractABI is the input ABI used to generate the binding from.
const KNWTokenContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newVotingAddress\",\"type\":\"address\"}],\"name\":\"setVotingAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KNWTokenContractBin is the compiled bytecode used for deploying new contracts.
const KNWTokenContractBin = `0x`

// DeployKNWTokenContract deploys a new Ethereum contract, binding an instance of KNWTokenContract to it.
func DeployKNWTokenContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KNWTokenContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWTokenContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KNWTokenContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KNWTokenContract{KNWTokenContractCaller: KNWTokenContractCaller{contract: contract}, KNWTokenContractTransactor: KNWTokenContractTransactor{contract: contract}, KNWTokenContractFilterer: KNWTokenContractFilterer{contract: contract}}, nil
}

// KNWTokenContract is an auto generated Go binding around an Ethereum contract.
type KNWTokenContract struct {
	KNWTokenContractCaller     // Read-only binding to the contract
	KNWTokenContractTransactor // Write-only binding to the contract
	KNWTokenContractFilterer   // Log filterer for contract events
}

// KNWTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KNWTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KNWTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KNWTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KNWTokenContractSession struct {
	Contract     *KNWTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KNWTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KNWTokenContractCallerSession struct {
	Contract *KNWTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KNWTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KNWTokenContractTransactorSession struct {
	Contract     *KNWTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KNWTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KNWTokenContractRaw struct {
	Contract *KNWTokenContract // Generic contract binding to access the raw methods on
}

// KNWTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KNWTokenContractCallerRaw struct {
	Contract *KNWTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// KNWTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KNWTokenContractTransactorRaw struct {
	Contract *KNWTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKNWTokenContract creates a new instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContract(address common.Address, backend bind.ContractBackend) (*KNWTokenContract, error) {
	contract, err := bindKNWTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContract{KNWTokenContractCaller: KNWTokenContractCaller{contract: contract}, KNWTokenContractTransactor: KNWTokenContractTransactor{contract: contract}, KNWTokenContractFilterer: KNWTokenContractFilterer{contract: contract}}, nil
}

// NewKNWTokenContractCaller creates a new read-only instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContractCaller(address common.Address, caller bind.ContractCaller) (*KNWTokenContractCaller, error) {
	contract, err := bindKNWTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContractCaller{contract: contract}, nil
}

// NewKNWTokenContractTransactor creates a new write-only instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*KNWTokenContractTransactor, error) {
	contract, err := bindKNWTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContractTransactor{contract: contract}, nil
}

// NewKNWTokenContractFilterer creates a new log filterer instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*KNWTokenContractFilterer, error) {
	contract, err := bindKNWTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContractFilterer{contract: contract}, nil
}

// bindKNWTokenContract binds a generic wrapper to an already deployed contract.
func bindKNWTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWTokenContract *KNWTokenContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWTokenContract.Contract.KNWTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWTokenContract *KNWTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.KNWTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWTokenContract *KNWTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.KNWTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWTokenContract *KNWTokenContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWTokenContract *KNWTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWTokenContract *KNWTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.contract.Transact(opts, method, params...)
}

// SetVotingAddress is a paid mutator transaction binding the contract method 0x7a6cfcab.
//
// Solidity: function setVotingAddress(address _newVotingAddress) returns()
func (_KNWTokenContract *KNWTokenContractTransactor) SetVotingAddress(opts *bind.TransactOpts, _newVotingAddress common.Address) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "setVotingAddress", _newVotingAddress)
}

// SetVotingAddress is a paid mutator transaction binding the contract method 0x7a6cfcab.
//
// Solidity: function setVotingAddress(address _newVotingAddress) returns()
func (_KNWTokenContract *KNWTokenContractSession) SetVotingAddress(_newVotingAddress common.Address) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.SetVotingAddress(&_KNWTokenContract.TransactOpts, _newVotingAddress)
}

// SetVotingAddress is a paid mutator transaction binding the contract method 0x7a6cfcab.
//
// Solidity: function setVotingAddress(address _newVotingAddress) returns()
func (_KNWTokenContract *KNWTokenContractTransactorSession) SetVotingAddress(_newVotingAddress common.Address) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.SetVotingAddress(&_KNWTokenContract.TransactOpts, _newVotingAddress)
}

// KNWVotingContractABI is the input ABI used to generate the binding from.
const KNWVotingContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newKNWTokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"}],\"name\":\"commitVote\",\"outputs\":[{\"name\":\"numVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_knowledgeLabel\",\"type\":\"string\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"},{\"name\":\"_proposersStake\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRepository\",\"type\":\"bytes32\"},{\"name\":\"_majority\",\"type\":\"uint256\"},{\"name\":\"_mintingMethod\",\"type\":\"uint256\"},{\"name\":\"_burningMethod\",\"type\":\"uint256\"}],\"name\":\"addNewRepository\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"resolveVote\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"resolvePoll\",\"outputs\":[{\"name\":\"votePassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCoordinatorAddress\",\"type\":\"address\"}],\"name\":\"setCoordinatorAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KNWVotingContractBin is the compiled bytecode used for deploying new contracts.
const KNWVotingContractBin = `0x`

// DeployKNWVotingContract deploys a new Ethereum contract, binding an instance of KNWVotingContract to it.
func DeployKNWVotingContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KNWVotingContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWVotingContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KNWVotingContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KNWVotingContract{KNWVotingContractCaller: KNWVotingContractCaller{contract: contract}, KNWVotingContractTransactor: KNWVotingContractTransactor{contract: contract}, KNWVotingContractFilterer: KNWVotingContractFilterer{contract: contract}}, nil
}

// KNWVotingContract is an auto generated Go binding around an Ethereum contract.
type KNWVotingContract struct {
	KNWVotingContractCaller     // Read-only binding to the contract
	KNWVotingContractTransactor // Write-only binding to the contract
	KNWVotingContractFilterer   // Log filterer for contract events
}

// KNWVotingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KNWVotingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KNWVotingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KNWVotingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KNWVotingContractSession struct {
	Contract     *KNWVotingContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KNWVotingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KNWVotingContractCallerSession struct {
	Contract *KNWVotingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// KNWVotingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KNWVotingContractTransactorSession struct {
	Contract     *KNWVotingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// KNWVotingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KNWVotingContractRaw struct {
	Contract *KNWVotingContract // Generic contract binding to access the raw methods on
}

// KNWVotingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KNWVotingContractCallerRaw struct {
	Contract *KNWVotingContractCaller // Generic read-only contract binding to access the raw methods on
}

// KNWVotingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KNWVotingContractTransactorRaw struct {
	Contract *KNWVotingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKNWVotingContract creates a new instance of KNWVotingContract, bound to a specific deployed contract.
func NewKNWVotingContract(address common.Address, backend bind.ContractBackend) (*KNWVotingContract, error) {
	contract, err := bindKNWVotingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KNWVotingContract{KNWVotingContractCaller: KNWVotingContractCaller{contract: contract}, KNWVotingContractTransactor: KNWVotingContractTransactor{contract: contract}, KNWVotingContractFilterer: KNWVotingContractFilterer{contract: contract}}, nil
}

// NewKNWVotingContractCaller creates a new read-only instance of KNWVotingContract, bound to a specific deployed contract.
func NewKNWVotingContractCaller(address common.Address, caller bind.ContractCaller) (*KNWVotingContractCaller, error) {
	contract, err := bindKNWVotingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KNWVotingContractCaller{contract: contract}, nil
}

// NewKNWVotingContractTransactor creates a new write-only instance of KNWVotingContract, bound to a specific deployed contract.
func NewKNWVotingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*KNWVotingContractTransactor, error) {
	contract, err := bindKNWVotingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KNWVotingContractTransactor{contract: contract}, nil
}

// NewKNWVotingContractFilterer creates a new log filterer instance of KNWVotingContract, bound to a specific deployed contract.
func NewKNWVotingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*KNWVotingContractFilterer, error) {
	contract, err := bindKNWVotingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KNWVotingContractFilterer{contract: contract}, nil
}

// bindKNWVotingContract binds a generic wrapper to an already deployed contract.
func bindKNWVotingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWVotingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWVotingContract *KNWVotingContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWVotingContract.Contract.KNWVotingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWVotingContract *KNWVotingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.KNWVotingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWVotingContract *KNWVotingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.KNWVotingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWVotingContract *KNWVotingContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWVotingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWVotingContract *KNWVotingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWVotingContract *KNWVotingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.contract.Transact(opts, method, params...)
}

// ResolveVote is a free data retrieval call binding the contract method 0xce729fd2.
//
// Solidity: function resolveVote(uint256 _pollID, uint256 _voteOption, address _address) constant returns(uint256 reward)
func (_KNWVotingContract *KNWVotingContractCaller) ResolveVote(opts *bind.CallOpts, _pollID *big.Int, _voteOption *big.Int, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVotingContract.contract.Call(opts, out, "resolveVote", _pollID, _voteOption, _address)
	return *ret0, err
}

// ResolveVote is a free data retrieval call binding the contract method 0xce729fd2.
//
// Solidity: function resolveVote(uint256 _pollID, uint256 _voteOption, address _address) constant returns(uint256 reward)
func (_KNWVotingContract *KNWVotingContractSession) ResolveVote(_pollID *big.Int, _voteOption *big.Int, _address common.Address) (*big.Int, error) {
	return _KNWVotingContract.Contract.ResolveVote(&_KNWVotingContract.CallOpts, _pollID, _voteOption, _address)
}

// ResolveVote is a free data retrieval call binding the contract method 0xce729fd2.
//
// Solidity: function resolveVote(uint256 _pollID, uint256 _voteOption, address _address) constant returns(uint256 reward)
func (_KNWVotingContract *KNWVotingContractCallerSession) ResolveVote(_pollID *big.Int, _voteOption *big.Int, _address common.Address) (*big.Int, error) {
	return _KNWVotingContract.Contract.ResolveVote(&_KNWVotingContract.CallOpts, _pollID, _voteOption, _address)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0xa3fba060.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority, uint256 _mintingMethod, uint256 _burningMethod) returns()
func (_KNWVotingContract *KNWVotingContractTransactor) AddNewRepository(opts *bind.TransactOpts, _newRepository [32]byte, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "addNewRepository", _newRepository, _majority, _mintingMethod, _burningMethod)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0xa3fba060.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority, uint256 _mintingMethod, uint256 _burningMethod) returns()
func (_KNWVotingContract *KNWVotingContractSession) AddNewRepository(_newRepository [32]byte, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.AddNewRepository(&_KNWVotingContract.TransactOpts, _newRepository, _majority, _mintingMethod, _burningMethod)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0xa3fba060.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority, uint256 _mintingMethod, uint256 _burningMethod) returns()
func (_KNWVotingContract *KNWVotingContractTransactorSession) AddNewRepository(_newRepository [32]byte, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.AddNewRepository(&_KNWVotingContract.TransactOpts, _newRepository, _majority, _mintingMethod, _burningMethod)
}

// CommitVote is a paid mutator transaction binding the contract method 0x7eb2ff52.
//
// Solidity: function commitVote(uint256 _pollID, address _address, bytes32 _secretHash) returns(uint256 numVotes)
func (_KNWVotingContract *KNWVotingContractTransactor) CommitVote(opts *bind.TransactOpts, _pollID *big.Int, _address common.Address, _secretHash [32]byte) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "commitVote", _pollID, _address, _secretHash)
}

// CommitVote is a paid mutator transaction binding the contract method 0x7eb2ff52.
//
// Solidity: function commitVote(uint256 _pollID, address _address, bytes32 _secretHash) returns(uint256 numVotes)
func (_KNWVotingContract *KNWVotingContractSession) CommitVote(_pollID *big.Int, _address common.Address, _secretHash [32]byte) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.CommitVote(&_KNWVotingContract.TransactOpts, _pollID, _address, _secretHash)
}

// CommitVote is a paid mutator transaction binding the contract method 0x7eb2ff52.
//
// Solidity: function commitVote(uint256 _pollID, address _address, bytes32 _secretHash) returns(uint256 numVotes)
func (_KNWVotingContract *KNWVotingContractTransactorSession) CommitVote(_pollID *big.Int, _address common.Address, _secretHash [32]byte) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.CommitVote(&_KNWVotingContract.TransactOpts, _pollID, _address, _secretHash)
}

// ResolvePoll is a paid mutator transaction binding the contract method 0xe74fef37.
//
// Solidity: function resolvePoll(uint256 _pollID) returns(bool votePassed)
func (_KNWVotingContract *KNWVotingContractTransactor) ResolvePoll(opts *bind.TransactOpts, _pollID *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "resolvePoll", _pollID)
}

// ResolvePoll is a paid mutator transaction binding the contract method 0xe74fef37.
//
// Solidity: function resolvePoll(uint256 _pollID) returns(bool votePassed)
func (_KNWVotingContract *KNWVotingContractSession) ResolvePoll(_pollID *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.ResolvePoll(&_KNWVotingContract.TransactOpts, _pollID)
}

// ResolvePoll is a paid mutator transaction binding the contract method 0xe74fef37.
//
// Solidity: function resolvePoll(uint256 _pollID) returns(bool votePassed)
func (_KNWVotingContract *KNWVotingContractTransactorSession) ResolvePoll(_pollID *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.ResolvePoll(&_KNWVotingContract.TransactOpts, _pollID)
}

// RevealVote is a paid mutator transaction binding the contract method 0x34f2f2d2.
//
// Solidity: function revealVote(uint256 _pollID, address _address, uint256 _voteOption, uint256 _salt) returns()
func (_KNWVotingContract *KNWVotingContractTransactor) RevealVote(opts *bind.TransactOpts, _pollID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "revealVote", _pollID, _address, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x34f2f2d2.
//
// Solidity: function revealVote(uint256 _pollID, address _address, uint256 _voteOption, uint256 _salt) returns()
func (_KNWVotingContract *KNWVotingContractSession) RevealVote(_pollID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.RevealVote(&_KNWVotingContract.TransactOpts, _pollID, _address, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x34f2f2d2.
//
// Solidity: function revealVote(uint256 _pollID, address _address, uint256 _voteOption, uint256 _salt) returns()
func (_KNWVotingContract *KNWVotingContractTransactorSession) RevealVote(_pollID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.RevealVote(&_KNWVotingContract.TransactOpts, _pollID, _address, _voteOption, _salt)
}

// SetCoordinatorAddress is a paid mutator transaction binding the contract method 0xf354b838.
//
// Solidity: function setCoordinatorAddress(address _newCoordinatorAddress) returns()
func (_KNWVotingContract *KNWVotingContractTransactor) SetCoordinatorAddress(opts *bind.TransactOpts, _newCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "setCoordinatorAddress", _newCoordinatorAddress)
}

// SetCoordinatorAddress is a paid mutator transaction binding the contract method 0xf354b838.
//
// Solidity: function setCoordinatorAddress(address _newCoordinatorAddress) returns()
func (_KNWVotingContract *KNWVotingContractSession) SetCoordinatorAddress(_newCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.SetCoordinatorAddress(&_KNWVotingContract.TransactOpts, _newCoordinatorAddress)
}

// SetCoordinatorAddress is a paid mutator transaction binding the contract method 0xf354b838.
//
// Solidity: function setCoordinatorAddress(address _newCoordinatorAddress) returns()
func (_KNWVotingContract *KNWVotingContractTransactorSession) SetCoordinatorAddress(_newCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.SetCoordinatorAddress(&_KNWVotingContract.TransactOpts, _newCoordinatorAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _newKNWTokenAddress) returns()
func (_KNWVotingContract *KNWVotingContractTransactor) SetTokenAddress(opts *bind.TransactOpts, _newKNWTokenAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "setTokenAddress", _newKNWTokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _newKNWTokenAddress) returns()
func (_KNWVotingContract *KNWVotingContractSession) SetTokenAddress(_newKNWTokenAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.SetTokenAddress(&_KNWVotingContract.TransactOpts, _newKNWTokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _newKNWTokenAddress) returns()
func (_KNWVotingContract *KNWVotingContractTransactorSession) SetTokenAddress(_newKNWTokenAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.SetTokenAddress(&_KNWVotingContract.TransactOpts, _newKNWTokenAddress)
}

// StartPoll is a paid mutator transaction binding the contract method 0x9156cd07.
//
// Solidity: function startPoll(bytes32 _repository, address _address, string _knowledgeLabel, uint256 _commitDuration, uint256 _revealDuration, uint256 _proposersStake) returns(uint256 pollID)
func (_KNWVotingContract *KNWVotingContractTransactor) StartPoll(opts *bind.TransactOpts, _repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "startPoll", _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// StartPoll is a paid mutator transaction binding the contract method 0x9156cd07.
//
// Solidity: function startPoll(bytes32 _repository, address _address, string _knowledgeLabel, uint256 _commitDuration, uint256 _revealDuration, uint256 _proposersStake) returns(uint256 pollID)
func (_KNWVotingContract *KNWVotingContractSession) StartPoll(_repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.StartPoll(&_KNWVotingContract.TransactOpts, _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// StartPoll is a paid mutator transaction binding the contract method 0x9156cd07.
//
// Solidity: function startPoll(bytes32 _repository, address _address, string _knowledgeLabel, uint256 _commitDuration, uint256 _revealDuration, uint256 _proposersStake) returns(uint256 pollID)
func (_KNWVotingContract *KNWVotingContractTransactorSession) StartPoll(_repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.StartPoll(&_KNWVotingContract.TransactOpts, _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"sqrt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x61016b610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100575763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663677342ce811461005c575b600080fd5b610067600435610079565b60408051918252519081900360200190f35b6000808083151561008d5760009250610138565b6001840184106100fe57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f466c6177656420696e70757420666f7220737172740000000000000000000000604482015290519081900360640190fd5b505060026001830104825b80821015610134575080600281808681151561012157fe5b040181151561012c57fe5b049150610109565b8192505b50509190505600a165627a7a72305820ce9eba5c2c770e6489e8634ca0a14db23dee696aa56e1247aac6317c7eab01bd0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// Sqrt is a free data retrieval call binding the contract method 0x677342ce.
//
// Solidity: function sqrt(uint256 a) constant returns(uint256)
func (_SafeMath *SafeMathCaller) Sqrt(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeMath.contract.Call(opts, out, "sqrt", a)
	return *ret0, err
}

// Sqrt is a free data retrieval call binding the contract method 0x677342ce.
//
// Solidity: function sqrt(uint256 a) constant returns(uint256)
func (_SafeMath *SafeMathSession) Sqrt(a *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Sqrt(&_SafeMath.CallOpts, a)
}

// Sqrt is a free data retrieval call binding the contract method 0x677342ce.
//
// Solidity: function sqrt(uint256 a) constant returns(uint256)
func (_SafeMath *SafeMathCallerSession) Sqrt(a *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Sqrt(&_SafeMath.CallOpts, a)
}

// DitCoordinatorABI is the input ABI used to generate the binding from.
const DitCoordinatorABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getKNWVoteIDFromProposalID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_knowledgeLabelIndex\",\"type\":\"uint256\"},{\"name\":\"_voteCommitDuration\",\"type\":\"uint256\"},{\"name\":\"_voteOpenDuration\",\"type\":\"uint256\"}],\"name\":\"proposeCommit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getCurrentProposalID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_voteSalt\",\"type\":\"uint256\"}],\"name\":\"openVoteOnProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"repositories\",\"outputs\":[{\"name\":\"votingMajority\",\"type\":\"uint256\"},{\"name\":\"mintingMethod\",\"type\":\"uint256\"},{\"name\":\"burningMethod\",\"type\":\"uint256\"},{\"name\":\"currentProposalID\",\"type\":\"uint256\"},{\"name\":\"minVoteCommitDuration\",\"type\":\"uint256\"},{\"name\":\"maxVoteCommitDuration\",\"type\":\"uint256\"},{\"name\":\"minVoteOpenDuration\",\"type\":\"uint256\"},{\"name\":\"maxVoteOpenDuration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"finalizeVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeKYC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalsOfRepository\",\"outputs\":[{\"name\":\"KNWVoteID\",\"type\":\"uint256\"},{\"name\":\"knowledgeLabel\",\"type\":\"string\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"isFinalized\",\"type\":\"bool\"},{\"name\":\"proposalAccepted\",\"type\":\"bool\"},{\"name\":\"individualStake\",\"type\":\"uint256\"},{\"name\":\"totalStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_label1\",\"type\":\"string\"},{\"name\":\"_label2\",\"type\":\"string\"},{\"name\":\"_label3\",\"type\":\"string\"},{\"name\":\"_voteSettings\",\"type\":\"uint256[7]\"}],\"name\":\"initRepository\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getIndividualStake\",\"outputs\":[{\"name\":\"individualStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeKYCValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalHasPassed\",\"outputs\":[{\"name\":\"hasPassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_knowledgeLabelID\",\"type\":\"uint256\"}],\"name\":\"getKnowledgeLabels\",\"outputs\":[{\"name\":\"knowledgeLabel\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteHash\",\"type\":\"bytes32\"}],\"name\":\"voteOnProposal\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"passedKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addKYCValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"repositoryIsInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWVotingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"passKYC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_KNWTokenAddress\",\"type\":\"address\"},{\"name\":\"_KNWVotingAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"}],\"name\":\"ProposeCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"CommitVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accept\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"OpenVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"FinalizeVote\",\"type\":\"event\"}]"

// DitCoordinatorBin is the compiled bytecode used for deploying new contracts.
const DitCoordinatorBin = `0x608060405234801561001057600080fd5b506040516040806121f3833981016040528051602090910151600160a060020a038116158015906100495750600160a060020a03821615155b15156100dc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4b4e57566f74696e6720616e64204b4e57546f6b656e2061646472657373206360448201527f616e277420626520656d70747900000000000000000000000000000000000000606482015290519081900360840190fd5b60008054600160a060020a03928316600160a060020a031991821617808355600280548316918516919091179055600180549484169482169490941780855560038054909216931692909217909155338152600760205260409020805460ff191690911790556120a2806101516000396000f30060806040526004361061011c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306ee459681146101215780630aba86881461014e5780630bdc90e8146101645780630ee62ec01461017c5780631341f25c1461019d5780631f51fd71146101d25780632e71d0fb1461022b57806339ba645b146102465780633e029f631461026757806351f43c2414610335578063552edc9d1461037357806373b0dddd1461038e57806387c9203d146103af57806395332229146103ca578063985dbfc51461045a578063a34c299a1461048b578063ccd9aa681461049c578063d0c397ef146104bd578063ea6c6d0f146104de578063eb49fe94146104f6578063eb9310241461050b575b600080fd5b34801561012d57600080fd5b5061013c60043560243561052c565b60408051918252519081900360200190f35b610162600435602435604435606435610549565b005b34801561017057600080fd5b5061013c600435610c29565b34801561018857600080fd5b50610162600435602435604435606435610c3e565b3480156101a957600080fd5b506101be600160a060020a0360043516610e05565b604080519115158252519081900360200190f35b3480156101de57600080fd5b506101ea600435610e1a565b604080519889526020890197909752878701959095526060870193909352608086019190915260a085015260c084015260e083015251908190036101000190f35b34801561023757600080fd5b50610162600435602435610e62565b34801561025257600080fd5b50610162600160a060020a0360043516611373565b34801561027357600080fd5b506102826004356024356113b3565b60408051888152600160a060020a038716918101919091528415156060820152831515608082015260a0810183905260c0810182905260e0602080830182815289519284019290925288516101008401918a019080838360005b838110156102f45781810151838201526020016102dc565b50505050905090810190601f1680156103215780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b34801561034157600080fd5b506101be60048035906024803580820192908101359160443580820192908101359160643590810191013560846114a7565b34801561037f57600080fd5b5061013c6004356024356117f9565b34801561039a57600080fd5b50610162600160a060020a0360043516611819565b3480156103bb57600080fd5b506101be600435602435611859565b3480156103d657600080fd5b506103e5600435602435611913565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561041f578181015183820152602001610407565b50505050905090810190601f16801561044c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561046657600080fd5b5061046f6119c1565b60408051600160a060020a039092168252519081900360200190f35b6101626004356024356044356119d0565b3480156104a857600080fd5b506101be600160a060020a0360043516611d53565b3480156104c957600080fd5b50610162600160a060020a0360043516611d68565b3480156104ea57600080fd5b506101be600435611dab565b34801561050257600080fd5b5061046f611dc1565b34801561051757600080fd5b50610162600160a060020a0360043516611dd0565b600091825260056020908152604080842092845291905290205490565b3360008181526006602052604090205460ff16151561056757600080fd5b600034116105e5576040805160e560020a62461bcd02815260206004820152602860248201527f56616c7565206f6620746865207472616e73616374696f6e2063616e206e6f7460448201527f206265207a65726f000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600085815260046020526040812085600381106105fe57fe5b0154600260001961010060018416150201909116041161068d576040805160e560020a62461bcd028152602060048201526024808201527f4b6e6f776c656467652d4c6162656c20696e646578206973206e6f7420636f7260448201527f7265637400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008581526004602052604090206007015483108015906106bf57506000858152600460205260409020600801548311155b1515610715576040805160e560020a62461bcd02815260206004820152601c60248201527f566f746520636f6d6d6974206475726174696f6e20696e76616c696400000000604482015290519081900360640190fd5b600085815260046020526040902060090154821080159061074757506000858152600460205260409020600a01548211155b151561079d576040805160e560020a62461bcd02815260206004820152601a60248201527f566f7465206f70656e206475726174696f6e20696e76616c6964000000000000604482015290519081900360640190fd5b6000858152600460205260409020600601546107c090600163ffffffff611e1316565b600086815260046020526040908190206006810192909255805160e0810190915260025490918291600160a060020a031690639156cd0790899033908a6003811061080757fe5b6040517c010000000000000000000000000000000000000000000000000000000063ffffffff871602815260048101858152600160a060020a0385166024830152606482018d9052608482018c90523460a4830181905260c0604484019081529490930180546002610100600183161502600019019091160460c4840181905290948e948e9490939260e490910190879080156108e55780601f106108ba576101008083540402835291602001916108e5565b820191906000526020600020905b8154815290600101906020018083116108c857829003601f168201915b5050975050505050505050602060405180830381600087803b15801561090a57600080fd5b505af115801561091e573d6000803e3d6000fd5b505050506040513d602081101561093457600080fd5b505181526000878152600460209081526040909120910190866003811061095757fe5b01805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156109dc5780601f106109b1576101008083540402835291602001916109dc565b820191906000526020600020905b8154815290600101906020018083116109bf57829003601f168201915b505050918352505033602080830191909152600060408084018290526060840182905234608085015260a09093018190528881526005825282812060048352838220600601548252825291909120825181558282015180519192610a4892600185019290910190611f25565b5060408281015160028301805460608601516080870151151575010000000000000000000000000000000000000000000275ff0000000000000000000000000000000000000000001991151560a060020a0274ff000000000000000000000000000000000000000019600160a060020a0390961673ffffffffffffffffffffffffffffffffffffffff199094169390931794909416919091171691909117905560a0830151600383015560c0909201516004918201556000878152600560209081528382208382528483206006015483529052919091200154610b31903463ffffffff611e1316565b6000868152600560209081526040808320600480845282852060068101805487529285529285208101959095559289905292905254339187907f171fe77c3addce776991159eb3eb73b14d9187ebd06c1c34ea12355a84ddbd83908860038110610b9757fe5b6040805160208082529390920180546002600019610100600184161502019091160493830184905292829182019084908015610c145780601f10610be957610100808354040283529160200191610c14565b820191906000526020600020905b815481529060010190602001808311610bf757829003601f168201915b50509250505060405180910390a45050505050565b60009081526004602052604090206006015490565b3360008181526006602052604090205460ff161515610c5c57600080fd5b60025460008681526005602090815260408083208884529091528082205481517f34f2f2d2000000000000000000000000000000000000000000000000000000008152600481019190915233602482015260448101879052606481018690529051600160a060020a03909316926334f2f2d29260848084019391929182900301818387803b158015610ced57600080fd5b505af1158015610d01573d6000803e3d6000fd5b505050600086815260056020818152604080842089855282528084203380865293810183529381902060018082018a9055905482518a8314948101859052928301819052606080845295820180546002600019948216156101000294909401169290920495830186905293955089948b947f864c0d6987266fd72e7e37f1fbc98b6a3794b7187dae454c67a2a626628a72ab94929390918190608082019086908015610dee5780601f10610dc357610100808354040283529160200191610dee565b820191906000526020600020905b815481529060010190602001808311610dd157829003601f168201915b505094505050505060405180910390a45050505050565b60076020526000908152604090205460ff1681565b600460205280600052604060002060009150905080600301549080600401549080600501549080600601549080600701549080600801549080600901549080600a0154905088565b3360008181526006602052604081205490919060ff161515610e8357600080fd5b60008481526005602081815260408084208785528252808420338552909201905290206002015460ff1615610f28576040805160e560020a62461bcd02815260206004820152602760248201527f45616368207061727469636970616e742063616e206f6e6c792066696e616c6960448201527f7a65206f6e636500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000848152600560208181526040808420878552825280842033855290920190528120541180610f7c57506000848152600560209081526040808320868452909152902060020154600160a060020a031633145b1515610ff8576040805160e560020a62461bcd02815260206004820152603a60248201527f4f6e6c79207061727469636970616e7473206f662074686520766f746520617260448201527f652061626c6520746f207265736f6c76652074686520766f7465000000000000606482015290519081900360840190fd5b600084815260056020908152604080832086845290915290206002015460a060020a900460ff16151561120e57600254600085815260056020908152604080832087845282528083205481517fe74fef3700000000000000000000000000000000000000000000000000000000815260048101919091529051600160a060020a039094169363e74fef3793602480840194938390030190829087803b1580156110a057600080fd5b505af11580156110b4573d6000803e3d6000fd5b505050506040513d60208110156110ca57600080fd5b505160008581526005602090815260408083208784528252918290206002808201805474ff000000000000000000000000000000000000000019961515750100000000000000000000000000000000000000000090810275ff00000000000000000000000000000000000000000019909216919091179690961660a060020a1790819055845195900460ff16801515938601939093528385526001918201805492831615610100026000190190921604928401839052869388937f6bd2699645e0f6c5547bdf0d053280e48fef1ab21514bd02c88610b1279b942a93919081906060820190859080156111fe5780601f106111d3576101008083540402835291602001916111fe565b820191906000526020600020905b8154815290600101906020018083116111e157829003601f168201915b5050935050505060405180910390a35b600254600085815260056020818152604080842088855282528084208054338087529190940183528185206001015482517fce729fd200000000000000000000000000000000000000000000000000000000815260048101959095526024850152604484015251600160a060020a039094169363ce729fd2936064808501948390030190829087803b1580156112a357600080fd5b505af11580156112b7573d6000803e3d6000fd5b505050506040513d60208110156112cd57600080fd5b50519150600082111561130957604051339083156108fc029084906000818181858888f19350505050158015611307573d6000803e3d6000fd5b505b6000848152600560209081526040808320868452909152902060040154611336908363ffffffff611e9d16565b600094855260056020818152604080882096885295815285872060048101939093553387529101905250509020600201805460ff19166001179055565b3360008181526007602052604090205460ff16151561139157600080fd5b50600160a060020a03166000908152600660205260409020805460ff19169055565b60056020908152600092835260408084208252918352918190208054600180830180548551600293821615610100026000190190911692909204601f81018790048702830187019095528482529194929390928301828280156114575780601f1061142c57610100808354040283529160200191611457565b820191906000526020600020905b81548152906001019060200180831161143a57829003601f168201915b505050600284015460038501546004909501549394600160a060020a0382169460ff60a060020a840481169550750100000000000000000000000000000000000000000090930490921692509087565b3360008181526006602052604081205490919060ff1615156114c857600080fd5b891515611545576040805160e560020a62461bcd02815260206004820152602360248201527f5265706f7369746f72792064657363726970746f722063616e2774206265207a60448201527f65726f0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008a815260046020526040902060030154156115d2576040805160e560020a62461bcd02815260206004820152602760248201527f5265706f7369746f72792063616e206f6e6c7920626520696e697469616c697a60448201527f6564206f6e636500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008811806115e15750600086115b806115ec5750600084115b1515611667576040805160e560020a62461bcd028152602060048201526024808201527f50726f76696465206174206c65617374206f6e65204b6e6f776c65646765204c60448201527f6162656c00000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b604080516101a06020601f8c01819004028201810190925261018081018a8152909182916101208301918291908e908e908190870183828082843782019150505050505081526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437505050929093525050508152843560208083019190915285013560408083019190915285013560608201526000608082015260a00184600360209081029190910135825260808601358282015260a086013560408084019190915260c087013560609093019290925260008d815260049091522081516117959082906003611fa3565b506020820151600382015560408201516004820155606082015160058201556080820151600682015560a0820151600782015560c0820151600882015560e0820151600982015561010090910151600a909101555060019998505050505050505050565b600091825260056020908152604080842092845291905290206003015490565b3360008181526007602052604090205460ff16151561183757600080fd5b50600160a060020a03166000908152600760205260409020805460ff19169055565b600082815260056020908152604080832084845290915281206002015460a060020a900460ff1615156118d6576040805160e560020a62461bcd02815260206004820152601d60248201527f50726f706f73616c206861736e2774206265656e207265736f6c766564000000604482015290519081900360640190fd5b5060009182526005602090815260408084209284529190529020600201547501000000000000000000000000000000000000000000900460ff1690565b6000828152600460205260409020606090826003811061192f57fe5b01805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156119b45780601f10611989576101008083540402835291602001916119b4565b820191906000526020600020905b81548152906001019060200180831161199757829003601f168201915b5050505050905092915050565b600154600160a060020a031681565b3360008181526006602052604081205490919060ff1615156119f157600080fd5b60008581526005602090815260408083208784529091529020600301543414611a8a576040805160e560020a62461bcd02815260206004820152603960248201527f56616c7565206f6620746865207472616e73616374696f6e20646f65736e277460448201527f206d6174636820746865207265717569726564207374616b6500000000000000606482015290519081900360840190fd5b6000858152600560209081526040808320878452909152902060020154600160a060020a0316331415611b2d576040805160e560020a62461bcd02815260206004820152603160248201527f5468652070726f706f736572206973206e6f7420616c6c6f77656420746f207660448201527f6f746520696e20612070726f706f73616c000000000000000000000000000000606482015290519081900360840190fd5b6000858152600560209081526040808320878452909152902060040154611b5a903463ffffffff611e1316565b60008681526005602090815260408083208884528252808320600480820195909555600254905482517f7eb2ff5200000000000000000000000000000000000000000000000000000000815295860152336024860152604485018890529051600160a060020a0390911693637eb2ff529360648083019493928390030190829087803b158015611be957600080fd5b505af1158015611bfd573d6000803e3d6000fd5b505050506040513d6020811015611c1357600080fd5b5051915060008211611c95576040805160e560020a62461bcd02815260206004820152603360248201527f566f74696e6720636f6e74726163742072657475726e656420616e20696e766160448201527f6c696420616d6f756e74206f6620766f74657300000000000000000000000000606482015290519081900360840190fd5b600085815260056020818152604080842088855280835281852033808752948101845282862088905594899052825280513492810183905290810186905260608082526001948501805460026101009782161597909702600019011695909504908201819052929388938a937fa01eea487bb3ec75528c167ccf90452d4164ddda7b13c55b2a89751a8dc5fbc19390918991908190608082019086908015610dee5780601f10610dc357610100808354040283529160200191610dee565b60066020526000908152604090205460ff1681565b3360008181526007602052604090205460ff161515611d8657600080fd5b50600160a060020a03166000908152600760205260409020805460ff19166001179055565b6000908152600460205260408120600301541190565b600054600160a060020a031681565b3360008181526007602052604090205460ff161515611dee57600080fd5b50600160a060020a03166000908152600660205260409020805460ff19166001179055565b600082820183811015611e96576040805160e560020a62461bcd02815260206004820152602a60248201527f526573756c742068617320746f20626520626967676572207468616e20626f7460448201527f682073756d6d616e647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b9392505050565b60008083831115611f1e576040805160e560020a62461bcd02815260206004820152603560248201527f43616e27742073756274726163742061206e756d6265722066726f6d2061207360448201527f6d616c6c6572206f6e6520776974682075696e74730000000000000000000000606482015290519081900360840190fd5b5050900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611f6657805160ff1916838001178555611f93565b82800160010185558215611f93579182015b82811115611f93578251825591602001919060010190611f78565b50611f9f929150611fef565b5090565b8260038101928215611fe3579160200282015b82811115611fe35782518051611fd3918491602090910190611f25565b5091602001919060010190611fb6565b50611f9f92915061200c565b61200991905b80821115611f9f5760008155600101611ff5565b90565b61200991905b80821115611f9f576000612026828261202f565b50600101612012565b50805460018160011615610100020316600290046000825580601f106120555750612073565b601f0160209004906000526020600020908101906120739190611fef565b505600a165627a7a72305820dfe23797f55902824dd698c2dd6d6a698dafbbb4e55266e33047a06088a7048b0029`

// DeployDitCoordinator deploys a new Ethereum contract, binding an instance of DitCoordinator to it.
func DeployDitCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _KNWTokenAddress common.Address, _KNWVotingAddress common.Address) (common.Address, *types.Transaction, *DitCoordinator, error) {
	parsed, err := abi.JSON(strings.NewReader(DitCoordinatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DitCoordinatorBin), backend, _KNWTokenAddress, _KNWVotingAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DitCoordinator{DitCoordinatorCaller: DitCoordinatorCaller{contract: contract}, DitCoordinatorTransactor: DitCoordinatorTransactor{contract: contract}, DitCoordinatorFilterer: DitCoordinatorFilterer{contract: contract}}, nil
}

// DitCoordinator is an auto generated Go binding around an Ethereum contract.
type DitCoordinator struct {
	DitCoordinatorCaller     // Read-only binding to the contract
	DitCoordinatorTransactor // Write-only binding to the contract
	DitCoordinatorFilterer   // Log filterer for contract events
}

// DitCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DitCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DitCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DitCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DitCoordinatorSession struct {
	Contract     *DitCoordinator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DitCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DitCoordinatorCallerSession struct {
	Contract *DitCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DitCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DitCoordinatorTransactorSession struct {
	Contract     *DitCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DitCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DitCoordinatorRaw struct {
	Contract *DitCoordinator // Generic contract binding to access the raw methods on
}

// DitCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DitCoordinatorCallerRaw struct {
	Contract *DitCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// DitCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DitCoordinatorTransactorRaw struct {
	Contract *DitCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDitCoordinator creates a new instance of DitCoordinator, bound to a specific deployed contract.
func NewDitCoordinator(address common.Address, backend bind.ContractBackend) (*DitCoordinator, error) {
	contract, err := bindDitCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DitCoordinator{DitCoordinatorCaller: DitCoordinatorCaller{contract: contract}, DitCoordinatorTransactor: DitCoordinatorTransactor{contract: contract}, DitCoordinatorFilterer: DitCoordinatorFilterer{contract: contract}}, nil
}

// NewDitCoordinatorCaller creates a new read-only instance of DitCoordinator, bound to a specific deployed contract.
func NewDitCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*DitCoordinatorCaller, error) {
	contract, err := bindDitCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorCaller{contract: contract}, nil
}

// NewDitCoordinatorTransactor creates a new write-only instance of DitCoordinator, bound to a specific deployed contract.
func NewDitCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DitCoordinatorTransactor, error) {
	contract, err := bindDitCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorTransactor{contract: contract}, nil
}

// NewDitCoordinatorFilterer creates a new log filterer instance of DitCoordinator, bound to a specific deployed contract.
func NewDitCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DitCoordinatorFilterer, error) {
	contract, err := bindDitCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorFilterer{contract: contract}, nil
}

// bindDitCoordinator binds a generic wrapper to an already deployed contract.
func bindDitCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DitCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DitCoordinator *DitCoordinatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DitCoordinator.Contract.DitCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DitCoordinator *DitCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DitCoordinator.Contract.DitCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DitCoordinator *DitCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DitCoordinator.Contract.DitCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DitCoordinator *DitCoordinatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DitCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DitCoordinator *DitCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DitCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DitCoordinator *DitCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DitCoordinator.Contract.contract.Transact(opts, method, params...)
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorCaller) KNWTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "KNWTokenAddress")
	return *ret0, err
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorSession) KNWTokenAddress() (common.Address, error) {
	return _DitCoordinator.Contract.KNWTokenAddress(&_DitCoordinator.CallOpts)
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorCallerSession) KNWTokenAddress() (common.Address, error) {
	return _DitCoordinator.Contract.KNWTokenAddress(&_DitCoordinator.CallOpts)
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorCaller) KNWVotingAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "KNWVotingAddress")
	return *ret0, err
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorSession) KNWVotingAddress() (common.Address, error) {
	return _DitCoordinator.Contract.KNWVotingAddress(&_DitCoordinator.CallOpts)
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorCallerSession) KNWVotingAddress() (common.Address, error) {
	return _DitCoordinator.Contract.KNWVotingAddress(&_DitCoordinator.CallOpts)
}

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(bytes32 _repository) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCaller) GetCurrentProposalID(opts *bind.CallOpts, _repository [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getCurrentProposalID", _repository)
	return *ret0, err
}

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(bytes32 _repository) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorSession) GetCurrentProposalID(_repository [32]byte) (*big.Int, error) {
	return _DitCoordinator.Contract.GetCurrentProposalID(&_DitCoordinator.CallOpts, _repository)
}

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(bytes32 _repository) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) GetCurrentProposalID(_repository [32]byte) (*big.Int, error) {
	return _DitCoordinator.Contract.GetCurrentProposalID(&_DitCoordinator.CallOpts, _repository)
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(bytes32 _repository, uint256 _proposalID) constant returns(uint256 individualStake)
func (_DitCoordinator *DitCoordinatorCaller) GetIndividualStake(opts *bind.CallOpts, _repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getIndividualStake", _repository, _proposalID)
	return *ret0, err
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(bytes32 _repository, uint256 _proposalID) constant returns(uint256 individualStake)
func (_DitCoordinator *DitCoordinatorSession) GetIndividualStake(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetIndividualStake(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(bytes32 _repository, uint256 _proposalID) constant returns(uint256 individualStake)
func (_DitCoordinator *DitCoordinatorCallerSession) GetIndividualStake(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetIndividualStake(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(bytes32 _repository, uint256 _proposalID) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCaller) GetKNWVoteIDFromProposalID(opts *bind.CallOpts, _repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getKNWVoteIDFromProposalID", _repository, _proposalID)
	return *ret0, err
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(bytes32 _repository, uint256 _proposalID) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorSession) GetKNWVoteIDFromProposalID(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetKNWVoteIDFromProposalID(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(bytes32 _repository, uint256 _proposalID) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) GetKNWVoteIDFromProposalID(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetKNWVoteIDFromProposalID(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(bytes32 _repository, uint256 _knowledgeLabelID) constant returns(string knowledgeLabel)
func (_DitCoordinator *DitCoordinatorCaller) GetKnowledgeLabels(opts *bind.CallOpts, _repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getKnowledgeLabels", _repository, _knowledgeLabelID)
	return *ret0, err
}

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(bytes32 _repository, uint256 _knowledgeLabelID) constant returns(string knowledgeLabel)
func (_DitCoordinator *DitCoordinatorSession) GetKnowledgeLabels(_repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	return _DitCoordinator.Contract.GetKnowledgeLabels(&_DitCoordinator.CallOpts, _repository, _knowledgeLabelID)
}

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(bytes32 _repository, uint256 _knowledgeLabelID) constant returns(string knowledgeLabel)
func (_DitCoordinator *DitCoordinatorCallerSession) GetKnowledgeLabels(_repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	return _DitCoordinator.Contract.GetKnowledgeLabels(&_DitCoordinator.CallOpts, _repository, _knowledgeLabelID)
}

// IsKYCValidator is a free data retrieval call binding the contract method 0x1341f25c.
//
// Solidity: function isKYCValidator(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCaller) IsKYCValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "isKYCValidator", arg0)
	return *ret0, err
}

// IsKYCValidator is a free data retrieval call binding the contract method 0x1341f25c.
//
// Solidity: function isKYCValidator(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorSession) IsKYCValidator(arg0 common.Address) (bool, error) {
	return _DitCoordinator.Contract.IsKYCValidator(&_DitCoordinator.CallOpts, arg0)
}

// IsKYCValidator is a free data retrieval call binding the contract method 0x1341f25c.
//
// Solidity: function isKYCValidator(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCallerSession) IsKYCValidator(arg0 common.Address) (bool, error) {
	return _DitCoordinator.Contract.IsKYCValidator(&_DitCoordinator.CallOpts, arg0)
}

// PassedKYC is a free data retrieval call binding the contract method 0xccd9aa68.
//
// Solidity: function passedKYC(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCaller) PassedKYC(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "passedKYC", arg0)
	return *ret0, err
}

// PassedKYC is a free data retrieval call binding the contract method 0xccd9aa68.
//
// Solidity: function passedKYC(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorSession) PassedKYC(arg0 common.Address) (bool, error) {
	return _DitCoordinator.Contract.PassedKYC(&_DitCoordinator.CallOpts, arg0)
}

// PassedKYC is a free data retrieval call binding the contract method 0xccd9aa68.
//
// Solidity: function passedKYC(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCallerSession) PassedKYC(arg0 common.Address) (bool, error) {
	return _DitCoordinator.Contract.PassedKYC(&_DitCoordinator.CallOpts, arg0)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(bytes32 _repository, uint256 _proposalID) constant returns(bool hasPassed)
func (_DitCoordinator *DitCoordinatorCaller) ProposalHasPassed(opts *bind.CallOpts, _repository [32]byte, _proposalID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "proposalHasPassed", _repository, _proposalID)
	return *ret0, err
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(bytes32 _repository, uint256 _proposalID) constant returns(bool hasPassed)
func (_DitCoordinator *DitCoordinatorSession) ProposalHasPassed(_repository [32]byte, _proposalID *big.Int) (bool, error) {
	return _DitCoordinator.Contract.ProposalHasPassed(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(bytes32 _repository, uint256 _proposalID) constant returns(bool hasPassed)
func (_DitCoordinator *DitCoordinatorCallerSession) ProposalHasPassed(_repository [32]byte, _proposalID *big.Int) (bool, error) {
	return _DitCoordinator.Contract.ProposalHasPassed(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(uint256 KNWVoteID, string knowledgeLabel, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitCoordinator *DitCoordinatorCaller) ProposalsOfRepository(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
	Proposer         common.Address
	IsFinalized      bool
	ProposalAccepted bool
	IndividualStake  *big.Int
	TotalStake       *big.Int
}, error) {
	ret := new(struct {
		KNWVoteID        *big.Int
		KnowledgeLabel   string
		Proposer         common.Address
		IsFinalized      bool
		ProposalAccepted bool
		IndividualStake  *big.Int
		TotalStake       *big.Int
	})
	out := ret
	err := _DitCoordinator.contract.Call(opts, out, "proposalsOfRepository", arg0, arg1)
	return *ret, err
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(uint256 KNWVoteID, string knowledgeLabel, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitCoordinator *DitCoordinatorSession) ProposalsOfRepository(arg0 [32]byte, arg1 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
	Proposer         common.Address
	IsFinalized      bool
	ProposalAccepted bool
	IndividualStake  *big.Int
	TotalStake       *big.Int
}, error) {
	return _DitCoordinator.Contract.ProposalsOfRepository(&_DitCoordinator.CallOpts, arg0, arg1)
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(uint256 KNWVoteID, string knowledgeLabel, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitCoordinator *DitCoordinatorCallerSession) ProposalsOfRepository(arg0 [32]byte, arg1 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
	Proposer         common.Address
	IsFinalized      bool
	ProposalAccepted bool
	IndividualStake  *big.Int
	TotalStake       *big.Int
}, error) {
	return _DitCoordinator.Contract.ProposalsOfRepository(&_DitCoordinator.CallOpts, arg0, arg1)
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories(bytes32 ) constant returns(uint256 votingMajority, uint256 mintingMethod, uint256 burningMethod, uint256 currentProposalID, uint256 minVoteCommitDuration, uint256 maxVoteCommitDuration, uint256 minVoteOpenDuration, uint256 maxVoteOpenDuration)
func (_DitCoordinator *DitCoordinatorCaller) Repositories(opts *bind.CallOpts, arg0 [32]byte) (struct {
	VotingMajority        *big.Int
	MintingMethod         *big.Int
	BurningMethod         *big.Int
	CurrentProposalID     *big.Int
	MinVoteCommitDuration *big.Int
	MaxVoteCommitDuration *big.Int
	MinVoteOpenDuration   *big.Int
	MaxVoteOpenDuration   *big.Int
}, error) {
	ret := new(struct {
		VotingMajority        *big.Int
		MintingMethod         *big.Int
		BurningMethod         *big.Int
		CurrentProposalID     *big.Int
		MinVoteCommitDuration *big.Int
		MaxVoteCommitDuration *big.Int
		MinVoteOpenDuration   *big.Int
		MaxVoteOpenDuration   *big.Int
	})
	out := ret
	err := _DitCoordinator.contract.Call(opts, out, "repositories", arg0)
	return *ret, err
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories(bytes32 ) constant returns(uint256 votingMajority, uint256 mintingMethod, uint256 burningMethod, uint256 currentProposalID, uint256 minVoteCommitDuration, uint256 maxVoteCommitDuration, uint256 minVoteOpenDuration, uint256 maxVoteOpenDuration)
func (_DitCoordinator *DitCoordinatorSession) Repositories(arg0 [32]byte) (struct {
	VotingMajority        *big.Int
	MintingMethod         *big.Int
	BurningMethod         *big.Int
	CurrentProposalID     *big.Int
	MinVoteCommitDuration *big.Int
	MaxVoteCommitDuration *big.Int
	MinVoteOpenDuration   *big.Int
	MaxVoteOpenDuration   *big.Int
}, error) {
	return _DitCoordinator.Contract.Repositories(&_DitCoordinator.CallOpts, arg0)
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories(bytes32 ) constant returns(uint256 votingMajority, uint256 mintingMethod, uint256 burningMethod, uint256 currentProposalID, uint256 minVoteCommitDuration, uint256 maxVoteCommitDuration, uint256 minVoteOpenDuration, uint256 maxVoteOpenDuration)
func (_DitCoordinator *DitCoordinatorCallerSession) Repositories(arg0 [32]byte) (struct {
	VotingMajority        *big.Int
	MintingMethod         *big.Int
	BurningMethod         *big.Int
	CurrentProposalID     *big.Int
	MinVoteCommitDuration *big.Int
	MaxVoteCommitDuration *big.Int
	MinVoteOpenDuration   *big.Int
	MaxVoteOpenDuration   *big.Int
}, error) {
	return _DitCoordinator.Contract.Repositories(&_DitCoordinator.CallOpts, arg0)
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(bytes32 _repository) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCaller) RepositoryIsInitialized(opts *bind.CallOpts, _repository [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "repositoryIsInitialized", _repository)
	return *ret0, err
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(bytes32 _repository) constant returns(bool)
func (_DitCoordinator *DitCoordinatorSession) RepositoryIsInitialized(_repository [32]byte) (bool, error) {
	return _DitCoordinator.Contract.RepositoryIsInitialized(&_DitCoordinator.CallOpts, _repository)
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(bytes32 _repository) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCallerSession) RepositoryIsInitialized(_repository [32]byte) (bool, error) {
	return _DitCoordinator.Contract.RepositoryIsInitialized(&_DitCoordinator.CallOpts, _repository)
}

// AddKYCValidator is a paid mutator transaction binding the contract method 0xd0c397ef.
//
// Solidity: function addKYCValidator(address _address) returns()
func (_DitCoordinator *DitCoordinatorTransactor) AddKYCValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "addKYCValidator", _address)
}

// AddKYCValidator is a paid mutator transaction binding the contract method 0xd0c397ef.
//
// Solidity: function addKYCValidator(address _address) returns()
func (_DitCoordinator *DitCoordinatorSession) AddKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.AddKYCValidator(&_DitCoordinator.TransactOpts, _address)
}

// AddKYCValidator is a paid mutator transaction binding the contract method 0xd0c397ef.
//
// Solidity: function addKYCValidator(address _address) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) AddKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.AddKYCValidator(&_DitCoordinator.TransactOpts, _address)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(bytes32 _repository, uint256 _proposalID) returns()
func (_DitCoordinator *DitCoordinatorTransactor) FinalizeVote(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "finalizeVote", _repository, _proposalID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(bytes32 _repository, uint256 _proposalID) returns()
func (_DitCoordinator *DitCoordinatorSession) FinalizeVote(_repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.FinalizeVote(&_DitCoordinator.TransactOpts, _repository, _proposalID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(bytes32 _repository, uint256 _proposalID) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) FinalizeVote(_repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.FinalizeVote(&_DitCoordinator.TransactOpts, _repository, _proposalID)
}

// InitRepository is a paid mutator transaction binding the contract method 0x51f43c24.
//
// Solidity: function initRepository(bytes32 _repository, string _label1, string _label2, string _label3, uint256[7] _voteSettings) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) InitRepository(opts *bind.TransactOpts, _repository [32]byte, _label1 string, _label2 string, _label3 string, _voteSettings [7]*big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "initRepository", _repository, _label1, _label2, _label3, _voteSettings)
}

// InitRepository is a paid mutator transaction binding the contract method 0x51f43c24.
//
// Solidity: function initRepository(bytes32 _repository, string _label1, string _label2, string _label3, uint256[7] _voteSettings) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) InitRepository(_repository [32]byte, _label1 string, _label2 string, _label3 string, _voteSettings [7]*big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.InitRepository(&_DitCoordinator.TransactOpts, _repository, _label1, _label2, _label3, _voteSettings)
}

// InitRepository is a paid mutator transaction binding the contract method 0x51f43c24.
//
// Solidity: function initRepository(bytes32 _repository, string _label1, string _label2, string _label3, uint256[7] _voteSettings) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) InitRepository(_repository [32]byte, _label1 string, _label2 string, _label3 string, _voteSettings [7]*big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.InitRepository(&_DitCoordinator.TransactOpts, _repository, _label1, _label2, _label3, _voteSettings)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(bytes32 _repository, uint256 _proposalID, uint256 _voteOption, uint256 _voteSalt) returns()
func (_DitCoordinator *DitCoordinatorTransactor) OpenVoteOnProposal(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "openVoteOnProposal", _repository, _proposalID, _voteOption, _voteSalt)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(bytes32 _repository, uint256 _proposalID, uint256 _voteOption, uint256 _voteSalt) returns()
func (_DitCoordinator *DitCoordinatorSession) OpenVoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.OpenVoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteOption, _voteSalt)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(bytes32 _repository, uint256 _proposalID, uint256 _voteOption, uint256 _voteSalt) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) OpenVoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.OpenVoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteOption, _voteSalt)
}

// PassKYC is a paid mutator transaction binding the contract method 0xeb931024.
//
// Solidity: function passKYC(address _address) returns()
func (_DitCoordinator *DitCoordinatorTransactor) PassKYC(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "passKYC", _address)
}

// PassKYC is a paid mutator transaction binding the contract method 0xeb931024.
//
// Solidity: function passKYC(address _address) returns()
func (_DitCoordinator *DitCoordinatorSession) PassKYC(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.PassKYC(&_DitCoordinator.TransactOpts, _address)
}

// PassKYC is a paid mutator transaction binding the contract method 0xeb931024.
//
// Solidity: function passKYC(address _address) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) PassKYC(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.PassKYC(&_DitCoordinator.TransactOpts, _address)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x0aba8688.
//
// Solidity: function proposeCommit(bytes32 _repository, uint256 _knowledgeLabelIndex, uint256 _voteCommitDuration, uint256 _voteOpenDuration) returns()
func (_DitCoordinator *DitCoordinatorTransactor) ProposeCommit(opts *bind.TransactOpts, _repository [32]byte, _knowledgeLabelIndex *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "proposeCommit", _repository, _knowledgeLabelIndex, _voteCommitDuration, _voteOpenDuration)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x0aba8688.
//
// Solidity: function proposeCommit(bytes32 _repository, uint256 _knowledgeLabelIndex, uint256 _voteCommitDuration, uint256 _voteOpenDuration) returns()
func (_DitCoordinator *DitCoordinatorSession) ProposeCommit(_repository [32]byte, _knowledgeLabelIndex *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ProposeCommit(&_DitCoordinator.TransactOpts, _repository, _knowledgeLabelIndex, _voteCommitDuration, _voteOpenDuration)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x0aba8688.
//
// Solidity: function proposeCommit(bytes32 _repository, uint256 _knowledgeLabelIndex, uint256 _voteCommitDuration, uint256 _voteOpenDuration) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) ProposeCommit(_repository [32]byte, _knowledgeLabelIndex *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ProposeCommit(&_DitCoordinator.TransactOpts, _repository, _knowledgeLabelIndex, _voteCommitDuration, _voteOpenDuration)
}

// RemoveKYCValidator is a paid mutator transaction binding the contract method 0x73b0dddd.
//
// Solidity: function removeKYCValidator(address _address) returns()
func (_DitCoordinator *DitCoordinatorTransactor) RemoveKYCValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "removeKYCValidator", _address)
}

// RemoveKYCValidator is a paid mutator transaction binding the contract method 0x73b0dddd.
//
// Solidity: function removeKYCValidator(address _address) returns()
func (_DitCoordinator *DitCoordinatorSession) RemoveKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.RemoveKYCValidator(&_DitCoordinator.TransactOpts, _address)
}

// RemoveKYCValidator is a paid mutator transaction binding the contract method 0x73b0dddd.
//
// Solidity: function removeKYCValidator(address _address) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) RemoveKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.RemoveKYCValidator(&_DitCoordinator.TransactOpts, _address)
}

// RevokeKYC is a paid mutator transaction binding the contract method 0x39ba645b.
//
// Solidity: function revokeKYC(address _address) returns()
func (_DitCoordinator *DitCoordinatorTransactor) RevokeKYC(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "revokeKYC", _address)
}

// RevokeKYC is a paid mutator transaction binding the contract method 0x39ba645b.
//
// Solidity: function revokeKYC(address _address) returns()
func (_DitCoordinator *DitCoordinatorSession) RevokeKYC(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.RevokeKYC(&_DitCoordinator.TransactOpts, _address)
}

// RevokeKYC is a paid mutator transaction binding the contract method 0x39ba645b.
//
// Solidity: function revokeKYC(address _address) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) RevokeKYC(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.RevokeKYC(&_DitCoordinator.TransactOpts, _address)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xa34c299a.
//
// Solidity: function voteOnProposal(bytes32 _repository, uint256 _proposalID, bytes32 _voteHash) returns()
func (_DitCoordinator *DitCoordinatorTransactor) VoteOnProposal(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "voteOnProposal", _repository, _proposalID, _voteHash)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xa34c299a.
//
// Solidity: function voteOnProposal(bytes32 _repository, uint256 _proposalID, bytes32 _voteHash) returns()
func (_DitCoordinator *DitCoordinatorSession) VoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitCoordinator.Contract.VoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteHash)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xa34c299a.
//
// Solidity: function voteOnProposal(bytes32 _repository, uint256 _proposalID, bytes32 _voteHash) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) VoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitCoordinator.Contract.VoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteHash)
}

// DitCoordinatorCommitVoteIterator is returned from FilterCommitVote and is used to iterate over the raw logs and unpacked data for CommitVote events raised by the DitCoordinator contract.
type DitCoordinatorCommitVoteIterator struct {
	Event *DitCoordinatorCommitVote // Event containing the contract specifics and raw log

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
func (it *DitCoordinatorCommitVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitCoordinatorCommitVote)
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
		it.Event = new(DitCoordinatorCommitVote)
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
func (it *DitCoordinatorCommitVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitCoordinatorCommitVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitCoordinatorCommitVote represents a CommitVote event raised by the DitCoordinator contract.
type DitCoordinatorCommitVote struct {
	Repository    [32]byte
	Proposal      *big.Int
	Who           common.Address
	Label         string
	Stake         *big.Int
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitVote is a free log retrieval operation binding the contract event 0xa01eea487bb3ec75528c167ccf90452d4164ddda7b13c55b2a89751a8dc5fbc1.
//
// Solidity: event CommitVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, uint256 stake, uint256 numberOfVotes)
func (_DitCoordinator *DitCoordinatorFilterer) FilterCommitVote(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitCoordinatorCommitVoteIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "CommitVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorCommitVoteIterator{contract: _DitCoordinator.contract, event: "CommitVote", logs: logs, sub: sub}, nil
}

// WatchCommitVote is a free log subscription operation binding the contract event 0xa01eea487bb3ec75528c167ccf90452d4164ddda7b13c55b2a89751a8dc5fbc1.
//
// Solidity: event CommitVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, uint256 stake, uint256 numberOfVotes)
func (_DitCoordinator *DitCoordinatorFilterer) WatchCommitVote(opts *bind.WatchOpts, sink chan<- *DitCoordinatorCommitVote, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "CommitVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitCoordinatorCommitVote)
				if err := _DitCoordinator.contract.UnpackLog(event, "CommitVote", log); err != nil {
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

// DitCoordinatorFinalizeVoteIterator is returned from FilterFinalizeVote and is used to iterate over the raw logs and unpacked data for FinalizeVote events raised by the DitCoordinator contract.
type DitCoordinatorFinalizeVoteIterator struct {
	Event *DitCoordinatorFinalizeVote // Event containing the contract specifics and raw log

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
func (it *DitCoordinatorFinalizeVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitCoordinatorFinalizeVote)
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
		it.Event = new(DitCoordinatorFinalizeVote)
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
func (it *DitCoordinatorFinalizeVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitCoordinatorFinalizeVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitCoordinatorFinalizeVote represents a FinalizeVote event raised by the DitCoordinator contract.
type DitCoordinatorFinalizeVote struct {
	Repository [32]byte
	Proposal   *big.Int
	Label      string
	Accepted   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinalizeVote is a free log retrieval operation binding the contract event 0x6bd2699645e0f6c5547bdf0d053280e48fef1ab21514bd02c88610b1279b942a.
//
// Solidity: event FinalizeVote(bytes32 indexed repository, uint256 indexed proposal, string label, bool accepted)
func (_DitCoordinator *DitCoordinatorFilterer) FilterFinalizeVote(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int) (*DitCoordinatorFinalizeVoteIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "FinalizeVote", repositoryRule, proposalRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorFinalizeVoteIterator{contract: _DitCoordinator.contract, event: "FinalizeVote", logs: logs, sub: sub}, nil
}

// WatchFinalizeVote is a free log subscription operation binding the contract event 0x6bd2699645e0f6c5547bdf0d053280e48fef1ab21514bd02c88610b1279b942a.
//
// Solidity: event FinalizeVote(bytes32 indexed repository, uint256 indexed proposal, string label, bool accepted)
func (_DitCoordinator *DitCoordinatorFilterer) WatchFinalizeVote(opts *bind.WatchOpts, sink chan<- *DitCoordinatorFinalizeVote, repository [][32]byte, proposal []*big.Int) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "FinalizeVote", repositoryRule, proposalRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitCoordinatorFinalizeVote)
				if err := _DitCoordinator.contract.UnpackLog(event, "FinalizeVote", log); err != nil {
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

// DitCoordinatorOpenVoteIterator is returned from FilterOpenVote and is used to iterate over the raw logs and unpacked data for OpenVote events raised by the DitCoordinator contract.
type DitCoordinatorOpenVoteIterator struct {
	Event *DitCoordinatorOpenVote // Event containing the contract specifics and raw log

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
func (it *DitCoordinatorOpenVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitCoordinatorOpenVote)
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
		it.Event = new(DitCoordinatorOpenVote)
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
func (it *DitCoordinatorOpenVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitCoordinatorOpenVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitCoordinatorOpenVote represents a OpenVote event raised by the DitCoordinator contract.
type DitCoordinatorOpenVote struct {
	Repository    [32]byte
	Proposal      *big.Int
	Who           common.Address
	Label         string
	Accept        bool
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenVote is a free log retrieval operation binding the contract event 0x864c0d6987266fd72e7e37f1fbc98b6a3794b7187dae454c67a2a626628a72ab.
//
// Solidity: event OpenVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, bool accept, uint256 numberOfVotes)
func (_DitCoordinator *DitCoordinatorFilterer) FilterOpenVote(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitCoordinatorOpenVoteIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "OpenVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorOpenVoteIterator{contract: _DitCoordinator.contract, event: "OpenVote", logs: logs, sub: sub}, nil
}

// WatchOpenVote is a free log subscription operation binding the contract event 0x864c0d6987266fd72e7e37f1fbc98b6a3794b7187dae454c67a2a626628a72ab.
//
// Solidity: event OpenVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, bool accept, uint256 numberOfVotes)
func (_DitCoordinator *DitCoordinatorFilterer) WatchOpenVote(opts *bind.WatchOpts, sink chan<- *DitCoordinatorOpenVote, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "OpenVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitCoordinatorOpenVote)
				if err := _DitCoordinator.contract.UnpackLog(event, "OpenVote", log); err != nil {
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

// DitCoordinatorProposeCommitIterator is returned from FilterProposeCommit and is used to iterate over the raw logs and unpacked data for ProposeCommit events raised by the DitCoordinator contract.
type DitCoordinatorProposeCommitIterator struct {
	Event *DitCoordinatorProposeCommit // Event containing the contract specifics and raw log

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
func (it *DitCoordinatorProposeCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitCoordinatorProposeCommit)
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
		it.Event = new(DitCoordinatorProposeCommit)
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
func (it *DitCoordinatorProposeCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitCoordinatorProposeCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitCoordinatorProposeCommit represents a ProposeCommit event raised by the DitCoordinator contract.
type DitCoordinatorProposeCommit struct {
	Repository [32]byte
	Proposal   *big.Int
	Who        common.Address
	Label      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposeCommit is a free log retrieval operation binding the contract event 0x171fe77c3addce776991159eb3eb73b14d9187ebd06c1c34ea12355a84ddbd83.
//
// Solidity: event ProposeCommit(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label)
func (_DitCoordinator *DitCoordinatorFilterer) FilterProposeCommit(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitCoordinatorProposeCommitIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "ProposeCommit", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorProposeCommitIterator{contract: _DitCoordinator.contract, event: "ProposeCommit", logs: logs, sub: sub}, nil
}

// WatchProposeCommit is a free log subscription operation binding the contract event 0x171fe77c3addce776991159eb3eb73b14d9187ebd06c1c34ea12355a84ddbd83.
//
// Solidity: event ProposeCommit(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label)
func (_DitCoordinator *DitCoordinatorFilterer) WatchProposeCommit(opts *bind.WatchOpts, sink chan<- *DitCoordinatorProposeCommit, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "ProposeCommit", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitCoordinatorProposeCommit)
				if err := _DitCoordinator.contract.UnpackLog(event, "ProposeCommit", log); err != nil {
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
