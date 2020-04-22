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

// PgcverifierABI is the input ABI used to generate the binding from.
const PgcverifierABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"params_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dleSigmaVerifier_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rangeProofVerifier_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggVerifier_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sigmaVerifier_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"aggRangeProofVerifier\",\"outputs\":[{\"internalType\":\"contractAggRangeProofVerifier\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dleSigmaVerifier\",\"outputs\":[{\"internalType\":\"contractDLESigmaVerifier\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"internalType\":\"contractPublicParams\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rangeProofVerifier\",\"outputs\":[{\"internalType\":\"contractRangeProofVerifier\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sigmaVerifier\",\"outputs\":[{\"internalType\":\"contractSigmaVerifier\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[36]\",\"name\":\"points\",\"type\":\"uint256[36]\"},{\"internalType\":\"uint256[10]\",\"name\":\"scalar\",\"type\":\"uint256[10]\"},{\"internalType\":\"uint256[14]\",\"name\":\"l\",\"type\":\"uint256[14]\"},{\"internalType\":\"uint256[14]\",\"name\":\"r\",\"type\":\"uint256[14]\"},{\"internalType\":\"uint256[4]\",\"name\":\"ub\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"verifyAggTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"publicKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"proof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"ub\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[]\",\"name\":\"input\",\"type\":\"uint256[]\"}],\"name\":\"verifyBurn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PgcverifierBin is the compiled bytecode used for deploying new contracts.
var PgcverifierBin = "0x60806040523480156200001157600080fd5b50604051620025ed380380620025ed833981810160405260a08110156200003757600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919050505084600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620001bf620004bb565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff1660e01b8152600401604080518083038186803b1580156200022757600080fd5b505afa1580156200023c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200026257600080fd5b8101908091905050905062000276620004bb565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff1660e01b8152600401604080518083038186803b158015620002de57600080fd5b505afa158015620002f3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200031957600080fd5b81019080919050509050816000600281106200033157fe5b60200201516000800181905550816001600281106200034c57fe5b6020020151600060010181905550806000600281106200036857fe5b6020020151600260000181905550806001600281106200038457fe5b6020020151600260010181905550600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da8972246040518163ffffffff1660e01b815260040160206040518083038186803b158015620003fb57600080fd5b505afa15801562000410573d6000803e3d6000fd5b505050506040513d60208110156200042757600080fd5b8101908080519060200190929190505050604014620004ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f62697473697a65206e6f7420657175616c00000000000000000000000000000081525060200191505060405180910390fd5b50505050505050620004dd565b6040518060400160405280600290602082028038833980820191505090505090565b61210080620004ed6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80633e8d3764116100715780633e8d3764146104cb57806361e5af5e146104e957806381b7bc7314610533578063b8c9d3651461057d578063cff0ab96146105a2578063e2179b8e146105ec576100b4565b806326f9031b146100b9578063293ec7b4146102615780632e52d606146102ab578063361eb474146102c95780633a4f6999146103135780633db3595c14610331575b600080fd5b61024760048036036101a08110156100d057600080fd5b810190808035906020019092919080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f820116905080830192505050505050919291929080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001906401000000008111156101c457600080fd5b8201836020820111156101d657600080fd5b803590602001918460208302840111640100000000831117156101f857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610611565b604051808215151515815260200191505060405180910390f35b61026961075c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102b3610782565b6040518082815260200191505060405180910390f35b6102d1610787565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61031b6107ad565b6040518082815260200191505060405180910390f35b6104b16004803603610a0081101561034857600080fd5b810190808061048001906024806020026040519081016040528092919082602460200280828437600081840152601f19601f8201169050808301925050505050509192919290806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f8201169050808301925050505050509192919290806101c00190600e806020026040519081016040528092919082600e60200280828437600081840152601f19601f8201169050808301925050505050509192919290806101c00190600e806020026040519081016040528092919082600e60200280828437600081840152601f19601f820116905080830192505050505050919291929080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f820116905080830192505050505050919291929080359060200190929190803590602001909291905050506107b5565b604051808215151515815260200191505060405180910390f35b6104d3611391565b6040518082815260200191505060405180910390f35b6104f1611396565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61053b6113bc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105856113e2565b604051808381526020018281526020019250505060405180910390f35b6105aa6113f4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105f461141a565b604051808381526020018281526020019250505060405180910390f35b600061061b611d75565b8360006004811061062857fe5b6020020151816000015160000181815250508360016004811061064757fe5b6020020151816000015160200181815250508360026004811061066657fe5b6020020151816020015160000181815250508360036004811061068557fe5b6020020151816020015160200181815250506106db888260405180604001604052808b6000600281106106b457fe5b602002015181526020018b6001600281106106cb57fe5b602002015181525089898861142c565b61074d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f646c65207369676d6120766572696679206661696c656400000000000000000081525060200191505060405180910390fd5b60019150509695505050505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604060020a81565b60006107bf611d75565b846000600481106107cc57fe5b602002015181600001516000018181525050846001600481106107eb57fe5b6020020151816000015160200181815250508460026004811061080a57fe5b6020020151816020015160000181815250508460036004811061082957fe5b602002015181602001516020018181525050610843611da2565b60008160000181815250505b6010816000015110156108a3578981600001516024811061086c57fe5b6020020151816040015182600001516010811061088557fe5b6020020181815250508060000180518091906001018152505061084f565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e1efe76682604001518b6000600a81106108f357fe5b60200201518c6001600a811061090557fe5b60200201516040518463ffffffff1660e01b81526004018084601060200280838360005b83811015610944578082015181840152602081019050610929565b505050509050018381526020018281526020019350505050602060405180830381600087803b15801561097657600080fd5b505af115801561098a573d6000803e3d6000fd5b505050506040513d60208110156109a057600080fd5b8101908080519060200190929190505050610a23576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f70746520657175616c2070726f6f6620696e76616c696400000000000000000081525060200191505060405180910390fd5b60405180604001604052808b600460248110610a3b57fe5b602002015181526020018b600560248110610a5257fe5b602002015181525081608001516000018190525060405180604001604052808b600860248110610a7e57fe5b602002015181526020018b600960248110610a9557fe5b6020020151815250816080015160200181905250610ad0610abd826080015160000151611788565b836000015161181890919063ffffffff16565b81610100015160000181905250610b04610af1826080015160200151611788565b836020015161181890919063ffffffff16565b8161010001516020018190525060405180604001604052808b601060248110610b2957fe5b602002015181526020018b601160248110610b4057fe5b60200201518152508161012001516000018190525060405180604001604052808b601260248110610b6d57fe5b602002015181526020018b601360248110610b8457fe5b60200201518152508161012001516020018190525060405180604001604052808b601860248110610bb157fe5b602002015181526020018b601960248110610bc857fe5b6020020151815250816101400151600060028110610be257fe5b602002018190525060405180604001604052808b601a60248110610c0257fe5b602002015181526020018b601b60248110610c1957fe5b6020020151815250816101400151600160028110610c3357fe5b6020020181905250600c604051908082528060200260200182016040528015610c6b5781602001602082028038833980820191505090505b50816102e0018190525084816102e00151600081518110610c8857fe5b60200260200101818152505083816102e00151600181518110610ca757fe5b60200260200101818152505060008160000181815250505b600a81600001511015610d1a5789816000015160248110610cdc57fe5b6020020151816102e00151600283600001510181518110610cf957fe5b60200260200101818152505080600001805180919060010181525050610cbf565b610d728161010001518261012001518361014001518d600060248110610d3c57fe5b60200201518e600160248110610d4e57fe5b60200201518e6004600a8110610d6057fe5b60200201516001886102e001516118b0565b610de4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f646c65207369676d612070726f6f66206661696c65640000000000000000000081525060200191505060405180910390fd5b89600060248110610df157fe5b602002015181606001516000600a8110610e0757fe5b60200201818152505089600160248110610e1d57fe5b602002015181606001516001600a8110610e3357fe5b60200201818152505060108160000181815250505b601881600001511015610ea25789816000015160248110610e6557fe5b602002015181606001516010836000015160020103600a8110610e8457fe5b60200201818152505080600001805180919060010181525050610e48565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663991395ca82606001518b6002600a8110610ef257fe5b60200201518c6003600a8110610f0457fe5b60200201516040518463ffffffff1660e01b81526004018084600a60200280838360005b83811015610f43578082015181840152602081019050610f28565b505050509050018381526020018281526020019350505050602060405180830381600087803b158015610f7557600080fd5b505af1158015610f89573d6000803e3d6000fd5b505050506040513d6020811015610f9f57600080fd5b8101908080519060200190929190505050611022576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f63742076616c69642070726f6f6620696e76616c69640000000000000000000081525060200191505060405180910390fd5b60008160000181815250505b60088160000151101561108657898160000151601c016024811061104e57fe5b60200201518161022001518260000151600c811061106857fe5b6020020181815250508060000180518091906001018152505061102e565b8960086024811061109357fe5b60200201518161022001516008600c81106110aa57fe5b602002018181525050896009602481106110c057fe5b60200201518161022001516009600c81106110d757fe5b602002018181525050896012602481106110ed57fe5b6020020151816102200151600a600c811061110457fe5b6020020181815250508960136024811061111a57fe5b6020020151816102200151600b600c811061113157fe5b60200201818152505060008160000181815250505b60058160000151101561119e57886005826000015101600a811061116657fe5b602002015181610240015182600001516005811061118057fe5b60200201818152505080600001805180919060010181525050611146565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e5d34a638261022001518361024001518b8b6040518563ffffffff1660e01b81526004018085600c60200280838360005b83811015611224578082015181840152602081019050611209565b5050505090500184600560200280838360005b83811015611252578082015181840152602081019050611237565b5050505090500183600e60200280838360005b83811015611280578082015181840152602081019050611265565b5050505090500182600e60200280838360005b838110156112ae578082015181840152602081019050611293565b5050505090500194505050505060206040518083038186803b1580156112d357600080fd5b505afa1580156112e7573d6000803e3d6000fd5b505050506040513d60208110156112fd57600080fd5b8101908080519060200190929190505050611380576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f616767726174652072616e67652070726f6f6620696e76616c6964000000000081525060200191505060405180910390fd5b600192505050979650505050505050565b604081565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060000154908060010154905082565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60028060000154908060010154905082565b6000611436611f06565b61148761147461146f8a600260405180604001604052908160008201548152602001600182015481525050611cd090919063ffffffff16565b611788565b886020015161181890919063ffffffff16565b9050611491611da2565b8560006004811061149e57fe5b60200201518161016001516000600c81106114b557fe5b602002018181525050856001600481106114cb57fe5b60200201518161016001516001600c81106114e257fe5b602002018181525050856002600481106114f857fe5b60200201518161016001516002600c811061150f57fe5b6020020181815250508560036004811061152557fe5b60200201518161016001516003600c811061153c57fe5b60200201818152505081600001518161016001516004600c811061155c57fe5b60200201818152505081602001518161016001516005600c811061157c57fe5b6020020181815250508760000151600001518161016001516006600c81106115a057fe5b6020020181815250508760000151602001518161016001516007600c81106115c457fe5b60200201818152505060008001548161016001516008600c81106115e457fe5b6020020181815250506000600101548161016001516009600c811061160557fe5b6020020181815250508660000151816101600151600a600c811061162557fe5b6020020181815250508660200151816101600151600b600c811061164557fe5b602002018181525050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e26d0b882610160015187876040518463ffffffff1660e01b81526004018084600c60200280838360005b838110156116ce5780820151818401526020810190506116b3565b5050505090500183815260200180602001828103825283818151815260200191508051906020019060200280838360005b8381101561171a5780820151818401526020810190506116ff565b5050505090500194505050505060206040518083038186803b15801561173f57600080fd5b505afa158015611753573d6000803e3d6000fd5b505050506040513d602081101561176957600080fd5b8101908080519060200190929190505050925050509695505050505050565b611790611f06565b600082600001511480156117a8575060008260200151145b156117cb5760405180604001604052806000815260200160008152509050611813565b60007f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd479050604051806040016040528084600001518152602001846020015183038152509150505b919050565b611820611f06565b611828611f20565b83600001518160006004811061183a57fe5b60200201818152505083602001518160016004811061185557fe5b60200201818152505082600001518160026004811061187057fe5b60200201818152505082602001518160036004811061188b57fe5b6020020181815250506040826080836006600019fa6118a957600080fd5b5092915050565b60006118ba611f06565b6118dd6118ca8b60200151611788565b8a6020015161181890919063ffffffff16565b90506118e7611f06565b61190a6118f78c60000151611788565b8b6000015161181890919063ffffffff16565b9050611914611f42565b8960006002811061192157fe5b602002015160000151816000600c811061193757fe5b6020020181815250508960006002811061194d57fe5b602002015160200151816001600c811061196357fe5b6020020181815250508960016002811061197957fe5b602002015160000151816002600c811061198f57fe5b602002018181525050896001600281106119a557fe5b602002015160200151816003600c81106119bb57fe5b6020020181815250508260000151816004600c81106119d657fe5b6020020181815250508260200151816005600c81106119f157fe5b6020020181815250508160000151816006600c8110611a0c57fe5b6020020181815250508160200151816007600c8110611a2757fe5b6020020181815250506000800154816008600c8110611a4257fe5b602002018181525050600060010154816009600c8110611a5e57fe5b6020020181815250508881600a600c8110611a7557fe5b6020020181815250508781600b600c8110611a8c57fe5b6020020181815250506000861415611b8657600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639751cb1382896040518363ffffffff1660e01b81526004018083600c60200280838360005b83811015611b18578082015181840152602081019050611afd565b505050509050018281526020019250505060206040518083038186803b158015611b4157600080fd5b505afa158015611b55573d6000803e3d6000fd5b505050506040513d6020811015611b6b57600080fd5b81019080805190602001909291905050509350505050611cc4565b6001861415611cc057600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e26d0b88289886040518463ffffffff1660e01b81526004018084600c60200280838360005b83811015611c0a578082015181840152602081019050611bef565b5050505090500183815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015611c56578082015181840152602081019050611c3b565b5050505090500194505050505060206040518083038186803b158015611c7b57600080fd5b505afa158015611c8f573d6000803e3d6000fd5b505050506040513d6020811015611ca557600080fd5b81019080805190602001909291905050509350505050611cc4565b5050505b98975050505050505050565b611cd8611f06565b6001821415611ce957829050611d6f565b6002821415611d0357611cfc8384611818565b9050611d6f565b611d0b611f65565b836000015181600060038110611d1d57fe5b602002018181525050836020015181600160038110611d3857fe5b6020020181815250508281600260038110611d4f57fe5b6020020181815250506040826060836007600019fa611d6d57600080fd5b505b92915050565b6040518060600160405280611d88611f87565b8152602001611d95611f87565b8152602001600081525090565b60405180610360016040528060008152602001611dbd611fa1565b8152602001611dca611fc4565b8152602001611dd7611fe7565b8152602001611de461200a565b8152602001611df161200a565b8152602001611dfe612037565b8152602001611e0b612037565b8152602001611e1861200a565b8152602001611e2561200a565b8152602001611e32612059565b8152602001611e3f612086565b8152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001611e77612037565b8152602001611e84611fe7565b8152602001611e91612086565b8152602001611e9e6120a9565b8152602001611eab612086565b8152602001611eb8612086565b8152602001611ec5612086565b8152602001611ed2612086565b815260200160608152602001611ee6612037565b8152602001611ef3612037565b8152602001611f00612037565b81525090565b604051806040016040528060008152602001600081525090565b6040518060800160405280600490602082028038833980820191505090505090565b604051806101800160405280600c90602082028038833980820191505090505090565b6040518060600160405280600390602082028038833980820191505090505090565b604051806040016040528060008152602001600081525090565b604051806102800160405280601490602082028038833980820191505090505090565b604051806102000160405280601090602082028038833980820191505090505090565b604051806101400160405280600a90602082028038833980820191505090505090565b604051806060016040528061201d611f87565b815260200161202a611f87565b8152602001600081525090565b6040518060800160405280600490602082028038833980820191505090505090565b60405180604001604052806002905b612070611f87565b8152602001906001900390816120685790505090565b604051806101800160405280600c90602082028038833980820191505090505090565b6040518060a0016040528060059060208202803883398082019150509050509056fea265627a7a72315820fb1b7c9cbabb79f232216e778eabe6560fc6561e6715357705fb4c28f7f42a6264736f6c63430005100032"

