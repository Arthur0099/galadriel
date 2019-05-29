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

// RangeproofverifierABI is the input ABI used to generate the binding from.
const RangeproofverifierABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"uBase\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rpParams\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ipVerifier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hvBase\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gvBase\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"params\",\"type\":\"address\"},{\"name\":\"ip\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[10]\"},{\"name\":\"scalar\",\"type\":\"uint256[5]\"},{\"name\":\"l\",\"type\":\"uint256[8]\"},{\"name\":\"r\",\"type\":\"uint256[8]\"}],\"name\":\"verifyRangeProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RangeproofverifierBin is the compiled bytecode used for deploying new contracts.
const RangeproofverifierBin = `0x60806040523480156200001157600080fd5b5060405160408062002528833981018060405260408110156200003357600080fd5b810190808051906020019092919080519060200190929190505050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da8972246040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200015357600080fd5b505afa15801562000168573d6000803e3d6000fd5b505050506040513d60208110156200017f57600080fd5b8101908080519060200190929190505050601014151562000208576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6269732073697a65206e6f7420657175616c000000000000000000000000000081525060200191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633e9552256040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200028c57600080fd5b505afa158015620002a1573d6000803e3d6000fd5b505050506040513d6020811015620002b857600080fd5b8101908080519060200190929190505050600414151562000341576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6e756d626572206f66206c2c72206e6f7420657175616c00000000000000000081525060200191505060405180910390fd5b6200034b6200090f565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b158015620003ce57600080fd5b505afa158015620003e3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200040957600080fd5b810190809190505090506200041d6200090f565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b158015620004a057600080fd5b505afa158015620004b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620004db57600080fd5b81019080919050509050620004ef6200090f565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324d6147d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200057257600080fd5b505afa15801562000587573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620005ad57600080fd5b81019080919050509050620005c162000931565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ffe237f06040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016104006040518083038186803b1580156200064657600080fd5b505afa1580156200065b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506104008110156200068257600080fd5b810190809190505090506200069662000931565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663483767f06040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016104006040518083038186803b1580156200071b57600080fd5b505afa15801562000730573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506104008110156200075757600080fd5b810190809190505090508460006002811015156200077157fe5b60200201516002600001819055508460016002811015156200078f57fe5b6020020151600260010181905550836000600281101515620007ad57fe5b6020020151600460000181905550836001600281101515620007cb57fe5b6020020151600460010181905550826000600281101515620007e957fe5b60200201516006600001819055508260016002811015156200080757fe5b602002015160066001018190555060008090505b6010811015620009015782816002026020811015156200083757fe5b60200201516008826010811015156200084c57fe5b600202016000018190555082600182600202016020811015156200086c57fe5b60200201516008826010811015156200088157fe5b600202016001018190555081816002026020811015156200089e57fe5b6020020151602882601081101515620008b357fe5b60020201600001819055508160018260020201602081101515620008d357fe5b6020020151602882601081101515620008e857fe5b600202016001018190555080806001019150506200081b565b505050505050505062000955565b6040805190810160405280600290602082028038833980820191505090505090565b61040060405190810160405280602090602082028038833980820191505090505090565b611bc380620009656000396000f3fe6080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631213e201146100a957806317876ecb146100db57806326ff9c68146101325780632e52d606146101895780633e8d3764146101b45780635eacae28146101df578063654474ee14610330578063b7479f5f14610386578063b8c9d365146103dc578063e2179b8e1461040e575b600080fd5b3480156100b557600080fd5b506100be610440565b604051808381526020018281526020019250505060405180910390f35b3480156100e757600080fd5b506100f0610452565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561013e57600080fd5b50610147610477565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561019557600080fd5b5061019e61049d565b6040518082815260200191505060405180910390f35b3480156101c057600080fd5b506101c96104a2565b6040518082815260200191505060405180910390f35b3480156101eb57600080fd5b5061031660048036036103e081101561020357600080fd5b81019080806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f82011690508083019250505050505091929192908060a001906005806020026040519081016040528092919082600560200280828437600081840152601f19601f82011690508083019250505050505091929192908061010001906008806020026040519081016040528092919082600860200280828437600081840152601f19601f82011690508083019250505050505091929192908061010001906008806020026040519081016040528092919082600860200280828437600081840152601f19601f82011690508083019250505050505091929192905050506104a7565b604051808215151515815260200191505060405180910390f35b34801561033c57600080fd5b506103696004803603602081101561035357600080fd5b81019080803590602001909291905050506106bf565b604051808381526020018281526020019250505060405180910390f35b34801561039257600080fd5b506103bf600480360360208110156103a957600080fd5b81019080803590602001909291905050506106e8565b604051808381526020018281526020019250505060405180910390f35b3480156103e857600080fd5b506103f1610711565b604051808381526020018281526020019250505060405180910390f35b34801561041a57600080fd5b50610423610723565b604051808381526020018281526020019250505060405180910390f35b60068060000154908060010154905082565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600481565b601081565b60006104b1611863565b6040805190810160405280876000600a811015156104cb57fe5b60200201518152602001876001600a811015156104e457fe5b602002015181525081600001819052506040805190810160405280876002600a8110151561050e57fe5b60200201518152602001876003600a8110151561052757fe5b602002015181525081602001819052506040805190810160405280876004600a8110151561055157fe5b60200201518152602001876005600a8110151561056a57fe5b602002015181525081604001819052506040805190810160405280876006600a8110151561059457fe5b60200201518152602001876007600a811015156105ad57fe5b602002015181525081606001819052508460006005811015156105cc57fe5b60200201518160800181815250508460016005811015156105e957fe5b60200201518160a001818152505084600260058110151561060657fe5b60200201518160c0018181525050838160e0015160000181905250828160e001516020018190525084600360058110151561063d57fe5b60200201518160e00151604001818152505084600460058110151561065e57fe5b60200201518160e0015160600181815250506106b46040805190810160405280886008600a8110151561068d57fe5b60200201518152602001886009600a811015156106a657fe5b602002015181525082610735565b915050949350505050565b6028816010811015156106ce57fe5b600202016000915090508060000154908060010154905082565b6008816010811015156106f757fe5b600202016000915090508060000154908060010154905082565b60048060000154908060010154905082565b60028060000154908060010154905082565b600061073f6118c7565b610747611136565b90506107516118c7565b6107596111ce565b90506107636118f6565b61078f856000015160000151866000015160200151876020015160000151886020015160200151611266565b8160e00181815250506107a58160e001516112bd565b816101400181815250506107bc8160e001516112f9565b8161010001819052506107da6107d58260e0015161139e565b6112f9565b8161012001819052506107f181610140015161146e565b8161018001818152505061082061081b82610140015183610140015161149a90919063ffffffff16565b6114d6565b8161016001818152505061083460026112f9565b816101a00181905250610869856040015160000151866040015160200151876060015160000151886060015160200151611266565b8160a001818152505061088d8160a001518260a0015161149a90919063ffffffff16565b8160c00181815250506109066108b48260c00151876060015161151090919063ffffffff16565b6108f86108d28460a00151896040015161151090919063ffffffff16565b6108ea8561016001518b61151090919063ffffffff16565b6115bd90919063ffffffff16565b6115bd90919063ffffffff16565b816101e0018190525061099861094e610923836101a0015161165f565b61094084610140015185610160015161149a90919063ffffffff16565b61149a90919063ffffffff16565b61098a61095f84610100015161165f565b61097c8561016001518661014001516116c690919063ffffffff16565b61149a90919063ffffffff16565b6116c690919063ffffffff16565b81610220018181525050610a386109dc8660a0015160046040805190810160405290816000820154815260200160018201548152505061151090919063ffffffff16565b610a2a6109fb84610220015189608001516116c690919063ffffffff16565b60026040805190810160405290816000820154815260200160018201548152505061151090919063ffffffff16565b6115bd90919063ffffffff16565b81610200018190525080610200015160000151816101e0015160000151141580610a74575080610200015160200151816101e001516020015114155b15610a855760009350505050611130565b610a8d611a05565b610abe610aab8360a00151886020015161151090919063ffffffff16565b87600001516115bd90919063ffffffff16565b905060008090505b6004811015610cdd576000610b578860e001516000015183600202600881101515610aed57fe5b60200201518960e001516000015160018560020201600881101515610b0e57fe5b60200201518a60e001516020015185600202600881101515610b2c57fe5b60200201518b60e001516020015160018760020201600881101515610b4d57fe5b6020020151611266565b90506000610b648261139e565b90508185610260015184600481101515610b7a57fe5b6020020181815250508085610280015184600481101515610b9757fe5b60200201818152505060408051908101604052808a60e001516000015185600202600881101515610bc457fe5b602002015181526020018a60e001516000015160018660020201600881101515610bea57fe5b60200201518152508560800181905250610c35610c26610c13848561149a90919063ffffffff16565b876080015161151090919063ffffffff16565b856115bd90919063ffffffff16565b935060408051908101604052808a60e001516020015185600202600881101515610c5b57fe5b602002015181526020018a60e001516020015160018660020201600881101515610c8157fe5b60200201518152508560800181905250610ccc610cbd610caa838461149a90919063ffffffff16565b876080015161151090919063ffffffff16565b856115bd90919063ffffffff16565b935050508080600101915050610ac6565b5060008090505b6010811015610fb65760008090505b6004811015610dd1576000610d0a8383600461170b565b15610d2e5784610260015182600481101515610d2257fe5b60200201519050610d49565b84610280015182600481101515610d4157fe5b602002015190505b6000821415610d745780856102a0015184601081101515610d6657fe5b602002018181525050610dc3565b610da6610da182876102a0015186601081101515610d8e57fe5b602002015161149a90919063ffffffff16565b6114d6565b856102a0015184601081101515610db957fe5b6020020181815250505b508080600101915050610cf3565b50610df2836102a0015182601081101515610de857fe5b602002015161139e565b836102c0015182601081101515610e0557fe5b602002018181525050610e57836101400151610e498960e0015160400151866102a0015185601081101515610e3657fe5b602002015161149a90919063ffffffff16565b61176990919063ffffffff16565b836102a0015182601081101515610e6a57fe5b602002018181525050610ea58760e0015160600151846102c0015183601081101515610e9257fe5b602002015161149a90919063ffffffff16565b836102c0015182601081101515610eb857fe5b602002018181525050610f19610ef3846101a0015183601081101515610eda57fe5b602002015185610160015161149a90919063ffffffff16565b846102c0015183601081101515610f0657fe5b60200201516116c690919063ffffffff16565b836102c0015182601081101515610f2c57fe5b602002018181525050610f8d836101400151610f7f85610120015184601081101515610f5457fe5b6020020151866102c0015185601081101515610f6c57fe5b602002015161149a90919063ffffffff16565b6116c690919063ffffffff16565b836102c0015182601081101515610fa057fe5b6020020181815250508080600101915050610ce4565b506000610ff087608001516040516020018082815260200191505060405160208183030381529060405280519060200120600190046114d6565b90506110f761102c8860c0015160046040805190810160405290816000820154815260200160018201548152505061151090919063ffffffff16565b6110e96110ac61107d61106e8c608001516110608e60e00151606001518f60e001516040015161149a90919063ffffffff16565b6116c690919063ffffffff16565b8661149a90919063ffffffff16565b60066040805190810160405290816000820154815260200160018201548152505061151090919063ffffffff16565b6110db6110be89896102c001516117a5565b6110cd8b8a6102a001516117a5565b6115bd90919063ffffffff16565b6115bd90919063ffffffff16565b6115bd90919063ffffffff16565b8361020001819052508160000151836102000151600001511480156111285750816020015183610200015160200151145b955050505050505b92915050565b61113e6118c7565b6111466118c7565b60008090505b60108110156111c65760088160108110151561116457fe5b6002020160000154828260108110151561117a57fe5b6020020151600001818152505060088160108110151561119657fe5b600202016001015482826010811015156111ac57fe5b60200201516020018181525050808060010191505061114c565b508091505090565b6111d66118c7565b6111de6118c7565b60008090505b601081101561125e576028816010811015156111fc57fe5b6002020160000154828260108110151561121257fe5b6020020151600001818152505060288160108110151561122e57fe5b6002020160010154828260108110151561124457fe5b6020020151602001818152505080806001019150506111e4565b508091505090565b60006112b3858585856040516020018085815260200184815260200183815260200182815260200194505050505060405160208183030381529060405280519060200120600190046114d6565b9050949350505050565b60006112f2826040516020018082815260200191505060405160208183030381529060405280519060200120600190046114d6565b9050919050565b611301611a1f565b600181600060108110151561131257fe5b6020020181815250508181600160108110151561132b57fe5b6020020181815250506000600290505b60108110156113985761137461136f84846001850360108110151561135c57fe5b602002015161149a90919063ffffffff16565b6114d6565b828260108110151561138257fe5b602002018181525050808060010191505061133b565b50919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050600083905060008114156113dd57600092505050611469565b818111156113f45781818115156113f057fe5b0690505b600080600190506000849050600084905060005b60008214151561144157818381151561141d57fe5b04905083848202860383848402860380955081965082975083985050505050611408565b600085121561145e57846000038703975050505050505050611469565b849750505050505050505b919050565b6000817f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001039050919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808015156114ca57fe5b83850991505092915050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381151561150757fe5b06915050919050565b611518611a05565b6001821415611529578290506115b7565b60028214156115435761153c83846115bd565b90506115b7565b61154b611a43565b836000015181600060038110151561155f57fe5b602002018181525050836020015181600160038110151561157c57fe5b6020020181815250508281600260038110151561159557fe5b6020020181815250506040826060836007600019fa15156115b557600080fd5b505b92915050565b6115c5611a05565b6115cd611a66565b83600001518160006004811015156115e157fe5b60200201818152505083602001518160016004811015156115fe57fe5b602002018181525050826000015181600260048110151561161b57fe5b602002018181525050826020015181600360048110151561163857fe5b6020020181815250506040826080836006600019fa151561165857600080fd5b5092915050565b60008082600060108110151561167157fe5b602002015190506000600190505b60108110156116bc576116ad848260108110151561169957fe5b60200201518361176990919063ffffffff16565b9150808060010191505061167f565b5080915050919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050828410156116fe578383820301611702565b8284035b91505092915050565b6000806001830360019060020a02905060008090505b8481101561174557600182908060020a820491505091508080600101915050611721565b50600081861614151561175c576001915050611762565b60009150505b9392505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508080151561179957fe5b83850891505092915050565b6117ad611a05565b6117b5611a05565b6117ee8360006010811015156117c757fe5b60200201518560006010811015156117db57fe5b602002015161151090919063ffffffff16565b90506000600190505b60108110156118585761184961183a858360108110151561181457fe5b6020020151878460108110151561182757fe5b602002015161151090919063ffffffff16565b836115bd90919063ffffffff16565b915080806001019150506117f7565b508091505092915050565b6103a060405190810160405280611878611a89565b8152602001611885611a89565b8152602001611892611a89565b815260200161189f611a89565b81526020016000815260200160008152602001600081526020016118c1611aa3565b81525090565b610400604051908101604052806010905b6118e0611a89565b8152602001906001900390816118d85790505090565b611ca06040519081016040528061190b611ad9565b8152602001611918611ad9565b8152602001611925611afd565b8152602001611932611afd565b815260200161193f611a89565b8152602001600081526020016000815260200160008152602001611961611b21565b815260200161196e611b21565b8152602001600081526020016000815260200160008152602001611990611b21565b815260200161199d611b45565b81526020016119aa611a89565b81526020016119b7611a89565b8152602001600081526020016119cb611b21565b81526020016119d8611b74565b81526020016119e5611b74565b81526020016119f2611b21565b81526020016119ff611b21565b81525090565b604080519081016040528060008152602001600081525090565b61020060405190810160405280601090602082028038833980820191505090505090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b61024060405190810160405280611ab8611afd565b8152602001611ac5611afd565b815260200160008152602001600081525090565b61040060405190810160405280602090602082028038833980820191505090505090565b61010060405190810160405280600890602082028038833980820191505090505090565b61020060405190810160405280601090602082028038833980820191505090505090565b610400604051908101604052806010905b611b5e611a89565b815260200190600190039081611b565790505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a723058203c6a40315a72b2aad26205c7c080fdc5f734bb5a6d1a7295e4fbf13ad7459eb60029`

