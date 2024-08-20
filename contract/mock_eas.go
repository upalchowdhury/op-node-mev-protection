// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttestationCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"encryptedData\",\"type\":\"string\"}],\"name\":\"attestTransactionWithPrivateData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"attestations\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"encryptedData\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"encryptedData\",\"type\":\"string\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610f0d8061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806306dda2d31461004e578063249d026b1461007e5780638d5b8361146100b0578063d607daf6146100cc575b5f80fd5b61006860048036038101906100639190610843565b6100fe565b60405161007591906108d3565b60405180910390f35b610098600480360381019061009391906108ec565b6101e2565b6040516100a7939291906109ec565b60405180910390f35b6100ca60048036038101906100c59190610843565b61034a565b005b6100e660048036038101906100e191906108ec565b6104ca565b6040516100f5939291906109ec565b60405180910390f35b5f805f8460405161010f9190610a69565b90815260200160405180910390205f01805461012a90610aac565b90500361016c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016390610b26565b60405180910390fd5b8160405160200161017d9190610a69565b604051602081830303815290604052805190602001205f846040516101a29190610a69565b90815260200160405180910390206001016040516020016101c39190610bd6565b6040516020818303038152906040528051906020012014905092915050565b5f818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461021890610aac565b80601f016020809104026020016040519081016040528092919081815260200182805461024490610aac565b801561028f5780601f106102665761010080835404028352916020019161028f565b820191905f5260205f20905b81548152906001019060200180831161027257829003601f168201915b5050505050908060010180546102a490610aac565b80601f01602080910402602001604051908101604052809291908181526020018280546102d090610aac565b801561031b5780601f106102f25761010080835404028352916020019161031b565b820191905f5260205f20905b8154815290600101906020018083116102fe57829003601f168201915b505050505090806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b5f808360405161035a9190610a69565b90815260200160405180910390205f01805461037590610aac565b9050146103b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ae90610c36565b60405180910390fd5b60405180606001604052808381526020018281526020013373ffffffffffffffffffffffffffffffffffffffff168152505f836040516103f79190610a69565b90815260200160405180910390205f820151815f0190816104189190610de8565b50602082015181600101908161042e9190610de8565b506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff167f121ed3452c21306e8ea9b632f37c9aaf822b11f2a6aa980480124ba8b6eb0ac8836040516104be9190610eb7565b60405180910390a25050565b6060805f805f856040516104de9190610a69565b90815260200160405180910390205f0180546104f990610aac565b90500361053b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053290610b26565b60405180910390fd5b5f808560405161054b9190610a69565b90815260200160405180910390206040518060600160405290815f8201805461057390610aac565b80601f016020809104026020016040519081016040528092919081815260200182805461059f90610aac565b80156105ea5780601f106105c1576101008083540402835291602001916105ea565b820191905f5260205f20905b8154815290600101906020018083116105cd57829003601f168201915b5050505050815260200160018201805461060390610aac565b80601f016020809104026020016040519081016040528092919081815260200182805461062f90610aac565b801561067a5780601f106106515761010080835404028352916020019161067a565b820191905f5260205f20905b81548152906001019060200180831161065d57829003601f168201915b50505050508152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050805f015181602001518260400151935093509350509193909250565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6107558261070f565b810181811067ffffffffffffffff821117156107745761077361071f565b5b80604052505050565b5f6107866106f6565b9050610792828261074c565b919050565b5f67ffffffffffffffff8211156107b1576107b061071f565b5b6107ba8261070f565b9050602081019050919050565b828183375f83830152505050565b5f6107e76107e284610797565b61077d565b9050828152602081018484840111156108035761080261070b565b5b61080e8482856107c7565b509392505050565b5f82601f83011261082a57610829610707565b5b813561083a8482602086016107d5565b91505092915050565b5f8060408385031215610859576108586106ff565b5b5f83013567ffffffffffffffff81111561087657610875610703565b5b61088285828601610816565b925050602083013567ffffffffffffffff8111156108a3576108a2610703565b5b6108af85828601610816565b9150509250929050565b5f8115159050919050565b6108cd816108b9565b82525050565b5f6020820190506108e65f8301846108c4565b92915050565b5f60208284031215610901576109006106ff565b5b5f82013567ffffffffffffffff81111561091e5761091d610703565b5b61092a84828501610816565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561096a57808201518184015260208101905061094f565b5f8484015250505050565b5f61097f82610933565b610989818561093d565b935061099981856020860161094d565b6109a28161070f565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6109d6826109ad565b9050919050565b6109e6816109cc565b82525050565b5f6060820190508181035f830152610a048186610975565b90508181036020830152610a188185610975565b9050610a2760408301846109dd565b949350505050565b5f81905092915050565b5f610a4382610933565b610a4d8185610a2f565b9350610a5d81856020860161094d565b80840191505092915050565b5f610a748284610a39565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610ac357607f821691505b602082108103610ad657610ad5610a7f565b5b50919050565b7f4174746573746174696f6e20646f6573206e6f742065786973740000000000005f82015250565b5f610b10601a8361093d565b9150610b1b82610adc565b602082019050919050565b5f6020820190508181035f830152610b3d81610b04565b9050919050565b5f819050815f5260205f209050919050565b5f8154610b6281610aac565b610b6c8186610a2f565b9450600182165f8114610b865760018114610b9b57610bcd565b60ff1983168652811515820286019350610bcd565b610ba485610b44565b5f5b83811015610bc557815481890152600182019150602081019050610ba6565b838801955050505b50505092915050565b5f610be18284610b56565b915081905092915050565b7f4174746573746174696f6e20616c7265616479206578697374730000000000005f82015250565b5f610c20601a8361093d565b9150610c2b82610bec565b602082019050919050565b5f6020820190508181035f830152610c4d81610c14565b9050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610c9e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610c63565b610ca88683610c63565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f610cec610ce7610ce284610cc0565b610cc9565b610cc0565b9050919050565b5f819050919050565b610d0583610cd2565b610d19610d1182610cf3565b848454610c6f565b825550505050565b5f90565b610d2d610d21565b610d38818484610cfc565b505050565b5b81811015610d5b57610d505f82610d25565b600181019050610d3e565b5050565b601f821115610da057610d7181610b44565b610d7a84610c54565b81016020851015610d89578190505b610d9d610d9585610c54565b830182610d3d565b50505b505050565b5f82821c905092915050565b5f610dc05f1984600802610da5565b1980831691505092915050565b5f610dd88383610db1565b9150826002028217905092915050565b610df182610933565b67ffffffffffffffff811115610e0a57610e0961071f565b5b610e148254610aac565b610e1f828285610d5f565b5f60209050601f831160018114610e50575f8415610e3e578287015190505b610e488582610dcd565b865550610eaf565b601f198416610e5e86610b44565b5f5b82811015610e8557848901518255600182019150602085019450602081019050610e60565b86831015610ea25784890151610e9e601f891682610db1565b8355505b6001600288020188555050505b505050505050565b5f6020820190508181035f830152610ecf8184610975565b90509291505056fea2646970667358221220f8f379f825fabf12affcbcd61b74198f1b5250e221149d31cbcff4d66940d05364736f6c63430008180033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Attestations is a free data retrieval call binding the contract method 0x249d026b.
//
// Solidity: function attestations(string ) view returns(string txHash, string encryptedData, address attester)
func (_Contract *ContractCaller) Attestations(opts *bind.CallOpts, arg0 string) (struct {
	TxHash        string
	EncryptedData string
	Attester      common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestations", arg0)

	outstruct := new(struct {
		TxHash        string
		EncryptedData string
		Attester      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TxHash = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.EncryptedData = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Attester = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Attestations is a free data retrieval call binding the contract method 0x249d026b.
//
// Solidity: function attestations(string ) view returns(string txHash, string encryptedData, address attester)
func (_Contract *ContractSession) Attestations(arg0 string) (struct {
	TxHash        string
	EncryptedData string
	Attester      common.Address
}, error) {
	return _Contract.Contract.Attestations(&_Contract.CallOpts, arg0)
}

// Attestations is a free data retrieval call binding the contract method 0x249d026b.
//
// Solidity: function attestations(string ) view returns(string txHash, string encryptedData, address attester)
func (_Contract *ContractCallerSession) Attestations(arg0 string) (struct {
	TxHash        string
	EncryptedData string
	Attester      common.Address
}, error) {
	return _Contract.Contract.Attestations(&_Contract.CallOpts, arg0)
}

// GetAttestation is a free data retrieval call binding the contract method 0xd607daf6.
//
// Solidity: function getAttestation(string txHash) view returns(string, string, address)
func (_Contract *ContractCaller) GetAttestation(opts *bind.CallOpts, txHash string) (string, string, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAttestation", txHash)

	if err != nil {
		return *new(string), *new(string), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// GetAttestation is a free data retrieval call binding the contract method 0xd607daf6.
//
// Solidity: function getAttestation(string txHash) view returns(string, string, address)
func (_Contract *ContractSession) GetAttestation(txHash string) (string, string, common.Address, error) {
	return _Contract.Contract.GetAttestation(&_Contract.CallOpts, txHash)
}

// GetAttestation is a free data retrieval call binding the contract method 0xd607daf6.
//
// Solidity: function getAttestation(string txHash) view returns(string, string, address)
func (_Contract *ContractCallerSession) GetAttestation(txHash string) (string, string, common.Address, error) {
	return _Contract.Contract.GetAttestation(&_Contract.CallOpts, txHash)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x06dda2d3.
//
// Solidity: function verifyAttestation(string txHash, string encryptedData) view returns(bool)
func (_Contract *ContractCaller) VerifyAttestation(opts *bind.CallOpts, txHash string, encryptedData string) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifyAttestation", txHash, encryptedData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x06dda2d3.
//
// Solidity: function verifyAttestation(string txHash, string encryptedData) view returns(bool)
func (_Contract *ContractSession) VerifyAttestation(txHash string, encryptedData string) (bool, error) {
	return _Contract.Contract.VerifyAttestation(&_Contract.CallOpts, txHash, encryptedData)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x06dda2d3.
//
// Solidity: function verifyAttestation(string txHash, string encryptedData) view returns(bool)
func (_Contract *ContractCallerSession) VerifyAttestation(txHash string, encryptedData string) (bool, error) {
	return _Contract.Contract.VerifyAttestation(&_Contract.CallOpts, txHash, encryptedData)
}

// AttestTransactionWithPrivateData is a paid mutator transaction binding the contract method 0x8d5b8361.
//
// Solidity: function attestTransactionWithPrivateData(string txHash, string encryptedData) returns()
func (_Contract *ContractTransactor) AttestTransactionWithPrivateData(opts *bind.TransactOpts, txHash string, encryptedData string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "attestTransactionWithPrivateData", txHash, encryptedData)
}

// AttestTransactionWithPrivateData is a paid mutator transaction binding the contract method 0x8d5b8361.
//
// Solidity: function attestTransactionWithPrivateData(string txHash, string encryptedData) returns()
func (_Contract *ContractSession) AttestTransactionWithPrivateData(txHash string, encryptedData string) (*types.Transaction, error) {
	return _Contract.Contract.AttestTransactionWithPrivateData(&_Contract.TransactOpts, txHash, encryptedData)
}

// AttestTransactionWithPrivateData is a paid mutator transaction binding the contract method 0x8d5b8361.
//
// Solidity: function attestTransactionWithPrivateData(string txHash, string encryptedData) returns()
func (_Contract *ContractTransactorSession) AttestTransactionWithPrivateData(txHash string, encryptedData string) (*types.Transaction, error) {
	return _Contract.Contract.AttestTransactionWithPrivateData(&_Contract.TransactOpts, txHash, encryptedData)
}

// ContractAttestationCreatedIterator is returned from FilterAttestationCreated and is used to iterate over the raw logs and unpacked data for AttestationCreated events raised by the Contract contract.
type ContractAttestationCreatedIterator struct {
	Event *ContractAttestationCreated // Event containing the contract specifics and raw log

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
func (it *ContractAttestationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationCreated)
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
		it.Event = new(ContractAttestationCreated)
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
func (it *ContractAttestationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationCreated represents a AttestationCreated event raised by the Contract contract.
type ContractAttestationCreated struct {
	TxHash   string
	Attester common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAttestationCreated is a free log retrieval operation binding the contract event 0x121ed3452c21306e8ea9b632f37c9aaf822b11f2a6aa980480124ba8b6eb0ac8.
//
// Solidity: event AttestationCreated(string txHash, address indexed attester)
func (_Contract *ContractFilterer) FilterAttestationCreated(opts *bind.FilterOpts, attester []common.Address) (*ContractAttestationCreatedIterator, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AttestationCreated", attesterRule)
	if err != nil {
		return nil, err
	}
	return &ContractAttestationCreatedIterator{contract: _Contract.contract, event: "AttestationCreated", logs: logs, sub: sub}, nil
}

// WatchAttestationCreated is a free log subscription operation binding the contract event 0x121ed3452c21306e8ea9b632f37c9aaf822b11f2a6aa980480124ba8b6eb0ac8.
//
// Solidity: event AttestationCreated(string txHash, address indexed attester)
func (_Contract *ContractFilterer) WatchAttestationCreated(opts *bind.WatchOpts, sink chan<- *ContractAttestationCreated, attester []common.Address) (event.Subscription, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AttestationCreated", attesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationCreated)
				if err := _Contract.contract.UnpackLog(event, "AttestationCreated", log); err != nil {
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

// ParseAttestationCreated is a log parse operation binding the contract event 0x121ed3452c21306e8ea9b632f37c9aaf822b11f2a6aa980480124ba8b6eb0ac8.
//
// Solidity: event AttestationCreated(string txHash, address indexed attester)
func (_Contract *ContractFilterer) ParseAttestationCreated(log types.Log) (*ContractAttestationCreated, error) {
	event := new(ContractAttestationCreated)
	if err := _Contract.contract.UnpackLog(event, "AttestationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