// DeployPgcverifier deploys a new Ethereum contract, binding an instance of Pgcverifier to it.
func DeployPgcverifier(auth *bind.TransactOpts, backend bind.ContractBackend, params_ common.Address, dleSigmaVerifier_ common.Address, rangeProofVerifier_ common.Address, aggVerifier_ common.Address, sigmaVerifier_ common.Address) (common.Address, *types.Transaction, *Pgcverifier, error) {
	parsed, err := abi.JSON(strings.NewReader(PgcverifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PgcverifierBin), backend, params_, dleSigmaVerifier_, rangeProofVerifier_, aggVerifier_, sigmaVerifier_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pgcverifier{PgcverifierCaller: PgcverifierCaller{contract: contract}, PgcverifierTransactor: PgcverifierTransactor{contract: contract}, PgcverifierFilterer: PgcverifierFilterer{contract: contract}}, nil
}

// Pgcverifier is an auto generated Go binding around an Ethereum contract.
type Pgcverifier struct {
	PgcverifierCaller     // Read-only binding to the contract
	PgcverifierTransactor // Write-only binding to the contract
	PgcverifierFilterer   // Log filterer for contract events
}

// PgcverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type PgcverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgcverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PgcverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgcverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PgcverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgcverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PgcverifierSession struct {
	Contract     *Pgcverifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PgcverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PgcverifierCallerSession struct {
	Contract *PgcverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PgcverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PgcverifierTransactorSession struct {
	Contract     *PgcverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PgcverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type PgcverifierRaw struct {
	Contract *Pgcverifier // Generic contract binding to access the raw methods on
}

// PgcverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PgcverifierCallerRaw struct {
	Contract *PgcverifierCaller // Generic read-only contract binding to access the raw methods on
}

// PgcverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PgcverifierTransactorRaw struct {
	Contract *PgcverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPgcverifier creates a new instance of Pgcverifier, bound to a specific deployed contract.
func NewPgcverifier(address common.Address, backend bind.ContractBackend) (*Pgcverifier, error) {
	contract, err := bindPgcverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pgcverifier{PgcverifierCaller: PgcverifierCaller{contract: contract}, PgcverifierTransactor: PgcverifierTransactor{contract: contract}, PgcverifierFilterer: PgcverifierFilterer{contract: contract}}, nil
}

// NewPgcverifierCaller creates a new read-only instance of Pgcverifier, bound to a specific deployed contract.
func NewPgcverifierCaller(address common.Address, caller bind.ContractCaller) (*PgcverifierCaller, error) {
	contract, err := bindPgcverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PgcverifierCaller{contract: contract}, nil
}

// NewPgcverifierTransactor creates a new write-only instance of Pgcverifier, bound to a specific deployed contract.
func NewPgcverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*PgcverifierTransactor, error) {
	contract, err := bindPgcverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PgcverifierTransactor{contract: contract}, nil
}

// NewPgcverifierFilterer creates a new log filterer instance of Pgcverifier, bound to a specific deployed contract.
func NewPgcverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*PgcverifierFilterer, error) {
	contract, err := bindPgcverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PgcverifierFilterer{contract: contract}, nil
}