// DeployRangeproofverifier deploys a new Ethereum contract, binding an instance of Rangeproofverifier to it.
func DeployRangeproofverifier(auth *bind.TransactOpts, backend bind.ContractBackend, params common.Address, ip common.Address) (common.Address, *types.Transaction, *Rangeproofverifier, error) {
	parsed, err := abi.JSON(strings.NewReader(RangeproofverifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RangeproofverifierBin), backend, params, ip)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rangeproofverifier{RangeproofverifierCaller: RangeproofverifierCaller{contract: contract}, RangeproofverifierTransactor: RangeproofverifierTransactor{contract: contract}, RangeproofverifierFilterer: RangeproofverifierFilterer{contract: contract}}, nil
}

// Rangeproofverifier is an auto generated Go binding around an Ethereum contract.
type Rangeproofverifier struct {
	RangeproofverifierCaller     // Read-only binding to the contract
	RangeproofverifierTransactor // Write-only binding to the contract
	RangeproofverifierFilterer   // Log filterer for contract events
}

// RangeproofverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type RangeproofverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RangeproofverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RangeproofverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RangeproofverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RangeproofverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RangeproofverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RangeproofverifierSession struct {
	Contract     *Rangeproofverifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RangeproofverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RangeproofverifierCallerSession struct {
	Contract *RangeproofverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RangeproofverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RangeproofverifierTransactorSession struct {
	Contract     *RangeproofverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RangeproofverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type RangeproofverifierRaw struct {
	Contract *Rangeproofverifier // Generic contract binding to access the raw methods on
}

// RangeproofverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RangeproofverifierCallerRaw struct {
	Contract *RangeproofverifierCaller // Generic read-only contract binding to access the raw methods on
}

// RangeproofverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RangeproofverifierTransactorRaw struct {
	Contract *RangeproofverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRangeproofverifier creates a new instance of Rangeproofverifier, bound to a specific deployed contract.
func NewRangeproofverifier(address common.Address, backend bind.ContractBackend) (*Rangeproofverifier, error) {
	contract, err := bindRangeproofverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rangeproofverifier{RangeproofverifierCaller: RangeproofverifierCaller{contract: contract}, RangeproofverifierTransactor: RangeproofverifierTransactor{contract: contract}, RangeproofverifierFilterer: RangeproofverifierFilterer{contract: contract}}, nil
}

// NewRangeproofverifierCaller creates a new read-only instance of Rangeproofverifier, bound to a specific deployed contract.
func NewRangeproofverifierCaller(address common.Address, caller bind.ContractCaller) (*RangeproofverifierCaller, error) {
	contract, err := bindRangeproofverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RangeproofverifierCaller{contract: contract}, nil
}

// NewRangeproofverifierTransactor creates a new write-only instance of Rangeproofverifier, bound to a specific deployed contract.
func NewRangeproofverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*RangeproofverifierTransactor, error) {
	contract, err := bindRangeproofverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RangeproofverifierTransactor{contract: contract}, nil
}

