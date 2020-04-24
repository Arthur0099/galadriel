// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipverifier

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

// IpverifierABI is the input ABI used to generate the binding from.
const IpverifierABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"gv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"hv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"u\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"l\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"r\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"optimizedVerifyIPProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"gv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"hv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"u\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"l\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"r\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"verifyIPProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IpverifierBin is the compiled bytecode used for deploying new contracts.
var IpverifierBin = "0x608060405234801561001057600080fd5b5061189d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637b7fc9ad1461003b578063976168631461006b575b600080fd5b610055600480360361005091908101906115da565b61009b565b6040516100629190611791565b60405180910390f35b610085600480360361008091908101906115da565b6100fb565b6040516100929190611791565b60405180910390f35b60006100a561138b565b60006100b88c8c8c8c8c8c8c8c8c61015b565b8093508192505050806100d0576000925050506100ee565b6100e98c8c8460000151856020015186604001516103b7565b925050505b9998505050505050505050565b600061010561138b565b60006101188c8c8c8c8c8c8c8c8c61015b565b8093508192505050806101305760009250505061014e565b6101498c8c8460000151856020015186604001516109bb565b925050505b9998505050505050505050565b600061016561138b565b61016d61138b565b8a518c5114158061018057508551875114155b156101925760008192509250506103a9565b60028c518161019d57fe5b0460028851816101a957fe5b0460020a146101bf5760008192509250506103a9565b86516040519080825280602002602001820160405280156101ef5781602001602082028038833980820191505090505b50816040015160000181905250865160405190808252806020026020018201604052801561022c5781602001602082028038833980820191505090505b5081604001516020018190525060008090505b87518110156102bd5787818151811061025457fe5b6020026020010151826040015160000151828151811061027057fe5b60200260200101818152505086818151811061028857fe5b602002602001015182604001516020015182815181106102a457fe5b602002602001018181525050808060010191505061023f565b508481604001516040018181525050838160400151606001818152505060006102e589610e87565b90506103308160405180604001604052808d60006002811061030357fe5b602002015181526020018d60016002811061031a57fe5b6020020151815250610ec290919063ffffffff16565b826000018190525061039760405180604001604052808d60006002811061035357fe5b602002015181526020018d60016002811061036a57fe5b60200201518152506103898b8560000151610ec290919063ffffffff16565b610f6790919063ffffffff16565b82602001819052506001829350935050505b995099975050505050505050565b60006103c16113be565b60006002905060008090505b6002856000015151816103dc57fe5b048110156108675760405180604001604052808660000151836002028151811061040257fe5b602002602001015181526020018660000151600184600202018151811061042557fe5b6020026020010151815250836000018190525060405180604001604052808660200151836002028151811061045657fe5b602002602001015181526020018660200151600184600202018151811061047957fe5b602002602001015181525083602001819052506000610508866000015183600202815181106104a457fe5b6020026020010151876000015160018560020201815181106104c257fe5b6020026020010151886020015185600202815181106104dd57fe5b6020026020010151896020015160018760020201815181106104fb57fe5b6020026020010151610fff565b9050600061051582611043565b905060008090505b8460028d518161052957fe5b048161053157fe5b048110156107ca57808560028e518161054657fe5b048161054e57fe5b04018660c001818152505060405180604001604052808d836002028151811061057357fe5b602002602001015181526020018d600184600202018151811061059257fe5b6020026020010151815250866040018190525060405180604001604052808d60028960c0015102815181106105c357fe5b602002602001015181526020018d600160028a60c001510201815181106105e657fe5b60200260200101518152508660600181905250610638610613848860600151610ec290919063ffffffff16565b61062a848960400151610ec290919063ffffffff16565b610f6790919063ffffffff16565b86608001819052508560800151600001518c600283028151811061065857fe5b6020026020010181815250508560800151602001518c600160028402018151811061067f57fe5b60200260200101818152505060405180604001604052808c83600202815181106106a557fe5b602002602001015181526020018c60018460020201815181106106c457fe5b6020026020010151815250866040018190525060405180604001604052808c60028960c0015102815181106106f557fe5b602002602001015181526020018c600160028a60c0015102018151811061071857fe5b6020026020010151815250866060018190525061076a610745838860600151610ec290919063ffffffff16565b61075c858960400151610ec290919063ffffffff16565b610f6790919063ffffffff16565b8660a001819052508560a00151600001518b600283028151811061078a57fe5b6020026020010181815250508560a00151602001518b60016002840201815181106107b157fe5b602002602001018181525050808060010191505061051d565b50610850886108426108036107f06107eb868761110d90919063ffffffff16565b611147565b8960200151610ec290919063ffffffff16565b61083461082161081c888961110d90919063ffffffff16565b611147565b8a60000151610ec290919063ffffffff16565b610f6790919063ffffffff16565b610f6790919063ffffffff16565b9750600284029350505080806001019150506103cd565b50600061088d6108888660600151876040015161110d90919063ffffffff16565b611147565b905060405180604001604052808a6000815181106108a757fe5b602002602001015181526020018a6001815181106108c157fe5b602002602001015181525083604001819052506040518060400160405280896000815181106108ec57fe5b602002602001015181526020018960018151811061090657fe5b6020026020010151815250836060018190525061092161142e565b61098c61093f87606001518660600151610ec290919063ffffffff16565b61097e61095d89604001518860400151610ec290919063ffffffff16565b610970868d610ec290919063ffffffff16565b610f6790919063ffffffff16565b610f6790919063ffffffff16565b9050806000015187600001511480156109ac575080602001518760200151145b94505050505095945050505050565b60006109c56113be565b6002836000015151816109d457fe5b04604051908082528060200260200182016040528015610a035781602001602082028038833980820191505090505b508160e00181905250600283600001515181610a1b57fe5b04604051908082528060200260200182016040528015610a4a5781602001602082028038833980820191505090505b5081610100018190525060008090505b600284600001515181610a6957fe5b04811015610c6557604051806040016040528085600001518360020281518110610a8f57fe5b6020026020010151815260200185600001516001846002020181518110610ab257fe5b60200260200101518152508260000181905250604051806040016040528085602001518360020281518110610ae357fe5b6020026020010151815260200185602001516001846002020181518110610b0657fe5b602002602001015181525082602001819052506000610b9585600001518360020281518110610b3157fe5b602002602001015186600001516001856002020181518110610b4f57fe5b602002602001015187602001518560020281518110610b6a57fe5b602002602001015188602001516001876002020181518110610b8857fe5b6020026020010151610fff565b90506000610ba282611043565b9050818460e001518481518110610bb557fe5b602002602001018181525050808461010001518481518110610bd357fe5b602002602001018181525050610c54610c0b610bf8838461110d90919063ffffffff16565b8660200151610ec290919063ffffffff16565b610c46610c37610c24868761110d90919063ffffffff16565b8860000151610ec290919063ffffffff16565b8a610f6790919063ffffffff16565b610f6790919063ffffffff16565b965050508080600101915050610a5a565b5060606002885181610c7357fe5b04604051908082528060200260200182016040528015610ca25781602001602082028038833980820191505090505b50905060008090505b6002895181610cb657fe5b04811015610dc15760008090505b600286600001515181610cd357fe5b04811015610db3576000610cf7838360028a600001515181610cf157fe5b0461117f565b15610d1b578460e001518281518110610d0c57fe5b60200260200101519050610d37565b8461010001518281518110610d2c57fe5b602002602001015190505b6000821415610d5e5780848481518110610d4d57fe5b602002602001018181525050610da5565b610d8c610d8782868681518110610d7157fe5b602002602001015161110d90919063ffffffff16565b611147565b848481518110610d9857fe5b6020026020010181815250505b508080600101915050610cc4565b508080600101915050610cab565b50610dca61142e565b610e59610dfa610deb8760600151886040015161110d90919063ffffffff16565b89610ec290919063ffffffff16565b610e4b610e1d8860600151610e0f8d886111d0565b610ec290919063ffffffff16565b610e3d8960400151610e2f8f89611271565b610ec290919063ffffffff16565b610f6790919063ffffffff16565b610f6790919063ffffffff16565b905085600001518160000151148015610e79575085602001518160200151145b935050505095945050505050565b6000610ebb82604051602001610e9d9190611728565b6040516020818303038152906040528051906020012060001c611147565b9050919050565b610eca61142e565b6001821415610edb57829050610f61565b6002821415610ef557610eee8384610f67565b9050610f61565b610efd611448565b836000015181600060038110610f0f57fe5b602002018181525050836020015181600160038110610f2a57fe5b6020020181815250508281600260038110610f4157fe5b6020020181815250506040826060836007600019fa610f5f57600080fd5b505b92915050565b610f6f61142e565b610f7761146a565b836000015181600060048110610f8957fe5b602002018181525050836020015181600160048110610fa457fe5b602002018181525050826000015181600260048110610fbf57fe5b602002018181525050826020015181600360048110610fda57fe5b6020020181815250506040826080836006600019fa610ff857600080fd5b5092915050565b60006110398585858560405160200161101b9493929190611743565b6040516020818303038152906040528051906020012060001c611147565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190506000839050600081141561108257600092505050611108565b818111156110975781818161109357fe5b0690505b600080600190506000849050600084905060005b600082146110e0578183816110bc57fe5b049050838482028603838484028603809550819650829750839850505050506110ab565b60008512156110fd57846000038703975050505050505050611108565b849750505050505050505b919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808061113b57fe5b83850991505092915050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161117657fe5b06915050919050565b600080600183036001901b905060008090505b848110156111ae57600182901c91508080600101915050611192565b506000818616146111c35760019150506111c9565b60009150505b9392505050565b6111d861142e565b6060825160405190808252806020026020018201604052801561120a5781602001602082028038833980820191505090505b50905060008090505b835181101561125d5761123884828151811061122b57fe5b6020026020010151611043565b82828151811061124457fe5b6020026020010181815250508080600101915050611213565b506112688482611271565b91505092915050565b61127961142e565b61128161142e565b6112e48360008151811061129157fe5b60200260200101516040518060400160405280876000815181106112b157fe5b60200260200101518152602001876001815181106112cb57fe5b6020026020010151815250610ec290919063ffffffff16565b90506000600190505b83518110156113805761137161136285838151811061130857fe5b6020026020010151604051806040016040528089866002028151811061132a57fe5b6020026020010151815260200189600187600202018151811061134957fe5b6020026020010151815250610ec290919063ffffffff16565b83610f6790919063ffffffff16565b915080806001019150506112ed565b508091505092915050565b604051806060016040528061139e61148c565b81526020016113ab61148c565b81526020016113b86114a6565b81525090565b6040518061012001604052806113d261148c565b81526020016113df61148c565b81526020016113ec61148c565b81526020016113f961148c565b815260200161140661148c565b815260200161141361148c565b81526020016000815260200160608152602001606081525090565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b604051806040016040528060008152602001600081525090565b6040518060800160405280606081526020016060815260200160008152602001600081525090565b600082601f8301126114df57600080fd5b60026114f26114ed826117d9565b6117ac565b9150818385602084028201111561150857600080fd5b60005b83811015611538578161151e88826115c5565b84526020840193506020830192505060018101905061150b565b5050505092915050565b600082601f83011261155357600080fd5b8135611566611561826117fb565b6117ac565b9150818183526020840193506020810190508385602084028201111561158b57600080fd5b60005b838110156115bb57816115a188826115c5565b84526020840193506020830192505060018101905061158e565b5050505092915050565b6000813590506115d481611843565b92915050565b60008060008060008060008060006101608a8c0312156115f957600080fd5b60008a013567ffffffffffffffff81111561161357600080fd5b61161f8c828d01611542565b99505060208a013567ffffffffffffffff81111561163c57600080fd5b6116488c828d01611542565b98505060406116598c828d016114ce565b975050608061166a8c828d016114ce565b96505060c061167b8c828d016115c5565b95505060e08a013567ffffffffffffffff81111561169857600080fd5b6116a48c828d01611542565b9450506101008a013567ffffffffffffffff8111156116c257600080fd5b6116ce8c828d01611542565b9350506101206116e08c828d016115c5565b9250506101406116f28c828d016115c5565b9150509295985092959850929598565b61170b81611823565b82525050565b61172261171d8261182f565b611839565b82525050565b60006117348284611711565b60208201915081905092915050565b600061174f8287611711565b60208201915061175f8286611711565b60208201915061176f8285611711565b60208201915061177f8284611711565b60208201915081905095945050505050565b60006020820190506117a66000830184611702565b92915050565b6000604051905081810181811067ffffffffffffffff821117156117cf57600080fd5b8060405250919050565b600067ffffffffffffffff8211156117f057600080fd5b602082029050919050565b600067ffffffffffffffff82111561181257600080fd5b602082029050602081019050919050565b60008115159050919050565b6000819050919050565b6000819050919050565b61184c8161182f565b811461185757600080fd5b5056fea365627a7a72315820755878c10923e5e8f9b2be53c57557af413161a4fcb96d4325e7c09d767494bb6c6578706572696d656e74616cf564736f6c63430005100040"

