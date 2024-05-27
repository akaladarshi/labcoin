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
	FormLink        string
	SpreadSheetID   string
	SheetID         *big.Int
	MaxDataSetCount *big.Int
}

// ResearchMetaData contains all meta data concerning the Research contract.
var ResearchMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxDataSetCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"formLink\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"researcher\",\"type\":\"address\"}],\"name\":\"ResearchRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAllFormsDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"formLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"spreadSheetID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sheetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataSetCount\",\"type\":\"uint256\"}],\"internalType\":\"structResearch.FormData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_formLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_spreadSheetID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_sheetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDataSetCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeDuration\",\"type\":\"uint256\"}],\"name\":\"registerResearch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"researchCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"researches\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timeDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"researcher\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"formLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"spreadSheetID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sheetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataSetCount\",\"type\":\"uint256\"}],\"internalType\":\"structResearch.FormData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506116768061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806308e45a321461004e5780635f7d07041461006c5780638f8f4174146100a0578063d7696ff1146100be575b5f80fd5b6100566100da565b6040516100639190610970565b60405180910390f35b610086600480360381019061008191906109c4565b6100df565b604051610097959493929190610b76565b60405180910390f35b6100a8610379565b6040516100b59190610cfe565b60405180910390f35b6100d860048036038101906100d39190610e4a565b610578565b005b5f5481565b6001602052805f5260405f205f91509050805f0180546100fe90610f84565b80601f016020809104026020016040519081016040528092919081815260200182805461012a90610f84565b80156101755780601f1061014c57610100808354040283529160200191610175565b820191905f5260205f20905b81548152906001019060200180831161015857829003601f168201915b50505050509080600101805461018a90610f84565b80601f01602080910402602001604051908101604052809291908181526020018280546101b690610f84565b80156102015780601f106101d857610100808354040283529160200191610201565b820191905f5260205f20905b8154815290600101906020018083116101e457829003601f168201915b505050505090806002015490806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004016040518060800160405290815f8201805461025090610f84565b80601f016020809104026020016040519081016040528092919081815260200182805461027c90610f84565b80156102c75780601f1061029e576101008083540402835291602001916102c7565b820191905f5260205f20905b8154815290600101906020018083116102aa57829003601f168201915b505050505081526020016001820180546102e090610f84565b80601f016020809104026020016040519081016040528092919081815260200182805461030c90610f84565b80156103575780601f1061032e57610100808354040283529160200191610357565b820191905f5260205f20905b81548152906001019060200180831161033a57829003601f168201915b5050505050815260200160028201548152602001600382015481525050905085565b60605f805467ffffffffffffffff81111561039757610396610d26565b5b6040519080825280602002602001820160405280156103d057816020015b6103bd610932565b8152602001906001900390816103b55790505b5090505f600190505b5f5481116105705760015f8281526020019081526020015f206004016040518060800160405290815f8201805461040f90610f84565b80601f016020809104026020016040519081016040528092919081815260200182805461043b90610f84565b80156104865780601f1061045d57610100808354040283529160200191610486565b820191905f5260205f20905b81548152906001019060200180831161046957829003601f168201915b5050505050815260200160018201805461049f90610f84565b80601f01602080910402602001604051908101604052809291908181526020018280546104cb90610f84565b80156105165780601f106104ed57610100808354040283529160200191610516565b820191905f5260205f20905b8154815290600101906020018083116104f957829003601f168201915b5050505050815260200160028201548152602001600382015481525050826001836105419190610fe1565b8151811061055257610551611014565b5b6020026020010181905250808061056890611041565b9150506103d9565b508091505090565b5f8081548092919061058990611041565b91905055505f6040518060a001604052808981526020018881526020018381526020013373ffffffffffffffffffffffffffffffffffffffff16815260200160405180608001604052808981526020018881526020018781526020018681525081525090506105f781610733565b8060015f805481526020019081526020015f205f820151815f01908161061d9190611225565b5060208201518160010190816106339190611225565b50604082015181600201556060820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015f820151815f0190816106a09190611225565b5060208201518160010190816106b69190611225565b50604082015181600201556060820151816003015550509050503373ffffffffffffffffffffffffffffffffffffffff167fab440c8e20c74e7307954c096ee1222edbdd8dde5ba095fec9bf49ce36aeb9e85f548a8a86888c604051610721969594939291906112f4565b60405180910390a25050505050505050565b5f815f01515111610779576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610770906113b2565b60405180910390fd5b5f816020015151116107c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b79061141a565b60405180910390fd5b5f816040015111610806576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107fd90611482565b60405180910390fd5b5f81608001516060015111610850576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610847906114ea565b60405180910390fd5b5f81608001515f0151511161089a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089190611552565b60405180910390fd5b5f816080015160400151116108e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108db906115ba565b60405180910390fd5b5f816080015160200151511161092f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092690611622565b60405180910390fd5b50565b604051806080016040528060608152602001606081526020015f81526020015f81525090565b5f819050919050565b61096a81610958565b82525050565b5f6020820190506109835f830184610961565b92915050565b5f604051905090565b5f80fd5b5f80fd5b6109a381610958565b81146109ad575f80fd5b50565b5f813590506109be8161099a565b92915050565b5f602082840312156109d9576109d8610992565b5b5f6109e6848285016109b0565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610a26578082015181840152602081019050610a0b565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610a4b826109ef565b610a5581856109f9565b9350610a65818560208601610a09565b610a6e81610a31565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610aa282610a79565b9050919050565b610ab281610a98565b82525050565b5f82825260208201905092915050565b5f610ad2826109ef565b610adc8185610ab8565b9350610aec818560208601610a09565b610af581610a31565b840191505092915050565b610b0981610958565b82525050565b5f608083015f8301518482035f860152610b298282610ac8565b91505060208301518482036020860152610b438282610ac8565b9150506040830151610b586040860182610b00565b506060830151610b6b6060860182610b00565b508091505092915050565b5f60a0820190508181035f830152610b8e8188610a41565b90508181036020830152610ba28187610a41565b9050610bb16040830186610961565b610bbe6060830185610aa9565b8181036080830152610bd08184610b0f565b90509695505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f608083015f8301518482035f860152610c1f8282610ac8565b91505060208301518482036020860152610c398282610ac8565b9150506040830151610c4e6040860182610b00565b506060830151610c616060860182610b00565b508091505092915050565b5f610c778383610c05565b905092915050565b5f602082019050919050565b5f610c9582610bdc565b610c9f8185610be6565b935083602082028501610cb185610bf6565b805f5b85811015610cec5784840389528151610ccd8582610c6c565b9450610cd883610c7f565b925060208a01995050600181019050610cb4565b50829750879550505050505092915050565b5f6020820190508181035f830152610d168184610c8b565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610d5c82610a31565b810181811067ffffffffffffffff82111715610d7b57610d7a610d26565b5b80604052505050565b5f610d8d610989565b9050610d998282610d53565b919050565b5f67ffffffffffffffff821115610db857610db7610d26565b5b610dc182610a31565b9050602081019050919050565b828183375f83830152505050565b5f610dee610de984610d9e565b610d84565b905082815260208101848484011115610e0a57610e09610d22565b5b610e15848285610dce565b509392505050565b5f82601f830112610e3157610e30610d1e565b5b8135610e41848260208601610ddc565b91505092915050565b5f805f805f805f60e0888a031215610e6557610e64610992565b5b5f88013567ffffffffffffffff811115610e8257610e81610996565b5b610e8e8a828b01610e1d565b975050602088013567ffffffffffffffff811115610eaf57610eae610996565b5b610ebb8a828b01610e1d565b965050604088013567ffffffffffffffff811115610edc57610edb610996565b5b610ee88a828b01610e1d565b955050606088013567ffffffffffffffff811115610f0957610f08610996565b5b610f158a828b01610e1d565b9450506080610f268a828b016109b0565b93505060a0610f378a828b016109b0565b92505060c0610f488a828b016109b0565b91505092959891949750929550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610f9b57607f821691505b602082108103610fae57610fad610f57565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610feb82610958565b9150610ff683610958565b925082820390508181111561100e5761100d610fb4565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61104b82610958565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361107d5761107c610fb4565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026110e47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826110a9565b6110ee86836110a9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61112961112461111f84610958565b611106565b610958565b9050919050565b5f819050919050565b6111428361110f565b61115661114e82611130565b8484546110b5565b825550505050565b5f90565b61116a61115e565b611175818484611139565b505050565b5b818110156111985761118d5f82611162565b60018101905061117b565b5050565b601f8211156111dd576111ae81611088565b6111b78461109a565b810160208510156111c6578190505b6111da6111d28561109a565b83018261117a565b50505b505050565b5f82821c905092915050565b5f6111fd5f19846008026111e2565b1980831691505092915050565b5f61121583836111ee565b9150826002028217905092915050565b61122e826109ef565b67ffffffffffffffff81111561124757611246610d26565b5b6112518254610f84565b61125c82828561119c565b5f60209050601f83116001811461128d575f841561127b578287015190505b611285858261120a565b8655506112ec565b601f19841661129b86611088565b5f5b828110156112c25784890151825560018201915060208501945060208101905061129d565b868310156112df57848901516112db601f8916826111ee565b8355505b6001600288020188555050505b505050505050565b5f60c0820190506113075f830189610961565b81810360208301526113198188610a41565b9050818103604083015261132d8187610a41565b905061133c6060830186610961565b6113496080830185610961565b81810360a083015261135b8184610a41565b9050979650505050505050565b7f5469746c652069732072657175697265640000000000000000000000000000005f82015250565b5f61139c6011836109f9565b91506113a782611368565b602082019050919050565b5f6020820190508181035f8301526113c981611390565b9050919050565b7f4465736372697074696f6e2069732072657175697265640000000000000000005f82015250565b5f6114046017836109f9565b915061140f826113d0565b602082019050919050565b5f6020820190508181035f830152611431816113f8565b9050919050565b7f54696d65206475726174696f6e206973207265717569726564000000000000005f82015250565b5f61146c6019836109f9565b915061147782611438565b602082019050919050565b5f6020820190508181035f83015261149981611460565b9050919050565b7f4d617820446174612073657420636f756e7420697320726571756972656400005f82015250565b5f6114d4601e836109f9565b91506114df826114a0565b602082019050919050565b5f6020820190508181035f830152611501816114c8565b9050919050565b7f466f726d204c696e6b20697320726571756972656400000000000000000000005f82015250565b5f61153c6015836109f9565b915061154782611508565b602082019050919050565b5f6020820190508181035f83015261156981611530565b9050919050565b7f53686565742049442069732072657175697265640000000000000000000000005f82015250565b5f6115a46014836109f9565b91506115af82611570565b602082019050919050565b5f6020820190508181035f8301526115d181611598565b9050919050565b7f53707265616420736865657420696420697320726571756972656400000000005f82015250565b5f61160c601b836109f9565b9150611617826115d8565b602082019050919050565b5f6020820190508181035f83015261163981611600565b905091905056fea2646970667358221220bf95f1872f0227d89363a46cdef83166a7a35db986e72dab047bf559224b554164736f6c63430008150033",
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