// NewRangeproofverifierFilterer creates a new log filterer instance of Rangeproofverifier, bound to a specific deployed contract.
func NewRangeproofverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*RangeproofverifierFilterer, error) {
	contract, err := bindRangeproofverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RangeproofverifierFilterer{contract: contract}, nil
}

// bindRangeproofverifier binds a generic wrapper to an already deployed contract.
func bindRangeproofverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RangeproofverifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rangeproofverifier *RangeproofverifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rangeproofverifier.Contract.RangeproofverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rangeproofverifier *RangeproofverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rangeproofverifier.Contract.RangeproofverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rangeproofverifier *RangeproofverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rangeproofverifier.Contract.RangeproofverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rangeproofverifier *RangeproofverifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rangeproofverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rangeproofverifier *RangeproofverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rangeproofverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rangeproofverifier *RangeproofverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rangeproofverifier.Contract.contract.Transact(opts, method, params...)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Rangeproofverifier *RangeproofverifierCaller) BitSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rangeproofverifier.contract.Call(opts, out, "bitSize")
	return *ret0, err
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Rangeproofverifier *RangeproofverifierSession) BitSize() (*big.Int, error) {
	return _Rangeproofverifier.Contract.BitSize(&_Rangeproofverifier.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Rangeproofverifier *RangeproofverifierCallerSession) BitSize() (*big.Int, error) {
	return _Rangeproofverifier.Contract.BitSize(&_Rangeproofverifier.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierCaller) G(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Rangeproofverifier.contract.Call(opts, out, "g")
	return *ret, err
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Rangeproofverifier.Contract.G(&_Rangeproofverifier.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierCallerSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Rangeproofverifier.Contract.G(&_Rangeproofverifier.CallOpts)
}

// GvBase is a free data retrieval call binding the contract method 0xb7479f5f.
//
// Solidity: function gvBase(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierCaller) GvBase(opts *bind.CallOpts, arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Rangeproofverifier.contract.Call(opts, out, "gvBase", arg0)
	return *ret, err
}

// GvBase is a free data retrieval call binding the contract method 0xb7479f5f.
//
// Solidity: function gvBase(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierSession) GvBase(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Rangeproofverifier.Contract.GvBase(&_Rangeproofverifier.CallOpts, arg0)
}

// GvBase is a free data retrieval call binding the contract method 0xb7479f5f.
//
// Solidity: function gvBase(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierCallerSession) GvBase(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Rangeproofverifier.Contract.GvBase(&_Rangeproofverifier.CallOpts, arg0)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierCaller) H(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Rangeproofverifier.contract.Call(opts, out, "h")
	return *ret, err
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Rangeproofverifier.Contract.H(&_Rangeproofverifier.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierCallerSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Rangeproofverifier.Contract.H(&_Rangeproofverifier.CallOpts)
}

// HvBase is a free data retrieval call binding the contract method 0x654474ee.
//
// Solidity: function hvBase(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierCaller) HvBase(opts *bind.CallOpts, arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Rangeproofverifier.contract.Call(opts, out, "hvBase", arg0)
	return *ret, err
}

