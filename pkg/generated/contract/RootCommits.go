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

// RootCommitsContractABI is the input ABI used to generate the binding from.
const RootCommitsContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blockN\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"setRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_blockN\",\"type\":\"uint64\"}],\"name\":\"getRootByBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint64\"}],\"name\":\"getRootByTime\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RootCommitsContractBin is the compiled bytecode used for deploying new contracts.
const RootCommitsContractBin = `0x608060405234801561001057600080fd5b50610869806100206000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663079cf76e811461006657806318c9990b14610099578063ab481d23146100c7578063dab5f340146100f5575b600080fd5b34801561007257600080fd5b50610087600160a060020a036004351661010f565b60408051918252519081900360200190f35b3480156100a557600080fd5b50610087600160a060020a036004351667ffffffffffffffff60243516610150565b3480156100d357600080fd5b50610087600160a060020a036004351667ffffffffffffffff602435166103b5565b34801561010157600080fd5b5061010d600435610673565b005b600160a060020a03811660009081526020819052604081208054600019810190811061013757fe5b9060005260206000209060020201600101549050919050565b60008080804367ffffffffffffffff8616106101cd57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6572724e6f467574757265416c6c6f7765640000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0386166000908152602081905260408120549093506000190191505b8183116103ac5750600160a060020a038516600090815260208190526040902080546002848401049167ffffffffffffffff8716918390811061022f57fe5b600091825260209091206002909102015467ffffffffffffffff16141561028d57600160a060020a038616600090815260208190526040902080548290811061027457fe5b90600052602060002090600202016001015493506103ac565b600160a060020a03861660009081526020819052604090208054829081106102b157fe5b600091825260209091206002909102015467ffffffffffffffff9081169086161180156103225750600160a060020a038616600090815260208190526040902080546001830190811061030057fe5b600091825260209091206002909102015467ffffffffffffffff908116908616105b1561034b57600160a060020a038616600090815260208190526040902080548290811061027457fe5b600160a060020a038616600090815260208190526040902080548290811061036f57fe5b600091825260209091206002909102015467ffffffffffffffff90811690861611156103a0578060010192506103a7565b6001810391505b6101f0565b50505092915050565b6000808080804267ffffffffffffffff87161061043357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6572724e6f467574757265416c6c6f7765640000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038716600090815260208190526040902054151561045a57839450610669565b600160a060020a0387166000908152602081905260408120549093506000190191505b8183116106695750600160a060020a038616600090815260208190526040902080546002848401049167ffffffffffffffff881691839081106104bc57fe5b600091825260209091206002909102015468010000000000000000900467ffffffffffffffff16141561052657600160a060020a038716600090815260208190526040902080548290811061050d57fe5b9060005260206000209060020201600101549450610669565b600160a060020a038716600090815260208190526040902080548290811061054a57fe5b600091825260209091206002909102015467ffffffffffffffff6801000000000000000090910481169087161180156105d35750600160a060020a03871660009081526020819052604090208054600183019081106105a557fe5b600091825260209091206002909102015467ffffffffffffffff680100000000000000009091048116908716105b156105fc57600160a060020a038716600090815260208190526040902080548290811061050d57fe5b600160a060020a038716600090815260208190526040902080548290811061062057fe5b600091825260209091206002909102015467ffffffffffffffff680100000000000000009091048116908716111561065d57806001019250610664565b6001810391505b61047d565b5050505092915050565b33600090815260208190526040812054111561075957336000908152602081905260409020805443919060001981019081106106ab57fe5b600091825260209091206002909102015467ffffffffffffffff16141561075957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f6e6f206d756c7469706c652073657420696e207468652073616d6520626c6f6360448201527f6b00000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b3360008181526020818152604080832081516060818101845267ffffffffffffffff4381168084524282168488018181528588018c8152875460018181018a55988c529a8a902096516002909b0290960180549151851668010000000000000000026fffffffffffffffff0000000000000000199b90951667ffffffffffffffff1990921691909117999099169290921788559251969093019590955582519586529285019290925283810191909152908201839052517f1d3d3c252bc04ff83e8069c09279b4d8acb5264996a680b6d7dcf20db2b7417b9181900360800190a1505600a165627a7a723058203496765d6cc454f9da18ab87fed00841e3f92390fe7c8090281faea5f964f5bc0029`

// DeployRootCommitsContract deploys a new Ethereum contract, binding an instance of RootCommitsContract to it.
func DeployRootCommitsContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootCommitsContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootCommitsContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootCommitsContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootCommitsContract{RootCommitsContractCaller: RootCommitsContractCaller{contract: contract}, RootCommitsContractTransactor: RootCommitsContractTransactor{contract: contract}, RootCommitsContractFilterer: RootCommitsContractFilterer{contract: contract}}, nil
}