// DeployIpverifier deploys a new Ethereum contract, binding an instance of Ipverifier to it.
func DeployIpverifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ipverifier, error) {
	parsed, err := abi.JSON(strings.NewReader(IpverifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IpverifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ipverifier{IpverifierCaller: IpverifierCaller{contract: contract}, IpverifierTransactor: IpverifierTransactor{contract: contract}, IpverifierFilterer: IpverifierFilterer{contract: contract}}, nil
}

// Ipverifier is an auto generated Go binding around an Ethereum contract.
type Ipverifier struct {
	IpverifierCaller     // Read-only binding to the contract
	IpverifierTransactor // Write-only binding to the contract
	IpverifierFilterer   // Log filterer for contract events
}

// IpverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpverifierSession struct {
	Contract     *Ipverifier       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpverifierCallerSession struct {
	Contract *IpverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IpverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpverifierTransactorSession struct {
	Contract     *IpverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IpverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpverifierRaw struct {
	Contract *Ipverifier // Generic contract binding to access the raw methods on
}

// IpverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpverifierCallerRaw struct {
	Contract *IpverifierCaller // Generic read-only contract binding to access the raw methods on
}

// IpverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpverifierTransactorRaw struct {
	Contract *IpverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpverifier creates a new instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifier(address common.Address, backend bind.ContractBackend) (*Ipverifier, error) {
	contract, err := bindIpverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipverifier{IpverifierCaller: IpverifierCaller{contract: contract}, IpverifierTransactor: IpverifierTransactor{contract: contract}, IpverifierFilterer: IpverifierFilterer{contract: contract}}, nil
}

// NewIpverifierCaller creates a new read-only instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifierCaller(address common.Address, caller bind.ContractCaller) (*IpverifierCaller, error) {
	contract, err := bindIpverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpverifierCaller{contract: contract}, nil
}

// NewIpverifierTransactor creates a new write-only instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IpverifierTransactor, error) {
	contract, err := bindIpverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpverifierTransactor{contract: contract}, nil
}

// NewIpverifierFilterer creates a new log filterer instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IpverifierFilterer, error) {
	contract, err := bindIpverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpverifierFilterer{contract: contract}, nil
}

