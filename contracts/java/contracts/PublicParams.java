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
    private static final String BINARY = "0x60806040523480156200001157600080fd5b5060016004600001819055506002600460010181905550620000776040805190810160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c670000000000008152506200017e640100000000026401000000009004565b600080820151816000015560208201518160010155905050620000de6040805190810160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c670000000000008152506200017e640100000000026401000000009004565b60026000820151816000015560208201518160010155905050620001466040805190810160405280601b81526020017f752067656e657261746f72206f6620696e6e657270726f6475637400000000008152506200017e640100000000026401000000009004565b60066000820151816000015560208201518160010155905050620001786200026f640100000000026401000000009004565b62000719565b62000188620006b9565b600062000222836040516020018082805190602001908083835b602083101515620001c95780518252602082019150602081019050602083039250620001a2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060019004620005096401000000000262000a18176401000000009004565b90506200026781600460408051908101604052908160008201548152602001600182015481525050620005446401000000000262000a52179091906401000000009004565b915050919050565b600080620002db60405160200180807f677673000000000000000000000000000000000000000000000000000000000081525060030190506040516020818303038152906040528051906020012060019004620005096401000000000262000a18176401000000009004565b905060008090505b6002602002811015620003b9576200033b818301604051602001808281526020019150506040516020818303038152906040528051906020012060019004620005096401000000000262000a18176401000000009004565b92506200038083600460408051908101604052908160008201548152602001600182015481525050620005446401000000000262000a52179091906401000000009004565b6008826040811015156200039057fe5b6002020160008201518160000155602082015181600101559050508080600101915050620002e3565b5060006200042560405160200180807f687673000000000000000000000000000000000000000000000000000000000081525060030190506040516020818303038152906040528051906020012060019004620005096401000000000262000a18176401000000009004565b905060008090505b6002602002811015620005035762000485828201604051602001808281526020019150506040516020818303038152906040528051906020012060019004620005096401000000000262000a18176401000000009004565b9350620004ca84600460408051908101604052908160008201548152602001600182015481525050620005446401000000000262000a52179091906401000000009004565b608882604081101515620004da57fe5b60020201600082015181600001556020820151816001015590505080806001019150506200042d565b50505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156200053b57fe5b06915050919050565b6200054e620006b9565b6001821415620005615782905062000608565b60028214156200058e576200058683846200060e640100000000026401000000009004565b905062000608565b62000598620006d3565b8360000151816000600381101515620005ad57fe5b6020020181815250508360200151816001600381101515620005cb57fe5b60200201818152505082816002600381101515620005e557fe5b6020020181815250506040826060836007600019fa15156200060657600080fd5b505b92915050565b62000618620006b9565b62000622620006f6565b83600001518160006004811015156200063757fe5b60200201818152505083602001518160016004811015156200065557fe5b60200201818152505082600001518160026004811015156200067357fe5b60200201818152505082602001518160036004811015156200069157fe5b6020020181815250506040826080836006600019fa1515620006b257600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b610c2d80620007296000396000f3fe6080604052600436106100fc576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304c09ce9146101015780630c00f8a01461015457806324d6147d146101aa5780632e52d606146101fd5780633e8d3764146102285780633e95522514610253578063483767f01461027e5780634db11875146102d157806352c9b9651461032457806356e2736c146103565780637982ebcc146103ac57806382529fdb146103d7578063b8c9d3651461042a578063c6a898c51461045c578063da2b99ce1461048e578063da897224146104e1578063e2179b8e1461050c578063ffe237f01461053e575b600080fd5b34801561010d57600080fd5b50610116610591565b6040518082600260200280838360005b83811015610141578082015181840152602081019050610126565b5050505090500191505060405180910390f35b34801561016057600080fd5b5061018d6004803603602081101561017757600080fd5b81019080803590602001909291905050506105e3565b604051808381526020018281526020019250505060405180910390f35b3480156101b657600080fd5b506101bf61060c565b6040518082600260200280838360005b838110156101ea5780820151818401526020810190506101cf565b5050505090500191505060405180910390f35b34801561020957600080fd5b5061021261065f565b6040518082815260200191505060405180910390f35b34801561023457600080fd5b5061023d610664565b6040518082815260200191505060405180910390f35b34801561025f57600080fd5b50610268610669565b6040518082815260200191505060405180910390f35b34801561028a57600080fd5b50610293610672565b6040518082604060200280838360005b838110156102be5780820151818401526020810190506102a3565b5050505090500191505060405180910390f35b3480156102dd57600080fd5b506102e661070b565b6040518082608060200280838360005b838110156103115780820151818401526020810190506102f6565b5050505090500191505060405180910390f35b34801561033057600080fd5b506103396107a7565b604051808381526020018281526020019250505060405180910390f35b34801561036257600080fd5b5061038f6004803603602081101561037957600080fd5b81019080803590602001909291905050506107b9565b604051808381526020018281526020019250505060405180910390f35b3480156103b857600080fd5b506103c16107e2565b6040518082815260200191505060405180910390f35b3480156103e357600080fd5b506103ec6107e7565b6040518082600260200280838360005b838110156104175780820151818401526020810190506103fc565b5050505090500191505060405180910390f35b34801561043657600080fd5b5061043f61083a565b604051808381526020018281526020019250505060405180910390f35b34801561046857600080fd5b5061047161084c565b604051808381526020018281526020019250505060405180910390f35b34801561049a57600080fd5b506104a361085e565b6040518082608060200280838360005b838110156104ce5780820151818401526020810190506104b3565b5050505090500191505060405180910390f35b3480156104ed57600080fd5b506104f66108fa565b6040518082815260200191505060405180910390f35b34801561051857600080fd5b50610521610903565b604051808381526020018281526020019250505060405180910390f35b34801561054a57600080fd5b50610553610915565b6040518082604060200280838360005b8381101561057e578082015181840152602081019050610563565b5050505090500191505060405180910390f35b6105996109ae565b6105a16109ae565b60008001548160006002811015156105b557fe5b6020020181815250506000600101548160016002811015156105d357fe5b6020020181815250508091505090565b6008816040811015156105f257fe5b600202016000915090508060000154908060010154905082565b6106146109ae565b61061c6109ae565b60066000015481600060028110151561063157fe5b60200201818152505060066001015481600160028110151561064f57fe5b6020020181815250508091505090565b600581565b602081565b60006005905090565b61067a6109d0565b6106826109d0565b60008090505b6020811015610703576088816040811015156106a057fe5b600202016000015482826002026040811015156106b957fe5b6020020181815250506088816040811015156106d157fe5b600202016001015482600183600202016040811015156106ed57fe5b6020020181815250508080600101915050610688565b508091505090565b6107136109f4565b61071b6109f4565b60008090505b600260200281101561079f5760088160408110151561073c57fe5b6002020160000154828260020260808110151561075557fe5b60200201818152505060088160408110151561076d57fe5b6002020160010154826001836002020160808110151561078957fe5b6020020181815250508080600101915050610721565b508091505090565b60048060000154908060010154905082565b6088816040811015156107c857fe5b600202016000915090508060000154908060010154905082565b600281565b6107ef6109ae565b6107f76109ae565b60026000015481600060028110151561080c57fe5b60200201818152505060026001015481600160028110151561082a57fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b6108666109f4565b61086e6109f4565b60008090505b60026020028110156108f25760888160408110151561088f57fe5b600202016000015482826002026080811015156108a857fe5b6020020181815250506088816040811015156108c057fe5b600202016001015482600183600202016080811015156108dc57fe5b6020020181815250508080600101915050610874565b508091505090565b60006020905090565b60008060000154908060010154905082565b61091d6109d0565b6109256109d0565b60008090505b60208110156109a65760088160408110151561094357fe5b6002020160000154828260020260408110151561095c57fe5b60200201818152505060088160408110151561097457fe5b6002020160010154826001836002020160408110151561099057fe5b602002018181525050808060010191505061092b565b508091505090565b6040805190810160405280600290602082028038833980820191505090505090565b61080060405190810160405280604090602082028038833980820191505090505090565b61100060405190810160405280608090602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508083811515610a4957fe5b06915050919050565b610a5a610ba1565b6001821415610a6b57829050610af9565b6002821415610a8557610a7e8384610aff565b9050610af9565b610a8d610bbb565b8360000151816000600381101515610aa157fe5b6020020181815250508360200151816001600381101515610abe57fe5b60200201818152505082816002600381101515610ad757fe5b6020020181815250506040826060836007600019fa1515610af757600080fd5b505b92915050565b610b07610ba1565b610b0f610bde565b8360000151816000600481101515610b2357fe5b6020020181815250508360200151816001600481101515610b4057fe5b6020020181815250508260000151816002600481101515610b5d57fe5b6020020181815250508260200151816003600481101515610b7a57fe5b6020020181815250506040826080836006600019fa1515610b9a57600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a723058209b228fa2c7dacc626d699224c6f9aa426fc51ea11afef107e3048c03a5a96c610029";

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
