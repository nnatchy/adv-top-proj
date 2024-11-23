// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simple_storage

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

// SimplestorageMetaData contains all meta data concerning the Simplestorage contract.
var SimplestorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumber\",\"type\":\"uint256\"}],\"name\":\"NumberUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b5061017a8061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80633fb5c1cb14610038578063f2c9ecd814610054575b5f80fd5b610052600480360381019061004d91906100f1565b610072565b005b61005c6100b2565b604051610069919061012b565b60405180910390f35b805f819055507f939f2783f4b99373bd9b6912aff8f7316ea141b98eb02e47c8a96bdae0d654b3816040516100a7919061012b565b60405180910390a150565b5f8054905090565b5f80fd5b5f819050919050565b6100d0816100be565b81146100da575f80fd5b50565b5f813590506100eb816100c7565b92915050565b5f60208284031215610106576101056100ba565b5b5f610113848285016100dd565b91505092915050565b610125816100be565b82525050565b5f60208201905061013e5f83018461011c565b9291505056fea264697066735822122093d888350cdf42ac17f5dd373ca523975d97bb10995936359fc8fa42a44e7eb364736f6c634300081a0033",
}

// SimplestorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SimplestorageMetaData.ABI instead.
var SimplestorageABI = SimplestorageMetaData.ABI

// SimplestorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimplestorageMetaData.Bin instead.
var SimplestorageBin = SimplestorageMetaData.Bin

// DeploySimplestorage deploys a new Ethereum contract, binding an instance of Simplestorage to it.
func DeploySimplestorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Simplestorage, error) {
	parsed, err := SimplestorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimplestorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simplestorage{SimplestorageCaller: SimplestorageCaller{contract: contract}, SimplestorageTransactor: SimplestorageTransactor{contract: contract}, SimplestorageFilterer: SimplestorageFilterer{contract: contract}}, nil
}

// Simplestorage is an auto generated Go binding around an Ethereum contract.
type Simplestorage struct {
	SimplestorageCaller     // Read-only binding to the contract
	SimplestorageTransactor // Write-only binding to the contract
	SimplestorageFilterer   // Log filterer for contract events
}

// SimplestorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplestorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplestorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplestorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplestorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplestorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplestorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplestorageSession struct {
	Contract     *Simplestorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimplestorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplestorageCallerSession struct {
	Contract *SimplestorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimplestorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplestorageTransactorSession struct {
	Contract     *SimplestorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimplestorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplestorageRaw struct {
	Contract *Simplestorage // Generic contract binding to access the raw methods on
}

// SimplestorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplestorageCallerRaw struct {
	Contract *SimplestorageCaller // Generic read-only contract binding to access the raw methods on
}

// SimplestorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplestorageTransactorRaw struct {
	Contract *SimplestorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplestorage creates a new instance of Simplestorage, bound to a specific deployed contract.
func NewSimplestorage(address common.Address, backend bind.ContractBackend) (*Simplestorage, error) {
	contract, err := bindSimplestorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simplestorage{SimplestorageCaller: SimplestorageCaller{contract: contract}, SimplestorageTransactor: SimplestorageTransactor{contract: contract}, SimplestorageFilterer: SimplestorageFilterer{contract: contract}}, nil
}

// NewSimplestorageCaller creates a new read-only instance of Simplestorage, bound to a specific deployed contract.
func NewSimplestorageCaller(address common.Address, caller bind.ContractCaller) (*SimplestorageCaller, error) {
	contract, err := bindSimplestorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplestorageCaller{contract: contract}, nil
}

// NewSimplestorageTransactor creates a new write-only instance of Simplestorage, bound to a specific deployed contract.
func NewSimplestorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplestorageTransactor, error) {
	contract, err := bindSimplestorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplestorageTransactor{contract: contract}, nil
}

// NewSimplestorageFilterer creates a new log filterer instance of Simplestorage, bound to a specific deployed contract.
func NewSimplestorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplestorageFilterer, error) {
	contract, err := bindSimplestorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplestorageFilterer{contract: contract}, nil
}

// bindSimplestorage binds a generic wrapper to an already deployed contract.
func bindSimplestorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimplestorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplestorage *SimplestorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplestorage.Contract.SimplestorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplestorage *SimplestorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplestorage.Contract.SimplestorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplestorage *SimplestorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplestorage.Contract.SimplestorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplestorage *SimplestorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplestorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplestorage *SimplestorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplestorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplestorage *SimplestorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplestorage.Contract.contract.Transact(opts, method, params...)
}

