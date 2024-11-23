// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package user_nft

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

// SolidityMetaData contains all meta data concerning the Solidity contract.
var SolidityMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50336040518060400160405280600781526020017f557365724e4654000000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f554e465400000000000000000000000000000000000000000000000000000000815250815f908161008b9190610427565b50806001908161009b9190610427565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361010e575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016101059190610535565b60405180910390fd5b61011d8161012a60201b60201c565b505f60078190555061054e565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061026857607f821691505b60208210810361027b5761027a610224565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026102dd7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826102a2565b6102e786836102a2565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61032b610326610321846102ff565b610308565b6102ff565b9050919050565b5f819050919050565b61034483610311565b61035861035082610332565b8484546102ae565b825550505050565b5f90565b61036c610360565b61037781848461033b565b505050565b5b8181101561039a5761038f5f82610364565b60018101905061037d565b5050565b601f8211156103df576103b081610281565b6103b984610293565b810160208510156103c8578190505b6103dc6103d485610293565b83018261037c565b50505b505050565b5f82821c905092915050565b5f6103ff5f19846008026103e4565b1980831691505092915050565b5f61041783836103f0565b9150826002028217905092915050565b610430826101ed565b67ffffffffffffffff811115610449576104486101f7565b5b6104538254610251565b61045e82828561039e565b5f60209050601f83116001811461048f575f841561047d578287015190505b610487858261040c565b8655506104ee565b601f19841661049d86610281565b5f5b828110156104c45784890151825560018201915060208501945060208101905061049f565b868310156104e157848901516104dd601f8916826103f0565b8355505b6001600288020188555050505b505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61051f826104f6565b9050919050565b61052f81610515565b82525050565b5f6020820190506105485f830184610526565b92915050565b6122848061055b5f395ff3fe608060405234801561000f575f80fd5b506004361061011f575f3560e01c806370a08231116100ab578063a22cb4651161006f578063a22cb46514610305578063b88d4fde14610321578063c87b56dd1461033d578063e985e9c51461036d578063f2fde38b1461039d5761011f565b806370a0823114610271578063715018a6146102a157806375794a3c146102ab5780638da5cb5b146102c957806395d89b41146102e75761011f565b806323b872dd116100f257806323b872dd146101bd57806342842e0e146101d957806342966c68146101f55780636352211e146102115780636a627842146102415761011f565b806301ffc9a71461012357806306fdde0314610153578063081812fc14610171578063095ea7b3146101a1575b5f80fd5b61013d60048036038101906101389190611a2d565b6103b9565b60405161014a9190611a72565b60405180910390f35b61015b61049a565b6040516101689190611afb565b60405180910390f35b61018b60048036038101906101869190611b4e565b610529565b6040516101989190611bb8565b60405180910390f35b6101bb60048036038101906101b69190611bfb565b610544565b005b6101d760048036038101906101d29190611c39565b61055a565b005b6101f360048036038101906101ee9190611c39565b610659565b005b61020f600480360381019061020a9190611b4e565b610678565b005b61022b60048036038101906102269190611b4e565b610737565b6040516102389190611bb8565b60405180910390f35b61025b60048036038101906102569190611c89565b610748565b6040516102689190611cc3565b60405180910390f35b61028b60048036038101906102869190611c89565b610781565b6040516102989190611cc3565b60405180910390f35b6102a9610837565b005b6102b361084a565b6040516102c09190611cc3565b60405180910390f35b6102d1610850565b6040516102de9190611bb8565b60405180910390f35b6102ef610878565b6040516102fc9190611afb565b60405180910390f35b61031f600480360381019061031a9190611d06565b610908565b005b61033b60048036038101906103369190611e70565b61091e565b005b61035760048036038101906103529190611b4e565b610943565b6040516103649190611afb565b60405180910390f35b61038760048036038101906103829190611ef0565b6109a9565b6040516103949190611a72565b60405180910390f35b6103b760048036038101906103b29190611c89565b610a37565b005b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061048357507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610493575061049282610abb565b5b9050919050565b60605f80546104a890611f5b565b80601f01602080910402602001604051908101604052809291908181526020018280546104d490611f5b565b801561051f5780601f106104f65761010080835404028352916020019161051f565b820191905f5260205f20905b81548152906001019060200180831161050257829003601f168201915b5050505050905090565b5f61053382610b24565b5061053d82610baa565b9050919050565b6105568282610551610be3565b610bea565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105ca575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016105c19190611bb8565b60405180910390fd5b5f6105dd83836105d8610be3565b610bfc565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610653578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161064a93929190611f8b565b60405180910390fd5b50505050565b61067383838360405180602001604052805f81525061091e565b505050565b3373ffffffffffffffffffffffffffffffffffffffff1661069882610737565b73ffffffffffffffffffffffffffffffffffffffff1614806106ec57506106bd610850565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b61072b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072290612030565b60405180910390fd5b61073481610e07565b50565b5f61074182610b24565b9050919050565b5f610751610e89565b5f60075490506107618382610f10565b60075f8154809291906107739061207b565b919050555080915050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107f2575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016107e99190611bb8565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b61083f610e89565b6108485f610f2d565b565b60075481565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606001805461088790611f5b565b80601f01602080910402602001604051908101604052809291908181526020018280546108b390611f5b565b80156108fe5780601f106108d5576101008083540402835291602001916108fe565b820191905f5260205f20905b8154815290600101906020018083116108e157829003601f168201915b5050505050905090565b61091a610913610be3565b8383610ff0565b5050565b61092984848461055a565b61093d610934610be3565b85858585611159565b50505050565b606061094e82610b24565b505f610958611305565b90505f8151116109765760405180602001604052805f8152506109a1565b806109808461131b565b6040516020016109919291906120fc565b6040516020818303038152906040525b915050919050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b610a3f610e89565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610aaf575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610aa69190611bb8565b60405180910390fd5b610ab881610f2d565b50565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f80610b2f836113e5565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ba157826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610b989190611cc3565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610bf7838383600161141e565b505050565b5f80610c07846113e5565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610c4857610c478184866115dd565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610cd357610c875f855f8061141e565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610d5257600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b5f610e135f835f610bfc565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610e8557816040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610e7c9190611cc3565b60405180910390fd5b5050565b610e91610be3565b73ffffffffffffffffffffffffffffffffffffffff16610eaf610850565b73ffffffffffffffffffffffffffffffffffffffff1614610f0e57610ed2610be3565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610f059190611bb8565b60405180910390fd5b565b610f29828260405180602001604052805f8152506116a0565b5050565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361106057816040517f5b08ba180000000000000000000000000000000000000000000000000000000081526004016110579190611bb8565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c318360405161114c9190611a72565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b11156112fe578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b81526004016111b79493929190612171565b6020604051808303815f875af19250505080156111f257506040513d601f19601f820116820180604052508101906111ef91906121cf565b60015b611273573d805f8114611220576040519150601f19603f3d011682016040523d82523d5f602084013e611225565b606091505b505f81510361126b57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016112629190611bb8565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146112fc57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016112f39190611bb8565b60405180910390fd5b505b5050505050565b606060405180602001604052805f815250905090565b60605f6001611329846116c3565b0190505f8167ffffffffffffffff81111561134757611346611d4c565b5b6040519080825280601f01601f1916602001820160405280156113795781602001600182028036833780820191505090505b5090505f82602001820190505b6001156113da578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85816113cf576113ce6121fa565b5b0494505f8503611386575b819350505050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061145657505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15611588575f61146584610b24565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156114cf57508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156114e257506114e081846109a9565b155b1561152457826040517fa9fbf51f00000000000000000000000000000000000000000000000000000000815260040161151b9190611bb8565b60405180910390fd5b811561158657838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6115e8838383611814565b61169b575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361165c57806040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016116539190611cc3565b60405180910390fd5b81816040517f177e802f000000000000000000000000000000000000000000000000000000008152600401611692929190612227565b60405180910390fd5b505050565b6116aa83836118d4565b6116be6116b5610be3565b5f858585611159565b505050565b5f805f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061171f577a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008381611715576117146121fa565b5b0492506040810190505b6d04ee2d6d415b85acef8100000000831061175c576d04ee2d6d415b85acef81000000008381611752576117516121fa565b5b0492506020810190505b662386f26fc10000831061178b57662386f26fc100008381611781576117806121fa565b5b0492506010810190505b6305f5e10083106117b4576305f5e10083816117aa576117a96121fa565b5b0492506008810190505b61271083106117d95761271083816117cf576117ce6121fa565b5b0492506004810190505b606483106117fc57606483816117f2576117f16121fa565b5b0492506002810190505b600a831061180b576001810190505b80915050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156118cb57508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061188c575061188b84846109a9565b5b806118ca57508273ffffffffffffffffffffffffffffffffffffffff166118b283610baa565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611944575f6040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161193b9190611bb8565b60405180910390fd5b5f61195083835f610bfc565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146119c2575f6040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081526004016119b99190611bb8565b60405180910390fd5b505050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611a0c816119d8565b8114611a16575f80fd5b50565b5f81359050611a2781611a03565b92915050565b5f60208284031215611a4257611a416119d0565b5b5f611a4f84828501611a19565b91505092915050565b5f8115159050919050565b611a6c81611a58565b82525050565b5f602082019050611a855f830184611a63565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611acd82611a8b565b611ad78185611a95565b9350611ae7818560208601611aa5565b611af081611ab3565b840191505092915050565b5f6020820190508181035f830152611b138184611ac3565b905092915050565b5f819050919050565b611b2d81611b1b565b8114611b37575f80fd5b50565b5f81359050611b4881611b24565b92915050565b5f60208284031215611b6357611b626119d0565b5b5f611b7084828501611b3a565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611ba282611b79565b9050919050565b611bb281611b98565b82525050565b5f602082019050611bcb5f830184611ba9565b92915050565b611bda81611b98565b8114611be4575f80fd5b50565b5f81359050611bf581611bd1565b92915050565b5f8060408385031215611c1157611c106119d0565b5b5f611c1e85828601611be7565b9250506020611c2f85828601611b3a565b9150509250929050565b5f805f60608486031215611c5057611c4f6119d0565b5b5f611c5d86828701611be7565b9350506020611c6e86828701611be7565b9250506040611c7f86828701611b3a565b9150509250925092565b5f60208284031215611c9e57611c9d6119d0565b5b5f611cab84828501611be7565b91505092915050565b611cbd81611b1b565b82525050565b5f602082019050611cd65f830184611cb4565b92915050565b611ce581611a58565b8114611cef575f80fd5b50565b5f81359050611d0081611cdc565b92915050565b5f8060408385031215611d1c57611d1b6119d0565b5b5f611d2985828601611be7565b9250506020611d3a85828601611cf2565b9150509250929050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611d8282611ab3565b810181811067ffffffffffffffff82111715611da157611da0611d4c565b5b80604052505050565b5f611db36119c7565b9050611dbf8282611d79565b919050565b5f67ffffffffffffffff821115611dde57611ddd611d4c565b5b611de782611ab3565b9050602081019050919050565b828183375f83830152505050565b5f611e14611e0f84611dc4565b611daa565b905082815260208101848484011115611e3057611e2f611d48565b5b611e3b848285611df4565b509392505050565b5f82601f830112611e5757611e56611d44565b5b8135611e67848260208601611e02565b91505092915050565b5f805f8060808587031215611e8857611e876119d0565b5b5f611e9587828801611be7565b9450506020611ea687828801611be7565b9350506040611eb787828801611b3a565b925050606085013567ffffffffffffffff811115611ed857611ed76119d4565b5b611ee487828801611e43565b91505092959194509250565b5f8060408385031215611f0657611f056119d0565b5b5f611f1385828601611be7565b9250506020611f2485828601611be7565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611f7257607f821691505b602082108103611f8557611f84611f2e565b5b50919050565b5f606082019050611f9e5f830186611ba9565b611fab6020830185611cb4565b611fb86040830184611ba9565b949350505050565b7f4f6e6c792074686520746f6b656e206f776e6572206f7220636f6e74726163745f8201527f206f776e65722063616e206275726e2074686520746f6b656e00000000000000602082015250565b5f61201a603983611a95565b915061202582611fc0565b604082019050919050565b5f6020820190508181035f8301526120478161200e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61208582611b1b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036120b7576120b661204e565b5b600182019050919050565b5f81905092915050565b5f6120d682611a8b565b6120e081856120c2565b93506120f0818560208601611aa5565b80840191505092915050565b5f61210782856120cc565b915061211382846120cc565b91508190509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f6121438261211f565b61214d8185612129565b935061215d818560208601611aa5565b61216681611ab3565b840191505092915050565b5f6080820190506121845f830187611ba9565b6121916020830186611ba9565b61219e6040830185611cb4565b81810360608301526121b08184612139565b905095945050505050565b5f815190506121c981611a03565b92915050565b5f602082840312156121e4576121e36119d0565b5b5f6121f1848285016121bb565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f60408201905061223a5f830185611ba9565b6122476020830184611cb4565b939250505056fea264697066735822122030f3e9aa0ea1a7e8bcb5291dca75315586e823721bd1b295b0eec0895b10dfd364736f6c634300081a0033",
}

