// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voting

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

// VotingSystemVote is an auto generated low-level Go binding around an user-defined struct.
type VotingSystemVote struct {
	StudentId    string
	VoterAddress common.Address
	VoteChoice   string
}

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"studentId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"voteChoice\",\"type\":\"string\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"studentId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"voteChoice\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VotingEnded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allVotes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"studentId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"voteChoice\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllVotes\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"studentId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"voteChoice\",\"type\":\"string\"}],\"internalType\":\"structVotingSystem.Vote[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001805f6101000a81548160ff0219169083151502179055503480156027575f80fd5b50336001806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611163806100755f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c8063851b6ef2146100645780638da5cb5b14610082578063a95824b4146100a0578063c3403ddf146100be578063de63f432146100c8578063fa057225146100fa575b5f80fd5b61006c610116565b6040516100799190610925565b60405180910390f35b61008a6102e5565b6040516100979190610954565b60405180910390f35b6100a8610309565b6040516100b59190610987565b60405180910390f35b6100c661031b565b005b6100e260048036038101906100dd91906109e4565b61043e565b6040516100f193929190610a57565b60405180910390f35b610114600480360381019061010f9190610bf0565b61059d565b005b60605f805480602002602001604051908101604052809291908181526020015f905b828210156102dc578382905f5260205f2090600302016040518060600160405290815f8201805461016890610ca5565b80601f016020809104026020016040519081016040528092919081815260200182805461019490610ca5565b80156101df5780601f106101b6576101008083540402835291602001916101df565b820191905f5260205f20905b8154815290600101906020018083116101c257829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461024d90610ca5565b80601f016020809104026020016040519081016040528092919081815260200182805461027990610ca5565b80156102c45780601f1061029b576101008083540402835291602001916102c4565b820191905f5260205f20905b8154815290600101906020018083116102a757829003601f168201915b50505050508152505081526020019060010190610138565b50505050905090565b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900460ff1681565b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a090610d45565b60405180910390fd5b60015f9054906101000a900460ff166103f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ee90610dad565b60405180910390fd5b5f60015f6101000a81548160ff0219169083151502179055507f7a19ed057db79e3c2fa0b97a54b43bef4fce74b31bb6c01af514b9a18a7f70ab60405160405180910390a1565b5f818154811061044c575f80fd5b905f5260205f2090600302015f91509050805f01805461046b90610ca5565b80601f016020809104026020016040519081016040528092919081815260200182805461049790610ca5565b80156104e25780601f106104b9576101008083540402835291602001916104e2565b820191905f5260205f20905b8154815290600101906020018083116104c557829003601f168201915b505050505090806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600201805461051c90610ca5565b80601f016020809104026020016040519081016040528092919081815260200182805461054890610ca5565b80156105935780601f1061056a57610100808354040283529160200191610593565b820191905f5260205f20905b81548152906001019060200180831161057657829003601f168201915b5050505050905083565b60015f9054906101000a900460ff166105eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e290610e15565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614610659576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065090610ea3565b60405180910390fd5b5f60405180606001604052808581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381525090505f81908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f0190816106c8919061105e565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019081610724919061105e565b5050507ea57c72212c37dbaffdcb5de556c1fcb9cb77e578fbf2138f88d107e699692084848460405161075993929190610a57565b60405180910390a150505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6107d282610790565b6107dc818561079a565b93506107ec8185602086016107aa565b6107f5816107b8565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61082982610800565b9050919050565b6108398161081f565b82525050565b5f606083015f8301518482035f86015261085982826107c8565b915050602083015161086e6020860182610830565b506040830151848203604086015261088682826107c8565b9150508091505092915050565b5f61089e838361083f565b905092915050565b5f602082019050919050565b5f6108bc82610767565b6108c68185610771565b9350836020820285016108d885610781565b805f5b8581101561091357848403895281516108f48582610893565b94506108ff836108a6565b925060208a019950506001810190506108db565b50829750879550505050505092915050565b5f6020820190508181035f83015261093d81846108b2565b905092915050565b61094e8161081f565b82525050565b5f6020820190506109675f830184610945565b92915050565b5f8115159050919050565b6109818161096d565b82525050565b5f60208201905061099a5f830184610978565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6109c3816109b1565b81146109cd575f80fd5b50565b5f813590506109de816109ba565b92915050565b5f602082840312156109f9576109f86109a9565b5b5f610a06848285016109d0565b91505092915050565b5f82825260208201905092915050565b5f610a2982610790565b610a338185610a0f565b9350610a438185602086016107aa565b610a4c816107b8565b840191505092915050565b5f6060820190508181035f830152610a6f8186610a1f565b9050610a7e6020830185610945565b8181036040830152610a908184610a1f565b9050949350505050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610ad8826107b8565b810181811067ffffffffffffffff82111715610af757610af6610aa2565b5b80604052505050565b5f610b096109a0565b9050610b158282610acf565b919050565b5f67ffffffffffffffff821115610b3457610b33610aa2565b5b610b3d826107b8565b9050602081019050919050565b828183375f83830152505050565b5f610b6a610b6584610b1a565b610b00565b905082815260208101848484011115610b8657610b85610a9e565b5b610b91848285610b4a565b509392505050565b5f82601f830112610bad57610bac610a9a565b5b8135610bbd848260208601610b58565b91505092915050565b610bcf8161081f565b8114610bd9575f80fd5b50565b5f81359050610bea81610bc6565b92915050565b5f805f60608486031215610c0757610c066109a9565b5b5f84013567ffffffffffffffff811115610c2457610c236109ad565b5b610c3086828701610b99565b9350506020610c4186828701610bdc565b925050604084013567ffffffffffffffff811115610c6257610c616109ad565b5b610c6e86828701610b99565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610cbc57607f821691505b602082108103610ccf57610cce610c78565b5b50919050565b7f4f6e6c7920746865206f776e65722063616e20706572666f726d2074686973205f8201527f616374696f6e2e00000000000000000000000000000000000000000000000000602082015250565b5f610d2f602783610a0f565b9150610d3a82610cd5565b604082019050919050565b5f6020820190508181035f830152610d5c81610d23565b9050919050565b7f566f74696e6720697320616c726561647920636c6f7365642e000000000000005f82015250565b5f610d97601983610a0f565b9150610da282610d63565b602082019050919050565b5f6020820190508181035f830152610dc481610d8b565b9050919050565b7f566f74696e672068617320656e6465642e0000000000000000000000000000005f82015250565b5f610dff601183610a0f565b9150610e0a82610dcb565b602082019050919050565b5f6020820190508181035f830152610e2c81610df3565b9050919050565b7f566f7465722061646472657373206d757374206d61746368207468652063616c5f8201527f6c6572277320616464726573732e000000000000000000000000000000000000602082015250565b5f610e8d602e83610a0f565b9150610e9882610e33565b604082019050919050565b5f6020820190508181035f830152610eba81610e81565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610f1d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610ee2565b610f278683610ee2565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610f62610f5d610f58846109b1565b610f3f565b6109b1565b9050919050565b5f819050919050565b610f7b83610f48565b610f8f610f8782610f69565b848454610eee565b825550505050565b5f90565b610fa3610f97565b610fae818484610f72565b505050565b5b81811015610fd157610fc65f82610f9b565b600181019050610fb4565b5050565b601f82111561101657610fe781610ec1565b610ff084610ed3565b81016020851015610fff578190505b61101361100b85610ed3565b830182610fb3565b50505b505050565b5f82821c905092915050565b5f6110365f198460080261101b565b1980831691505092915050565b5f61104e8383611027565b9150826002028217905092915050565b61106782610790565b67ffffffffffffffff8111156110805761107f610aa2565b5b61108a8254610ca5565b611095828285610fd5565b5f60209050601f8311600181146110c6575f84156110b4578287015190505b6110be8582611043565b865550611125565b601f1984166110d486610ec1565b5f5b828110156110fb578489015182556001820191506020850194506020810190506110d6565b868310156111185784890151611114601f891682611027565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220904cf6319fa312dac70b3f6206b5f0129fc26751e0a014fd7971c09aa430958e64736f6c634300081a0033",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// VotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotingMetaData.Bin instead.
var VotingBin = VotingMetaData.Bin

