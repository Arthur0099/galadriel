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
    private static final String BINARY = "0x60806040523480156200001157600080fd5b5060016004600001819055506002600460010181905550620000776040805190810160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c670000000000008152506200017e640100000000026401000000009004565b600080820151816000015560208201518160010155905050620000de6040805190810160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c670000000000008152506200017e640100000000026401000000009004565b60026000820151816000015560208201518160010155905050620001466040805190810160405280601b81526020017f752067656e657261746f72206f6620696e6e657270726f6475637400000000008152506200017e640100000000026401000000009004565b60066000820151816000015560208201518160010155905050620001786200026f640100000000026401000000009004565b6200071a565b62000188620006ba565b600062000222836040516020018082805190602001908083835b602083101515620001c95780518252602082019150602081019050602083039250620001a2565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120600190046200050a6401000000000262000a24176401000000009004565b90506200026781600460408051908101604052908160008201548152602001600182015481525050620005456401000000000262000a5e179091906401000000009004565b915050919050565b600080620002db60405160200180807f6776730000000000000000000000000000000000000000000000000000000000815250600301905060405160208183030381529060405280519060200120600190046200050a6401000000000262000a24176401000000009004565b905060008090505b6002604002811015620003b9576200033b8183016040516020018082815260200191505060405160208183030381529060405280519060200120600190046200050a6401000000000262000a24176401000000009004565b92506200038083600460408051908101604052908160008201548152602001600182015481525050620005456401000000000262000a5e179091906401000000009004565b6008826080811015156200039057fe5b6002020160008201518160000155602082015181600101559050508080600101915050620002e3565b5060006200042560405160200180807f6876730000000000000000000000000000000000000000000000000000000000815250600301905060405160208183030381529060405280519060200120600190046200050a6401000000000262000a24176401000000009004565b905060008090505b60026040028110156200050457620004858282016040516020018082815260200191505060405160208183030381529060405280519060200120600190046200050a6401000000000262000a24176401000000009004565b9350620004ca84600460408051908101604052908160008201548152602001600182015481525050620005456401000000000262000a5e179091906401000000009004565b61010882608081101515620004db57fe5b60020201600082015181600001556020820151816001015590505080806001019150506200042d565b50505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156200053c57fe5b06915050919050565b6200054f620006ba565b6001821415620005625782905062000609565b60028214156200058f576200058783846200060f640100000000026401000000009004565b905062000609565b62000599620006d4565b8360000151816000600381101515620005ae57fe5b6020020181815250508360200151816001600381101515620005cc57fe5b60200201818152505082816002600381101515620005e657fe5b6020020181815250506040826060836007600019fa15156200060757600080fd5b505b92915050565b62000619620006ba565b62000623620006f7565b83600001518160006004811015156200063857fe5b60200201818152505083602001518160016004811015156200065657fe5b60200201818152505082600001518160026004811015156200067457fe5b60200201818152505082602001518160036004811015156200069257fe5b6020020181815250506040826080836006600019fa1515620006b357600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b610c39806200072a6000396000f3fe6080604052600436106100fc576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304c09ce9146101015780630c00f8a01461015457806324d6147d146101aa5780632e52d606146101fd5780633e8d3764146102285780633e95522514610253578063483767f01461027e5780634db11875146102d157806352c9b9651461032557806356e2736c146103575780637982ebcc146103ad57806382529fdb146103d8578063b8c9d3651461042b578063c6a898c51461045d578063da2b99ce1461048f578063da897224146104e3578063e2179b8e1461050e578063ffe237f014610540575b600080fd5b34801561010d57600080fd5b50610116610593565b6040518082600260200280838360005b83811015610141578082015181840152602081019050610126565b5050505090500191505060405180910390f35b34801561016057600080fd5b5061018d6004803603602081101561017757600080fd5b81019080803590602001909291905050506105e5565b604051808381526020018281526020019250505060405180910390f35b3480156101b657600080fd5b506101bf61060e565b6040518082600260200280838360005b838110156101ea5780820151818401526020810190506101cf565b5050505090500191505060405180910390f35b34801561020957600080fd5b50610212610661565b6040518082815260200191505060405180910390f35b34801561023457600080fd5b5061023d610666565b6040518082815260200191505060405180910390f35b34801561025f57600080fd5b5061026861066b565b6040518082815260200191505060405180910390f35b34801561028a57600080fd5b50610293610674565b6040518082608060200280838360005b838110156102be5780820151818401526020810190506102a3565b5050505090500191505060405180910390f35b3480156102dd57600080fd5b506102e661070f565b604051808261010060200280838360005b838110156103125780820151818401526020810190506102f7565b5050505090500191505060405180910390f35b34801561033157600080fd5b5061033a6107ad565b604051808381526020018281526020019250505060405180910390f35b34801561036357600080fd5b506103906004803603602081101561037a57600080fd5b81019080803590602001909291905050506107bf565b604051808381526020018281526020019250505060405180910390f35b3480156103b957600080fd5b506103c26107e9565b6040518082815260200191505060405180910390f35b3480156103e457600080fd5b506103ed6107ee565b6040518082600260200280838360005b838110156104185780820151818401526020810190506103fd565b5050505090500191505060405180910390f35b34801561043757600080fd5b50610440610841565b604051808381526020018281526020019250505060405180910390f35b34801561046957600080fd5b50610472610853565b604051808381526020018281526020019250505060405180910390f35b34801561049b57600080fd5b506104a4610865565b604051808261010060200280838360005b838110156104d05780820151818401526020810190506104b5565b5050505090500191505060405180910390f35b3480156104ef57600080fd5b506104f8610905565b6040518082815260200191505060405180910390f35b34801561051a57600080fd5b5061052361090e565b604051808381526020018281526020019250505060405180910390f35b34801561054c57600080fd5b50610555610920565b6040518082608060200280838360005b83811015610580578082015181840152602081019050610565565b5050505090500191505060405180910390f35b61059b6109b9565b6105a36109b9565b60008001548160006002811015156105b757fe5b6020020181815250506000600101548160016002811015156105d557fe5b6020020181815250508091505090565b6008816080811015156105f457fe5b600202016000915090508060000154908060010154905082565b6106166109b9565b61061e6109b9565b60066000015481600060028110151561063357fe5b60200201818152505060066001015481600160028110151561065157fe5b6020020181815250508091505090565b600681565b604081565b60006006905090565b61067c6109db565b6106846109db565b60008090505b604081101561070757610108816080811015156106a357fe5b600202016000015482826002026080811015156106bc57fe5b602002018181525050610108816080811015156106d557fe5b600202016001015482600183600202016080811015156106f157fe5b602002018181525050808060010191505061068a565b508091505090565b6107176109ff565b61071f6109ff565b60008090505b60026040028110156107a55760088160808110151561074057fe5b600202016000015482826002026101008110151561075a57fe5b60200201818152505060088160808110151561077257fe5b600202016001015482600183600202016101008110151561078f57fe5b6020020181815250508080600101915050610725565b508091505090565b60048060000154908060010154905082565b610108816080811015156107cf57fe5b600202016000915090508060000154908060010154905082565b600281565b6107f66109b9565b6107fe6109b9565b60026000015481600060028110151561081357fe5b60200201818152505060026001015481600160028110151561083157fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b61086d6109ff565b6108756109ff565b60008090505b60026040028110156108fd576101088160808110151561089757fe5b60020201600001548282600202610100811015156108b157fe5b602002018181525050610108816080811015156108ca57fe5b60020201600101548260018360020201610100811015156108e757fe5b602002018181525050808060010191505061087b565b508091505090565b60006040905090565b60008060000154908060010154905082565b6109286109db565b6109306109db565b60008090505b60408110156109b15760088160808110151561094e57fe5b6002020160000154828260020260808110151561096757fe5b60200201818152505060088160808110151561097f57fe5b6002020160010154826001836002020160808110151561099b57fe5b6020020181815250508080600101915050610936565b508091505090565b6040805190810160405280600290602082028038833980820191505090505090565b61100060405190810160405280608090602082028038833980820191505090505090565b6120006040519081016040528061010090602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508083811515610a5557fe5b06915050919050565b610a66610bad565b6001821415610a7757829050610b05565b6002821415610a9157610a8a8384610b0b565b9050610b05565b610a99610bc7565b8360000151816000600381101515610aad57fe5b6020020181815250508360200151816001600381101515610aca57fe5b60200201818152505082816002600381101515610ae357fe5b6020020181815250506040826060836007600019fa1515610b0357600080fd5b505b92915050565b610b13610bad565b610b1b610bea565b8360000151816000600481101515610b2f57fe5b6020020181815250508360200151816001600481101515610b4c57fe5b6020020181815250508260000151816002600481101515610b6957fe5b6020020181815250508260200151816003600481101515610b8657fe5b6020020181815250506040826080836006600019fa1515610ba657600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a7230582021eaab2949956f42de51fea10cba06158adbc245f9121c871f2847c8fac218290029";

    public static final String FUNC_GVECTOR = "gVector";

    public static final String FUNC_N = "n";

    public static final String FUNC_BITSIZE = "bitSize";

    public static final String FUNC_GBASE = "gBase";

    public static final String FUNC_HVECTOR = "hVector";

    public static final String FUNC_AGGSIZE = "aggSize";

    public static final String FUNC_H = "h";

    public static final String FUNC_U = "u";

    public static final String FUNC_G = "g";

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

    public RemoteCall<BigInteger> aggSize() {
        final Function function = new Function(FUNC_AGGSIZE, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
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
