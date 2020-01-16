package contracts;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.concurrent.Callable;
import org.web3j.abi.FunctionEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Address;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tuples.generated.Tuple2;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.ContractGasProvider;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 3.6.0.
 */
public class RangeProofVerifier extends Contract {
    private static final String BINARY = "0x60806040523480156200001157600080fd5b506040516040806200275c833981018060405260408110156200003357600080fd5b810190808051906020019092919080519060200190929190505050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da8972246040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200015357600080fd5b505afa15801562000168573d6000803e3d6000fd5b505050506040513d60208110156200017f57600080fd5b8101908080519060200190929190505050602014151562000208576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6269732073697a65206e6f7420657175616c000000000000000000000000000081525060200191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633e9552256040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200028c57600080fd5b505afa158015620002a1573d6000803e3d6000fd5b505050506040513d6020811015620002b857600080fd5b8101908080519060200190929190505050600514151562000341576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6e756d626572206f66206c2c72206e6f7420657175616c00000000000000000081525060200191505060405180910390fd5b6200034b6200090f565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b158015620003ce57600080fd5b505afa158015620003e3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200040957600080fd5b810190809190505090506200041d6200090f565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b158015620004a057600080fd5b505afa158015620004b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620004db57600080fd5b81019080919050509050620004ef6200090f565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324d6147d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200057257600080fd5b505afa15801562000587573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620005ad57600080fd5b81019080919050509050620005c162000931565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ffe237f06040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016108006040518083038186803b1580156200064657600080fd5b505afa1580156200065b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506108008110156200068257600080fd5b810190809190505090506200069662000931565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663483767f06040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016108006040518083038186803b1580156200071b57600080fd5b505afa15801562000730573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506108008110156200075757600080fd5b810190809190505090508460006002811015156200077157fe5b60200201516002600001819055508460016002811015156200078f57fe5b6020020151600260010181905550836000600281101515620007ad57fe5b6020020151600460000181905550836001600281101515620007cb57fe5b6020020151600460010181905550826000600281101515620007e957fe5b60200201516006600001819055508260016002811015156200080757fe5b602002015160066001018190555060008090505b6020811015620009015782816002026040811015156200083757fe5b60200201516008826020811015156200084c57fe5b600202016000018190555082600182600202016040811015156200086c57fe5b60200201516008826020811015156200088157fe5b600202016001018190555081816002026040811015156200089e57fe5b6020020151604882602081101515620008b357fe5b60020201600001819055508160018260020201604081101515620008d357fe5b6020020151604882602081101515620008e857fe5b600202016001018190555080806001019150506200081b565b505050505050505062000955565b6040805190810160405280600290602082028038833980820191505090505090565b61080060405190810160405280604090602082028038833980820191505090505090565b611df780620009656000396000f3fe6080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631213e201146100a957806317876ecb146100db57806326ff9c68146101325780632e52d606146101895780633e8d3764146101b457806341476a5e146101df578063654474ee14610330578063b7479f5f14610386578063b8c9d365146103dc578063e2179b8e1461040e575b600080fd5b3480156100b557600080fd5b506100be610440565b604051808381526020018281526020019250505060405180910390f35b3480156100e757600080fd5b506100f0610452565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561013e57600080fd5b50610147610477565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561019557600080fd5b5061019e61049d565b6040518082815260200191505060405180910390f35b3480156101c057600080fd5b506101c96104a2565b6040518082815260200191505060405180910390f35b3480156101eb57600080fd5b50610316600480360361046081101561020357600080fd5b81019080806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f82011690508083019250505050505091929192908060a001906005806020026040519081016040528092919082600560200280828437600081840152601f19601f8201169050808301925050505050509192919290806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f8201169050808301925050505050509192919290806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f82011690508083019250505050505091929192905050506104a7565b604051808215151515815260200191505060405180910390f35b34801561033c57600080fd5b506103696004803603602081101561035357600080fd5b81019080803590602001909291905050506106bf565b604051808381526020018281526020019250505060405180910390f35b34801561039257600080fd5b506103bf600480360360208110156103a957600080fd5b81019080803590602001909291905050506106e8565b604051808381526020018281526020019250505060405180910390f35b3480156103e857600080fd5b506103f1610711565b604051808381526020018281526020019250505060405180910390f35b34801561041a57600080fd5b50610423610723565b604051808381526020018281526020019250505060405180910390f35b60068060000154908060010154905082565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600581565b602081565b60006104b1611a63565b6040805190810160405280876000600a811015156104cb57fe5b60200201518152602001876001600a811015156104e457fe5b602002015181525081600001819052506040805190810160405280876002600a8110151561050e57fe5b60200201518152602001876003600a8110151561052757fe5b602002015181525081602001819052506040805190810160405280876004600a8110151561055157fe5b60200201518152602001876005600a8110151561056a57fe5b602002015181525081604001819052506040805190810160405280876006600a8110151561059457fe5b60200201518152602001876007600a811015156105ad57fe5b602002015181525081606001819052508460006005811015156105cc57fe5b60200201518160800181815250508460016005811015156105e957fe5b60200201518160a001818152505084600260058110151561060657fe5b60200201518160c0018181525050838160e0015160000181905250828160e001516020018190525084600360058110151561063d57fe5b60200201518160e00151604001818152505084600460058110151561065e57fe5b60200201518160e0015160600181815250506106b46040805190810160405280886008600a8110151561068d57fe5b60200201518152602001886009600a811015156106a657fe5b602002015181525082610735565b915050949350505050565b6048816020811015156106ce57fe5b600202016000915090508060000154908060010154905082565b6008816020811015156106f757fe5b600202016000915090508060000154908060010154905082565b60048060000154908060010154905082565b60028060000154908060010154905082565b600061073f611ac7565b6107476112fc565b9050610751611ac7565b610759611394565b9050610763611af6565b61078f85600001516000015186600001516020015187602001516000015188602001516020015161142c565b8160e00181815250506107a58160e00151611483565b816101400181815250506107bc8160e001516114bf565b8161010001819052506107da6107d58260e00151611564565b6114bf565b8161012001819052506107f1816101400151611634565b8161018001818152505061082061081b82610140015183610140015161166090919063ffffffff16565b61169c565b8161016001818152505061083460026114bf565b816101a0018190525061086985604001516000015186604001516020015187606001516000015188606001516020015161142c565b8160a001818152505061088d8160a001518260a0015161166090919063ffffffff16565b8160c00181815250506109066108b48260c0015187606001516116d690919063ffffffff16565b6108f86108d28460a0015189604001516116d690919063ffffffff16565b6108ea8561016001518b6116d690919063ffffffff16565b61178390919063ffffffff16565b61178390919063ffffffff16565b816101e0018190525061099861094e610923836101a00151611825565b61094084610140015185610160015161166090919063ffffffff16565b61166090919063ffffffff16565b61098a61095f846101000151611825565b61097c85610160015186610140015161188c90919063ffffffff16565b61166090919063ffffffff16565b61188c90919063ffffffff16565b81610220018181525050610a386109dc8660a001516004604080519081016040529081600082015481526020016001820154815250506116d690919063ffffffff16565b610a2a6109fb846102200151896080015161188c90919063ffffffff16565b6002604080519081016040529081600082015481526020016001820154815250506116d690919063ffffffff16565b61178390919063ffffffff16565b81610200018190525080610200015160000151816101e0015160000151141580610a74575080610200015160200151816101e001516020015114155b15610a8557600093505050506112f6565b610a8d611c39565b610abe610aab8360a0015188602001516116d690919063ffffffff16565b876000015161178390919063ffffffff16565b905060008090505b6005811015610d2f576000610b578860e001516000015183600202600a81101515610aed57fe5b60200201518960e001516000015160018560020201600a81101515610b0e57fe5b60200201518a60e001516020015185600202600a81101515610b2c57fe5b60200201518b60e001516020015160018760020201600a81101515610b4d57fe5b602002015161142c565b90508084610260015183600581101515610b6d57fe5b602002018181525050610b91610b8c828361166090919063ffffffff16565b61169c565b846102a0015183600581101515610ba457fe5b602002018181525050610bcd846102a0015183600581101515610bc357fe5b6020020151611564565b846102c0015183600581101515610be057fe5b60200201818152505060408051908101604052808960e001516000015184600202600a81101515610c0d57fe5b602002015181526020018960e001516000015160018560020201600a81101515610c3357fe5b60200201518152508460800181905250610c83610c74856102a0015184600581101515610c5c57fe5b602002015186608001516116d690919063ffffffff16565b8461178390919063ffffffff16565b925060408051908101604052808960e001516020015184600202600a81101515610ca957fe5b602002015181526020018960e001516020015160018560020201600a81101515610ccf57fe5b60200201518152508460800181905250610d1f610d10856102c0015184600581101515610cf857fe5b602002015186608001516116d690919063ffffffff16565b8461178390919063ffffffff16565b9250508080600101915050610ac6565b5060008090505b602081101561117c576000811415610e715760008090505b6005811015610dfb57600084610260015182600581101515610d6c57fe5b602002015190506000821415610d9e578085610300015184602081101515610d9057fe5b602002018181525050610ded565b610dd0610dcb8287610300015186602081101515610db857fe5b602002015161166090919063ffffffff16565b61169c565b85610300015184602081101515610de357fe5b6020020181815250505b508080600101915050610d4e565b5082610300015181602081101515610e0f57fe5b602002015183610340015182602081101515610e2757fe5b602002018181525050610e5083610300015182602081101515610e4657fe5b6020020151611564565b83610300015182602081101515610e6357fe5b602002018181525050610f6c565b6000610e7e8260056118d1565b9050610ed9610ed4856102a0015183600503600581101515610e9c57fe5b6020020151866103000151610eb360018603611921565b8603602081101515610ec157fe5b602002015161166090919063ffffffff16565b61169c565b84610300015183602081101515610eec57fe5b602002018181525050610f4e610f49856102c0015183600503600581101515610f1157fe5b6020020151866103400151610f2860018603611921565b8603602081101515610f3657fe5b602002015161166090919063ffffffff16565b61169c565b84610340015183602081101515610f6157fe5b602002018181525050505b82610300015181602081101515610f7f57fe5b6020020151836102e0015182602081101515610f9757fe5b60200201818152505082610340015181602081101515610fb357fe5b602002015183610320015182602081101515610fcb57fe5b60200201818152505061101d83610140015161100f8960e0015160400151866102e0015185602081101515610ffc57fe5b602002015161166090919063ffffffff16565b61196990919063ffffffff16565b836102e001518260208110151561103057fe5b60200201818152505061106b8760e00151606001518461032001518360208110151561105857fe5b602002015161166090919063ffffffff16565b8361032001518260208110151561107e57fe5b6020020181815250506110df6110b9846101a00151836020811015156110a057fe5b602002015185610160015161166090919063ffffffff16565b846103200151836020811015156110cc57fe5b602002015161188c90919063ffffffff16565b836103200151826020811015156110f257fe5b6020020181815250506111538361014001516111458561012001518460208110151561111a57fe5b60200201518661032001518560208110151561113257fe5b602002015161166090919063ffffffff16565b61188c90919063ffffffff16565b8361032001518260208110151561116657fe5b6020020181815250508080600101915050610d36565b5060006111b6876080015160405160200180828152602001915050604051602081830303815290604052805190602001206001900461169c565b90506112bd6111f28860c001516004604080519081016040529081600082015481526020016001820154815250506116d690919063ffffffff16565b6112af6112726112436112348c608001516112268e60e00151606001518f60e001516040015161166090919063ffffffff16565b61188c90919063ffffffff16565b8661166090919063ffffffff16565b6006604080519081016040529081600082015481526020016001820154815250506116d690919063ffffffff16565b6112a1611284898961032001516119a5565b6112938b8a6102e001516119a5565b61178390919063ffffffff16565b61178390919063ffffffff16565b61178390919063ffffffff16565b8361020001819052508160000151836102000151600001511480156112ee5750816020015183610200015160200151145b955050505050505b92915050565b611304611ac7565b61130c611ac7565b60008090505b602081101561138c5760088160208110151561132a57fe5b6002020160000154828260208110151561134057fe5b6020020151600001818152505060088160208110151561135c57fe5b6002020160010154828260208110151561137257fe5b602002015160200181815250508080600101915050611312565b508091505090565b61139c611ac7565b6113a4611ac7565b60008090505b6020811015611424576048816020811015156113c257fe5b600202016000015482826020811015156113d857fe5b602002015160000181815250506048816020811015156113f457fe5b6002020160010154828260208110151561140a57fe5b6020020151602001818152505080806001019150506113aa565b508091505090565b60006114798585858560405160200180858152602001848152602001838152602001828152602001945050505050604051602081830303815290604052805190602001206001900461169c565b9050949350505050565b60006114b88260405160200180828152602001915050604051602081830303815290604052805190602001206001900461169c565b9050919050565b6114c7611c53565b60018160006020811015156114d857fe5b602002018181525050818160016020811015156114f157fe5b6020020181815250506000600290505b602081101561155e5761153a61153584846001850360208110151561152257fe5b602002015161166090919063ffffffff16565b61169c565b828260208110151561154857fe5b6020020181815250508080600101915050611501565b50919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050600083905060008114156115a35760009250505061162f565b818111156115ba5781818115156115b657fe5b0690505b600080600190506000849050600084905060005b6000821415156116075781838115156115e357fe5b049050838482028603838484028603809550819650829750839850505050506115ce565b60008512156116245784600003870397505050505050505061162f565b849750505050505050505b919050565b6000817f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001039050919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508080151561169057fe5b83850991505092915050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156116cd57fe5b06915050919050565b6116de611c39565b60018214156116ef5782905061177d565b6002821415611709576117028384611783565b905061177d565b611711611c77565b836000015181600060038110151561172557fe5b602002018181525050836020015181600160038110151561174257fe5b6020020181815250508281600260038110151561175b57fe5b6020020181815250506040826060836007600019fa151561177b57600080fd5b505b92915050565b61178b611c39565b611793611c9a565b83600001518160006004811015156117a757fe5b60200201818152505083602001518160016004811015156117c457fe5b60200201818152505082600001518160026004811015156117e157fe5b60200201818152505082602001518160036004811015156117fe57fe5b6020020181815250506040826080836006600019fa151561181e57600080fd5b5092915050565b60008082600060208110151561183757fe5b602002015190506000600190505b602081101561188257611873848260208110151561185f57fe5b60200201518361196990919063ffffffff16565b91508080600101915050611845565b5080915050919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050828410156118c45783838203016118c8565b8284035b91505092915050565b6000808260019060020a02905060005b81851080156118f05750600082115b1561191157600182908060020a8204915050915080806001019150506118e1565b8084600101039250505092915050565b600080829050600081141561193a576001915050611964565b6000600190505b600082111561195e57600281029050818060019003925050611941565b80925050505b919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508080151561199957fe5b83850891505092915050565b6119ad611c39565b6119b5611c39565b6119ee8360006020811015156119c757fe5b60200201518560006020811015156119db57fe5b60200201516116d690919063ffffffff16565b90506000600190505b6020811015611a5857611a49611a3a8583602081101515611a1457fe5b60200201518784602081101515611a2757fe5b60200201516116d690919063ffffffff16565b8361178390919063ffffffff16565b915080806001019150506119f7565b508091505092915050565b61042060405190810160405280611a78611cbd565b8152602001611a85611cbd565b8152602001611a92611cbd565b8152602001611a9f611cbd565b8152602001600081526020016000815260200160008152602001611ac1611cd7565b81525090565b610800604051908101604052806020905b611ae0611cbd565b815260200190600190039081611ad85790505090565b613ea060405190810160405280611b0b611d0d565b8152602001611b18611d0d565b8152602001611b25611d31565b8152602001611b32611d31565b8152602001611b3f611cbd565b8152602001600081526020016000815260200160008152602001611b61611d55565b8152602001611b6e611d55565b8152602001600081526020016000815260200160008152602001611b90611d55565b8152602001611b9d611d79565b8152602001611baa611cbd565b8152602001611bb7611cbd565b815260200160008152602001611bcb611d55565b8152602001611bd8611da8565b8152602001611be5611da8565b8152602001611bf2611da8565b8152602001611bff611da8565b8152602001611c0c611d55565b8152602001611c19611d55565b8152602001611c26611d55565b8152602001611c33611d55565b81525090565b604080519081016040528060008152602001600081525090565b61040060405190810160405280602090602082028038833980820191505090505090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b6102c060405190810160405280611cec611d31565b8152602001611cf9611d31565b815260200160008152602001600081525090565b61080060405190810160405280604090602082028038833980820191505090505090565b61014060405190810160405280600a90602082028038833980820191505090505090565b61040060405190810160405280602090602082028038833980820191505090505090565b610800604051908101604052806020905b611d92611cbd565b815260200190600190039081611d8a5790505090565b60a06040519081016040528060059060208202803883398082019150509050509056fea165627a7a7230582084843188e67109bc6fccb7d56cad10597457f65b5fb479495a8dab1e8530e0ec0029";