// bindPgcverifier binds a generic wrapper to an already deployed contract.
func bindPgcverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PgcverifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pgcverifier *PgcverifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pgcverifier.Contract.PgcverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pgcverifier *PgcverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pgcverifier.Contract.PgcverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pgcverifier *PgcverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pgcverifier.Contract.PgcverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pgcverifier *PgcverifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pgcverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pgcverifier *PgcverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pgcverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pgcverifier *PgcverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pgcverifier.Contract.contract.Transact(opts, method, params...)
}

// AggRangeProofVerifier is a free data retrieval call binding the contract method 0x293ec7b4.
//
// Solidity: function aggRangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCaller) AggRangeProofVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "aggRangeProofVerifier")
	return *ret0, err
}

// AggRangeProofVerifier is a free data retrieval call binding the contract method 0x293ec7b4.
//
// Solidity: function aggRangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierSession) AggRangeProofVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.AggRangeProofVerifier(&_Pgcverifier.CallOpts)
}

// AggRangeProofVerifier is a free data retrieval call binding the contract method 0x293ec7b4.
//
// Solidity: function aggRangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCallerSession) AggRangeProofVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.AggRangeProofVerifier(&_Pgcverifier.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Pgcverifier *PgcverifierCaller) BitSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "bitSize")
	return *ret0, err
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Pgcverifier *PgcverifierSession) BitSize() (*big.Int, error) {
	return _Pgcverifier.Contract.BitSize(&_Pgcverifier.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Pgcverifier *PgcverifierCallerSession) BitSize() (*big.Int, error) {
	return _Pgcverifier.Contract.BitSize(&_Pgcverifier.CallOpts)
}

// DleSigmaVerifier is a free data retrieval call binding the contract method 0x361eb474.
//
// Solidity: function dleSigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCaller) DleSigmaVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "dleSigmaVerifier")
	return *ret0, err
}

