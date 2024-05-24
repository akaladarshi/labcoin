// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ResearchFormData is an auto generated low-level Go binding around an user-defined struct.
type ResearchFormData struct {
	Link     string
	DataLink string
}

// ResearchMetaData contains all meta data concerning the Research contract.
var ResearchMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dataSetCount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"researcher\",\"type\":\"address\"}],\"name\":\"ResearchRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAllResearches\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"link\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataLink\",\"type\":\"string\"}],\"internalType\":\"structResearch.FormData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_link\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dataLink\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_dataSetCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeDuration\",\"type\":\"uint256\"}],\"name\":\"registerResearch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"researchCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"researches\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataSetCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"researcher\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"link\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataLink\",\"type\":\"string\"}],\"internalType\":\"structResearch.FormData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061150a8061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806308e45a321461004e5780632f3f4b921461006c5780635f7d07041461008a578063d876c21b146100bf575b5f80fd5b6100566100db565b60405161006391906108e2565b60405180910390f35b6100746100e0565b6040516100819190610a81565b60405180910390f35b6100a4600480360381019061009f9190610adc565b6102cb565b6040516100b696959493929190610bcf565b60405180910390f35b6100d960048036038101906100d49190610d6f565b610557565b005b5f5481565b60605f805467ffffffffffffffff8111156100fe576100fd610c4b565b5b60405190808252806020026020018201604052801561013757816020015b6101246108b0565b81526020019060019003908161011c5790505b5090505f600190505b5f5481116102c35760015f8281526020019081526020015f206005016040518060400160405290815f8201805461017690610e95565b80601f01602080910402602001604051908101604052809291908181526020018280546101a290610e95565b80156101ed5780601f106101c4576101008083540402835291602001916101ed565b820191905f5260205f20905b8154815290600101906020018083116101d057829003601f168201915b5050505050815260200160018201805461020690610e95565b80601f016020809104026020016040519081016040528092919081815260200182805461023290610e95565b801561027d5780601f106102545761010080835404028352916020019161027d565b820191905f5260205f20905b81548152906001019060200180831161026057829003601f168201915b505050505081525050826001836102949190610ef2565b815181106102a5576102a4610f25565b5b602002602001018190525080806102bb90610f52565b915050610140565b508091505090565b6001602052805f5260405f205f91509050805f0180546102ea90610e95565b80601f016020809104026020016040519081016040528092919081815260200182805461031690610e95565b80156103615780601f1061033857610100808354040283529160200191610361565b820191905f5260205f20905b81548152906001019060200180831161034457829003601f168201915b50505050509080600101805461037690610e95565b80601f01602080910402602001604051908101604052809291908181526020018280546103a290610e95565b80156103ed5780601f106103c4576101008083540402835291602001916103ed565b820191905f5260205f20905b8154815290600101906020018083116103d057829003601f168201915b505050505090806002015490806003015490806004015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005016040518060400160405290815f8201805461044290610e95565b80601f016020809104026020016040519081016040528092919081815260200182805461046e90610e95565b80156104b95780601f10610490576101008083540402835291602001916104b9565b820191905f5260205f20905b81548152906001019060200180831161049c57829003601f168201915b505050505081526020016001820180546104d290610e95565b80601f01602080910402602001604051908101604052809291908181526020018280546104fe90610e95565b80156105495780601f1061052057610100808354040283529160200191610549565b820191905f5260205f20905b81548152906001019060200180831161052c57829003601f168201915b505050505081525050905086565b5f8081548092919061056890610f52565b91905055505f6040518060c001604052808881526020018781526020018381526020018481526020013373ffffffffffffffffffffffffffffffffffffffff16815260200160405180604001604052808881526020018781525081525090506105d0816106ff565b8060015f805481526020019081526020015f205f820151815f0190816105f69190611136565b50602082015181600101908161060c9190611136565b5060408201518160020155606082015181600301556080820151816004015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816005015f820151815f0190816106839190611136565b5060208201518160010190816106999190611136565b5050509050503373ffffffffffffffffffffffffffffffffffffffff167fd0e656448754c6135e780b2e4793c864bc1064195f25a3d1e07ee989b90447015f54898986886040516106ee959493929190611205565b60405180910390a250505050505050565b5f815f01515111610745576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073c906112ae565b60405180910390fd5b5f8160200151511161078c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078390611316565b60405180910390fd5b5f8160400151116107d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c99061137e565b60405180910390fd5b5f816060015111610818576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080f906113e6565b60405180910390fd5b5f8160a001515f01515111610862576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108599061144e565b60405180910390fd5b5f8160a001516020015151116108ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a4906114b6565b60405180910390fd5b50565b604051806040016040528060608152602001606081525090565b5f819050919050565b6108dc816108ca565b82525050565b5f6020820190506108f55f8301846108d3565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561095b578082015181840152602081019050610940565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61098082610924565b61098a818561092e565b935061099a81856020860161093e565b6109a381610966565b840191505092915050565b5f604083015f8301518482035f8601526109c88282610976565b915050602083015184820360208601526109e28282610976565b9150508091505092915050565b5f6109fa83836109ae565b905092915050565b5f602082019050919050565b5f610a18826108fb565b610a228185610905565b935083602082028501610a3485610915565b805f5b85811015610a6f5784840389528151610a5085826109ef565b9450610a5b83610a02565b925060208a01995050600181019050610a37565b50829750879550505050505092915050565b5f6020820190508181035f830152610a998184610a0e565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b610abb816108ca565b8114610ac5575f80fd5b50565b5f81359050610ad681610ab2565b92915050565b5f60208284031215610af157610af0610aaa565b5b5f610afe84828501610ac8565b91505092915050565b5f82825260208201905092915050565b5f610b2182610924565b610b2b8185610b07565b9350610b3b81856020860161093e565b610b4481610966565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b7882610b4f565b9050919050565b610b8881610b6e565b82525050565b5f604083015f8301518482035f860152610ba88282610976565b91505060208301518482036020860152610bc28282610976565b9150508091505092915050565b5f60c0820190508181035f830152610be78189610b17565b90508181036020830152610bfb8188610b17565b9050610c0a60408301876108d3565b610c1760608301866108d3565b610c246080830185610b7f565b81810360a0830152610c368184610b8e565b9050979650505050505050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610c8182610966565b810181811067ffffffffffffffff82111715610ca057610c9f610c4b565b5b80604052505050565b5f610cb2610aa1565b9050610cbe8282610c78565b919050565b5f67ffffffffffffffff821115610cdd57610cdc610c4b565b5b610ce682610966565b9050602081019050919050565b828183375f83830152505050565b5f610d13610d0e84610cc3565b610ca9565b905082815260208101848484011115610d2f57610d2e610c47565b5b610d3a848285610cf3565b509392505050565b5f82601f830112610d5657610d55610c43565b5b8135610d66848260208601610d01565b91505092915050565b5f805f805f8060c08789031215610d8957610d88610aaa565b5b5f87013567ffffffffffffffff811115610da657610da5610aae565b5b610db289828a01610d42565b965050602087013567ffffffffffffffff811115610dd357610dd2610aae565b5b610ddf89828a01610d42565b955050604087013567ffffffffffffffff811115610e0057610dff610aae565b5b610e0c89828a01610d42565b945050606087013567ffffffffffffffff811115610e2d57610e2c610aae565b5b610e3989828a01610d42565b9350506080610e4a89828a01610ac8565b92505060a0610e5b89828a01610ac8565b9150509295509295509295565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610eac57607f821691505b602082108103610ebf57610ebe610e68565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610efc826108ca565b9150610f07836108ca565b9250828203905081811115610f1f57610f1e610ec5565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f610f5c826108ca565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f8e57610f8d610ec5565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610ff57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610fba565b610fff8683610fba565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61103a611035611030846108ca565b611017565b6108ca565b9050919050565b5f819050919050565b61105383611020565b61106761105f82611041565b848454610fc6565b825550505050565b5f90565b61107b61106f565b61108681848461104a565b505050565b5b818110156110a95761109e5f82611073565b60018101905061108c565b5050565b601f8211156110ee576110bf81610f99565b6110c884610fab565b810160208510156110d7578190505b6110eb6110e385610fab565b83018261108b565b50505b505050565b5f82821c905092915050565b5f61110e5f19846008026110f3565b1980831691505092915050565b5f61112683836110ff565b9150826002028217905092915050565b61113f82610924565b67ffffffffffffffff81111561115857611157610c4b565b5b6111628254610e95565b61116d8282856110ad565b5f60209050601f83116001811461119e575f841561118c578287015190505b611196858261111b565b8655506111fd565b601f1984166111ac86610f99565b5f5b828110156111d3578489015182556001820191506020850194506020810190506111ae565b868310156111f057848901516111ec601f8916826110ff565b8355505b6001600288020188555050505b505050505050565b5f60a0820190506112185f8301886108d3565b818103602083015261122a8187610b17565b9050818103604083015261123e8186610b17565b905061124d60608301856108d3565b61125a60808301846108d3565b9695505050505050565b7f5469746c652069732072657175697265640000000000000000000000000000005f82015250565b5f611298601183610b07565b91506112a382611264565b602082019050919050565b5f6020820190508181035f8301526112c58161128c565b9050919050565b7f4465736372697074696f6e2069732072657175697265640000000000000000005f82015250565b5f611300601783610b07565b915061130b826112cc565b602082019050919050565b5f6020820190508181035f83015261132d816112f4565b9050919050565b7f54696d65206475726174696f6e206973207265717569726564000000000000005f82015250565b5f611368601983610b07565b915061137382611334565b602082019050919050565b5f6020820190508181035f8301526113958161135c565b9050919050565b7f446174612073657420636f756e742069732072657175697265640000000000005f82015250565b5f6113d0601a83610b07565b91506113db8261139c565b602082019050919050565b5f6020820190508181035f8301526113fd816113c4565b9050919050565b7f4c696e6b206973207265717569726564000000000000000000000000000000005f82015250565b5f611438601083610b07565b915061144382611404565b602082019050919050565b5f6020820190508181035f8301526114658161142c565b9050919050565b7f44617461206c696e6b20697320726571756972656400000000000000000000005f82015250565b5f6114a0601583610b07565b91506114ab8261146c565b602082019050919050565b5f6020820190508181035f8301526114cd81611494565b905091905056fea264697066735822122020f886a27e1fa987579aa9962bf2ee080dde5649c3abe84253546de30dd5d51d64736f6c63430008150033",
}

