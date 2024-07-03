// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_favouriteNumber\",\"type\":\"uint256\"}],\"name\":\"addPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nameToFavouriteNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"people\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"favouriteNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_favouriteNumber\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506108658061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80632e64cec1146100595780636057361d146100775780636f760f41146100935780639e7a13ad146100af578063b2ac62ef146100e0575b5f80fd5b610061610110565b60405161006e919061027c565b60405180910390f35b610091600480360381019061008c91906102d0565b610118565b005b6100ad60048036038101906100a89190610437565b610121565b005b6100c960048036038101906100c491906102d0565b610182565b6040516100d79291906104f1565b60405180910390f35b6100fa60048036038101906100f5919061051f565b610237565b604051610107919061027c565b60405180910390f35b5f8054905090565b805f8190555050565b6002604051806040016040528083815260200184815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f0155602082015181600101908161017b9190610760565b5050505050565b60028181548110610191575f80fd5b905f5260205f2090600202015f91509050805f0154908060010180546101b690610593565b80601f01602080910402602001604051908101604052809291908181526020018280546101e290610593565b801561022d5780601f106102045761010080835404028352916020019161022d565b820191905f5260205f20905b81548152906001019060200180831161021057829003601f168201915b5050505050905082565b6001818051602081018201805184825260208301602085012081835280955050505050505f915090505481565b5f819050919050565b61027681610264565b82525050565b5f60208201905061028f5f83018461026d565b92915050565b5f604051905090565b5f80fd5b5f80fd5b6102af81610264565b81146102b9575f80fd5b50565b5f813590506102ca816102a6565b92915050565b5f602082840312156102e5576102e461029e565b5b5f6102f2848285016102bc565b91505092915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61034982610303565b810181811067ffffffffffffffff8211171561036857610367610313565b5b80604052505050565b5f61037a610295565b90506103868282610340565b919050565b5f67ffffffffffffffff8211156103a5576103a4610313565b5b6103ae82610303565b9050602081019050919050565b828183375f83830152505050565b5f6103db6103d68461038b565b610371565b9050828152602081018484840111156103f7576103f66102ff565b5b6104028482856103bb565b509392505050565b5f82601f83011261041e5761041d6102fb565b5b813561042e8482602086016103c9565b91505092915050565b5f806040838503121561044d5761044c61029e565b5b5f83013567ffffffffffffffff81111561046a576104696102a2565b5b6104768582860161040a565b9250506020610487858286016102bc565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6104c382610491565b6104cd818561049b565b93506104dd8185602086016104ab565b6104e681610303565b840191505092915050565b5f6040820190506105045f83018561026d565b818103602083015261051681846104b9565b90509392505050565b5f602082840312156105345761053361029e565b5b5f82013567ffffffffffffffff811115610551576105506102a2565b5b61055d8482850161040a565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806105aa57607f821691505b6020821081036105bd576105bc610566565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261061f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826105e4565b61062986836105e4565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61066461065f61065a84610264565b610641565b610264565b9050919050565b5f819050919050565b61067d8361064a565b6106916106898261066b565b8484546105f0565b825550505050565b5f90565b6106a5610699565b6106b0818484610674565b505050565b5b818110156106d3576106c85f8261069d565b6001810190506106b6565b5050565b601f821115610718576106e9816105c3565b6106f2846105d5565b81016020851015610701578190505b61071561070d856105d5565b8301826106b5565b50505b505050565b5f82821c905092915050565b5f6107385f198460080261071d565b1980831691505092915050565b5f6107508383610729565b9150826002028217905092915050565b61076982610491565b67ffffffffffffffff81111561078257610781610313565b5b61078c8254610593565b6107978282856106d7565b5f60209050601f8311600181146107c8575f84156107b6578287015190505b6107c08582610745565b865550610827565b601f1984166107d6866105c3565b5f5b828110156107fd578489015182556001820191506020850194506020810190506107d8565b8683101561081a5784890151610816601f891682610729565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220b86696d8bedddca98f7cc43ab7301475586cde323d028d136abd7ede8fd36d2464736f6c634300081a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// NameToFavouriteNumber is a free data retrieval call binding the contract method 0xb2ac62ef.
//
// Solidity: function nameToFavouriteNumber(string ) view returns(uint256)
func (_Api *ApiCaller) NameToFavouriteNumber(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "nameToFavouriteNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameToFavouriteNumber is a free data retrieval call binding the contract method 0xb2ac62ef.
//
// Solidity: function nameToFavouriteNumber(string ) view returns(uint256)
func (_Api *ApiSession) NameToFavouriteNumber(arg0 string) (*big.Int, error) {
	return _Api.Contract.NameToFavouriteNumber(&_Api.CallOpts, arg0)
}

// NameToFavouriteNumber is a free data retrieval call binding the contract method 0xb2ac62ef.
//
// Solidity: function nameToFavouriteNumber(string ) view returns(uint256)
func (_Api *ApiCallerSession) NameToFavouriteNumber(arg0 string) (*big.Int, error) {
	return _Api.Contract.NameToFavouriteNumber(&_Api.CallOpts, arg0)
}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favouriteNumber, string name)
func (_Api *ApiCaller) People(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FavouriteNumber *big.Int
	Name            string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "people", arg0)

	outstruct := new(struct {
		FavouriteNumber *big.Int
		Name            string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FavouriteNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favouriteNumber, string name)
func (_Api *ApiSession) People(arg0 *big.Int) (struct {
	FavouriteNumber *big.Int
	Name            string
}, error) {
	return _Api.Contract.People(&_Api.CallOpts, arg0)
}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favouriteNumber, string name)
func (_Api *ApiCallerSession) People(arg0 *big.Int) (struct {
	FavouriteNumber *big.Int
	Name            string
}, error) {
	return _Api.Contract.People(&_Api.CallOpts, arg0)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Api *ApiCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Api *ApiSession) Retrieve() (*big.Int, error) {
	return _Api.Contract.Retrieve(&_Api.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Api *ApiCallerSession) Retrieve() (*big.Int, error) {
	return _Api.Contract.Retrieve(&_Api.CallOpts)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favouriteNumber) returns()
func (_Api *ApiTransactor) AddPerson(opts *bind.TransactOpts, _name string, _favouriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addPerson", _name, _favouriteNumber)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favouriteNumber) returns()
func (_Api *ApiSession) AddPerson(_name string, _favouriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddPerson(&_Api.TransactOpts, _name, _favouriteNumber)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favouriteNumber) returns()
func (_Api *ApiTransactorSession) AddPerson(_name string, _favouriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddPerson(&_Api.TransactOpts, _name, _favouriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favouriteNumber) returns()
func (_Api *ApiTransactor) Store(opts *bind.TransactOpts, _favouriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "store", _favouriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favouriteNumber) returns()
func (_Api *ApiSession) Store(_favouriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Store(&_Api.TransactOpts, _favouriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favouriteNumber) returns()
func (_Api *ApiTransactorSession) Store(_favouriteNumber *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Store(&_Api.TransactOpts, _favouriteNumber)
}