// DeployVoting deploys a new Ethereum contract, binding an instance of Voting to it.
func DeployVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Voting, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// AllVotes is a free data retrieval call binding the contract method 0xde63f432.
//
// Solidity: function allVotes(uint256 ) view returns(string studentId, address voterAddress, string voteChoice)
func (_Voting *VotingCaller) AllVotes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StudentId    string
	VoterAddress common.Address
	VoteChoice   string
}, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "allVotes", arg0)

	outstruct := new(struct {
		StudentId    string
		VoterAddress common.Address
		VoteChoice   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StudentId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VoterAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.VoteChoice = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// AllVotes is a free data retrieval call binding the contract method 0xde63f432.
//
// Solidity: function allVotes(uint256 ) view returns(string studentId, address voterAddress, string voteChoice)
func (_Voting *VotingSession) AllVotes(arg0 *big.Int) (struct {
	StudentId    string
	VoterAddress common.Address
	VoteChoice   string
}, error) {
	return _Voting.Contract.AllVotes(&_Voting.CallOpts, arg0)
}

// AllVotes is a free data retrieval call binding the contract method 0xde63f432.
//
// Solidity: function allVotes(uint256 ) view returns(string studentId, address voterAddress, string voteChoice)
func (_Voting *VotingCallerSession) AllVotes(arg0 *big.Int) (struct {
	StudentId    string
	VoterAddress common.Address
	VoteChoice   string
}, error) {
	return _Voting.Contract.AllVotes(&_Voting.CallOpts, arg0)
}

// GetAllVotes is a free data retrieval call binding the contract method 0x851b6ef2.
//
// Solidity: function getAllVotes() view returns((string,address,string)[])
func (_Voting *VotingCaller) GetAllVotes(opts *bind.CallOpts) ([]VotingSystemVote, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getAllVotes")

	if err != nil {
		return *new([]VotingSystemVote), err
	}

	out0 := *abi.ConvertType(out[0], new([]VotingSystemVote)).(*[]VotingSystemVote)

	return out0, err

}

// GetAllVotes is a free data retrieval call binding the contract method 0x851b6ef2.
//
// Solidity: function getAllVotes() view returns((string,address,string)[])
func (_Voting *VotingSession) GetAllVotes() ([]VotingSystemVote, error) {
	return _Voting.Contract.GetAllVotes(&_Voting.CallOpts)
}

// GetAllVotes is a free data retrieval call binding the contract method 0x851b6ef2.
//
// Solidity: function getAllVotes() view returns((string,address,string)[])
func (_Voting *VotingCallerSession) GetAllVotes() ([]VotingSystemVote, error) {
	return _Voting.Contract.GetAllVotes(&_Voting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voting *VotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voting *VotingSession) Owner() (common.Address, error) {
	return _Voting.Contract.Owner(&_Voting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voting *VotingCallerSession) Owner() (common.Address, error) {
	return _Voting.Contract.Owner(&_Voting.CallOpts)
}

// VotingOpen is a free data retrieval call binding the contract method 0xa95824b4.
//
// Solidity: function votingOpen() view returns(bool)
func (_Voting *VotingCaller) VotingOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "votingOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VotingOpen is a free data retrieval call binding the contract method 0xa95824b4.
//
// Solidity: function votingOpen() view returns(bool)
func (_Voting *VotingSession) VotingOpen() (bool, error) {
	return _Voting.Contract.VotingOpen(&_Voting.CallOpts)
}

// VotingOpen is a free data retrieval call binding the contract method 0xa95824b4.
//
// Solidity: function votingOpen() view returns(bool)
func (_Voting *VotingCallerSession) VotingOpen() (bool, error) {
	return _Voting.Contract.VotingOpen(&_Voting.CallOpts)
}

// CastVote is a paid mutator transaction binding the contract method 0xfa057225.
//
// Solidity: function castVote(string studentId, address voterAddress, string voteChoice) returns()
func (_Voting *VotingTransactor) CastVote(opts *bind.TransactOpts, studentId string, voterAddress common.Address, voteChoice string) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "castVote", studentId, voterAddress, voteChoice)
}

// CastVote is a paid mutator transaction binding the contract method 0xfa057225.
//
// Solidity: function castVote(string studentId, address voterAddress, string voteChoice) returns()
func (_Voting *VotingSession) CastVote(studentId string, voterAddress common.Address, voteChoice string) (*types.Transaction, error) {
	return _Voting.Contract.CastVote(&_Voting.TransactOpts, studentId, voterAddress, voteChoice)
}

// CastVote is a paid mutator transaction binding the contract method 0xfa057225.
//
// Solidity: function castVote(string studentId, address voterAddress, string voteChoice) returns()
func (_Voting *VotingTransactorSession) CastVote(studentId string, voterAddress common.Address, voteChoice string) (*types.Transaction, error) {
	return _Voting.Contract.CastVote(&_Voting.TransactOpts, studentId, voterAddress, voteChoice)
}

// EndVoting is a paid mutator transaction binding the contract method 0xc3403ddf.
//
// Solidity: function endVoting() returns()
func (_Voting *VotingTransactor) EndVoting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "endVoting")
}