    public static final String FUNC_UBASE = "uBase";

    public static final String FUNC_RPPARAMS = "rpParams";

    public static final String FUNC_IPVERIFIER = "ipVerifier";

    public static final String FUNC_N = "n";

    public static final String FUNC_BITSIZE = "bitSize";

    public static final String FUNC_HVBASE = "hvBase";

    public static final String FUNC_GVBASE = "gvBase";

    public static final String FUNC_H = "h";

    public static final String FUNC_G = "g";

    public static final String FUNC_VERIFYRANGEPROOF = "verifyRangeProof";

    @Deprecated
    protected RangeProofVerifier(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected RangeProofVerifier(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected RangeProofVerifier(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected RangeProofVerifier(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> uBase() {
        final Function function = new Function(FUNC_UBASE, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple2<BigInteger, BigInteger>>(
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(), 
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteCall<String> rpParams() {
        final Function function = new Function(FUNC_RPPARAMS, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<String> ipVerifier() {
        final Function function = new Function(FUNC_IPVERIFIER, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<BigInteger> n() {
        final Function function = new Function(FUNC_N, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<BigInteger> bitSize() {
        final Function function = new Function(FUNC_BITSIZE, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> hvBase(BigInteger param0) {
        final Function function = new Function(FUNC_HVBASE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(param0)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple2<BigInteger, BigInteger>>(
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(), 
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> gvBase(BigInteger param0) {
        final Function function = new Function(FUNC_GVBASE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(param0)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple2<BigInteger, BigInteger>>(
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(), 
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> h() {
        final Function function = new Function(FUNC_H, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple2<BigInteger, BigInteger>>(
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(), 
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> g() {
        final Function function = new Function(FUNC_G, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple2<BigInteger, BigInteger>>(
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(), 
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public static RemoteCall<RangeProofVerifier> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider, String params, String ip) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params), 
                new org.web3j.abi.datatypes.Address(ip)));
        return deployRemoteCall(RangeProofVerifier.class, web3j, credentials, contractGasProvider, BINARY, encodedConstructor);
    }

    public static RemoteCall<RangeProofVerifier> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider, String params, String ip) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params), 
                new org.web3j.abi.datatypes.Address(ip)));
        return deployRemoteCall(RangeProofVerifier.class, web3j, transactionManager, contractGasProvider, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<RangeProofVerifier> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String params, String ip) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params), 
                new org.web3j.abi.datatypes.Address(ip)));
        return deployRemoteCall(RangeProofVerifier.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<RangeProofVerifier> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String params, String ip) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params), 
                new org.web3j.abi.datatypes.Address(ip)));
        return deployRemoteCall(RangeProofVerifier.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public RemoteCall<TransactionReceipt> verifyRangeProof(List<BigInteger> points, List<BigInteger> scalar, List<BigInteger> l, List<BigInteger> r) {
        final Function function = new Function(
                FUNC_VERIFYRANGEPROOF, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.StaticArray10<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(points, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.StaticArray5<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(scalar, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.StaticArray10<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(l, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.StaticArray10<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(r, org.web3j.abi.datatypes.generated.Uint256.class))), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    @Deprecated
    public static RangeProofVerifier load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new RangeProofVerifier(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static RangeProofVerifier load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new RangeProofVerifier(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static RangeProofVerifier load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new RangeProofVerifier(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static RangeProofVerifier load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new RangeProofVerifier(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