// HvBase is a free data retrieval call binding the contract method 0x654474ee.
//
// Solidity: function hvBase(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierSession) HvBase(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Rangeproofverifier.Contract.HvBase(&_Rangeproofverifier.CallOpts, arg0)
}

// HvBase is a free data retrieval call binding the contract method 0x654474ee.
//
// Solidity: function hvBase(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierCallerSession) HvBase(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Rangeproofverifier.Contract.HvBase(&_Rangeproofverifier.CallOpts, arg0)
}

// IpVerifier is a free data retrieval call binding the contract method 0x26ff9c68.
//
// Solidity: function ipVerifier() constant returns(address)
func (_Rangeproofverifier *RangeproofverifierCaller) IpVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rangeproofverifier.contract.Call(opts, out, "ipVerifier")
	return *ret0, err
}

// IpVerifier is a free data retrieval call binding the contract method 0x26ff9c68.
//
// Solidity: function ipVerifier() constant returns(address)
func (_Rangeproofverifier *RangeproofverifierSession) IpVerifier() (common.Address, error) {
	return _Rangeproofverifier.Contract.IpVerifier(&_Rangeproofverifier.CallOpts)
}

// IpVerifier is a free data retrieval call binding the contract method 0x26ff9c68.
//
// Solidity: function ipVerifier() constant returns(address)
func (_Rangeproofverifier *RangeproofverifierCallerSession) IpVerifier() (common.Address, error) {
	return _Rangeproofverifier.Contract.IpVerifier(&_Rangeproofverifier.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Rangeproofverifier *RangeproofverifierCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rangeproofverifier.contract.Call(opts, out, "n")
	return *ret0, err
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Rangeproofverifier *RangeproofverifierSession) N() (*big.Int, error) {
	return _Rangeproofverifier.Contract.N(&_Rangeproofverifier.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Rangeproofverifier *RangeproofverifierCallerSession) N() (*big.Int, error) {
	return _Rangeproofverifier.Contract.N(&_Rangeproofverifier.CallOpts)
}

// RpParams is a free data retrieval call binding the contract method 0x17876ecb.
//
// Solidity: function rpParams() constant returns(address)
func (_Rangeproofverifier *RangeproofverifierCaller) RpParams(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rangeproofverifier.contract.Call(opts, out, "rpParams")
	return *ret0, err
}

// RpParams is a free data retrieval call binding the contract method 0x17876ecb.
//
// Solidity: function rpParams() constant returns(address)
func (_Rangeproofverifier *RangeproofverifierSession) RpParams() (common.Address, error) {
	return _Rangeproofverifier.Contract.RpParams(&_Rangeproofverifier.CallOpts)
}

// RpParams is a free data retrieval call binding the contract method 0x17876ecb.
//
// Solidity: function rpParams() constant returns(address)
func (_Rangeproofverifier *RangeproofverifierCallerSession) RpParams() (common.Address, error) {
	return _Rangeproofverifier.Contract.RpParams(&_Rangeproofverifier.CallOpts)
}

// UBase is a free data retrieval call binding the contract method 0x1213e201.
//
// Solidity: function uBase() constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierCaller) UBase(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Rangeproofverifier.contract.Call(opts, out, "uBase")
	return *ret, err
}

// UBase is a free data retrieval call binding the contract method 0x1213e201.
//
// Solidity: function uBase() constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierSession) UBase() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Rangeproofverifier.Contract.UBase(&_Rangeproofverifier.CallOpts)
}