// ResearchABI is the input ABI used to generate the binding from.
// Deprecated: Use ResearchMetaData.ABI instead.
var ResearchABI = ResearchMetaData.ABI

// ResearchBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ResearchMetaData.Bin instead.
var ResearchBin = ResearchMetaData.Bin

// DeployResearch deploys a new Ethereum contract, binding an instance of Research to it.
func DeployResearch(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Research, error) {
	parsed, err := ResearchMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ResearchBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Research{ResearchCaller: ResearchCaller{contract: contract}, ResearchTransactor: ResearchTransactor{contract: contract}, ResearchFilterer: ResearchFilterer{contract: contract}}, nil
}

// Research is an auto generated Go binding around an Ethereum contract.
type Research struct {
	ResearchCaller     // Read-only binding to the contract
	ResearchTransactor // Write-only binding to the contract
	ResearchFilterer   // Log filterer for contract events
}

// ResearchCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResearchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResearchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResearchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResearchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResearchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResearchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResearchSession struct {
	Contract     *Research         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResearchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResearchCallerSession struct {
	Contract *ResearchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ResearchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResearchTransactorSession struct {
	Contract     *ResearchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ResearchRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResearchRaw struct {
	Contract *Research // Generic contract binding to access the raw methods on
}

// ResearchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResearchCallerRaw struct {
	Contract *ResearchCaller // Generic read-only contract binding to access the raw methods on
}

// ResearchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResearchTransactorRaw struct {
	Contract *ResearchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResearch creates a new instance of Research, bound to a specific deployed contract.
func NewResearch(address common.Address, backend bind.ContractBackend) (*Research, error) {
	contract, err := bindResearch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Research{ResearchCaller: ResearchCaller{contract: contract}, ResearchTransactor: ResearchTransactor{contract: contract}, ResearchFilterer: ResearchFilterer{contract: contract}}, nil
}

// NewResearchCaller creates a new read-only instance of Research, bound to a specific deployed contract.
func NewResearchCaller(address common.Address, caller bind.ContractCaller) (*ResearchCaller, error) {
	contract, err := bindResearch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResearchCaller{contract: contract}, nil
}

// NewResearchTransactor creates a new write-only instance of Research, bound to a specific deployed contract.
func NewResearchTransactor(address common.Address, transactor bind.ContractTransactor) (*ResearchTransactor, error) {
	contract, err := bindResearch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResearchTransactor{contract: contract}, nil
}

// NewResearchFilterer creates a new log filterer instance of Research, bound to a specific deployed contract.
func NewResearchFilterer(address common.Address, filterer bind.ContractFilterer) (*ResearchFilterer, error) {
	contract, err := bindResearch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResearchFilterer{contract: contract}, nil
}

// bindResearch binds a generic wrapper to an already deployed contract.
func bindResearch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResearchMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Research *ResearchRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Research.Contract.ResearchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Research *ResearchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Research.Contract.ResearchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Research *ResearchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Research.Contract.ResearchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Research *ResearchCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Research.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Research *ResearchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Research.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Research *ResearchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Research.Contract.contract.Transact(opts, method, params...)
}