// DleSigmaVerifier is a free data retrieval call binding the contract method 0x361eb474.
//
// Solidity: function dleSigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierSession) DleSigmaVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.DleSigmaVerifier(&_Pgcverifier.CallOpts)
}

// DleSigmaVerifier is a free data retrieval call binding the contract method 0x361eb474.
//
// Solidity: function dleSigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCallerSession) DleSigmaVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.DleSigmaVerifier(&_Pgcverifier.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierCaller) G(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Pgcverifier.contract.Call(opts, out, "g")
	return *ret, err
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Pgcverifier.Contract.G(&_Pgcverifier.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierCallerSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Pgcverifier.Contract.G(&_Pgcverifier.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierCaller) H(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Pgcverifier.contract.Call(opts, out, "h")
	return *ret, err
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Pgcverifier.Contract.H(&_Pgcverifier.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierCallerSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Pgcverifier.Contract.H(&_Pgcverifier.CallOpts)
}

// MaxNumber is a free data retrieval call binding the contract method 0x3a4f6999.
//
// Solidity: function maxNumber() constant returns(uint256)
func (_Pgcverifier *PgcverifierCaller) MaxNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "maxNumber")
	return *ret0, err
}

// MaxNumber is a free data retrieval call binding the contract method 0x3a4f6999.
//
// Solidity: function maxNumber() constant returns(uint256)
func (_Pgcverifier *PgcverifierSession) MaxNumber() (*big.Int, error) {
	return _Pgcverifier.Contract.MaxNumber(&_Pgcverifier.CallOpts)
}

