# 系统环境

1. 操作系统: ubuntu 16.04, amd64, 4.15.0-46-generic
2. CPU: 1个, 4核, (4  Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz)
3. 内存: 8GB
4. 编程语言: go version go1.13 linux/amd64
5. 测试方法: go语言内置Benchmark测试框架

# 依赖库

1. secp256k1: "github.com/btcsuite/btcd/btcec" S256()
2. Hash: "github.com/ethereum/go-ethereum/crypto" Keccak256()
3. 随机数生成, big.Int等库: go官方crypto, math库.

# 测试数据
bitSize = 32

聚合range proof
| 操作         | 平均花费时间 | 总执行次数 | 总执行时间 |
|--------------|--------------|------------|------------|
| 生成ctx      | 156.79ms     | 7          | 1.298s     |
| 验证ctx      | 77.6ms       | 14         | 1.525s     |
| 生成+验证ctx | 234.17ms     | 5          | 2.403s     |

原两个range proof版本
| 操作         | 平均花费时间 | 总执行次数 | 总执行时间 |
|--------------|--------------|------------|------------|
| 生成ctx      | 302.37ms     | 4          | 2.514s     |
| 验证ctx      | 136.98ms     | 8          | 3.281s     |
| 生成+验证ctx | 430.52ms     | 3          | 2.637s     |

## 仅transfer

//