// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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

// SolMetaData contains all meta data concerning the Sol contract.
var SolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"Deposited\",\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":true,\"type\":\"address\",\"internalType\":\"address\",\"name\":\"_who\"},{\"indexed\":true,\"internalType\":\"address\",\"type\":\"address\",\"name\":\"_tokenAddress\"},{\"internalType\":\"uint256\",\"indexed\":true,\"type\":\"uint256\",\"name\":\"_amount\"}]},{\"anonymous\":false,\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"internalType\":\"address\",\"name\":\"previousOwner\",\"indexed\":true},{\"internalType\":\"address\",\"indexed\":true,\"type\":\"address\",\"name\":\"newOwner\"}]},{\"anonymous\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"indexed\":true,\"type\":\"address\"},{\"internalType\":\"address\",\"indexed\":true,\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"event\",\"name\":\"Withdrawn\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"name\":\"owner\"},{\"outputs\":[],\"stateMutability\":\"nonpayable\",\"inputs\":[],\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"internalType\":\"address\",\"name\":\"newOwner\"}],\"type\":\"function\",\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"inputs\":[{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"amount_\"}]},{\"outputs\":[],\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"token_\",\"internalType\":\"contractIERC20\"},{\"name\":\"recipient_\",\"type\":\"address\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"amount_\"}],\"type\":\"function\",\"name\":\"withdraw\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106100575760003560e01c806347e7ef241461005c578063715018a6146100785780638da5cb5b14610082578063d9caed12146100a0578063f2fde38b146100bc575b600080fd5b61007660048036038101906100719190610b1b565b6100d8565b005b6100806101e2565b005b61008a61026a565b6040516100979190610d37565b60405180910390f35b6100ba60048036038101906100b59190610acc565b610293565b005b6100d660048036038101906100d19190610a7a565b610547565b005b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610148576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013f90610e8b565b60405180910390fd5b61017c61015361063f565b30838573ffffffffffffffffffffffffffffffffffffffff16610647909392919063ffffffff16565b808273ffffffffffffffffffffffffffffffffffffffff1661019c61063f565b73ffffffffffffffffffffffffffffffffffffffff167f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a760405160405180910390a45050565b6101ea61063f565b73ffffffffffffffffffffffffffffffffffffffff1661020861026a565b73ffffffffffffffffffffffffffffffffffffffff161461025e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025590610e2b565b60405180910390fd5b61026860006106d0565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61029b61063f565b73ffffffffffffffffffffffffffffffffffffffff166102b961026a565b73ffffffffffffffffffffffffffffffffffffffff161461030f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030690610e2b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561037f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037690610e8b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156103ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e690610deb565b60405180910390fd5b808373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016104299190610d37565b60206040518083038186803b15801561044157600080fd5b505afa158015610455573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104799190610b57565b10156104ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b190610e0b565b60405180910390fd5b6104e73083838673ffffffffffffffffffffffffffffffffffffffff16610647909392919063ffffffff16565b808373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb60405160405180910390a4505050565b61054f61063f565b73ffffffffffffffffffffffffffffffffffffffff1661056d61026a565b73ffffffffffffffffffffffffffffffffffffffff16146105c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ba90610e2b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610633576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062a90610dab565b60405180910390fd5b61063c816106d0565b50565b600033905090565b6106ca846323b872dd60e01b85858560405160240161066893929190610d52565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610794565b50505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60006107f6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661085b9092919063ffffffff16565b905060008151111561085657808060200190518101906108169190610aa3565b610855576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084c90610e6b565b60405180910390fd5b5b505050565b606061086a8484600085610873565b90509392505050565b6060824710156108b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108af90610dcb565b60405180910390fd5b6108c185610987565b610900576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f790610e4b565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516109299190610d20565b60006040518083038185875af1925050503d8060008114610966576040519150601f19603f3d011682016040523d82523d6000602084013e61096b565b606091505b509150915061097b8282866109aa565b92505050949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b606083156109ba57829050610a0a565b6000835111156109cd5782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a019190610d89565b60405180910390fd5b9392505050565b600081359050610a2081611181565b92915050565b600081519050610a3581611198565b92915050565b600081359050610a4a816111af565b92915050565b600081359050610a5f816111c6565b92915050565b600081519050610a74816111c6565b92915050565b600060208284031215610a8c57600080fd5b6000610a9a84828501610a11565b91505092915050565b600060208284031215610ab557600080fd5b6000610ac384828501610a26565b91505092915050565b600080600060608486031215610ae157600080fd5b6000610aef86828701610a3b565b9350506020610b0086828701610a11565b9250506040610b1186828701610a50565b9150509250925092565b60008060408385031215610b2e57600080fd5b6000610b3c85828601610a3b565b9250506020610b4d85828601610a50565b9150509250929050565b600060208284031215610b6957600080fd5b6000610b7784828501610a65565b91505092915050565b610b8981610edd565b82525050565b6000610b9a82610eab565b610ba48185610ec1565b9350610bb4818560208601610f37565b80840191505092915050565b6000610bcb82610eb6565b610bd58185610ecc565b9350610be5818560208601610f37565b610bee81610f6a565b840191505092915050565b6000610c06602683610ecc565b9150610c1182610f7b565b604082019050919050565b6000610c29602683610ecc565b9150610c3482610fca565b604082019050919050565b6000610c4c602083610ecc565b9150610c5782611019565b602082019050919050565b6000610c6f602683610ecc565b9150610c7a82611042565b604082019050919050565b6000610c92602083610ecc565b9150610c9d82611091565b602082019050919050565b6000610cb5601d83610ecc565b9150610cc0826110ba565b602082019050919050565b6000610cd8602a83610ecc565b9150610ce3826110e3565b604082019050919050565b6000610cfb602383610ecc565b9150610d0682611132565b604082019050919050565b610d1a81610f2d565b82525050565b6000610d2c8284610b8f565b915081905092915050565b6000602082019050610d4c6000830184610b80565b92915050565b6000606082019050610d676000830186610b80565b610d746020830185610b80565b610d816040830184610d11565b949350505050565b60006020820190508181036000830152610da38184610bc0565b905092915050565b60006020820190508181036000830152610dc481610bf9565b9050919050565b60006020820190508181036000830152610de481610c1c565b9050919050565b60006020820190508181036000830152610e0481610c3f565b9050919050565b60006020820190508181036000830152610e2481610c62565b9050919050565b60006020820190508181036000830152610e4481610c85565b9050919050565b60006020820190508181036000830152610e6481610ca8565b9050919050565b60006020820190508181036000830152610e8481610ccb565b9050919050565b60006020820190508181036000830152610ea481610cee565b9050919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b6000610ee882610f0d565b9050919050565b60008115159050919050565b6000610f0682610edd565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015610f55578082015181840152602081019050610f3a565b83811115610f64576000848401525b50505050565b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f5468652061646472657373206f6620726563697069656e74206973206e756c6c600082015250565b7f54686973207661756c7420646f65736e2774206861766520656e6f756768206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5468652061646472657373206f662049455243323020746f6b656e206973206e60008201527f756c6c0000000000000000000000000000000000000000000000000000000000602082015250565b61118a81610edd565b811461119557600080fd5b50565b6111a181610eef565b81146111ac57600080fd5b50565b6111b881610efb565b81146111c357600080fd5b50565b6111cf81610f2d565b81146111da57600080fd5b5056fea26469706673582212202543a22d2a3d0712ea5ee41cd86f4ddd87254b03c564a4803fe8fc42efcdb37c64736f6c63430008040033",
}

