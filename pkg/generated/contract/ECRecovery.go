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

// ECRecoveryContractABI is the input ABI used to generate the binding from.
const ECRecoveryContractABI = "[]"

// ECRecoveryContractBin is the compiled bytecode used for deploying new contracts.
const ECRecoveryContractBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582070c431c61d617c0edfdeadd10b17167ea6948ea600cbbf91d806acec056c50920029`

// DeployECRecoveryContract deploys a new Ethereum contract, binding an instance of ECRecoveryContract to it.
func DeployECRecoveryContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECRecoveryContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECRecoveryContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECRecoveryContract{ECRecoveryContractCaller: ECRecoveryContractCaller{contract: contract}, ECRecoveryContractTransactor: ECRecoveryContractTransactor{contract: contract}, ECRecoveryContractFilterer: ECRecoveryContractFilterer{contract: contract}}, nil
}

// ECRecoveryContract is an auto generated Go binding around an Ethereum contract.
type ECRecoveryContract struct {
	ECRecoveryContractCaller     // Read-only binding to the contract
	ECRecoveryContractTransactor // Write-only binding to the contract
	ECRecoveryContractFilterer   // Log filterer for contract events
}

// ECRecoveryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoveryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoveryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECRecoveryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECRecoveryContractSession struct {
	Contract     *ECRecoveryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ECRecoveryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECRecoveryContractCallerSession struct {
	Contract *ECRecoveryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ECRecoveryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECRecoveryContractTransactorSession struct {
	Contract     *ECRecoveryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ECRecoveryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECRecoveryContractRaw struct {
	Contract *ECRecoveryContract // Generic contract binding to access the raw methods on
}

// ECRecoveryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECRecoveryContractCallerRaw struct {
	Contract *ECRecoveryContractCaller // Generic read-only contract binding to access the raw methods on
}

// ECRecoveryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECRecoveryContractTransactorRaw struct {
	Contract *ECRecoveryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECRecoveryContract creates a new instance of ECRecoveryContract, bound to a specific deployed contract.
func NewECRecoveryContract(address common.Address, backend bind.ContractBackend) (*ECRecoveryContract, error) {
	contract, err := bindECRecoveryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryContract{ECRecoveryContractCaller: ECRecoveryContractCaller{contract: contract}, ECRecoveryContractTransactor: ECRecoveryContractTransactor{contract: contract}, ECRecoveryContractFilterer: ECRecoveryContractFilterer{contract: contract}}, nil
}

// NewECRecoveryContractCaller creates a new read-only instance of ECRecoveryContract, bound to a specific deployed contract.
func NewECRecoveryContractCaller(address common.Address, caller bind.ContractCaller) (*ECRecoveryContractCaller, error) {
	contract, err := bindECRecoveryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryContractCaller{contract: contract}, nil
}

// NewECRecoveryContractTransactor creates a new write-only instance of ECRecoveryContract, bound to a specific deployed contract.
func NewECRecoveryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoveryContractTransactor, error) {
	contract, err := bindECRecoveryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryContractTransactor{contract: contract}, nil
}

// NewECRecoveryContractFilterer creates a new log filterer instance of ECRecoveryContract, bound to a specific deployed contract.
func NewECRecoveryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ECRecoveryContractFilterer, error) {
	contract, err := bindECRecoveryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryContractFilterer{contract: contract}, nil
}

// bindECRecoveryContract binds a generic wrapper to an already deployed contract.
func bindECRecoveryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecoveryContract *ECRecoveryContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecoveryContract.Contract.ECRecoveryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecoveryContract *ECRecoveryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecoveryContract.Contract.ECRecoveryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecoveryContract *ECRecoveryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecoveryContract.Contract.ECRecoveryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecoveryContract *ECRecoveryContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecoveryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecoveryContract *ECRecoveryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecoveryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecoveryContract *ECRecoveryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecoveryContract.Contract.contract.Transact(opts, method, params...)
}
