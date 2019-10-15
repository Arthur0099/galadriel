package contracts;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.Callable;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.StaticArray;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.generated.StaticArray2;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
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
public class PublicParams extends Contract {
    private static final String BINARY = "0x60806040523480156200001157600080fd5b5060016004600001819055506002600460010181905550620000776040805190810160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c6700000000000081525062000133640100000000026401000000009004565b600080820151816000015560208201518160010155905050620000de6040805190810160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c6700000000000081525062000133640100000000026401000000009004565b6002600082015181600001556020820151816001015590505060008001546006600001819055506000600101546006600101819055506200012d62000224640100000000026401000000009004565b620005ec565b6200013d6200058c565b6000620001d7836040516020018082805190602001908083835b6020831015156200017e578051825260208201915060208101905060208303925062000157565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060019004620003dc64010000000002620007c5176401000000009004565b90506200021c816004604080519081016040529081600082015481526020016001820154815250506200041764010000000002620007ff179091906401000000009004565b915050919050565b600080600090505b6020811015620002fd576200027f81604051602001808281526020019150506040516020818303038152906040528051906020012060019004620003dc64010000000002620007c5176401000000009004565b9150620002c4826004604080519081016040529081600082015481526020016001820154815250506200041764010000000002620007ff179091906401000000009004565b600882602081101515620002d457fe5b60020201600082015181600001556020820151816001015590505080806001019150506200022c565b5060008090505b6020811015620003d8576200035a60208201604051602001808281526020019150506040516020818303038152906040528051906020012060019004620003dc64010000000002620007c5176401000000009004565b91506200039f826004604080519081016040529081600082015481526020016001820154815250506200041764010000000002620007ff179091906401000000009004565b604882602081101515620003af57fe5b600202016000820151816000015560208201518160010155905050808060010191505062000304565b5050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156200040e57fe5b06915050919050565b620004216200058c565b60018214156200043457829050620004db565b60028214156200046157620004598384620004e1640100000000026401000000009004565b9050620004db565b6200046b620005a6565b83600001518160006003811015156200048057fe5b60200201818152505083602001518160016003811015156200049e57fe5b60200201818152505082816002600381101515620004b857fe5b6020020181815250506040826060836007600019fa1515620004d957600080fd5b505b92915050565b620004eb6200058c565b620004f5620005c9565b83600001518160006004811015156200050a57fe5b60200201818152505083602001518160016004811015156200052857fe5b60200201818152505082600001518160026004811015156200054657fe5b60200201818152505082602001518160036004811015156200056457fe5b6020020181815250506040826080836006600019fa15156200058557600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b6109da80620005fc6000396000f3fe6080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304c09ce9146100e05780630c00f8a01461013357806324d6147d146101895780632e52d606146101dc5780633e8d3764146102075780633e95522514610232578063483767f01461025d57806352c9b965146102b057806356e2736c146102e257806382529fdb14610338578063b8c9d3651461038b578063c6a898c5146103bd578063da897224146103ef578063e2179b8e1461041a578063ffe237f01461044c575b600080fd5b3480156100ec57600080fd5b506100f561049f565b6040518082600260200280838360005b83811015610120578082015181840152602081019050610105565b5050505090500191505060405180910390f35b34801561013f57600080fd5b5061016c6004803603602081101561015657600080fd5b81019080803590602001909291905050506104f1565b604051808381526020018281526020019250505060405180910390f35b34801561019557600080fd5b5061019e61051a565b6040518082600260200280838360005b838110156101c95780820151818401526020810190506101ae565b5050505090500191505060405180910390f35b3480156101e857600080fd5b506101f161056d565b6040518082815260200191505060405180910390f35b34801561021357600080fd5b5061021c610572565b6040518082815260200191505060405180910390f35b34801561023e57600080fd5b50610247610577565b6040518082815260200191505060405180910390f35b34801561026957600080fd5b50610272610580565b6040518082604060200280838360005b8381101561029d578082015181840152602081019050610282565b5050505090500191505060405180910390f35b3480156102bc57600080fd5b506102c5610619565b604051808381526020018281526020019250505060405180910390f35b3480156102ee57600080fd5b5061031b6004803603602081101561030557600080fd5b810190808035906020019092919050505061062b565b604051808381526020018281526020019250505060405180910390f35b34801561034457600080fd5b5061034d610654565b6040518082600260200280838360005b8381101561037857808201518184015260208101905061035d565b5050505090500191505060405180910390f35b34801561039757600080fd5b506103a06106a7565b604051808381526020018281526020019250505060405180910390f35b3480156103c957600080fd5b506103d26106b9565b604051808381526020018281526020019250505060405180910390f35b3480156103fb57600080fd5b506104046106cb565b6040518082815260200191505060405180910390f35b34801561042657600080fd5b5061042f6106d4565b604051808381526020018281526020019250505060405180910390f35b34801561045857600080fd5b506104616106e6565b6040518082604060200280838360005b8381101561048c578082015181840152602081019050610471565b5050505090500191505060405180910390f35b6104a761077f565b6104af61077f565b60008001548160006002811015156104c357fe5b6020020181815250506000600101548160016002811015156104e157fe5b6020020181815250508091505090565b60088160208110151561050057fe5b600202016000915090508060000154908060010154905082565b61052261077f565b61052a61077f565b60066000015481600060028110151561053f57fe5b60200201818152505060066001015481600160028110151561055d57fe5b6020020181815250508091505090565b600581565b602081565b60006005905090565b6105886107a1565b6105906107a1565b60008090505b6020811015610611576048816020811015156105ae57fe5b600202016000015482826002026040811015156105c757fe5b6020020181815250506048816020811015156105df57fe5b600202016001015482600183600202016040811015156105fb57fe5b6020020181815250508080600101915050610596565b508091505090565b60048060000154908060010154905082565b60488160208110151561063a57fe5b600202016000915090508060000154908060010154905082565b61065c61077f565b61066461077f565b60026000015481600060028110151561067957fe5b60200201818152505060026001015481600160028110151561069757fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b60006020905090565b60008060000154908060010154905082565b6106ee6107a1565b6106f66107a1565b60008090505b60208110156107775760088160208110151561071457fe5b6002020160000154828260020260408110151561072d57fe5b60200201818152505060088160208110151561074557fe5b6002020160010154826001836002020160408110151561076157fe5b60200201818152505080806001019150506106fc565b508091505090565b6040805190810160405280600290602082028038833980820191505090505090565b61080060405190810160405280604090602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156107f657fe5b06915050919050565b61080761094e565b6001821415610818578290506108a6565b60028214156108325761082b83846108ac565b90506108a6565b61083a610968565b836000015181600060038110151561084e57fe5b602002018181525050836020015181600160038110151561086b57fe5b6020020181815250508281600260038110151561088457fe5b6020020181815250506040826060836007600019fa15156108a457600080fd5b505b92915050565b6108b461094e565b6108bc61098b565b83600001518160006004811015156108d057fe5b60200201818152505083602001518160016004811015156108ed57fe5b602002018181525050826000015181600260048110151561090a57fe5b602002018181525050826020015181600360048110151561092757fe5b6020020181815250506040826080836006600019fa151561094757600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a72305820222c6c0a9ef08751befa4f66866ca191afe142a8a057ecca044954bc0151372e0029";

