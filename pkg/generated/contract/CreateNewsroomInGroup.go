// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// CreateNewsroomInGroupContractABI is the input ABI used to generate the binding from.
const CreateNewsroomInGroupContractABI = "[{\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\"},{\"name\":\"_controller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"charterUri\",\"type\":\"string\"},{\"name\":\"charterHash\",\"type\":\"bytes32\"},{\"name\":\"initialOwners\",\"type\":\"address[]\"},{\"name\":\"initialRequired\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[{\"name\":\"newsroom\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CreateNewsroomInGroupContractBin is the compiled bytecode used for deploying new contracts.
const CreateNewsroomInGroupContractBin = `0x608060405234801561001057600080fd5b5060405160408061056383398101604052805160209091015160008054600160a060020a03938416600160a060020a031991821617909155600180549390921692169190911790556104fc806100676000396000f3006080604052600436106100405763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166360bb9dab8114610045575b600080fd5b34801561005157600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261011e94369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020808901358a01803580830284810184018652818552999c8b359c909b909a95019850929650810194509092508291908501908490808284375094975050933594506101479350505050565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60008083511115156101e057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f696e697469616c4f776e657273206d7573742068617665206174206c6561737460448201527f206f6e65206d656d626572000000000000000000000000000000000000000000606482015290519081900360840190fd5b600080546040517f60bb9dab000000000000000000000000000000000000000000000000000000008152604481018790526084810185905260a060048201908152895160a4830152895173ffffffffffffffffffffffffffffffffffffffff909316936360bb9dab938b938b938b938b938b9383926024820192606483019260c4019160208c0191908190849084905b83811015610288578181015183820152602001610270565b50505050905090810190601f1680156102b55780820380516001836020036101000a031916815260200191505b5084810383528851815288516020918201918a019080838360005b838110156102e85781810151838201526020016102d0565b50505050905090810190601f1680156103155780820380516001836020036101000a031916815260200191505b508481038252865181528651602091820191808901910280838360005b8381101561034a578181015183820152602001610332565b5050505090500198505050505050505050602060405180830381600087803b15801561037557600080fd5b505af1158015610389573d6000803e3d6000fd5b505050506040513d602081101561039f57600080fd5b5051600154604080517f8da5cb5b000000000000000000000000000000000000000000000000000000008152905192935073ffffffffffffffffffffffffffffffffffffffff9182169263607c60bb92851691638da5cb5b9160048083019260209291908290030181600087803b15801561041957600080fd5b505af115801561042d573d6000803e3d6000fd5b505050506040513d602081101561044357600080fd5b5051604080517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815273ffffffffffffffffffffffffffffffffffffffff909216600483015251602480830192600092919082900301818387803b1580156104af57600080fd5b505af11580156104c3573d6000803e3d6000fd5b50505050959450505050505600a165627a7a723058203579613bb9a745d08e565c3c4d04c1fafbfcaacdb3a495949c47f1e1e013f3670029`

// DeployCreateNewsroomInGroupContract deploys a new Ethereum contract, binding an instance of CreateNewsroomInGroupContract to it.
func DeployCreateNewsroomInGroupContract(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _controller common.Address) (common.Address, *types.Transaction, *CreateNewsroomInGroupContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreateNewsroomInGroupContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CreateNewsroomInGroupContractBin), backend, _factory, _controller)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CreateNewsroomInGroupContract{CreateNewsroomInGroupContractCaller: CreateNewsroomInGroupContractCaller{contract: contract}, CreateNewsroomInGroupContractTransactor: CreateNewsroomInGroupContractTransactor{contract: contract}, CreateNewsroomInGroupContractFilterer: CreateNewsroomInGroupContractFilterer{contract: contract}}, nil
}

// CreateNewsroomInGroupContract is an auto generated Go binding around an Ethereum contract.
type CreateNewsroomInGroupContract struct {
	CreateNewsroomInGroupContractCaller     // Read-only binding to the contract
	CreateNewsroomInGroupContractTransactor // Write-only binding to the contract
	CreateNewsroomInGroupContractFilterer   // Log filterer for contract events
}

// CreateNewsroomInGroupContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreateNewsroomInGroupContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateNewsroomInGroupContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreateNewsroomInGroupContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateNewsroomInGroupContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreateNewsroomInGroupContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateNewsroomInGroupContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreateNewsroomInGroupContractSession struct {
	Contract     *CreateNewsroomInGroupContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CreateNewsroomInGroupContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreateNewsroomInGroupContractCallerSession struct {
	Contract *CreateNewsroomInGroupContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// CreateNewsroomInGroupContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreateNewsroomInGroupContractTransactorSession struct {
	Contract     *CreateNewsroomInGroupContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// CreateNewsroomInGroupContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreateNewsroomInGroupContractRaw struct {
	Contract *CreateNewsroomInGroupContract // Generic contract binding to access the raw methods on
}

// CreateNewsroomInGroupContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreateNewsroomInGroupContractCallerRaw struct {
	Contract *CreateNewsroomInGroupContractCaller // Generic read-only contract binding to access the raw methods on
}

// CreateNewsroomInGroupContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreateNewsroomInGroupContractTransactorRaw struct {
	Contract *CreateNewsroomInGroupContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreateNewsroomInGroupContract creates a new instance of CreateNewsroomInGroupContract, bound to a specific deployed contract.
func NewCreateNewsroomInGroupContract(address common.Address, backend bind.ContractBackend) (*CreateNewsroomInGroupContract, error) {
	contract, err := bindCreateNewsroomInGroupContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreateNewsroomInGroupContract{CreateNewsroomInGroupContractCaller: CreateNewsroomInGroupContractCaller{contract: contract}, CreateNewsroomInGroupContractTransactor: CreateNewsroomInGroupContractTransactor{contract: contract}, CreateNewsroomInGroupContractFilterer: CreateNewsroomInGroupContractFilterer{contract: contract}}, nil
}

// NewCreateNewsroomInGroupContractCaller creates a new read-only instance of CreateNewsroomInGroupContract, bound to a specific deployed contract.
func NewCreateNewsroomInGroupContractCaller(address common.Address, caller bind.ContractCaller) (*CreateNewsroomInGroupContractCaller, error) {
	contract, err := bindCreateNewsroomInGroupContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreateNewsroomInGroupContractCaller{contract: contract}, nil
}

// NewCreateNewsroomInGroupContractTransactor creates a new write-only instance of CreateNewsroomInGroupContract, bound to a specific deployed contract.
func NewCreateNewsroomInGroupContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CreateNewsroomInGroupContractTransactor, error) {
	contract, err := bindCreateNewsroomInGroupContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreateNewsroomInGroupContractTransactor{contract: contract}, nil
}

// NewCreateNewsroomInGroupContractFilterer creates a new log filterer instance of CreateNewsroomInGroupContract, bound to a specific deployed contract.
func NewCreateNewsroomInGroupContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CreateNewsroomInGroupContractFilterer, error) {
	contract, err := bindCreateNewsroomInGroupContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreateNewsroomInGroupContractFilterer{contract: contract}, nil
}

// bindCreateNewsroomInGroupContract binds a generic wrapper to an already deployed contract.
func bindCreateNewsroomInGroupContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreateNewsroomInGroupContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreateNewsroomInGroupContract *CreateNewsroomInGroupContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CreateNewsroomInGroupContract.Contract.CreateNewsroomInGroupContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreateNewsroomInGroupContract *CreateNewsroomInGroupContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreateNewsroomInGroupContract.Contract.CreateNewsroomInGroupContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreateNewsroomInGroupContract *CreateNewsroomInGroupContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreateNewsroomInGroupContract.Contract.CreateNewsroomInGroupContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreateNewsroomInGroupContract *CreateNewsroomInGroupContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CreateNewsroomInGroupContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreateNewsroomInGroupContract *CreateNewsroomInGroupContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreateNewsroomInGroupContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreateNewsroomInGroupContract *CreateNewsroomInGroupContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreateNewsroomInGroupContract.Contract.contract.Transact(opts, method, params...)
}

// Create is a paid mutator transaction binding the contract method 0x60bb9dab.
//
// Solidity: function create(string name, string charterUri, bytes32 charterHash, address[] initialOwners, uint256 initialRequired) returns(address newsroom)
func (_CreateNewsroomInGroupContract *CreateNewsroomInGroupContractTransactor) Create(opts *bind.TransactOpts, name string, charterUri string, charterHash [32]byte, initialOwners []common.Address, initialRequired *big.Int) (*types.Transaction, error) {
	return _CreateNewsroomInGroupContract.contract.Transact(opts, "create", name, charterUri, charterHash, initialOwners, initialRequired)
}

// Create is a paid mutator transaction binding the contract method 0x60bb9dab.
//
// Solidity: function create(string name, string charterUri, bytes32 charterHash, address[] initialOwners, uint256 initialRequired) returns(address newsroom)
func (_CreateNewsroomInGroupContract *CreateNewsroomInGroupContractSession) Create(name string, charterUri string, charterHash [32]byte, initialOwners []common.Address, initialRequired *big.Int) (*types.Transaction, error) {
	return _CreateNewsroomInGroupContract.Contract.Create(&_CreateNewsroomInGroupContract.TransactOpts, name, charterUri, charterHash, initialOwners, initialRequired)
}

// Create is a paid mutator transaction binding the contract method 0x60bb9dab.
//
// Solidity: function create(string name, string charterUri, bytes32 charterHash, address[] initialOwners, uint256 initialRequired) returns(address newsroom)
func (_CreateNewsroomInGroupContract *CreateNewsroomInGroupContractTransactorSession) Create(name string, charterUri string, charterHash [32]byte, initialOwners []common.Address, initialRequired *big.Int) (*types.Transaction, error) {
	return _CreateNewsroomInGroupContract.Contract.Create(&_CreateNewsroomInGroupContract.TransactOpts, name, charterUri, charterHash, initialOwners, initialRequired)
}
