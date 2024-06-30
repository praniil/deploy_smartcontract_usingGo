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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_favouriteNumber\",\"type\":\"uint256\"}],\"name\":\"addPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nameToFavouriteNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"people\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fovouriteNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_favouriteNumber\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506104fa8061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80632e64cec1146100595780636057361d1461006f5780636f760f41146100835780639e7a13ad14610096578063b2ac62ef146100b7575b5f80fd5b5f545b6040519081526020015b60405180910390f35b61008161007d36600461021a565b5f55565b005b6100816100913660046102d0565b6100e2565b6100a96100a436600461021a565b610166565b604051610066929190610312565b61005c6100c536600461034e565b805160208183018101805160018252928201919093012091525481565b6040805180820190915281815260208101838152600280546001810182555f829052835191027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace8101918255915190917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf019061015f9082610409565b5050505050565b60028181548110610175575f80fd5b5f918252602090912060029091020180546001820180549193509061019990610388565b80601f01602080910402602001604051908101604052809291908181526020018280546101c590610388565b80156102105780601f106101e757610100808354040283529160200191610210565b820191905f5260205f20905b8154815290600101906020018083116101f357829003601f168201915b5050505050905082565b5f6020828403121561022a575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610254575f80fd5b813567ffffffffffffffff81111561026e5761026e610231565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561029d5761029d610231565b6040528181528382016020018510156102b4575f80fd5b816020850160208301375f918101602001919091529392505050565b5f80604083850312156102e1575f80fd5b823567ffffffffffffffff8111156102f7575f80fd5b61030385828601610245565b95602094909401359450505050565b828152604060208201525f82518060408401528060208501606085015e5f606082850101526060601f19601f8301168401019150509392505050565b5f6020828403121561035e575f80fd5b813567ffffffffffffffff811115610374575f80fd5b61038084828501610245565b949350505050565b600181811c9082168061039c57607f821691505b6020821081036103ba57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561040457805f5260205f20601f840160051c810160208510156103e55750805b601f840160051c820191505b8181101561015f575f81556001016103f1565b505050565b815167ffffffffffffffff81111561042357610423610231565b610437816104318454610388565b846103c0565b6020601f821160018114610469575f83156104525750848201515b5f19600385901b1c1916600184901b17845561015f565b5f84815260208120601f198516915b828110156104985787850151825560209485019460019092019101610478565b50848210156104b557868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fea264697066735822122063ae212f24d356352303a5b4064fd97a2a773a1ef0b502cefea03cbc8b6a9a5164736f6c634300081a0033",
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
// Solidity: function people(uint256 ) view returns(uint256 fovouriteNumber, string name)
func (_Api *ApiCaller) People(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FovouriteNumber *big.Int
	Name            string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "people", arg0)

	outstruct := new(struct {
		FovouriteNumber *big.Int
		Name            string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FovouriteNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 fovouriteNumber, string name)
func (_Api *ApiSession) People(arg0 *big.Int) (struct {
	FovouriteNumber *big.Int
	Name            string
}, error) {
	return _Api.Contract.People(&_Api.CallOpts, arg0)
}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 fovouriteNumber, string name)
func (_Api *ApiCallerSession) People(arg0 *big.Int) (struct {
	FovouriteNumber *big.Int
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
