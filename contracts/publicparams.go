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

// BN128G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN128G1Point struct {
	X *big.Int
	Y *big.Int
}

// PublicparamsABI is the input ABI used to generate the binding from.
const PublicparamsABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"aggSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gVector\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hVector\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"u\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getG\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getH\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getU\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAggGV\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN128.G1Point[128]\",\"name\":\"gv\",\"type\":\"tuple[128]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAggGVector\",\"outputs\":[{\"internalType\":\"uint256[256]\",\"name\":\"\",\"type\":\"uint256[256]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAggHV\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN128.G1Point[128]\",\"name\":\"hv\",\"type\":\"tuple[128]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAggHVector\",\"outputs\":[{\"internalType\":\"uint256[256]\",\"name\":\"\",\"type\":\"uint256[256]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGVector\",\"outputs\":[{\"internalType\":\"uint256[128]\",\"name\":\"\",\"type\":\"uint256[128]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHVector\",\"outputs\":[{\"internalType\":\"uint256[128]\",\"name\":\"\",\"type\":\"uint256[128]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PublicparamsBin is the compiled bytecode used for deploying new contracts.
var PublicparamsBin = "0x60806040523480156200001157600080fd5b50600160046000018190555060026004600101819055506200006e6040518060400160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c670000000000008152506200018660201b60201c565b600080820151816000015560208201518160010155905050620000cc6040518060400160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c670000000000008152506200018660201b60201c565b600260008201518160000155602082015181600101559050506200012b6040518060400160405280601b81526020017f752067656e657261746f72206f6620696e6e657270726f6475637400000000008152506200018660201b60201c565b600660008201518160000155602082015181600101559050506040801415620001685762000162600060406200021760201b60201c565b62000180565b6200017f600060026040026200021760201b60201c565b5b620007bd565b62000190620005c5565b6000620001d383604051602001620001a99190620006f9565b6040516020818303038152906040528051906020012060001c6200043260201b62000a9e1760201c565b90506200020f816004604051806040016040529081600082015481526020016001820154815250506200046b60201b62000ad61790919060201c565b915050919050565b600080620002596040516020016200022f9062000729565b6040516020818303038152906040528051906020012060001c6200043260201b62000a9e1760201c565b905060008490505b838110156200032057620002ad81830160405160200162000283919062000740565b6040516020818303038152906040528051906020012060001c6200043260201b62000a9e1760201c565b9250620002e9836004604051806040016040529081600082015481526020016001820154815250506200046b60201b62000ad61790919060201c565b60088260808110620002f757fe5b600202016000820151816000015560208201518160010155905050808060010191505062000261565b50600062000362604051602001620003389062000712565b6040516020818303038152906040528051906020012060001c6200043260201b62000a9e1760201c565b905060008590505b848110156200042a57620003b68282016040516020016200038c919062000740565b6040516020818303038152906040528051906020012060001c6200043260201b62000a9e1760201c565b9350620003f2846004604051806040016040529081600082015481526020016001820154815250506200046b60201b62000ad61790919060201c565b61010882608081106200040157fe5b60020201600082015181600001556020820151816001015590505080806001019150506200036a565b505050505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508083816200046257fe5b06915050919050565b62000475620005c5565b600182141562000488578290506200051e565b6002821415620004ac57620004a483846200052460201b60201c565b90506200051e565b620004b6620005df565b836000015181600060038110620004c957fe5b602002018181525050836020015181600160038110620004e557fe5b6020020181815250508281600260038110620004fd57fe5b6020020181815250506040826060836007600019fa6200051c57600080fd5b505b92915050565b6200052e620005c5565b6200053862000601565b8360000151816000600481106200054b57fe5b6020020181815250508360200151816001600481106200056757fe5b6020020181815250508260000151816002600481106200058357fe5b6020020181815250508260200151816003600481106200059f57fe5b6020020181815250506040826080836006600019fa620005be57600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b600062000630826200075d565b6200063c818562000768565b93506200064e8185602086016200077d565b80840191505092915050565b60006200066960038362000768565b91507f68767300000000000000000000000000000000000000000000000000000000006000830152600382019050919050565b6000620006ab60038362000768565b91507f67767300000000000000000000000000000000000000000000000000000000006000830152600382019050919050565b620006f3620006ed8262000773565b620007b3565b82525050565b600062000707828462000623565b915081905092915050565b60006200071f826200065a565b9150819050919050565b600062000736826200069c565b9150819050919050565b60006200074e8284620006de565b60208201915081905092915050565b600081519050919050565b600081905092915050565b6000819050919050565b60005b838110156200079d57808201518184015260208101905062000780565b83811115620007ad576000848401525b50505050565b6000819050919050565b6111f680620007cd6000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c8063713eb04d116100b8578063d498bef01161007c578063d498bef014610327578063da2b99ce14610345578063da89722414610363578063e1c7392a14610381578063e2179b8e1461038b578063ffe237f0146103aa57610137565b8063713eb04d1461028f5780637982ebcc146102ad57806382529fdb146102cb578063b8c9d365146102e9578063c6a898c51461030857610137565b80633e955225116100ff5780633e955225146101e5578063483767f0146102035780634db118751461022157806352c9b9651461023f57806356e2736c1461025e57610137565b806304c09ce91461013c5780630c00f8a01461015a57806324d6147d1461018b5780632e52d606146101a95780633e8d3764146101c7575b600080fd5b6101446103c8565b6040516101519190611069565b60405180910390f35b610174600480360361016f9190810190610d37565b610416565b60405161018292919061109f565b60405180910390f35b61019361043d565b6040516101a09190611069565b60405180910390f35b6101b161048c565b6040516101be9190611084565b60405180910390f35b6101cf610491565b6040516101dc9190611084565b60405180910390f35b6101ed610496565b6040516101fa9190611084565b60405180910390f35b61020b61049f565b6040516102189190611031565b60405180910390f35b610229610532565b604051610236919061104d565b60405180910390f35b6102476105c8565b60405161025592919061109f565b60405180910390f35b61027860048036036102739190810190610d37565b6105da565b60405161028692919061109f565b60405180910390f35b610297610602565b6040516102a49190611015565b60405180910390f35b6102b5610672565b6040516102c29190611084565b60405180910390f35b6102d3610677565b6040516102e09190611069565b60405180910390f35b6102f16106c6565b6040516102ff92919061109f565b60405180910390f35b6103106106d8565b60405161031e92919061109f565b60405180910390f35b61032f6106ea565b60405161033c9190611015565b60405180910390f35b61034d610759565b60405161035a919061104d565b60405180910390f35b61036b6107f1565b6040516103789190611084565b60405180910390f35b6103896107fa565b005b610393610830565b6040516103a192919061109f565b60405180910390f35b6103b2610842565b6040516103bf9190611031565b60405180910390f35b6103d0610c13565b6103d8610c13565b6000800154816000600281106103ea57fe5b6020020181815250506000600101548160016002811061040657fe5b6020020181815250508091505090565b6008816080811061042357fe5b600202016000915090508060000154908060010154905082565b610445610c13565b61044d610c13565b6006600001548160006002811061046057fe5b6020020181815250506006600101548160016002811061047c57fe5b6020020181815250508091505090565b600681565b604081565b60006006905090565b6104a7610c35565b6104af610c35565b60008090505b604081101561052a5761010881608081106104cc57fe5b60020201600001548282600202608081106104e357fe5b60200201818152505061010881608081106104fa57fe5b600202016001015482600183600202016080811061051457fe5b60200201818152505080806001019150506104b5565b508091505090565b61053a610c58565b610542610c58565b60008090505b60026040028110156105c0576008816080811061056157fe5b60020201600001548282600202610100811061057957fe5b6020020181815250506008816080811061058f57fe5b6002020160010154826001836002020161010081106105aa57fe5b6020020181815250508080600101915050610548565b508091505090565b60048060000154908060010154905082565b61010881608081106105e857fe5b600202016000915090508060000154908060010154905082565b61060a610c7c565b60008090505b600260400281101561066e57610108816080811061062a57fe5b600202016040518060400160405290816000820154815260200160018201548152505082826080811061065957fe5b60200201819052508080600101915050610610565b5090565b600281565b61067f610c13565b610687610c13565b6002600001548160006002811061069a57fe5b602002018181525050600260010154816001600281106106b657fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b6106f2610c7c565b60008090505b6002604002811015610755576008816080811061071157fe5b600202016040518060400160405290816000820154815260200160018201548152505082826080811061074057fe5b602002018190525080806001019150506106f8565b5090565b610761610c58565b610769610c58565b60008090505b60026040028110156107e957610108816080811061078957fe5b6002020160000154828260020261010081106107a157fe5b60200201818152505061010881608081106107b857fe5b6002020160010154826001836002020161010081106107d357fe5b602002018181525050808060010191505061076f565b508091505090565b60006040905090565b6000600860016002604002036080811061081057fe5b6002020160000154141561082e5761082d604060026040026108d3565b5b565b60008060000154908060010154905082565b61084a610c35565b610852610c35565b60008090505b60408110156108cb576008816080811061086e57fe5b600202016000015482826002026080811061088557fe5b6020020181815250506008816080811061089b57fe5b60020201600101548260018360020201608081106108b557fe5b6020020181815250508080600101915050610858565b508091505090565b6000806109066040516020016108e890610fe5565b6040516020818303038152906040528051906020012060001c610a9e565b905060008490505b838110156109b45761094a81830160405160200161092c9190610ffa565b6040516020818303038152906040528051906020012060001c610a9e565b925061097f83600460405180604001604052908160008201548152602001600182015481525050610ad690919063ffffffff16565b6008826080811061098c57fe5b600202016000820151816000015560208201518160010155905050808060010191505061090e565b5060006109e76040516020016109c990610fd0565b6040516020818303038152906040528051906020012060001c610a9e565b905060008590505b84811015610a9657610a2b828201604051602001610a0d9190610ffa565b6040516020818303038152906040528051906020012060001c610a9e565b9350610a6084600460405180604001604052908160008201548152602001600182015481525050610ad690919063ffffffff16565b6101088260808110610a6e57fe5b60020201600082015181600001556020820151816001015590505080806001019150506109ef565b505050505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381610acd57fe5b06915050919050565b610ade610caa565b6001821415610aef57829050610b75565b6002821415610b0957610b028384610b7b565b9050610b75565b610b11610cc4565b836000015181600060038110610b2357fe5b602002018181525050836020015181600160038110610b3e57fe5b6020020181815250508281600260038110610b5557fe5b6020020181815250506040826060836007600019fa610b7357600080fd5b505b92915050565b610b83610caa565b610b8b610ce6565b836000015181600060048110610b9d57fe5b602002018181525050836020015181600160048110610bb857fe5b602002018181525050826000015181600260048110610bd357fe5b602002018181525050826020015181600360048110610bee57fe5b6020020181815250506040826080836006600019fa610c0c57600080fd5b5092915050565b6040518060400160405280600290602082028038833980820191505090505090565b604051806110000160405280608090602082028038833980820191505090505090565b60405180612000016040528061010090602082028038833980820191505090505090565b6040518061100001604052806080905b610c94610d08565b815260200190600190039081610c8c5790505090565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b604051806040016040528060008152602001600081525090565b600081359050610d318161119c565b92915050565b600060208284031215610d4957600080fd5b6000610d5784828501610d22565b91505092915050565b6000610d6c8383610f6c565b60408301905092915050565b6000610d848383610f9b565b60208301905092915050565b610d99816110f0565b610da38184611151565b9250610dae826110c8565b8060005b83811015610ddf578151610dc68782610d60565b9650610dd18361111d565b925050600181019050610db2565b505050505050565b610df0816110fb565b610dfa818461115c565b9250610e05826110d2565b8060005b83811015610e36578151610e1d8782610d78565b9650610e288361112a565b925050600181019050610e09565b505050505050565b610e4781611106565b610e518184611167565b9250610e5c826110dc565b8060005b83811015610e8d578151610e748782610d78565b9650610e7f83611137565b925050600181019050610e60565b505050505050565b610e9e81611112565b610ea88184611172565b9250610eb3826110e6565b8060005b83811015610ee4578151610ecb8782610d78565b9650610ed683611144565b925050600181019050610eb7565b505050505050565b6000610ef960038361117d565b91507f68767300000000000000000000000000000000000000000000000000000000006000830152600382019050919050565b6000610f3960038361117d565b91507f67767300000000000000000000000000000000000000000000000000000000006000830152600382019050919050565b604082016000820151610f826000850182610f9b565b506020820151610f956020850182610f9b565b50505050565b610fa481611188565b82525050565b610fb381611188565b82525050565b610fca610fc582611188565b611192565b82525050565b6000610fdb82610eec565b9150819050919050565b6000610ff082610f2c565b9150819050919050565b60006110068284610fb9565b60208201915081905092915050565b60006120008201905061102b6000830184610d90565b92915050565b6000611000820190506110476000830184610de7565b92915050565b6000612000820190506110636000830184610e3e565b92915050565b600060408201905061107e6000830184610e95565b92915050565b60006020820190506110996000830184610faa565b92915050565b60006040820190506110b46000830185610faa565b6110c16020830184610faa565b9392505050565b6000819050919050565b6000819050919050565b6000819050919050565b6000819050919050565b600060809050919050565b600060809050919050565b60006101009050919050565b600060029050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600081905092915050565b600081905092915050565b600081905092915050565b600081905092915050565b600081905092915050565b6000819050919050565b6000819050919050565b6111a581611188565b81146111b057600080fd5b5056fea365627a7a723158204a1db4e87eea061058dd97d98adc6052cd7ddb9ff1adaa4d2738c48002d5f1016c6578706572696d656e74616cf564736f6c63430005100040"