// EndVoting is a paid mutator transaction binding the contract method 0xc3403ddf.
//
// Solidity: function endVoting() returns()
func (_Voting *VotingSession) EndVoting() (*types.Transaction, error) {
	return _Voting.Contract.EndVoting(&_Voting.TransactOpts)
}

// EndVoting is a paid mutator transaction binding the contract method 0xc3403ddf.
//
// Solidity: function endVoting() returns()
func (_Voting *VotingTransactorSession) EndVoting() (*types.Transaction, error) {
	return _Voting.Contract.EndVoting(&_Voting.TransactOpts)
}

// VotingVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Voting contract.
type VotingVoteCastIterator struct {
	Event *VotingVoteCast // Event containing the contract specifics and raw log

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
func (it *VotingVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingVoteCast)
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
		it.Event = new(VotingVoteCast)
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
func (it *VotingVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingVoteCast represents a VoteCast event raised by the Voting contract.
type VotingVoteCast struct {
	StudentId    string
	VoterAddress common.Address
	VoteChoice   string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0x00a57c72212c37dbaffdcb5de556c1fcb9cb77e578fbf2138f88d107e6996920.
//
// Solidity: event VoteCast(string studentId, address voterAddress, string voteChoice)
func (_Voting *VotingFilterer) FilterVoteCast(opts *bind.FilterOpts) (*VotingVoteCastIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return &VotingVoteCastIterator{contract: _Voting.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0x00a57c72212c37dbaffdcb5de556c1fcb9cb77e578fbf2138f88d107e6996920.
//
// Solidity: event VoteCast(string studentId, address voterAddress, string voteChoice)
func (_Voting *VotingFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *VotingVoteCast) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingVoteCast)
				if err := _Voting.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0x00a57c72212c37dbaffdcb5de556c1fcb9cb77e578fbf2138f88d107e6996920.
//
// Solidity: event VoteCast(string studentId, address voterAddress, string voteChoice)
func (_Voting *VotingFilterer) ParseVoteCast(log types.Log) (*VotingVoteCast, error) {
	event := new(VotingVoteCast)
	if err := _Voting.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingVotingEndedIterator is returned from FilterVotingEnded and is used to iterate over the raw logs and unpacked data for VotingEnded events raised by the Voting contract.
type VotingVotingEndedIterator struct {
	Event *VotingVotingEnded // Event containing the contract specifics and raw log

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
func (it *VotingVotingEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingVotingEnded)
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
		it.Event = new(VotingVotingEnded)
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
func (it *VotingVotingEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingVotingEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingVotingEnded represents a VotingEnded event raised by the Voting contract.
type VotingVotingEnded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVotingEnded is a free log retrieval operation binding the contract event 0x7a19ed057db79e3c2fa0b97a54b43bef4fce74b31bb6c01af514b9a18a7f70ab.
//
// Solidity: event VotingEnded()
func (_Voting *VotingFilterer) FilterVotingEnded(opts *bind.FilterOpts) (*VotingVotingEndedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "VotingEnded")
	if err != nil {
		return nil, err
	}
	return &VotingVotingEndedIterator{contract: _Voting.contract, event: "VotingEnded", logs: logs, sub: sub}, nil
}

// WatchVotingEnded is a free log subscription operation binding the contract event 0x7a19ed057db79e3c2fa0b97a54b43bef4fce74b31bb6c01af514b9a18a7f70ab.
//
// Solidity: event VotingEnded()
func (_Voting *VotingFilterer) WatchVotingEnded(opts *bind.WatchOpts, sink chan<- *VotingVotingEnded) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "VotingEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingVotingEnded)
				if err := _Voting.contract.UnpackLog(event, "VotingEnded", log); err != nil {
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

// ParseVotingEnded is a log parse operation binding the contract event 0x7a19ed057db79e3c2fa0b97a54b43bef4fce74b31bb6c01af514b9a18a7f70ab.
//
// Solidity: event VotingEnded()
func (_Voting *VotingFilterer) ParseVotingEnded(log types.Log) (*VotingVotingEnded, error) {
	event := new(VotingVotingEnded)
	if err := _Voting.contract.UnpackLog(event, "VotingEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