// SolABI is the input ABI used to generate the binding from.
// Deprecated: Use SolMetaData.ABI instead.
var SolABI = SolMetaData.ABI

// SolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolMetaData.Bin instead.
var SolBin = SolMetaData.Bin

// DeploySol deploys a new Ethereum contract, binding an instance of Sol to it.
func DeploySol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sol, error) {
	parsed, err := SolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sol{SolCaller: SolCaller{contract: contract}, SolTransactor: SolTransactor{contract: contract}, SolFilterer: SolFilterer{contract: contract}}, nil
}

// Sol is an auto generated Go binding around an Ethereum contract.
type Sol struct {
	SolCaller     // Read-only binding to the contract
	SolTransactor // Write-only binding to the contract
	SolFilterer   // Log filterer for contract events
}

// SolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolSession struct {
	Contract     *Sol              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolCallerSession struct {
	Contract *SolCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolTransactorSession struct {
	Contract     *SolTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolRaw struct {
	Contract *Sol // Generic contract binding to access the raw methods on
}

// SolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolCallerRaw struct {
	Contract *SolCaller // Generic read-only contract binding to access the raw methods on
}

// SolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolTransactorRaw struct {
	Contract *SolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSol creates a new instance of Sol, bound to a specific deployed contract.
func NewSol(address common.Address, backend bind.ContractBackend) (*Sol, error) {
	contract, err := bindSol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sol{SolCaller: SolCaller{contract: contract}, SolTransactor: SolTransactor{contract: contract}, SolFilterer: SolFilterer{contract: contract}}, nil
}

// NewSolCaller creates a new read-only instance of Sol, bound to a specific deployed contract.
func NewSolCaller(address common.Address, caller bind.ContractCaller) (*SolCaller, error) {
	contract, err := bindSol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolCaller{contract: contract}, nil
}

// NewSolTransactor creates a new write-only instance of Sol, bound to a specific deployed contract.
func NewSolTransactor(address common.Address, transactor bind.ContractTransactor) (*SolTransactor, error) {
	contract, err := bindSol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolTransactor{contract: contract}, nil
}

// NewSolFilterer creates a new log filterer instance of Sol, bound to a specific deployed contract.
func NewSolFilterer(address common.Address, filterer bind.ContractFilterer) (*SolFilterer, error) {
	contract, err := bindSol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolFilterer{contract: contract}, nil
}

// bindSol binds a generic wrapper to an already deployed contract.
func bindSol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sol *SolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sol.Contract.SolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sol *SolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sol.Contract.SolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sol *SolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sol.Contract.SolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sol *SolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sol *SolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sol *SolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sol.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sol *SolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sol.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sol *SolSession) Owner() (common.Address, error) {
	return _Sol.Contract.Owner(&_Sol.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sol *SolCallerSession) Owner() (common.Address, error) {
	return _Sol.Contract.Owner(&_Sol.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token_, uint256 amount_) returns()
func (_Sol *SolTransactor) Deposit(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Sol.contract.Transact(opts, "deposit", token_, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token_, uint256 amount_) returns()
func (_Sol *SolSession) Deposit(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.Deposit(&_Sol.TransactOpts, token_, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token_, uint256 amount_) returns()
func (_Sol *SolTransactorSession) Deposit(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.Deposit(&_Sol.TransactOpts, token_, amount_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sol *SolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sol.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sol *SolSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sol.Contract.RenounceOwnership(&_Sol.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sol *SolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sol.Contract.RenounceOwnership(&_Sol.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sol *SolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Sol.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sol *SolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sol.Contract.TransferOwnership(&_Sol.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sol *SolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sol.Contract.TransferOwnership(&_Sol.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token_, address recipient_, uint256 amount_) returns()
func (_Sol *SolTransactor) Withdraw(opts *bind.TransactOpts, token_ common.Address, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Sol.contract.Transact(opts, "withdraw", token_, recipient_, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token_, address recipient_, uint256 amount_) returns()
func (_Sol *SolSession) Withdraw(token_ common.Address, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.Withdraw(&_Sol.TransactOpts, token_, recipient_, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token_, address recipient_, uint256 amount_) returns()
func (_Sol *SolTransactorSession) Withdraw(token_ common.Address, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.Withdraw(&_Sol.TransactOpts, token_, recipient_, amount_)
}

// SolDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Sol contract.
type SolDepositedIterator struct {
	Event *SolDeposited // Event containing the contract specifics and raw log

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
func (it *SolDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolDeposited)
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
		it.Event = new(SolDeposited)
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
func (it *SolDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolDeposited represents a Deposited event raised by the Sol contract.
type SolDeposited struct {
	Who          common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed _who, address indexed _tokenAddress, uint256 indexed _amount)
func (_Sol *SolFilterer) FilterDeposited(opts *bind.FilterOpts, _who []common.Address, _tokenAddress []common.Address, _amount []*big.Int) (*SolDepositedIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Sol.contract.FilterLogs(opts, "Deposited", _whoRule, _tokenAddressRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &SolDepositedIterator{contract: _Sol.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed _who, address indexed _tokenAddress, uint256 indexed _amount)
func (_Sol *SolFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *SolDeposited, _who []common.Address, _tokenAddress []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Sol.contract.WatchLogs(opts, "Deposited", _whoRule, _tokenAddressRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolDeposited)
				if err := _Sol.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed _who, address indexed _tokenAddress, uint256 indexed _amount)
func (_Sol *SolFilterer) ParseDeposited(log types.Log) (*SolDeposited, error) {
	event := new(SolDeposited)
	if err := _Sol.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Sol contract.
type SolOwnershipTransferredIterator struct {
	Event *SolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolOwnershipTransferred)
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
		it.Event = new(SolOwnershipTransferred)
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
func (it *SolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolOwnershipTransferred represents a OwnershipTransferred event raised by the Sol contract.
type SolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sol *SolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sol.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolOwnershipTransferredIterator{contract: _Sol.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sol *SolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sol.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolOwnershipTransferred)
				if err := _Sol.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Sol *SolFilterer) ParseOwnershipTransferred(log types.Log) (*SolOwnershipTransferred, error) {
	event := new(SolOwnershipTransferred)
	if err := _Sol.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Sol contract.
type SolWithdrawnIterator struct {
	Event *SolWithdrawn // Event containing the contract specifics and raw log

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
func (it *SolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolWithdrawn)
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
		it.Event = new(SolWithdrawn)
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
func (it *SolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolWithdrawn represents a Withdrawn event raised by the Sol contract.
type SolWithdrawn struct {
	Who          common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed _who, address indexed _tokenAddress, uint256 indexed _amount)
func (_Sol *SolFilterer) FilterWithdrawn(opts *bind.FilterOpts, _who []common.Address, _tokenAddress []common.Address, _amount []*big.Int) (*SolWithdrawnIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Sol.contract.FilterLogs(opts, "Withdrawn", _whoRule, _tokenAddressRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &SolWithdrawnIterator{contract: _Sol.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed _who, address indexed _tokenAddress, uint256 indexed _amount)
func (_Sol *SolFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *SolWithdrawn, _who []common.Address, _tokenAddress []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Sol.contract.WatchLogs(opts, "Withdrawn", _whoRule, _tokenAddressRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolWithdrawn)
				if err := _Sol.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed _who, address indexed _tokenAddress, uint256 indexed _amount)
func (_Sol *SolFilterer) ParseWithdrawn(log types.Log) (*SolWithdrawn, error) {
	event := new(SolWithdrawn)
	if err := _Sol.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