// DeployPublicparams deploys a new Ethereum contract, binding an instance of Publicparams to it.
func DeployPublicparams(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Publicparams, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicparamsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PublicparamsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Publicparams{PublicparamsCaller: PublicparamsCaller{contract: contract}, PublicparamsTransactor: PublicparamsTransactor{contract: contract}, PublicparamsFilterer: PublicparamsFilterer{contract: contract}}, nil
}

// Publicparams is an auto generated Go binding around an Ethereum contract.
type Publicparams struct {
	PublicparamsCaller     // Read-only binding to the contract
	PublicparamsTransactor // Write-only binding to the contract
	PublicparamsFilterer   // Log filterer for contract events
}

// PublicparamsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicparamsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicparamsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicparamsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicparamsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicparamsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicparamsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicparamsSession struct {
	Contract     *Publicparams     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublicparamsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicparamsCallerSession struct {
	Contract *PublicparamsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PublicparamsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicparamsTransactorSession struct {
	Contract     *PublicparamsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PublicparamsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicparamsRaw struct {
	Contract *Publicparams // Generic contract binding to access the raw methods on
}

// PublicparamsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicparamsCallerRaw struct {
	Contract *PublicparamsCaller // Generic read-only contract binding to access the raw methods on
}

// PublicparamsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicparamsTransactorRaw struct {
	Contract *PublicparamsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicparams creates a new instance of Publicparams, bound to a specific deployed contract.
func NewPublicparams(address common.Address, backend bind.ContractBackend) (*Publicparams, error) {
	contract, err := bindPublicparams(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Publicparams{PublicparamsCaller: PublicparamsCaller{contract: contract}, PublicparamsTransactor: PublicparamsTransactor{contract: contract}, PublicparamsFilterer: PublicparamsFilterer{contract: contract}}, nil
}

// NewPublicparamsCaller creates a new read-only instance of Publicparams, bound to a specific deployed contract.
func NewPublicparamsCaller(address common.Address, caller bind.ContractCaller) (*PublicparamsCaller, error) {
	contract, err := bindPublicparams(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicparamsCaller{contract: contract}, nil
}

// NewPublicparamsTransactor creates a new write-only instance of Publicparams, bound to a specific deployed contract.
func NewPublicparamsTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicparamsTransactor, error) {
	contract, err := bindPublicparams(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicparamsTransactor{contract: contract}, nil
}

// NewPublicparamsFilterer creates a new log filterer instance of Publicparams, bound to a specific deployed contract.
func NewPublicparamsFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicparamsFilterer, error) {
	contract, err := bindPublicparams(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicparamsFilterer{contract: contract}, nil
}

// bindPublicparams binds a generic wrapper to an already deployed contract.
func bindPublicparams(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicparamsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Publicparams *PublicparamsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Publicparams.Contract.PublicparamsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Publicparams *PublicparamsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publicparams.Contract.PublicparamsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Publicparams *PublicparamsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Publicparams.Contract.PublicparamsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Publicparams *PublicparamsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Publicparams.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Publicparams *PublicparamsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publicparams.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Publicparams *PublicparamsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Publicparams.Contract.contract.Transact(opts, method, params...)
}

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) AggSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "aggSize")
	return *ret0, err
}

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() constant returns(uint256)
func (_Publicparams *PublicparamsSession) AggSize() (*big.Int, error) {
	return _Publicparams.Contract.AggSize(&_Publicparams.CallOpts)
}

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) AggSize() (*big.Int, error) {
	return _Publicparams.Contract.AggSize(&_Publicparams.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) BitSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "bitSize")
	return *ret0, err
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Publicparams *PublicparamsSession) BitSize() (*big.Int, error) {
	return _Publicparams.Contract.BitSize(&_Publicparams.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) BitSize() (*big.Int, error) {
	return _Publicparams.Contract.BitSize(&_Publicparams.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) G(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "g")
	return *ret, err
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.G(&_Publicparams.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.G(&_Publicparams.CallOpts)
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) GBase(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "gBase")
	return *ret, err
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) GBase() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GBase(&_Publicparams.CallOpts)
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) GBase() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GBase(&_Publicparams.CallOpts)
}

