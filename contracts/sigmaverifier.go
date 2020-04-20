// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// SigmaverifierABI is the input ABI used to generate the binding from.
const SigmaverifierABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"params\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[16]\",\"name\":\"points\",\"type\":\"uint256[16]\"},{\"internalType\":\"uint256\",\"name\":\"z1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"z2\",\"type\":\"uint256\"}],\"name\":\"verifyPTEProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[10]\",\"name\":\"points\",\"type\":\"uint256[10]\"},{\"internalType\":\"uint256\",\"name\":\"z1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"z2\",\"type\":\"uint256\"}],\"name\":\"verifyCTValidProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[20]\",\"name\":\"points\",\"type\":\"uint256[20]\"},{\"internalType\":\"uint256\",\"name\":\"z1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"z2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"z3\",\"type\":\"uint256\"}],\"name\":\"verifySigmaProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SigmaverifierBin is the compiled bytecode used for deploying new contracts.
const SigmaverifierBin = `0x60806040523480156200001157600080fd5b506040516200132638038062001326833981810160405260208110156200003757600080fd5b8101908080519060200190929190505050600081905062000057620001f0565b8173ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff1660e01b8152600401604080518083038186803b1580156200009d57600080fd5b505afa158015620000b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620000d857600080fd5b81019080919050509050620000ec620001f0565b8273ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff1660e01b8152600401604080518083038186803b1580156200013257600080fd5b505afa15801562000147573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200016d57600080fd5b81019080919050509050816000600281106200018557fe5b6020020151600080018190555081600160028110620001a057fe5b602002015160006001018190555080600060028110620001bc57fe5b602002015160026000018190555080600160028110620001d857fe5b60200201516002600101819055505050505062000212565b6040518060400160405280600290602082028038833980820191505090505090565b61110480620002226000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80636e26cc311461005c578063991395ca146100f9578063b8c9d3651461018c578063e1efe766146101b1578063e2179b8e14610244575b600080fd5b6100df60048036036102e081101561007357600080fd5b810190808061028001906014806020026040519081016040528092919082601460200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291908035906020019092919080359060200190929190505050610269565b604051808215151515815260200191505060405180910390f35b610172600480360361018081101561011057600080fd5b81019080806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080359060200190929190505050610694565b604051808215151515815260200191505060405180910390f35b6101946108a8565b604051808381526020018281526020019250505060405180910390f35b61022a60048036036102408110156101c857600080fd5b810190808061020001906010806020026040519081016040528092919082601060200280828437600081840152601f19601f820116905080830192505050505050919291929080359060200190929190803590602001909291905050506108ba565b604051808215151515815260200191505060405180910390f35b61024c610c10565b604051808381526020018281526020019250505060405180910390f35b6000806060600860405190808252806020026020018201604052801561029e5781602001602082028038833980820191505090505b50905060008090505b60088110156102e95787600c8201601481106102bf57fe5b60200201518282815181106102d057fe5b60200260200101818152505080806001019150506102a7565b506102f381610c22565b91506102fd610ffd565b60405180604001604052808960006014811061031557fe5b602002015181526020018960016014811061032c57fe5b60200201518152508160006003811061034157fe5b6020020181905250604051806040016040528089600c6014811061036157fe5b6020020151815260200189600d6014811061037857fe5b60200201518152508160016003811061038d57fe5b60200201819052506040518060400160405280896002601481106103ad57fe5b60200201518152602001896003601481106103c457fe5b6020020151815250816002600381106103d957fe5b60200201819052506103ec818885610c8a565b6103fc576000935050505061068c565b610404610ffd565b60405180604001604052808a60066014811061041c57fe5b602002015181526020018a60076014811061043357fe5b60200201518152508160006003811061044857fe5b602002018190525060405180604001604052808a600e6014811061046857fe5b602002015181526020018a600f6014811061047f57fe5b60200201518152508160016003811061049457fe5b602002018190525060405180604001604052808a6008601481106104b457fe5b602002015181526020018a6009601481106104cb57fe5b6020020151815250816002600381106104e057fe5b60200201819052506104f3818886610c8a565b61050457600094505050505061068c565b61050c61102a565b60405180604001604052808b60106014811061052457fe5b602002015181526020018b60116014811061053b57fe5b60200201518152508160006002811061055057fe5b602002018190525060405180604001604052808b60046014811061057057fe5b602002015181526020018b60056014811061058757fe5b60200201518152508160016002811061059c57fe5b60200201819052506105b0818a8988610d34565b6105c25760009550505050505061068c565b6105ca61102a565b60405180604001604052808c6012601481106105e257fe5b602002015181526020018c6013601481106105f957fe5b60200201518152508160006002811061060e57fe5b602002018190525060405180604001604052808c600a6014811061062e57fe5b602002015181526020018c600b6014811061064557fe5b60200201518152508160016002811061065a57fe5b602002018190525061066e818a8a89610d34565b610681576000965050505050505061068c565b600196505050505050505b949350505050565b6000806106e7856006600a81106106a757fe5b6020020151866007600a81106106b957fe5b6020020151876008600a81106106cb57fe5b6020020151886009600a81106106dd57fe5b6020020151610e32565b90506106f1610ffd565b6040518060400160405280876000600a811061070957fe5b60200201518152602001876001600a811061072057fe5b60200201518152508160006003811061073557fe5b60200201819052506040518060400160405280876006600a811061075557fe5b60200201518152602001876007600a811061076c57fe5b60200201518152508160016003811061078157fe5b60200201819052506040518060400160405280876002600a81106107a157fe5b60200201518152602001876003600a81106107b857fe5b6020020151815250816002600381106107cd57fe5b60200201819052506107e0818684610c8a565b6107ef576000925050506108a1565b6107f761102a565b6040518060400160405280886008600a811061080f57fe5b60200201518152602001886009600a811061082657fe5b60200201518152508160006002811061083b57fe5b60200201819052506040518060400160405280886004600a811061085b57fe5b60200201518152602001886005600a811061087257fe5b60200201518152508160016002811061088757fe5b602002018190525061089b81878786610d34565b93505050505b9392505050565b60028060000154908060010154905082565b600080606060066040519080825280602002602001820160405280156108ef5781602001602082028038833980820191505090505b50905060008090505b600681101561093a578681600a016010811061091057fe5b602002015182828151811061092157fe5b60200260200101818152505080806001019150506108f8565b5061094481610c22565b915061094e610ffd565b60405180604001604052808860006010811061096657fe5b602002015181526020018860016010811061097d57fe5b60200201518152508160006003811061099257fe5b6020020181905250604051806040016040528088600a601081106109b257fe5b6020020151815260200188600b601081106109c957fe5b6020020151815250816001600381106109de57fe5b60200201819052506040518060400160405280886004601081106109fe57fe5b6020020151815260200188600560108110610a1557fe5b602002015181525081600260038110610a2a57fe5b6020020181905250610a3d818785610c8a565b610a4d5760009350505050610c09565b610a55610ffd565b604051806040016040528089600260108110610a6d57fe5b6020020151815260200189600360108110610a8457fe5b602002015181525081600060038110610a9957fe5b6020020181905250604051806040016040528089600c60108110610ab957fe5b6020020151815260200189600d60108110610ad057fe5b602002015181525081600160038110610ae557fe5b6020020181905250604051806040016040528089600660108110610b0557fe5b6020020151815260200189600760108110610b1c57fe5b602002015181525081600260038110610b3157fe5b6020020181905250610b44818886610c8a565b610b55576000945050505050610c09565b610b5d61102a565b60405180604001604052808a600e60108110610b7557fe5b602002015181526020018a600f60108110610b8c57fe5b602002015181525081600060028110610ba157fe5b602002018190525060405180604001604052808a600860108110610bc157fe5b602002015181526020018a600960108110610bd857fe5b602002015181525081600160028110610bed57fe5b6020020181905250610c0181898988610d34565b955050505050505b9392505050565b60008060000154908060010154905082565b6000610c838260405160200180828051906020019060200280838360005b83811015610c5b578082015181840152602081019050610c40565b505050509050019150506040516020818303038152906040528051906020012060001c610e88565b9050919050565b6000610c94611057565b610cdb85600160038110610ca457fe5b6020020151610ccd8588600260038110610cba57fe5b6020020151610ec090919063ffffffff16565b610f6590919063ffffffff16565b9050610ce5611057565b610d098587600060038110610cf657fe5b6020020151610ec090919063ffffffff16565b905080600001518260000151148015610d29575080602001518260200151145b925050509392505050565b6000610d3e611057565b610db5610d7485600060405180604001604052908160008201548152602001600182015481525050610ec090919063ffffffff16565b610da787600260405180604001604052908160008201548152602001600182015481525050610ec090919063ffffffff16565b610f6590919063ffffffff16565b9050610dbf611057565b610e06610de68589600160028110610dd357fe5b6020020151610ec090919063ffffffff16565b88600060028110610df357fe5b6020020151610f6590919063ffffffff16565b905080600001518260000151148015610e26575080602001518260200151145b92505050949350505050565b6000610e7e85858585604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c610e88565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381610eb757fe5b06915050919050565b610ec8611057565b6001821415610ed957829050610f5f565b6002821415610ef357610eec8384610f65565b9050610f5f565b610efb611071565b836000015181600060038110610f0d57fe5b602002018181525050836020015181600160038110610f2857fe5b6020020181815250508281600260038110610f3f57fe5b6020020181815250506040826060836007600019fa610f5d57600080fd5b505b92915050565b610f6d611057565b610f75611093565b836000015181600060048110610f8757fe5b602002018181525050836020015181600160048110610fa257fe5b602002018181525050826000015181600260048110610fbd57fe5b602002018181525050826020015181600360048110610fd857fe5b6020020181815250506040826080836006600019fa610ff657600080fd5b5092915050565b60405180606001604052806003905b6110146110b5565b81526020019060019003908161100c5790505090565b60405180604001604052806002905b6110416110b5565b8152602001906001900390816110395790505090565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b60405180604001604052806000815260200160008152509056fea265627a7a7231582085d27d73bfb7fda515a790da7485030b36b636b13866a6fd6daa731b73e5f2c164736f6c63430005100032`