// MaxNumber is a free data retrieval call binding the contract method 0x3a4f6999.
//
// Solidity: function maxNumber() constant returns(uint256)
func (_Pgcverifier *PgcverifierCallerSession) MaxNumber() (*big.Int, error) {
	return _Pgcverifier.Contract.MaxNumber(&_Pgcverifier.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Pgcverifier *PgcverifierCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "n")
	return *ret0, err
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Pgcverifier *PgcverifierSession) N() (*big.Int, error) {
	return _Pgcverifier.Contract.N(&_Pgcverifier.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Pgcverifier *PgcverifierCallerSession) N() (*big.Int, error) {
	return _Pgcverifier.Contract.N(&_Pgcverifier.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() constant returns(address)
func (_Pgcverifier *PgcverifierCaller) Params(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "params")
	return *ret0, err
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() constant returns(address)
func (_Pgcverifier *PgcverifierSession) Params() (common.Address, error) {
	return _Pgcverifier.Contract.Params(&_Pgcverifier.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() constant returns(address)
func (_Pgcverifier *PgcverifierCallerSession) Params() (common.Address, error) {
	return _Pgcverifier.Contract.Params(&_Pgcverifier.CallOpts)
}

// RangeProofVerifier is a free data retrieval call binding the contract method 0x61e5af5e.
//
// Solidity: function rangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCaller) RangeProofVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "rangeProofVerifier")
	return *ret0, err
}

// RangeProofVerifier is a free data retrieval call binding the contract method 0x61e5af5e.
//
// Solidity: function rangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierSession) RangeProofVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.RangeProofVerifier(&_Pgcverifier.CallOpts)
}

// RangeProofVerifier is a free data retrieval call binding the contract method 0x61e5af5e.
//
// Solidity: function rangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCallerSession) RangeProofVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.RangeProofVerifier(&_Pgcverifier.CallOpts)
}

// SigmaVerifier is a free data retrieval call binding the contract method 0x81b7bc73.
//
// Solidity: function sigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCaller) SigmaVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "sigmaVerifier")
	return *ret0, err
}

// SigmaVerifier is a free data retrieval call binding the contract method 0x81b7bc73.
//
// Solidity: function sigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierSession) SigmaVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.SigmaVerifier(&_Pgcverifier.CallOpts)
}