// GVector is a free data retrieval call binding the contract method 0x0c00f8a0.
//
// Solidity: function gVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) GVector(opts *bind.CallOpts, arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "gVector", arg0)
	return *ret, err
}

// GVector is a free data retrieval call binding the contract method 0x0c00f8a0.
//
// Solidity: function gVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) GVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GVector(&_Publicparams.CallOpts, arg0)
}

// GVector is a free data retrieval call binding the contract method 0x0c00f8a0.
//
// Solidity: function gVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) GVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GVector(&_Publicparams.CallOpts, arg0)
}

// GetAggGV is a free data retrieval call binding the contract method 0xd498bef0.
//
// Solidity: function getAggGV() constant returns([128]BN128G1Point gv)
func (_Publicparams *PublicparamsCaller) GetAggGV(opts *bind.CallOpts) ([128]BN128G1Point, error) {
	var (
		ret0 = new([128]BN128G1Point)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getAggGV")
	return *ret0, err
}

// GetAggGV is a free data retrieval call binding the contract method 0xd498bef0.
//
// Solidity: function getAggGV() constant returns([128]BN128G1Point gv)
func (_Publicparams *PublicparamsSession) GetAggGV() ([128]BN128G1Point, error) {
	return _Publicparams.Contract.GetAggGV(&_Publicparams.CallOpts)
}

// GetAggGV is a free data retrieval call binding the contract method 0xd498bef0.
//
// Solidity: function getAggGV() constant returns([128]BN128G1Point gv)
func (_Publicparams *PublicparamsCallerSession) GetAggGV() ([128]BN128G1Point, error) {
	return _Publicparams.Contract.GetAggGV(&_Publicparams.CallOpts)
}

// GetAggGVector is a free data retrieval call binding the contract method 0x4db11875.
//
// Solidity: function getAggGVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsCaller) GetAggGVector(opts *bind.CallOpts) ([256]*big.Int, error) {
	var (
		ret0 = new([256]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getAggGVector")
	return *ret0, err
}

// GetAggGVector is a free data retrieval call binding the contract method 0x4db11875.
//
// Solidity: function getAggGVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsSession) GetAggGVector() ([256]*big.Int, error) {
	return _Publicparams.Contract.GetAggGVector(&_Publicparams.CallOpts)
}

// GetAggGVector is a free data retrieval call binding the contract method 0x4db11875.
//
// Solidity: function getAggGVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsCallerSession) GetAggGVector() ([256]*big.Int, error) {
	return _Publicparams.Contract.GetAggGVector(&_Publicparams.CallOpts)
}