// DeploySigmaverifier deploys a new Ethereum contract, binding an instance of Sigmaverifier to it.
func DeploySigmaverifier(auth *bind.TransactOpts, backend bind.ContractBackend, params common.Address) (common.Address, *types.Transaction, *Sigmaverifier, error) {
	parsed, err := abi.JSON(strings.NewReader(SigmaverifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigmaverifierBin), backend, params)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sigmaverifier{SigmaverifierCaller: SigmaverifierCaller{contract: contract}, SigmaverifierTransactor: SigmaverifierTransactor{contract: contract}, SigmaverifierFilterer: SigmaverifierFilterer{contract: contract}}, nil
}

// Sigmaverifier is an auto generated Go binding around an Ethereum contract.
type Sigmaverifier struct {
	SigmaverifierCaller     // Read-only binding to the contract
	SigmaverifierTransactor // Write-only binding to the contract
	SigmaverifierFilterer   // Log filterer for contract events
}

// SigmaverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigmaverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigmaverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigmaverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigmaverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigmaverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigmaverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigmaverifierSession struct {
	Contract     *Sigmaverifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigmaverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigmaverifierCallerSession struct {
	Contract *SigmaverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SigmaverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigmaverifierTransactorSession struct {
	Contract     *SigmaverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SigmaverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigmaverifierRaw struct {
	Contract *Sigmaverifier // Generic contract binding to access the raw methods on
}

// SigmaverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigmaverifierCallerRaw struct {
	Contract *SigmaverifierCaller // Generic read-only contract binding to access the raw methods on
}

// SigmaverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigmaverifierTransactorRaw struct {
	Contract *SigmaverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigmaverifier creates a new instance of Sigmaverifier, bound to a specific deployed contract.
func NewSigmaverifier(address common.Address, backend bind.ContractBackend) (*Sigmaverifier, error) {
	contract, err := bindSigmaverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sigmaverifier{SigmaverifierCaller: SigmaverifierCaller{contract: contract}, SigmaverifierTransactor: SigmaverifierTransactor{contract: contract}, SigmaverifierFilterer: SigmaverifierFilterer{contract: contract}}, nil
}

// NewSigmaverifierCaller creates a new read-only instance of Sigmaverifier, bound to a specific deployed contract.
func NewSigmaverifierCaller(address common.Address, caller bind.ContractCaller) (*SigmaverifierCaller, error) {
	contract, err := bindSigmaverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigmaverifierCaller{contract: contract}, nil
}

// NewSigmaverifierTransactor creates a new write-only instance of Sigmaverifier, bound to a specific deployed contract.
func NewSigmaverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*SigmaverifierTransactor, error) {
	contract, err := bindSigmaverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigmaverifierTransactor{contract: contract}, nil
}