// bindIpverifier binds a generic wrapper to an already deployed contract.
func bindIpverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IpverifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipverifier *IpverifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ipverifier.Contract.IpverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipverifier *IpverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipverifier.Contract.IpverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipverifier *IpverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipverifier.Contract.IpverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipverifier *IpverifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ipverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipverifier *IpverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipverifier *IpverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipverifier.Contract.contract.Transact(opts, method, params...)
}

// OptimizedVerifyIPProof is a free data retrieval call binding the contract method 0x97616863.
//
// Solidity: function optimizedVerifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierCaller) OptimizedVerifyIPProof(opts *bind.CallOpts, gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ipverifier.contract.Call(opts, out, "optimizedVerifyIPProof", gv, hv, p, u, c, l, r, a, b)
	return *ret0, err
}

// OptimizedVerifyIPProof is a free data retrieval call binding the contract method 0x97616863.
//
// Solidity: function optimizedVerifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierSession) OptimizedVerifyIPProof(gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.OptimizedVerifyIPProof(&_Ipverifier.CallOpts, gv, hv, p, u, c, l, r, a, b)
}

// OptimizedVerifyIPProof is a free data retrieval call binding the contract method 0x97616863.
//
// Solidity: function optimizedVerifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierCallerSession) OptimizedVerifyIPProof(gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.OptimizedVerifyIPProof(&_Ipverifier.CallOpts, gv, hv, p, u, c, l, r, a, b)
}

// VerifyIPProof is a free data retrieval call binding the contract method 0x7b7fc9ad.
//
// Solidity: function verifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierCaller) VerifyIPProof(opts *bind.CallOpts, gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ipverifier.contract.Call(opts, out, "verifyIPProof", gv, hv, p, u, c, l, r, a, b)
	return *ret0, err
}

// VerifyIPProof is a free data retrieval call binding the contract method 0x7b7fc9ad.
//
// Solidity: function verifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierSession) VerifyIPProof(gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProof(&_Ipverifier.CallOpts, gv, hv, p, u, c, l, r, a, b)
}

// VerifyIPProof is a free data retrieval call binding the contract method 0x7b7fc9ad.
//
// Solidity: function verifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierCallerSession) VerifyIPProof(gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProof(&_Ipverifier.CallOpts, gv, hv, p, u, c, l, r, a, b)
}