// GetAggHV is a free data retrieval call binding the contract method 0x713eb04d.
//
// Solidity: function getAggHV() constant returns([128]BN128G1Point hv)
func (_Publicparams *PublicparamsCaller) GetAggHV(opts *bind.CallOpts) ([128]BN128G1Point, error) {
	var (
		ret0 = new([128]BN128G1Point)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getAggHV")
	return *ret0, err
}

// GetAggHV is a free data retrieval call binding the contract method 0x713eb04d.
//
// Solidity: function getAggHV() constant returns([128]BN128G1Point hv)
func (_Publicparams *PublicparamsSession) GetAggHV() ([128]BN128G1Point, error) {
	return _Publicparams.Contract.GetAggHV(&_Publicparams.CallOpts)
}

// GetAggHV is a free data retrieval call binding the contract method 0x713eb04d.
//
// Solidity: function getAggHV() constant returns([128]BN128G1Point hv)
func (_Publicparams *PublicparamsCallerSession) GetAggHV() ([128]BN128G1Point, error) {
	return _Publicparams.Contract.GetAggHV(&_Publicparams.CallOpts)
}

// GetAggHVector is a free data retrieval call binding the contract method 0xda2b99ce.
//
// Solidity: function getAggHVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsCaller) GetAggHVector(opts *bind.CallOpts) ([256]*big.Int, error) {
	var (
		ret0 = new([256]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getAggHVector")
	return *ret0, err
}

// GetAggHVector is a free data retrieval call binding the contract method 0xda2b99ce.
//
// Solidity: function getAggHVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsSession) GetAggHVector() ([256]*big.Int, error) {
	return _Publicparams.Contract.GetAggHVector(&_Publicparams.CallOpts)
}

// GetAggHVector is a free data retrieval call binding the contract method 0xda2b99ce.
//
// Solidity: function getAggHVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsCallerSession) GetAggHVector() ([256]*big.Int, error) {
	return _Publicparams.Contract.GetAggHVector(&_Publicparams.CallOpts)
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) GetBitSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getBitSize")
	return *ret0, err
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() constant returns(uint256)
func (_Publicparams *PublicparamsSession) GetBitSize() (*big.Int, error) {
	return _Publicparams.Contract.GetBitSize(&_Publicparams.CallOpts)
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) GetBitSize() (*big.Int, error) {
	return _Publicparams.Contract.GetBitSize(&_Publicparams.CallOpts)
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() constant returns(uint256[2])
func (_Publicparams *PublicparamsCaller) GetG(opts *bind.CallOpts) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getG")
	return *ret0, err
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() constant returns(uint256[2])
func (_Publicparams *PublicparamsSession) GetG() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetG(&_Publicparams.CallOpts)
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() constant returns(uint256[2])
func (_Publicparams *PublicparamsCallerSession) GetG() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetG(&_Publicparams.CallOpts)
}