// NewSigmaverifierFilterer creates a new log filterer instance of Sigmaverifier, bound to a specific deployed contract.
func NewSigmaverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*SigmaverifierFilterer, error) {
	contract, err := bindSigmaverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigmaverifierFilterer{contract: contract}, nil
}

// bindSigmaverifier binds a generic wrapper to an already deployed contract.
func bindSigmaverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigmaverifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sigmaverifier *SigmaverifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sigmaverifier.Contract.SigmaverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sigmaverifier *SigmaverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.SigmaverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sigmaverifier *SigmaverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.SigmaverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sigmaverifier *SigmaverifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sigmaverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sigmaverifier *SigmaverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sigmaverifier *SigmaverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.contract.Transact(opts, method, params...)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierCaller) G(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Sigmaverifier.contract.Call(opts, out, "g")
	return *ret, err
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sigmaverifier.Contract.G(&_Sigmaverifier.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierCallerSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sigmaverifier.Contract.G(&_Sigmaverifier.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierCaller) H(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Sigmaverifier.contract.Call(opts, out, "h")
	return *ret, err
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sigmaverifier.Contract.H(&_Sigmaverifier.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierCallerSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sigmaverifier.Contract.H(&_Sigmaverifier.CallOpts)
}

// VerifySigmaProof is a free data retrieval call binding the contract method 0x6e26cc31.
//
// Solidity: function verifySigmaProof(uint256[20] points, uint256 z1, uint256 z2, uint256 z3) constant returns(bool)
func (_Sigmaverifier *SigmaverifierCaller) VerifySigmaProof(opts *bind.CallOpts, points [20]*big.Int, z1 *big.Int, z2 *big.Int, z3 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Sigmaverifier.contract.Call(opts, out, "verifySigmaProof", points, z1, z2, z3)
	return *ret0, err
}

// VerifySigmaProof is a free data retrieval call binding the contract method 0x6e26cc31.
//
// Solidity: function verifySigmaProof(uint256[20] points, uint256 z1, uint256 z2, uint256 z3) constant returns(bool)
func (_Sigmaverifier *SigmaverifierSession) VerifySigmaProof(points [20]*big.Int, z1 *big.Int, z2 *big.Int, z3 *big.Int) (bool, error) {
	return _Sigmaverifier.Contract.VerifySigmaProof(&_Sigmaverifier.CallOpts, points, z1, z2, z3)
}

// VerifySigmaProof is a free data retrieval call binding the contract method 0x6e26cc31.
//
// Solidity: function verifySigmaProof(uint256[20] points, uint256 z1, uint256 z2, uint256 z3) constant returns(bool)
func (_Sigmaverifier *SigmaverifierCallerSession) VerifySigmaProof(points [20]*big.Int, z1 *big.Int, z2 *big.Int, z3 *big.Int) (bool, error) {
	return _Sigmaverifier.Contract.VerifySigmaProof(&_Sigmaverifier.CallOpts, points, z1, z2, z3)
}

// VerifyCTValidProof is a paid mutator transaction binding the contract method 0x991395ca.
//
// Solidity: function verifyCTValidProof(uint256[10] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierTransactor) VerifyCTValidProof(opts *bind.TransactOpts, points [10]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.contract.Transact(opts, "verifyCTValidProof", points, z1, z2)
}

// VerifyCTValidProof is a paid mutator transaction binding the contract method 0x991395ca.
//
// Solidity: function verifyCTValidProof(uint256[10] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierSession) VerifyCTValidProof(points [10]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.VerifyCTValidProof(&_Sigmaverifier.TransactOpts, points, z1, z2)
}

// VerifyCTValidProof is a paid mutator transaction binding the contract method 0x991395ca.
//
// Solidity: function verifyCTValidProof(uint256[10] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierTransactorSession) VerifyCTValidProof(points [10]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.VerifyCTValidProof(&_Sigmaverifier.TransactOpts, points, z1, z2)
}

// VerifyPTEProof is a paid mutator transaction binding the contract method 0xe1efe766.
//
// Solidity: function verifyPTEProof(uint256[16] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierTransactor) VerifyPTEProof(opts *bind.TransactOpts, points [16]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.contract.Transact(opts, "verifyPTEProof", points, z1, z2)
}

// VerifyPTEProof is a paid mutator transaction binding the contract method 0xe1efe766.
//
// Solidity: function verifyPTEProof(uint256[16] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierSession) VerifyPTEProof(points [16]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.VerifyPTEProof(&_Sigmaverifier.TransactOpts, points, z1, z2)
}

// VerifyPTEProof is a paid mutator transaction binding the contract method 0xe1efe766.
//
// Solidity: function verifyPTEProof(uint256[16] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierTransactorSession) VerifyPTEProof(points [16]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.VerifyPTEProof(&_Sigmaverifier.TransactOpts, points, z1, z2)
}