// UBase is a free data retrieval call binding the contract method 0x1213e201.
//
// Solidity: function uBase() constant returns(uint256 X, uint256 Y)
func (_Rangeproofverifier *RangeproofverifierCallerSession) UBase() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Rangeproofverifier.Contract.UBase(&_Rangeproofverifier.CallOpts)
}

// VerifyRangeProof is a free data retrieval call binding the contract method 0x5eacae28.
//
// Solidity: function verifyRangeProof(uint256[10] points, uint256[5] scalar, uint256[8] l, uint256[8] r) constant returns(bool)
func (_Rangeproofverifier *RangeproofverifierCaller) VerifyRangeProof(opts *bind.CallOpts, points [10]*big.Int, scalar [5]*big.Int, l [8]*big.Int, r [8]*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Rangeproofverifier.contract.Call(opts, out, "verifyRangeProof", points, scalar, l, r)
	return *ret0, err
}

// VerifyRangeProof is a free data retrieval call binding the contract method 0x5eacae28.
//
// Solidity: function verifyRangeProof(uint256[10] points, uint256[5] scalar, uint256[8] l, uint256[8] r) constant returns(bool)
func (_Rangeproofverifier *RangeproofverifierSession) VerifyRangeProof(points [10]*big.Int, scalar [5]*big.Int, l [8]*big.Int, r [8]*big.Int) (bool, error) {
	return _Rangeproofverifier.Contract.VerifyRangeProof(&_Rangeproofverifier.CallOpts, points, scalar, l, r)
}

// VerifyRangeProof is a free data retrieval call binding the contract method 0x5eacae28.
//
// Solidity: function verifyRangeProof(uint256[10] points, uint256[5] scalar, uint256[8] l, uint256[8] r) constant returns(bool)
func (_Rangeproofverifier *RangeproofverifierCallerSession) VerifyRangeProof(points [10]*big.Int, scalar [5]*big.Int, l [8]*big.Int, r [8]*big.Int) (bool, error) {
	return _Rangeproofverifier.Contract.VerifyRangeProof(&_Rangeproofverifier.CallOpts, points, scalar, l, r)
}