// GetAllResearches is a free data retrieval call binding the contract method 0x2f3f4b92.
//
// Solidity: function getAllResearches() view returns((string,string)[])
func (_Research *ResearchCaller) GetAllResearches(opts *bind.CallOpts) ([]ResearchFormData, error) {
	var out []interface{}
	err := _Research.contract.Call(opts, &out, "getAllResearches")

	if err != nil {
		return *new([]ResearchFormData), err
	}

	out0 := *abi.ConvertType(out[0], new([]ResearchFormData)).(*[]ResearchFormData)

	return out0, err

}

// GetAllResearches is a free data retrieval call binding the contract method 0x2f3f4b92.
//
// Solidity: function getAllResearches() view returns((string,string)[])
func (_Research *ResearchSession) GetAllResearches() ([]ResearchFormData, error) {
	return _Research.Contract.GetAllResearches(&_Research.CallOpts)
}

// GetAllResearches is a free data retrieval call binding the contract method 0x2f3f4b92.
//
// Solidity: function getAllResearches() view returns((string,string)[])
func (_Research *ResearchCallerSession) GetAllResearches() ([]ResearchFormData, error) {
	return _Research.Contract.GetAllResearches(&_Research.CallOpts)
}

// ResearchCounter is a free data retrieval call binding the contract method 0x08e45a32.
//
// Solidity: function researchCounter() view returns(uint256)
func (_Research *ResearchCaller) ResearchCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Research.contract.Call(opts, &out, "researchCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResearchCounter is a free data retrieval call binding the contract method 0x08e45a32.
//
// Solidity: function researchCounter() view returns(uint256)
func (_Research *ResearchSession) ResearchCounter() (*big.Int, error) {
	return _Research.Contract.ResearchCounter(&_Research.CallOpts)
}

// ResearchCounter is a free data retrieval call binding the contract method 0x08e45a32.
//
// Solidity: function researchCounter() view returns(uint256)
func (_Research *ResearchCallerSession) ResearchCounter() (*big.Int, error) {
	return _Research.Contract.ResearchCounter(&_Research.CallOpts)
}

// Researches is a free data retrieval call binding the contract method 0x5f7d0704.
//
// Solidity: function researches(uint256 ) view returns(string title, string description, uint256 timeDuration, uint256 dataSetCount, address researcher, (string,string) data)
func (_Research *ResearchCaller) Researches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Title        string
	Description  string
	TimeDuration *big.Int
	DataSetCount *big.Int
	Researcher   common.Address
	Data         ResearchFormData
}, error) {
	var out []interface{}
	err := _Research.contract.Call(opts, &out, "researches", arg0)

	outstruct := new(struct {
		Title        string
		Description  string
		TimeDuration *big.Int
		DataSetCount *big.Int
		Researcher   common.Address
		Data         ResearchFormData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.TimeDuration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DataSetCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Researcher = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Data = *abi.ConvertType(out[5], new(ResearchFormData)).(*ResearchFormData)

	return *outstruct, err

}

// Researches is a free data retrieval call binding the contract method 0x5f7d0704.
//
// Solidity: function researches(uint256 ) view returns(string title, string description, uint256 timeDuration, uint256 dataSetCount, address researcher, (string,string) data)
func (_Research *ResearchSession) Researches(arg0 *big.Int) (struct {
	Title        string
	Description  string
	TimeDuration *big.Int
	DataSetCount *big.Int
	Researcher   common.Address
	Data         ResearchFormData
}, error) {
	return _Research.Contract.Researches(&_Research.CallOpts, arg0)
}

// Researches is a free data retrieval call binding the contract method 0x5f7d0704.
//
// Solidity: function researches(uint256 ) view returns(string title, string description, uint256 timeDuration, uint256 dataSetCount, address researcher, (string,string) data)
func (_Research *ResearchCallerSession) Researches(arg0 *big.Int) (struct {
	Title        string
	Description  string
	TimeDuration *big.Int
	DataSetCount *big.Int
	Researcher   common.Address
	Data         ResearchFormData
}, error) {
	return _Research.Contract.Researches(&_Research.CallOpts, arg0)
}

// RegisterResearch is a paid mutator transaction binding the contract method 0xd876c21b.
//
// Solidity: function registerResearch(string _title, string _description, string _link, string _dataLink, uint256 _dataSetCount, uint256 _timeDuration) returns()
func (_Research *ResearchTransactor) RegisterResearch(opts *bind.TransactOpts, _title string, _description string, _link string, _dataLink string, _dataSetCount *big.Int, _timeDuration *big.Int) (*types.Transaction, error) {
	return _Research.contract.Transact(opts, "registerResearch", _title, _description, _link, _dataLink, _dataSetCount, _timeDuration)
}

// RegisterResearch is a paid mutator transaction binding the contract method 0xd876c21b.
//
// Solidity: function registerResearch(string _title, string _description, string _link, string _dataLink, uint256 _dataSetCount, uint256 _timeDuration) returns()
func (_Research *ResearchSession) RegisterResearch(_title string, _description string, _link string, _dataLink string, _dataSetCount *big.Int, _timeDuration *big.Int) (*types.Transaction, error) {
	return _Research.Contract.RegisterResearch(&_Research.TransactOpts, _title, _description, _link, _dataLink, _dataSetCount, _timeDuration)
}

// RegisterResearch is a paid mutator transaction binding the contract method 0xd876c21b.
//
// Solidity: function registerResearch(string _title, string _description, string _link, string _dataLink, uint256 _dataSetCount, uint256 _timeDuration) returns()
func (_Research *ResearchTransactorSession) RegisterResearch(_title string, _description string, _link string, _dataLink string, _dataSetCount *big.Int, _timeDuration *big.Int) (*types.Transaction, error) {
	return _Research.Contract.RegisterResearch(&_Research.TransactOpts, _title, _description, _link, _dataLink, _dataSetCount, _timeDuration)
}

// ResearchResearchRegisteredIterator is returned from FilterResearchRegistered and is used to iterate over the raw logs and unpacked data for ResearchRegistered events raised by the Research contract.
type ResearchResearchRegisteredIterator struct {
	Event *ResearchResearchRegistered // Event containing the contract specifics and raw log

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
func (it *ResearchResearchRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResearchResearchRegistered)
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
		it.Event = new(ResearchResearchRegistered)
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
func (it *ResearchResearchRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResearchResearchRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResearchResearchRegistered represents a ResearchRegistered event raised by the Research contract.
type ResearchResearchRegistered struct {
	Id           *big.Int
	Title        string
	Description  string
	TimeDuration *big.Int
	DataSetCount *big.Int
	Researcher   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterResearchRegistered is a free log retrieval operation binding the contract event 0xd0e656448754c6135e780b2e4793c864bc1064195f25a3d1e07ee989b9044701.
//
// Solidity: event ResearchRegistered(uint256 id, string title, string description, uint256 timeDuration, uint256 dataSetCount, address indexed researcher)
func (_Research *ResearchFilterer) FilterResearchRegistered(opts *bind.FilterOpts, researcher []common.Address) (*ResearchResearchRegisteredIterator, error) {

	var researcherRule []interface{}
	for _, researcherItem := range researcher {
		researcherRule = append(researcherRule, researcherItem)
	}

	logs, sub, err := _Research.contract.FilterLogs(opts, "ResearchRegistered", researcherRule)
	if err != nil {
		return nil, err
	}
	return &ResearchResearchRegisteredIterator{contract: _Research.contract, event: "ResearchRegistered", logs: logs, sub: sub}, nil
}

// WatchResearchRegistered is a free log subscription operation binding the contract event 0xd0e656448754c6135e780b2e4793c864bc1064195f25a3d1e07ee989b9044701.
//
// Solidity: event ResearchRegistered(uint256 id, string title, string description, uint256 timeDuration, uint256 dataSetCount, address indexed researcher)
func (_Research *ResearchFilterer) WatchResearchRegistered(opts *bind.WatchOpts, sink chan<- *ResearchResearchRegistered, researcher []common.Address) (event.Subscription, error) {

	var researcherRule []interface{}
	for _, researcherItem := range researcher {
		researcherRule = append(researcherRule, researcherItem)
	}

	logs, sub, err := _Research.contract.WatchLogs(opts, "ResearchRegistered", researcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResearchResearchRegistered)
				if err := _Research.contract.UnpackLog(event, "ResearchRegistered", log); err != nil {
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

// ParseResearchRegistered is a log parse operation binding the contract event 0xd0e656448754c6135e780b2e4793c864bc1064195f25a3d1e07ee989b9044701.
//
// Solidity: event ResearchRegistered(uint256 id, string title, string description, uint256 timeDuration, uint256 dataSetCount, address indexed researcher)
func (_Research *ResearchFilterer) ParseResearchRegistered(log types.Log) (*ResearchResearchRegistered, error) {
	event := new(ResearchResearchRegistered)
	if err := _Research.contract.UnpackLog(event, "ResearchRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