// RootCommitsContract is an auto generated Go binding around an Ethereum contract.
type RootCommitsContract struct {
	RootCommitsContractCaller     // Read-only binding to the contract
	RootCommitsContractTransactor // Write-only binding to the contract
	RootCommitsContractFilterer   // Log filterer for contract events
}

// RootCommitsContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootCommitsContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootCommitsContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootCommitsContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootCommitsContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootCommitsContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootCommitsContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootCommitsContractSession struct {
	Contract     *RootCommitsContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootCommitsContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootCommitsContractCallerSession struct {
	Contract *RootCommitsContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RootCommitsContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootCommitsContractTransactorSession struct {
	Contract     *RootCommitsContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RootCommitsContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootCommitsContractRaw struct {
	Contract *RootCommitsContract // Generic contract binding to access the raw methods on
}

// RootCommitsContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootCommitsContractCallerRaw struct {
	Contract *RootCommitsContractCaller // Generic read-only contract binding to access the raw methods on
}

// RootCommitsContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootCommitsContractTransactorRaw struct {
	Contract *RootCommitsContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootCommitsContract creates a new instance of RootCommitsContract, bound to a specific deployed contract.
func NewRootCommitsContract(address common.Address, backend bind.ContractBackend) (*RootCommitsContract, error) {
	contract, err := bindRootCommitsContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootCommitsContract{RootCommitsContractCaller: RootCommitsContractCaller{contract: contract}, RootCommitsContractTransactor: RootCommitsContractTransactor{contract: contract}, RootCommitsContractFilterer: RootCommitsContractFilterer{contract: contract}}, nil
}

// NewRootCommitsContractCaller creates a new read-only instance of RootCommitsContract, bound to a specific deployed contract.
func NewRootCommitsContractCaller(address common.Address, caller bind.ContractCaller) (*RootCommitsContractCaller, error) {
	contract, err := bindRootCommitsContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootCommitsContractCaller{contract: contract}, nil
}

// NewRootCommitsContractTransactor creates a new write-only instance of RootCommitsContract, bound to a specific deployed contract.
func NewRootCommitsContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RootCommitsContractTransactor, error) {
	contract, err := bindRootCommitsContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootCommitsContractTransactor{contract: contract}, nil
}

// NewRootCommitsContractFilterer creates a new log filterer instance of RootCommitsContract, bound to a specific deployed contract.
func NewRootCommitsContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RootCommitsContractFilterer, error) {
	contract, err := bindRootCommitsContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootCommitsContractFilterer{contract: contract}, nil
}

// bindRootCommitsContract binds a generic wrapper to an already deployed contract.
func bindRootCommitsContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootCommitsContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootCommitsContract *RootCommitsContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootCommitsContract.Contract.RootCommitsContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootCommitsContract *RootCommitsContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootCommitsContract.Contract.RootCommitsContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootCommitsContract *RootCommitsContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootCommitsContract.Contract.RootCommitsContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootCommitsContract *RootCommitsContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootCommitsContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootCommitsContract *RootCommitsContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootCommitsContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootCommitsContract *RootCommitsContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootCommitsContract.Contract.contract.Transact(opts, method, params...)
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address _address) constant returns(bytes32)
func (_RootCommitsContract *RootCommitsContractCaller) GetRoot(opts *bind.CallOpts, _address common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootCommitsContract.contract.Call(opts, out, "getRoot", _address)
	return *ret0, err
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address _address) constant returns(bytes32)
func (_RootCommitsContract *RootCommitsContractSession) GetRoot(_address common.Address) ([32]byte, error) {
	return _RootCommitsContract.Contract.GetRoot(&_RootCommitsContract.CallOpts, _address)
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address _address) constant returns(bytes32)
func (_RootCommitsContract *RootCommitsContractCallerSession) GetRoot(_address common.Address) ([32]byte, error) {
	return _RootCommitsContract.Contract.GetRoot(&_RootCommitsContract.CallOpts, _address)
}

// GetRootByBlock is a free data retrieval call binding the contract method 0x18c9990b.
//
// Solidity: function getRootByBlock(address _address, uint64 _blockN) constant returns(bytes32)
func (_RootCommitsContract *RootCommitsContractCaller) GetRootByBlock(opts *bind.CallOpts, _address common.Address, _blockN uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootCommitsContract.contract.Call(opts, out, "getRootByBlock", _address, _blockN)
	return *ret0, err
}