    public static final String FUNC_GVECTOR = "gVector";

    public static final String FUNC_N = "n";

    public static final String FUNC_BITSIZE = "bitSize";

    public static final String FUNC_GBASE = "gBase";

    public static final String FUNC_HVECTOR = "hVector";

    public static final String FUNC_H = "h";

    public static final String FUNC_U = "u";

    public static final String FUNC_G = "g";

    public static final String FUNC_GETG = "getG";

    public static final String FUNC_GETH = "getH";

    public static final String FUNC_GETU = "getU";

    public static final String FUNC_GETGVECTOR = "getGVector";

    public static final String FUNC_GETHVECTOR = "getHVector";

    public static final String FUNC_GETBITSIZE = "getBitSize";

    public static final String FUNC_GETN = "getN";

    @Deprecated
    protected PublicParams(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected PublicParams(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected PublicParams(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected PublicParams(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> gVector(BigInteger param0) {
        final Function function = new Function(FUNC_GVECTOR, 
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

    public RemoteCall<Tuple2<BigInteger, BigInteger>> gBase() {
        final Function function = new Function(FUNC_GBASE, 
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

    public RemoteCall<Tuple2<BigInteger, BigInteger>> hVector(BigInteger param0) {
        final Function function = new Function(FUNC_HVECTOR, 
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

    public RemoteCall<Tuple2<BigInteger, BigInteger>> u() {
        final Function function = new Function(FUNC_U, 
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

    public static RemoteCall<PublicParams> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return deployRemoteCall(PublicParams.class, web3j, credentials, contractGasProvider, BINARY, "");
    }

    public static RemoteCall<PublicParams> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return deployRemoteCall(PublicParams.class, web3j, transactionManager, contractGasProvider, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<PublicParams> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(PublicParams.class, web3j, credentials, gasPrice, gasLimit, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<PublicParams> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(PublicParams.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, "");
    }

    public RemoteCall<List> getG() {
        final Function function = new Function(FUNC_GETG, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<StaticArray2<Uint256>>() {}));
        return new RemoteCall<List>(
                new Callable<List>() {
                    @Override
                    @SuppressWarnings("unchecked")
                    public List call() throws Exception {
                        List<Type> result = (List<Type>) executeCallSingleValueReturn(function, List.class);
                        return convertToNative(result);
                    }
                });
    }

    public RemoteCall<List> getH() {
        final Function function = new Function(FUNC_GETH, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<StaticArray2<Uint256>>() {}));
        return new RemoteCall<List>(
                new Callable<List>() {
                    @Override
                    @SuppressWarnings("unchecked")
                    public List call() throws Exception {
                        List<Type> result = (List<Type>) executeCallSingleValueReturn(function, List.class);
                        return convertToNative(result);
                    }
                });
    }

    public RemoteCall<List> getU() {
        final Function function = new Function(FUNC_GETU, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<StaticArray2<Uint256>>() {}));
        return new RemoteCall<List>(
                new Callable<List>() {
                    @Override
                    @SuppressWarnings("unchecked")
                    public List call() throws Exception {
                        List<Type> result = (List<Type>) executeCallSingleValueReturn(function, List.class);
                        return convertToNative(result);
                    }
                });
    }

    public RemoteCall<List> getGVector() {
        final Function function = new Function(FUNC_GETGVECTOR, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<StaticArray<Uint256>>() {}));
        return new RemoteCall<List>(
                new Callable<List>() {
                    @Override
                    @SuppressWarnings("unchecked")
                    public List call() throws Exception {
                        List<Type> result = (List<Type>) executeCallSingleValueReturn(function, List.class);
                        return convertToNative(result);
                    }
                });
    }

    public RemoteCall<List> getHVector() {
        final Function function = new Function(FUNC_GETHVECTOR, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<StaticArray<Uint256>>() {}));
        return new RemoteCall<List>(
                new Callable<List>() {
                    @Override
                    @SuppressWarnings("unchecked")
                    public List call() throws Exception {
                        List<Type> result = (List<Type>) executeCallSingleValueReturn(function, List.class);
                        return convertToNative(result);
                    }
                });
    }

    public RemoteCall<BigInteger> getBitSize() {
        final Function function = new Function(FUNC_GETBITSIZE, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<BigInteger> getN() {
        final Function function = new Function(FUNC_GETN, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    @Deprecated
    public static PublicParams load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new PublicParams(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static PublicParams load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new PublicParams(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static PublicParams load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new PublicParams(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static PublicParams load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new PublicParams(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