// SigmaVerifier is a free data retrieval call binding the contract method 0x81b7bc73.
//
// Solidity: function sigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCallerSession) SigmaVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.SigmaVerifier(&_Pgcverifier.CallOpts)
}

// VerifyAggTransfer is a paid mutator transaction binding the contract method 0x3db3595c.
//
// Solidity: function verifyAggTransfer(uint256[36] points, uint256[10] scalar, uint256[14] l, uint256[14] r, uint256[4] ub, uint256 nonce, uint256 token) returns(bool)
func (_Pgcverifier *PgcverifierTransactor) VerifyAggTransfer(opts *bind.TransactOpts, points [36]*big.Int, scalar [10]*big.Int, l [14]*big.Int, r [14]*big.Int, ub [4]*big.Int, nonce *big.Int, token *big.Int) (*types.Transaction, error) {
	return _Pgcverifier.contract.Transact(opts, "verifyAggTransfer", points, scalar, l, r, ub, nonce, token)
}

// VerifyAggTransfer is a paid mutator transaction binding the contract method 0x3db3595c.
//
// Solidity: function verifyAggTransfer(uint256[36] points, uint256[10] scalar, uint256[14] l, uint256[14] r, uint256[4] ub, uint256 nonce, uint256 token) returns(bool)
func (_Pgcverifier *PgcverifierSession) VerifyAggTransfer(points [36]*big.Int, scalar [10]*big.Int, l [14]*big.Int, r [14]*big.Int, ub [4]*big.Int, nonce *big.Int, token *big.Int) (*types.Transaction, error) {
	return _Pgcverifier.Contract.VerifyAggTransfer(&_Pgcverifier.TransactOpts, points, scalar, l, r, ub, nonce, token)
}