// GetGVector is a free data retrieval call binding the contract method 0xffe237f0.
//
// Solidity: function getGVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsCaller) GetGVector(opts *bind.CallOpts) ([128]*big.Int, error) {
	var (
		ret0 = new([128]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getGVector")
	return *ret0, err
}

// GetGVector is a free data retrieval call binding the contract method 0xffe237f0.
//
// Solidity: function getGVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsSession) GetGVector() ([128]*big.Int, error) {
	return _Publicparams.Contract.GetGVector(&_Publicparams.CallOpts)
}

// GetGVector is a free data retrieval call binding the contract method 0xffe237f0.
//
// Solidity: function getGVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsCallerSession) GetGVector() ([128]*big.Int, error) {
	return _Publicparams.Contract.GetGVector(&_Publicparams.CallOpts)
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() constant returns(uint256[2])
func (_Publicparams *PublicparamsCaller) GetH(opts *bind.CallOpts) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getH")
	return *ret0, err
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() constant returns(uint256[2])
func (_Publicparams *PublicparamsSession) GetH() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetH(&_Publicparams.CallOpts)
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() constant returns(uint256[2])
func (_Publicparams *PublicparamsCallerSession) GetH() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetH(&_Publicparams.CallOpts)
}

// GetHVector is a free data retrieval call binding the contract method 0x483767f0.
//
// Solidity: function getHVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsCaller) GetHVector(opts *bind.CallOpts) ([128]*big.Int, error) {
	var (
		ret0 = new([128]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getHVector")
	return *ret0, err
}

// GetHVector is a free data retrieval call binding the contract method 0x483767f0.
//
// Solidity: function getHVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsSession) GetHVector() ([128]*big.Int, error) {
	return _Publicparams.Contract.GetHVector(&_Publicparams.CallOpts)
}