// SolidityABI is the input ABI used to generate the binding from.
// Deprecated: Use SolidityMetaData.ABI instead.
var SolidityABI = SolidityMetaData.ABI

// SolidityBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolidityMetaData.Bin instead.
var SolidityBin = SolidityMetaData.Bin

// DeploySolidity deploys a new Ethereum contract, binding an instance of Solidity to it.
func DeploySolidity(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Solidity, error) {
	parsed, err := SolidityMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolidityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Solidity{SolidityCaller: SolidityCaller{contract: contract}, SolidityTransactor: SolidityTransactor{contract: contract}, SolidityFilterer: SolidityFilterer{contract: contract}}, nil
}

// Solidity is an auto generated Go binding around an Ethereum contract.
type Solidity struct {
	SolidityCaller     // Read-only binding to the contract
	SolidityTransactor // Write-only binding to the contract
	SolidityFilterer   // Log filterer for contract events
}

// SolidityCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolidityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolidityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolidityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolidityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolidityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoliditySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SoliditySession struct {
	Contract     *Solidity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolidityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolidityCallerSession struct {
	Contract *SolidityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SolidityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolidityTransactorSession struct {
	Contract     *SolidityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SolidityRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolidityRaw struct {
	Contract *Solidity // Generic contract binding to access the raw methods on
}

// SolidityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolidityCallerRaw struct {
	Contract *SolidityCaller // Generic read-only contract binding to access the raw methods on
}

// SolidityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolidityTransactorRaw struct {
	Contract *SolidityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolidity creates a new instance of Solidity, bound to a specific deployed contract.
func NewSolidity(address common.Address, backend bind.ContractBackend) (*Solidity, error) {
	contract, err := bindSolidity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Solidity{SolidityCaller: SolidityCaller{contract: contract}, SolidityTransactor: SolidityTransactor{contract: contract}, SolidityFilterer: SolidityFilterer{contract: contract}}, nil
}

// NewSolidityCaller creates a new read-only instance of Solidity, bound to a specific deployed contract.
func NewSolidityCaller(address common.Address, caller bind.ContractCaller) (*SolidityCaller, error) {
	contract, err := bindSolidity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolidityCaller{contract: contract}, nil
}

// NewSolidityTransactor creates a new write-only instance of Solidity, bound to a specific deployed contract.
func NewSolidityTransactor(address common.Address, transactor bind.ContractTransactor) (*SolidityTransactor, error) {
	contract, err := bindSolidity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolidityTransactor{contract: contract}, nil
}

// NewSolidityFilterer creates a new log filterer instance of Solidity, bound to a specific deployed contract.
func NewSolidityFilterer(address common.Address, filterer bind.ContractFilterer) (*SolidityFilterer, error) {
	contract, err := bindSolidity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolidityFilterer{contract: contract}, nil
}

// bindSolidity binds a generic wrapper to an already deployed contract.
func bindSolidity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolidityMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solidity *SolidityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solidity.Contract.SolidityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solidity *SolidityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.Contract.SolidityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solidity *SolidityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solidity.Contract.SolidityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solidity *SolidityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solidity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solidity *SolidityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solidity *SolidityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solidity.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Solidity *SolidityCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Solidity *SoliditySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Solidity.Contract.BalanceOf(&_Solidity.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Solidity *SolidityCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Solidity.Contract.BalanceOf(&_Solidity.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Solidity *SolidityCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Solidity *SoliditySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Solidity.Contract.GetApproved(&_Solidity.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Solidity *SolidityCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Solidity.Contract.GetApproved(&_Solidity.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Solidity *SolidityCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Solidity *SoliditySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Solidity.Contract.IsApprovedForAll(&_Solidity.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Solidity *SolidityCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Solidity.Contract.IsApprovedForAll(&_Solidity.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Solidity *SolidityCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Solidity *SoliditySession) Name() (string, error) {
	return _Solidity.Contract.Name(&_Solidity.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Solidity *SolidityCallerSession) Name() (string, error) {
	return _Solidity.Contract.Name(&_Solidity.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Solidity *SolidityCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Solidity *SoliditySession) NextTokenId() (*big.Int, error) {
	return _Solidity.Contract.NextTokenId(&_Solidity.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Solidity *SolidityCallerSession) NextTokenId() (*big.Int, error) {
	return _Solidity.Contract.NextTokenId(&_Solidity.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Solidity *SolidityCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Solidity *SoliditySession) Owner() (common.Address, error) {
	return _Solidity.Contract.Owner(&_Solidity.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Solidity *SolidityCallerSession) Owner() (common.Address, error) {
	return _Solidity.Contract.Owner(&_Solidity.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Solidity *SolidityCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Solidity *SoliditySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Solidity.Contract.OwnerOf(&_Solidity.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Solidity *SolidityCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Solidity.Contract.OwnerOf(&_Solidity.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Solidity *SolidityCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Solidity *SoliditySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Solidity.Contract.SupportsInterface(&_Solidity.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Solidity *SolidityCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Solidity.Contract.SupportsInterface(&_Solidity.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Solidity *SolidityCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Solidity *SoliditySession) Symbol() (string, error) {
	return _Solidity.Contract.Symbol(&_Solidity.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Solidity *SolidityCallerSession) Symbol() (string, error) {
	return _Solidity.Contract.Symbol(&_Solidity.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Solidity *SolidityCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Solidity *SoliditySession) TokenURI(tokenId *big.Int) (string, error) {
	return _Solidity.Contract.TokenURI(&_Solidity.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Solidity *SolidityCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Solidity.Contract.TokenURI(&_Solidity.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Solidity *SolidityTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Solidity *SoliditySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.Approve(&_Solidity.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Solidity *SolidityTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.Approve(&_Solidity.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Solidity *SolidityTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Solidity *SoliditySession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.Burn(&_Solidity.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Solidity *SolidityTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.Burn(&_Solidity.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_Solidity *SolidityTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_Solidity *SoliditySession) Mint(to common.Address) (*types.Transaction, error) {
	return _Solidity.Contract.Mint(&_Solidity.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_Solidity *SolidityTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Solidity.Contract.Mint(&_Solidity.TransactOpts, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Solidity *SolidityTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Solidity *SoliditySession) RenounceOwnership() (*types.Transaction, error) {
	return _Solidity.Contract.RenounceOwnership(&_Solidity.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Solidity *SolidityTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Solidity.Contract.RenounceOwnership(&_Solidity.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Solidity *SolidityTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Solidity *SoliditySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.SafeTransferFrom(&_Solidity.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Solidity *SolidityTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.SafeTransferFrom(&_Solidity.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Solidity *SolidityTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Solidity *SoliditySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Solidity.Contract.SafeTransferFrom0(&_Solidity.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Solidity *SolidityTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Solidity.Contract.SafeTransferFrom0(&_Solidity.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Solidity *SolidityTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Solidity *SoliditySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Solidity.Contract.SetApprovalForAll(&_Solidity.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Solidity *SolidityTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Solidity.Contract.SetApprovalForAll(&_Solidity.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Solidity *SolidityTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Solidity *SoliditySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.TransferFrom(&_Solidity.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Solidity *SolidityTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.TransferFrom(&_Solidity.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Solidity *SolidityTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Solidity *SoliditySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Solidity.Contract.TransferOwnership(&_Solidity.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Solidity *SolidityTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Solidity.Contract.TransferOwnership(&_Solidity.TransactOpts, newOwner)
}

// SolidityApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Solidity contract.
type SolidityApprovalIterator struct {
	Event *SolidityApproval // Event containing the contract specifics and raw log

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
func (it *SolidityApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolidityApproval)
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
		it.Event = new(SolidityApproval)
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
func (it *SolidityApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolidityApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolidityApproval represents a Approval event raised by the Solidity contract.
type SolidityApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Solidity *SolidityFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SolidityApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Solidity.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SolidityApprovalIterator{contract: _Solidity.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Solidity *SolidityFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SolidityApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Solidity.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolidityApproval)
				if err := _Solidity.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Solidity *SolidityFilterer) ParseApproval(log types.Log) (*SolidityApproval, error) {
	event := new(SolidityApproval)
	if err := _Solidity.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolidityApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Solidity contract.
type SolidityApprovalForAllIterator struct {
	Event *SolidityApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SolidityApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolidityApprovalForAll)
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
		it.Event = new(SolidityApprovalForAll)
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
func (it *SolidityApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolidityApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolidityApprovalForAll represents a ApprovalForAll event raised by the Solidity contract.
type SolidityApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Solidity *SolidityFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SolidityApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Solidity.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SolidityApprovalForAllIterator{contract: _Solidity.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Solidity *SolidityFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SolidityApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Solidity.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolidityApprovalForAll)
				if err := _Solidity.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Solidity *SolidityFilterer) ParseApprovalForAll(log types.Log) (*SolidityApprovalForAll, error) {
	event := new(SolidityApprovalForAll)
	if err := _Solidity.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolidityOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Solidity contract.
type SolidityOwnershipTransferredIterator struct {
	Event *SolidityOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SolidityOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolidityOwnershipTransferred)
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
		it.Event = new(SolidityOwnershipTransferred)
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
func (it *SolidityOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolidityOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolidityOwnershipTransferred represents a OwnershipTransferred event raised by the Solidity contract.
type SolidityOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Solidity *SolidityFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SolidityOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Solidity.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolidityOwnershipTransferredIterator{contract: _Solidity.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Solidity *SolidityFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SolidityOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Solidity.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolidityOwnershipTransferred)
				if err := _Solidity.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Solidity *SolidityFilterer) ParseOwnershipTransferred(log types.Log) (*SolidityOwnershipTransferred, error) {
	event := new(SolidityOwnershipTransferred)
	if err := _Solidity.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolidityTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Solidity contract.
type SolidityTransferIterator struct {
	Event *SolidityTransfer // Event containing the contract specifics and raw log

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
func (it *SolidityTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolidityTransfer)
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
		it.Event = new(SolidityTransfer)
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
func (it *SolidityTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolidityTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolidityTransfer represents a Transfer event raised by the Solidity contract.
type SolidityTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Solidity *SolidityFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SolidityTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Solidity.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SolidityTransferIterator{contract: _Solidity.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Solidity *SolidityFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SolidityTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Solidity.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolidityTransfer)
				if err := _Solidity.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Solidity *SolidityFilterer) ParseTransfer(log types.Log) (*SolidityTransfer, error) {
	event := new(SolidityTransfer)
	if err := _Solidity.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
