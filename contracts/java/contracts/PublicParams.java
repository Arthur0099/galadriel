package contracts;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.Callable;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.generated.StaticArray2;
import org.web3j.abi.datatypes.generated.StaticArray32;
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
    private static final String BINARY = "0x60806040523480156200001157600080fd5b5060016004600001819055506002600460010181905550620000776040805190810160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c6700000000000081525062000133640100000000026401000000009004565b600080820151816000015560208201518160010155905050620000de6040805190810160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c6700000000000081525062000133640100000000026401000000009004565b6002600082015181600001556020820151816001015590505060008001546006600001819055506000600101546006600101819055506200012d62000224640100000000026401000000009004565b620005ec565b6200013d6200058c565b6000620001d7836040516020018082805190602001908083835b6020831015156200017e578051825260208201915060208101905060208303925062000157565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060019004620003dc64010000000002620007c3176401000000009004565b90506200021c816004604080519081016040529081600082015481526020016001820154815250506200041764010000000002620007fd179091906401000000009004565b915050919050565b600080600090505b6010811015620002fd576200027f81604051602001808281526020019150506040516020818303038152906040528051906020012060019004620003dc64010000000002620007c3176401000000009004565b9150620002c4826004604080519081016040529081600082015481526020016001820154815250506200041764010000000002620007fd179091906401000000009004565b600882601081101515620002d457fe5b60020201600082015181600001556020820151816001015590505080806001019150506200022c565b5060008090505b6010811015620003d8576200035a60108201604051602001808281526020019150506040516020818303038152906040528051906020012060019004620003dc64010000000002620007c3176401000000009004565b91506200039f826004604080519081016040529081600082015481526020016001820154815250506200041764010000000002620007fd179091906401000000009004565b602882601081101515620003af57fe5b600202016000820151816000015560208201518160010155905050808060010191505062000304565b5050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156200040e57fe5b06915050919050565b620004216200058c565b60018214156200043457829050620004db565b60028214156200046157620004598384620004e1640100000000026401000000009004565b9050620004db565b6200046b620005a6565b83600001518160006003811015156200048057fe5b60200201818152505083602001518160016003811015156200049e57fe5b60200201818152505082816002600381101515620004b857fe5b6020020181815250506040826060836007600019fa1515620004d957600080fd5b505b92915050565b620004eb6200058c565b620004f5620005c9565b83600001518160006004811015156200050a57fe5b60200201818152505083602001518160016004811015156200052857fe5b60200201818152505082600001518160026004811015156200054657fe5b60200201818152505082602001518160036004811015156200056457fe5b6020020181815250506040826080836006600019fa15156200058557600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b6109d880620005fc6000396000f3fe6080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304c09ce9146100e05780630c00f8a01461013357806324d6147d146101895780632e52d606146101dc5780633e8d3764146102075780633e95522514610232578063483767f01461025d57806352c9b965146102af57806356e2736c146102e157806382529fdb14610337578063b8c9d3651461038a578063c6a898c5146103bc578063da897224146103ee578063e2179b8e14610419578063ffe237f01461044b575b600080fd5b3480156100ec57600080fd5b506100f561049d565b6040518082600260200280838360005b83811015610120578082015181840152602081019050610105565b5050505090500191505060405180910390f35b34801561013f57600080fd5b5061016c6004803603602081101561015657600080fd5b81019080803590602001909291905050506104ef565b604051808381526020018281526020019250505060405180910390f35b34801561019557600080fd5b5061019e610518565b6040518082600260200280838360005b838110156101c95780820151818401526020810190506101ae565b5050505090500191505060405180910390f35b3480156101e857600080fd5b506101f161056b565b6040518082815260200191505060405180910390f35b34801561021357600080fd5b5061021c610570565b6040518082815260200191505060405180910390f35b34801561023e57600080fd5b50610247610575565b6040518082815260200191505060405180910390f35b34801561026957600080fd5b5061027261057e565b60405180826020800280838360005b8381101561029c578082015181840152602081019050610281565b5050505090500191505060405180910390f35b3480156102bb57600080fd5b506102c4610617565b604051808381526020018281526020019250505060405180910390f35b3480156102ed57600080fd5b5061031a6004803603602081101561030457600080fd5b8101908080359060200190929190505050610629565b604051808381526020018281526020019250505060405180910390f35b34801561034357600080fd5b5061034c610652565b6040518082600260200280838360005b8381101561037757808201518184015260208101905061035c565b5050505090500191505060405180910390f35b34801561039657600080fd5b5061039f6106a5565b604051808381526020018281526020019250505060405180910390f35b3480156103c857600080fd5b506103d16106b7565b604051808381526020018281526020019250505060405180910390f35b3480156103fa57600080fd5b506104036106c9565b6040518082815260200191505060405180910390f35b34801561042557600080fd5b5061042e6106d2565b604051808381526020018281526020019250505060405180910390f35b34801561045757600080fd5b506104606106e4565b60405180826020800280838360005b8381101561048a57808201518184015260208101905061046f565b5050505090500191505060405180910390f35b6104a561077d565b6104ad61077d565b60008001548160006002811015156104c157fe5b6020020181815250506000600101548160016002811015156104df57fe5b6020020181815250508091505090565b6008816010811015156104fe57fe5b600202016000915090508060000154908060010154905082565b61052061077d565b61052861077d565b60066000015481600060028110151561053d57fe5b60200201818152505060066001015481600160028110151561055b57fe5b6020020181815250508091505090565b600481565b601081565b60006004905090565b61058661079f565b61058e61079f565b60008090505b601081101561060f576028816010811015156105ac57fe5b600202016000015482826002026020811015156105c557fe5b6020020181815250506028816010811015156105dd57fe5b600202016001015482600183600202016020811015156105f957fe5b6020020181815250508080600101915050610594565b508091505090565b60048060000154908060010154905082565b60288160108110151561063857fe5b600202016000915090508060000154908060010154905082565b61065a61077d565b61066261077d565b60026000015481600060028110151561067757fe5b60200201818152505060026001015481600160028110151561069557fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b60006010905090565b60008060000154908060010154905082565b6106ec61079f565b6106f461079f565b60008090505b60108110156107755760088160108110151561071257fe5b6002020160000154828260020260208110151561072b57fe5b60200201818152505060088160108110151561074357fe5b6002020160010154826001836002020160208110151561075f57fe5b60200201818152505080806001019150506106fa565b508091505090565b6040805190810160405280600290602082028038833980820191505090505090565b61040060405190810160405280602090602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156107f457fe5b06915050919050565b61080561094c565b6001821415610816578290506108a4565b60028214156108305761082983846108aa565b90506108a4565b610838610966565b836000015181600060038110151561084c57fe5b602002018181525050836020015181600160038110151561086957fe5b6020020181815250508281600260038110151561088257fe5b6020020181815250506040826060836007600019fa15156108a257600080fd5b505b92915050565b6108b261094c565b6108ba610989565b83600001518160006004811015156108ce57fe5b60200201818152505083602001518160016004811015156108eb57fe5b602002018181525050826000015181600260048110151561090857fe5b602002018181525050826020015181600360048110151561092557fe5b6020020181815250506040826080836006600019fa151561094557600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a7230582000b62347f7f9cb1462a11b6e8b1f782169c22c4b031fa0415cba7e1aeb87e41f0029";

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
                Arrays.<TypeReference<?>>asList(new TypeReference<StaticArray32<Uint256>>() {}));
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
                Arrays.<TypeReference<?>>asList(new TypeReference<StaticArray32<Uint256>>() {}));
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