// GetAllFormsDetails is a free data retrieval call binding the contract method 0x8f8f4174.
//
// Solidity: function getAllFormsDetails() view returns((string,string,uint256,uint256)[])
func (_Research *ResearchCaller) GetAllFormsDetails(opts *bind.CallOpts) ([]ResearchFormData, error) {
	var out []interface{}
	err := _Research.contract.Call(opts, &out, "getAllFormsDetails")

	if err != nil {
		return *new([]ResearchFormData), err
	}

	out0 := *abi.ConvertType(out[0], new([]ResearchFormData)).(*[]ResearchFormData)

	return out0, err

}

// GetAllFormsDetails is a free data retrieval call binding the contract method 0x8f8f4174.
//
// Solidity: function getAllFormsDetails() view returns((string,string,uint256,uint256)[])
func (_Research *ResearchSession) GetAllFormsDetails() ([]ResearchFormData, error) {
	return _Research.Contract.GetAllFormsDetails(&_Research.CallOpts)
}

// GetAllFormsDetails is a free data retrieval call binding the contract method 0x8f8f4174.
//
// Solidity: function getAllFormsDetails() view returns((string,string,uint256,uint256)[])
func (_Research *ResearchCallerSession) GetAllFormsDetails() ([]ResearchFormData, error) {
	return _Research.Contract.GetAllFormsDetails(&_Research.CallOpts)
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
// Solidity: function researches(uint256 ) view returns(string title, string description, uint256 timeDuration, address researcher, (string,string,uint256,uint256) data)
func (_Research *ResearchCaller) Researches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Title        string
	Description  string
	TimeDuration *big.Int
	Researcher   common.Address
	Data         ResearchFormData
}, error) {
	var out []interface{}
	err := _Research.contract.Call(opts, &out, "researches", arg0)

	outstruct := new(struct {
		Title        string
		Description  string
		TimeDuration *big.Int
		Researcher   common.Address
		Data         ResearchFormData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.TimeDuration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Researcher = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Data = *abi.ConvertType(out[4], new(ResearchFormData)).(*ResearchFormData)

	return *outstruct, err

}

// Researches is a free data retrieval call binding the contract method 0x5f7d0704.
//
// Solidity: function researches(uint256 ) view returns(string title, string description, uint256 timeDuration, address researcher, (string,string,uint256,uint256) data)
func (_Research *ResearchSession) Researches(arg0 *big.Int) (struct {
	Title        string
	Description  string
	TimeDuration *big.Int
	Researcher   common.Address
	Data         ResearchFormData
}, error) {
	return _Research.Contract.Researches(&_Research.CallOpts, arg0)
}

// Researches is a free data retrieval call binding the contract method 0x5f7d0704.
//
// Solidity: function researches(uint256 ) view returns(string title, string description, uint256 timeDuration, address researcher, (string,string,uint256,uint256) data)
func (_Research *ResearchCallerSession) Researches(arg0 *big.Int) (struct {
	Title        string
	Description  string
	TimeDuration *big.Int
	Researcher   common.Address
	Data         ResearchFormData
}, error) {
	return _Research.Contract.Researches(&_Research.CallOpts, arg0)
}

// RegisterResearch is a paid mutator transaction binding the contract method 0xd7696ff1.
//
// Solidity: function registerResearch(string _title, string _description, string _formLink, string _spreadSheetID, uint256 _sheetID, uint256 _maxDataSetCount, uint256 _timeDuration) returns()
func (_Research *ResearchTransactor) RegisterResearch(opts *bind.TransactOpts, _title string, _description string, _formLink string, _spreadSheetID string, _sheetID *big.Int, _maxDataSetCount *big.Int, _timeDuration *big.Int) (*types.Transaction, error) {
	return _Research.contract.Transact(opts, "registerResearch", _title, _description, _formLink, _spreadSheetID, _sheetID, _maxDataSetCount, _timeDuration)
}

// RegisterResearch is a paid mutator transaction binding the contract method 0xd7696ff1.
//
// Solidity: function registerResearch(string _title, string _description, string _formLink, string _spreadSheetID, uint256 _sheetID, uint256 _maxDataSetCount, uint256 _timeDuration) returns()
func (_Research *ResearchSession) RegisterResearch(_title string, _description string, _formLink string, _spreadSheetID string, _sheetID *big.Int, _maxDataSetCount *big.Int, _timeDuration *big.Int) (*types.Transaction, error) {
	return _Research.Contract.RegisterResearch(&_Research.TransactOpts, _title, _description, _formLink, _spreadSheetID, _sheetID, _maxDataSetCount, _timeDuration)
}

// RegisterResearch is a paid mutator transaction binding the contract method 0xd7696ff1.
//
// Solidity: function registerResearch(string _title, string _description, string _formLink, string _spreadSheetID, uint256 _sheetID, uint256 _maxDataSetCount, uint256 _timeDuration) returns()
func (_Research *ResearchTransactorSession) RegisterResearch(_title string, _description string, _formLink string, _spreadSheetID string, _sheetID *big.Int, _maxDataSetCount *big.Int, _timeDuration *big.Int) (*types.Transaction, error) {
	return _Research.Contract.RegisterResearch(&_Research.TransactOpts, _title, _description, _formLink, _spreadSheetID, _sheetID, _maxDataSetCount, _timeDuration)
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
	Id              *big.Int
	Title           string
	Description     string
	TimeDuration    *big.Int
	MaxDataSetCount *big.Int
	FormLink        string
	Researcher      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterResearchRegistered is a free log retrieval operation binding the contract event 0xab440c8e20c74e7307954c096ee1222edbdd8dde5ba095fec9bf49ce36aeb9e8.
//
// Solidity: event ResearchRegistered(uint256 id, string title, string description, uint256 timeDuration, uint256 maxDataSetCount, string formLink, address indexed researcher)
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

// WatchResearchRegistered is a free log subscription operation binding the contract event 0xab440c8e20c74e7307954c096ee1222edbdd8dde5ba095fec9bf49ce36aeb9e8.
//
// Solidity: event ResearchRegistered(uint256 id, string title, string description, uint256 timeDuration, uint256 maxDataSetCount, string formLink, address indexed researcher)
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

// ParseResearchRegistered is a log parse operation binding the contract event 0xab440c8e20c74e7307954c096ee1222edbdd8dde5ba095fec9bf49ce36aeb9e8.
//
// Solidity: event ResearchRegistered(uint256 id, string title, string description, uint256 timeDuration, uint256 maxDataSetCount, string formLink, address indexed researcher)
func (_Research *ResearchFilterer) ParseResearchRegistered(log types.Log) (*ResearchResearchRegistered, error) {
	event := new(ResearchResearchRegistered)
	if err := _Research.contract.UnpackLog(event, "ResearchRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