// GetNumber is a free data retrieval call binding the contract method 0xf2c9ecd8.
//
// Solidity: function getNumber() view returns(uint256)
func (_Simplestorage *SimplestorageCaller) GetNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simplestorage.contract.Call(opts, &out, "getNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumber is a free data retrieval call binding the contract method 0xf2c9ecd8.
//
// Solidity: function getNumber() view returns(uint256)
func (_Simplestorage *SimplestorageSession) GetNumber() (*big.Int, error) {
	return _Simplestorage.Contract.GetNumber(&_Simplestorage.CallOpts)
}

// GetNumber is a free data retrieval call binding the contract method 0xf2c9ecd8.
//
// Solidity: function getNumber() view returns(uint256)
func (_Simplestorage *SimplestorageCallerSession) GetNumber() (*big.Int, error) {
	return _Simplestorage.Contract.GetNumber(&_Simplestorage.CallOpts)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) returns()
func (_Simplestorage *SimplestorageTransactor) SetNumber(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _Simplestorage.contract.Transact(opts, "setNumber", _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) returns()
func (_Simplestorage *SimplestorageSession) SetNumber(_number *big.Int) (*types.Transaction, error) {
	return _Simplestorage.Contract.SetNumber(&_Simplestorage.TransactOpts, _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) returns()
func (_Simplestorage *SimplestorageTransactorSession) SetNumber(_number *big.Int) (*types.Transaction, error) {
	return _Simplestorage.Contract.SetNumber(&_Simplestorage.TransactOpts, _number)
}

// SimplestorageNumberUpdatedIterator is returned from FilterNumberUpdated and is used to iterate over the raw logs and unpacked data for NumberUpdated events raised by the Simplestorage contract.
type SimplestorageNumberUpdatedIterator struct {
	Event *SimplestorageNumberUpdated // Event containing the contract specifics and raw log

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
func (it *SimplestorageNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplestorageNumberUpdated)
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
		it.Event = new(SimplestorageNumberUpdated)
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
func (it *SimplestorageNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplestorageNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplestorageNumberUpdated represents a NumberUpdated event raised by the Simplestorage contract.
type SimplestorageNumberUpdated struct {
	NewNumber *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNumberUpdated is a free log retrieval operation binding the contract event 0x939f2783f4b99373bd9b6912aff8f7316ea141b98eb02e47c8a96bdae0d654b3.
//
// Solidity: event NumberUpdated(uint256 newNumber)
func (_Simplestorage *SimplestorageFilterer) FilterNumberUpdated(opts *bind.FilterOpts) (*SimplestorageNumberUpdatedIterator, error) {

	logs, sub, err := _Simplestorage.contract.FilterLogs(opts, "NumberUpdated")
	if err != nil {
		return nil, err
	}
	return &SimplestorageNumberUpdatedIterator{contract: _Simplestorage.contract, event: "NumberUpdated", logs: logs, sub: sub}, nil
}

// WatchNumberUpdated is a free log subscription operation binding the contract event 0x939f2783f4b99373bd9b6912aff8f7316ea141b98eb02e47c8a96bdae0d654b3.
//
// Solidity: event NumberUpdated(uint256 newNumber)
func (_Simplestorage *SimplestorageFilterer) WatchNumberUpdated(opts *bind.WatchOpts, sink chan<- *SimplestorageNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _Simplestorage.contract.WatchLogs(opts, "NumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplestorageNumberUpdated)
				if err := _Simplestorage.contract.UnpackLog(event, "NumberUpdated", log); err != nil {
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

// ParseNumberUpdated is a log parse operation binding the contract event 0x939f2783f4b99373bd9b6912aff8f7316ea141b98eb02e47c8a96bdae0d654b3.
//
// Solidity: event NumberUpdated(uint256 newNumber)
func (_Simplestorage *SimplestorageFilterer) ParseNumberUpdated(log types.Log) (*SimplestorageNumberUpdated, error) {
	event := new(SimplestorageNumberUpdated)
	if err := _Simplestorage.contract.UnpackLog(event, "NumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
