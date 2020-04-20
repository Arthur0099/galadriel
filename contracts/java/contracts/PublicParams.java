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
    private static final String BINARY = "0x60806040523480156200001157600080fd5b50600160046000018190555060026004600101819055506200006e6040518060400160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c670000000000008152506200015a60201b60201c565b600080820151816000015560208201518160010155905050620000cc6040518060400160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c670000000000008152506200015a60201b60201c565b600260008201518160000155602082015181600101559050506200012b6040518060400160405280601b81526020017f752067656e657261746f72206f6620696e6e657270726f6475637400000000008152506200015a60201b60201c565b60066000820151816000015560208201518160010155905050620001546200023660201b60201c565b62000684565b6200016462000626565b6000620001f2836040516020018082805190602001908083835b60208310620001a357805182526020820191506020810190506020830392506200017e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060001c6200049360201b620009221760201c565b90506200022e81600460405180604001604052908160008201548152602001600182015481525050620004cc60201b6200095a1790919060201c565b915050919050565b6000806200029860405160200180807f677673000000000000000000000000000000000000000000000000000000000081525060030190506040516020818303038152906040528051906020012060001c6200049360201b620009221760201c565b905060008090505b60026040028110156200036157620002ee818301604051602001808281526020019150506040516020818303038152906040528051906020012060001c6200049360201b620009221760201c565b92506200032a83600460405180604001604052908160008201548152602001600182015481525050620004cc60201b6200095a1790919060201c565b600882608081106200033857fe5b6002020160008201518160000155602082015181600101559050508080600101915050620002a0565b506000620003c360405160200180807f687673000000000000000000000000000000000000000000000000000000000081525060030190506040516020818303038152906040528051906020012060001c6200049360201b620009221760201c565b905060008090505b60026040028110156200048d5762000419828201604051602001808281526020019150506040516020818303038152906040528051906020012060001c6200049360201b620009221760201c565b93506200045584600460405180604001604052908160008201548152602001600182015481525050620004cc60201b6200095a1790919060201c565b61010882608081106200046457fe5b6002020160008201518160000155602082015181600101559050508080600101915050620003cb565b50505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381620004c357fe5b06915050919050565b620004d662000626565b6001821415620004e9578290506200057f565b60028214156200050d576200050583846200058560201b60201c565b90506200057f565b6200051762000640565b8360000151816000600381106200052a57fe5b6020020181815250508360200151816001600381106200054657fe5b60200201818152505082816002600381106200055e57fe5b6020020181815250506040826060836007600019fa6200057d57600080fd5b505b92915050565b6200058f62000626565b6200059962000662565b836000015181600060048110620005ac57fe5b602002018181525050836020015181600160048110620005c857fe5b602002018181525050826000015181600260048110620005e457fe5b6020020181815250508260200151816003600481106200060057fe5b6020020181815250506040826080836006600019fa6200061f57600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b610b2a80620006946000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806356e2736c116100a2578063c6a898c511610071578063c6a898c5146103ce578063da2b99ce146103f3578063da8972241461043a578063e2179b8e14610458578063ffe237f01461047d57610116565b806356e2736c146102fc5780637982ebcc1461034557806382529fdb14610363578063b8c9d365146103a957610116565b80633e8d3764116100e95780633e8d37641461020e5780633e9552251461022c578063483767f01461024a5780634db118751461029057806352c9b965146102d757610116565b806304c09ce91461011b5780630c00f8a01461016157806324d6147d146101aa5780632e52d606146101f0575b600080fd5b6101236104c3565b6040518082600260200280838360005b8381101561014e578082015181840152602081019050610133565b5050505090500191505060405180910390f35b61018d6004803603602081101561017757600080fd5b8101908080359060200190929190505050610511565b604051808381526020018281526020019250505060405180910390f35b6101b2610538565b6040518082600260200280838360005b838110156101dd5780820151818401526020810190506101c2565b5050505090500191505060405180910390f35b6101f8610587565b6040518082815260200191505060405180910390f35b61021661058c565b6040518082815260200191505060405180910390f35b610234610591565b6040518082815260200191505060405180910390f35b61025261059a565b6040518082608060200280838360005b8381101561027d578082015181840152602081019050610262565b5050505090500191505060405180910390f35b61029861062d565b604051808261010060200280838360005b838110156102c45780820151818401526020810190506102a9565b5050505090500191505060405180910390f35b6102df6106c3565b604051808381526020018281526020019250505060405180910390f35b6103286004803603602081101561031257600080fd5b81019080803590602001909291905050506106d5565b604051808381526020018281526020019250505060405180910390f35b61034d6106fd565b6040518082815260200191505060405180910390f35b61036b610702565b6040518082600260200280838360005b8381101561039657808201518184015260208101905061037b565b5050505090500191505060405180910390f35b6103b1610751565b604051808381526020018281526020019250505060405180910390f35b6103d6610763565b604051808381526020018281526020019250505060405180910390f35b6103fb610775565b604051808261010060200280838360005b8381101561042757808201518184015260208101905061040c565b5050505090500191505060405180910390f35b61044261080d565b6040518082815260200191505060405180910390f35b610460610816565b604051808381526020018281526020019250505060405180910390f35b610485610828565b6040518082608060200280838360005b838110156104b0578082015181840152602081019050610495565b5050505090500191505060405180910390f35b6104cb6108b9565b6104d36108b9565b6000800154816000600281106104e557fe5b6020020181815250506000600101548160016002811061050157fe5b6020020181815250508091505090565b6008816080811061051e57fe5b600202016000915090508060000154908060010154905082565b6105406108b9565b6105486108b9565b6006600001548160006002811061055b57fe5b6020020181815250506006600101548160016002811061057757fe5b6020020181815250508091505090565b600681565b604081565b60006006905090565b6105a26108db565b6105aa6108db565b60008090505b60408110156106255761010881608081106105c757fe5b60020201600001548282600202608081106105de57fe5b60200201818152505061010881608081106105f557fe5b600202016001015482600183600202016080811061060f57fe5b60200201818152505080806001019150506105b0565b508091505090565b6106356108fe565b61063d6108fe565b60008090505b60026040028110156106bb576008816080811061065c57fe5b60020201600001548282600202610100811061067457fe5b6020020181815250506008816080811061068a57fe5b6002020160010154826001836002020161010081106106a557fe5b6020020181815250508080600101915050610643565b508091505090565b60048060000154908060010154905082565b61010881608081106106e357fe5b600202016000915090508060000154908060010154905082565b600281565b61070a6108b9565b6107126108b9565b6002600001548160006002811061072557fe5b6020020181815250506002600101548160016002811061074157fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b61077d6108fe565b6107856108fe565b60008090505b60026040028110156108055761010881608081106107a557fe5b6002020160000154828260020261010081106107bd57fe5b60200201818152505061010881608081106107d457fe5b6002020160010154826001836002020161010081106107ef57fe5b602002018181525050808060010191505061078b565b508091505090565b60006040905090565b60008060000154908060010154905082565b6108306108db565b6108386108db565b60008090505b60408110156108b1576008816080811061085457fe5b600202016000015482826002026080811061086b57fe5b6020020181815250506008816080811061088157fe5b600202016001015482600183600202016080811061089b57fe5b602002018181525050808060010191505061083e565b508091505090565b6040518060400160405280600290602082028038833980820191505090505090565b604051806110000160405280608090602082028038833980820191505090505090565b60405180612000016040528061010090602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161095157fe5b06915050919050565b610962610a97565b6001821415610973578290506109f9565b600282141561098d5761098683846109ff565b90506109f9565b610995610ab1565b8360000151816000600381106109a757fe5b6020020181815250508360200151816001600381106109c257fe5b60200201818152505082816002600381106109d957fe5b6020020181815250506040826060836007600019fa6109f757600080fd5b505b92915050565b610a07610a97565b610a0f610ad3565b836000015181600060048110610a2157fe5b602002018181525050836020015181600160048110610a3c57fe5b602002018181525050826000015181600260048110610a5757fe5b602002018181525050826020015181600360048110610a7257fe5b6020020181815250506040826080836006600019fa610a9057600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b604051806080016040528060049060208202803883398082019150509050509056fea265627a7a723158208d873b904394af55d1e70458846e343074672366748dc8d0d1f22405d761ef7164736f6c63430005100032";

    public static final String FUNC_AGGSIZE = "aggSize";

    public static final String FUNC_BITSIZE = "bitSize";

    public static final String FUNC_G = "g";

    public static final String FUNC_GBASE = "gBase";

    public static final String FUNC_GVECTOR = "gVector";

    public static final String FUNC_H = "h";

    public static final String FUNC_HVECTOR = "hVector";

    public static final String FUNC_N = "n";

    public static final String FUNC_U = "u";

    public static final String FUNC_GETG = "getG";

    public static final String FUNC_GETH = "getH";

    public static final String FUNC_GETU = "getU";

    public static final String FUNC_GETAGGGVECTOR = "getAggGVector";

    public static final String FUNC_GETAGGHVECTOR = "getAggHVector";

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

    public RemoteCall<BigInteger> aggSize() {
        final Function function = new Function(FUNC_AGGSIZE, 
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

    public RemoteCall<BigInteger> n() {
        final Function function = new Function(FUNC_N, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
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

    public RemoteCall<List> getAggGVector() {
        final Function function = new Function(FUNC_GETAGGGVECTOR, 
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

    public RemoteCall<List> getAggHVector() {
        final Function function = new Function(FUNC_GETAGGHVECTOR, 
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