// GetHVector is a free data retrieval call binding the contract method 0x483767f0.
//
// Solidity: function getHVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsCallerSession) GetHVector() ([128]*big.Int, error) {
	return _Publicparams.Contract.GetHVector(&_Publicparams.CallOpts)
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) GetN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getN")
	return *ret0, err
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() constant returns(uint256)
func (_Publicparams *PublicparamsSession) GetN() (*big.Int, error) {
	return _Publicparams.Contract.GetN(&_Publicparams.CallOpts)
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) GetN() (*big.Int, error) {
	return _Publicparams.Contract.GetN(&_Publicparams.CallOpts)
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() constant returns(uint256[2])
func (_Publicparams *PublicparamsCaller) GetU(opts *bind.CallOpts) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getU")
	return *ret0, err
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() constant returns(uint256[2])
func (_Publicparams *PublicparamsSession) GetU() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetU(&_Publicparams.CallOpts)
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() constant returns(uint256[2])
func (_Publicparams *PublicparamsCallerSession) GetU() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetU(&_Publicparams.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) H(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "h")
	return *ret, err
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.H(&_Publicparams.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.H(&_Publicparams.CallOpts)
}

// HVector is a free data retrieval call binding the contract method 0x56e2736c.
//
// Solidity: function hVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) HVector(opts *bind.CallOpts, arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "hVector", arg0)
	return *ret, err
}

// HVector is a free data retrieval call binding the contract method 0x56e2736c.
//
// Solidity: function hVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) HVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.HVector(&_Publicparams.CallOpts, arg0)
}

// HVector is a free data retrieval call binding the contract method 0x56e2736c.
//
// Solidity: function hVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) HVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.HVector(&_Publicparams.CallOpts, arg0)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "n")
	return *ret0, err
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Publicparams *PublicparamsSession) N() (*big.Int, error) {
	return _Publicparams.Contract.N(&_Publicparams.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) N() (*big.Int, error) {
	return _Publicparams.Contract.N(&_Publicparams.CallOpts)
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) U(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "u")
	return *ret, err
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) U() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.U(&_Publicparams.CallOpts)
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) U() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.U(&_Publicparams.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Publicparams *PublicparamsTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publicparams.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Publicparams *PublicparamsSession) Init() (*types.Transaction, error) {
	return _Publicparams.Contract.Init(&_Publicparams.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Publicparams *PublicparamsTransactorSession) Init() (*types.Transaction, error) {
	return _Publicparams.Contract.Init(&_Publicparams.TransactOpts)
}