// GetRootByBlock is a free data retrieval call binding the contract method 0x18c9990b.
//
// Solidity: function getRootByBlock(address _address, uint64 _blockN) constant returns(bytes32)
func (_RootCommitsContract *RootCommitsContractSession) GetRootByBlock(_address common.Address, _blockN uint64) ([32]byte, error) {
	return _RootCommitsContract.Contract.GetRootByBlock(&_RootCommitsContract.CallOpts, _address, _blockN)
}

// GetRootByBlock is a free data retrieval call binding the contract method 0x18c9990b.
//
// Solidity: function getRootByBlock(address _address, uint64 _blockN) constant returns(bytes32)
func (_RootCommitsContract *RootCommitsContractCallerSession) GetRootByBlock(_address common.Address, _blockN uint64) ([32]byte, error) {
	return _RootCommitsContract.Contract.GetRootByBlock(&_RootCommitsContract.CallOpts, _address, _blockN)
}

// GetRootByTime is a free data retrieval call binding the contract method 0xab481d23.
//
// Solidity: function getRootByTime(address _address, uint64 _timestamp) constant returns(bytes32)
func (_RootCommitsContract *RootCommitsContractCaller) GetRootByTime(opts *bind.CallOpts, _address common.Address, _timestamp uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootCommitsContract.contract.Call(opts, out, "getRootByTime", _address, _timestamp)
	return *ret0, err
}

// GetRootByTime is a free data retrieval call binding the contract method 0xab481d23.
//
// Solidity: function getRootByTime(address _address, uint64 _timestamp) constant returns(bytes32)
func (_RootCommitsContract *RootCommitsContractSession) GetRootByTime(_address common.Address, _timestamp uint64) ([32]byte, error) {
	return _RootCommitsContract.Contract.GetRootByTime(&_RootCommitsContract.CallOpts, _address, _timestamp)
}

// GetRootByTime is a free data retrieval call binding the contract method 0xab481d23.
//
// Solidity: function getRootByTime(address _address, uint64 _timestamp) constant returns(bytes32)
func (_RootCommitsContract *RootCommitsContractCallerSession) GetRootByTime(_address common.Address, _timestamp uint64) ([32]byte, error) {
	return _RootCommitsContract.Contract.GetRootByTime(&_RootCommitsContract.CallOpts, _address, _timestamp)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 _root) returns()
func (_RootCommitsContract *RootCommitsContractTransactor) SetRoot(opts *bind.TransactOpts, _root [32]byte) (*types.Transaction, error) {
	return _RootCommitsContract.contract.Transact(opts, "setRoot", _root)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 _root) returns()
func (_RootCommitsContract *RootCommitsContractSession) SetRoot(_root [32]byte) (*types.Transaction, error) {
	return _RootCommitsContract.Contract.SetRoot(&_RootCommitsContract.TransactOpts, _root)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 _root) returns()
func (_RootCommitsContract *RootCommitsContractTransactorSession) SetRoot(_root [32]byte) (*types.Transaction, error) {
	return _RootCommitsContract.Contract.SetRoot(&_RootCommitsContract.TransactOpts, _root)
}

// RootCommitsContractRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the RootCommitsContract contract.
type RootCommitsContractRootUpdatedIterator struct {
	Event *RootCommitsContractRootUpdated // Event containing the contract specifics and raw log

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
func (it *RootCommitsContractRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootCommitsContractRootUpdated)
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
		it.Event = new(RootCommitsContractRootUpdated)
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
func (it *RootCommitsContractRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootCommitsContractRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootCommitsContractRootUpdated represents a RootUpdated event raised by the RootCommitsContract contract.
type RootCommitsContractRootUpdated struct {
	Addr      common.Address
	BlockN    uint64
	Timestamp uint64
	Root      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x1d3d3c252bc04ff83e8069c09279b4d8acb5264996a680b6d7dcf20db2b7417b.
//
// Solidity: event RootUpdated(address addr, uint64 blockN, uint64 timestamp, bytes32 root)
func (_RootCommitsContract *RootCommitsContractFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*RootCommitsContractRootUpdatedIterator, error) {

	logs, sub, err := _RootCommitsContract.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &RootCommitsContractRootUpdatedIterator{contract: _RootCommitsContract.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x1d3d3c252bc04ff83e8069c09279b4d8acb5264996a680b6d7dcf20db2b7417b.
//
// Solidity: event RootUpdated(address addr, uint64 blockN, uint64 timestamp, bytes32 root)
func (_RootCommitsContract *RootCommitsContractFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *RootCommitsContractRootUpdated) (event.Subscription, error) {

	logs, sub, err := _RootCommitsContract.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootCommitsContractRootUpdated)
				if err := _RootCommitsContract.contract.UnpackLog(event, "RootUpdated", log); err != nil {
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