// VerifyAggTransfer is a paid mutator transaction binding the contract method 0x3db3595c.
//
// Solidity: function verifyAggTransfer(uint256[36] points, uint256[10] scalar, uint256[14] l, uint256[14] r, uint256[4] ub, uint256 nonce, uint256 token) returns(bool)
func (_Pgcverifier *PgcverifierTransactorSession) VerifyAggTransfer(points [36]*big.Int, scalar [10]*big.Int, l [14]*big.Int, r [14]*big.Int, ub [4]*big.Int, nonce *big.Int, token *big.Int) (*types.Transaction, error) {
	return _Pgcverifier.Contract.VerifyAggTransfer(&_Pgcverifier.TransactOpts, points, scalar, l, r, ub, nonce, token)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x26f9031b.
//
// Solidity: function verifyBurn(uint256 amount, uint256[2] publicKey, uint256[4] proof, uint256 z, uint256[4] ub, uint256[] input) returns(bool)
func (_Pgcverifier *PgcverifierTransactor) VerifyBurn(opts *bind.TransactOpts, amount *big.Int, publicKey [2]*big.Int, proof [4]*big.Int, z *big.Int, ub [4]*big.Int, input []*big.Int) (*types.Transaction, error) {
	return _Pgcverifier.contract.Transact(opts, "verifyBurn", amount, publicKey, proof, z, ub, input)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x26f9031b.
//
// Solidity: function verifyBurn(uint256 amount, uint256[2] publicKey, uint256[4] proof, uint256 z, uint256[4] ub, uint256[] input) returns(bool)
func (_Pgcverifier *PgcverifierSession) VerifyBurn(amount *big.Int, publicKey [2]*big.Int, proof [4]*big.Int, z *big.Int, ub [4]*big.Int, input []*big.Int) (*types.Transaction, error) {
	return _Pgcverifier.Contract.VerifyBurn(&_Pgcverifier.TransactOpts, amount, publicKey, proof, z, ub, input)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x26f9031b.
//
// Solidity: function verifyBurn(uint256 amount, uint256[2] publicKey, uint256[4] proof, uint256 z, uint256[4] ub, uint256[] input) returns(bool)
func (_Pgcverifier *PgcverifierTransactorSession) VerifyBurn(amount *big.Int, publicKey [2]*big.Int, proof [4]*big.Int, z *big.Int, ub [4]*big.Int, input []*big.Int) (*types.Transaction, error) {
	return _Pgcverifier.Contract.VerifyBurn(&_Pgcverifier.TransactOpts, amount, publicKey, proof, z, ub, input)
}
