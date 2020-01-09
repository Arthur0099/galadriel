# PGC

PGC(Pretty Good Confidential Transaction System with Accountability)是一种基于账户体现的加密交易系统。

使用到的底层技术：
* Elgamal: 作用pgc账户体系的pke使用，其公钥对交易数据进行加密，提供隐私。**对原生的加密和解密的方法进行了修改，以适应bulletproof**
* BulletProof: 提供零知识证明相关功能。
* ECDSA: 提供签名和验签等相关功能，与以太坊和比特币兼容。

# 整体流程

## 公共参数生成



1. prover对其秘密数据v今次进行commit，获得m个变量，這些变量代表着秘密输入v.
2. prover进行以下步骤生成constraints。
   1. 


# 交易流程

交易开始前，Alice的账户上的余额为CA；Bob的账户上的余额为CB，其中，CA是使用了Alice公钥加密后的数据；CB是使用了Bob的工钥加密后的数据。
Cout=(X1, Y1); Cin=(X2, Y2)

例如：Alice向Bob转帐V个。
1. Alice使用自己的工钥和Bob的工钥分别对V进行加密，得到加密后的输出Cout和Cin。
2. 提供相应的零知识证明的相关数据。
   * L1(证明Cout中和Cin中加密的值一致)：使用Σ-protocol进行证明：(需要传入的数据如下（pk1, Cout， pk2, Cin）)
     * Prover随机选取a1,a2,b；并且计算A1=pk1^a1; A2=pk2^a2; B1=g^a1 * h^b; B2=g^a2 * h^b;并将其发送至Verifier。
     * Verifier随机选择e并将其发送给Prover作为一个挑战。
     * Prover计算z1=a1+e×r1; z2=a2+e×r2; z3=b+e*v(Alice自己知道V的值是多少)；并将z1,z2,z3发送给Verifier。
     * Verifier计算以下：pk1^z1==A1* X1^e; pk2^z2==A2 * X2^e; g^z1 * h^z3 == B1* Y1^e; g^z2 * h^z3 == B2 * Y2^e;
     * **因为在上诉的过程中，Prover和Verifier需要进行交互，所以使用Fiat－Shamir来替换其中e的选择**Prover自己使用Hash（g， h， ）参数待定来自行计算e的值，最后一次性将
     所有的数据发送给Verifier，因为Hash方法的不可预测性，所有Prover伪造证明的几率较小。Verifier可以根据同样的规则计算出e是多少值，从而进行验证。
  * L2(Alice转账额V处于一定范围)：使用bulletproof的rangeproof进行证明。
  * L3(Alice转账后的余额大于0,处于一定范围)：
    * 因为Alice并不知道CA-Cout值中对应的随机数是多少，所以她无法直接对CA-Cout提供一个零知识证明。但是，因为CA-Cout是用Alice的公钥加密的，所以Alice可以先对CA-Cout进行
    解密，然后使用一个新的随机数生成新的加密数据C’；然后证明以下内容：
    * CA-Cout和C’加密的是同样的数据
    * 使用零知识证明证明C‘是处于一定范围内的。 现在新的加密后的数据为（X‘， Y’）， X‘＝pk1^r’；Y‘＝g^r' * h^(m-v)； 原数据为（X1, Y1）, X1=pk1^r1; Y1=g^r1 * h^(m-v);
    所以等同于证明log Ỹ /Y ′ X̃/X ′ equals log g pk 1 with witness sk 1
    * Σ2:输入参数（Y’/Y1, X'/X1, g, pk1）（以g1, h1, g2, h2代替）, witness为sk1(Aclice的私钥)：
      * Prover随机选取a，计算A1=g1^a；A2=g2^a；并发送给Verifier；
      * Verifier随机选择e发送给Prover
      * Prover计算z＝a＋w×e；并将其发送至verifier（w即为sk1）
      * Verifier计算g1^z＝＝A1*h1^e；g2^z＝A2*h2^e； 相等则接受证明。
      * **同样使用Fiat－shamir来替换e的选择**
     
3. 
# PKE

PKE的相关功能由Elgamal算法提供，同时对其进行了调整以适应bulletproof。

# Bulletproof

bulletproof零知识证明。

## 基本架构

### rlcs_proof(constraint system)

### inner product protocol

用于压缩数据，减少协议中需要传输的数据量。



## 参考

* [百科](https://zh.wikipedia.org/wiki/ElGamal%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95)
* [bulletproof-rust](https://doc-internal.dalek.rs/bulletproofs/notes/index.html)

## zether对比

1. 6.2优化.(pgc也用了里面的sigle muti)
2. conftransfer中, sigmaproof和2个rangeproof共花费了156次加法和154次乘法.
3. 关于multi-exponent的说法, 在18页.?? 反而会增加gas消耗.(不适合solidity)

## pgc

# 32bit
transfer:
1. sigma proof: 10 mul, 6 add.
2. dle sigma proof: 4 mul, 4 add.
3. range proof: 2*(bitSize+n)+14 mul, 2*(bitSize+n)+9 add.
4. sig: 4 mul, 1 add.
5. 其他: 8 add.

1. mul: 40000 gas; add: 600 gas.

2. transfer 共用了1次sigma, 1次dlesigma, 2次range: **194 mul, 185 add**,总共消耗gas**9182528**(曲线消耗**7871000**, 其余消耗**1311528**)
3. zether 1次sigma,2次range, 154 mul, 156 add. (曲线消耗**6253600**, 总共**718.8万**, 其余消耗**934400**)

使用分叉后:
mul: 6000 gas; add: 150 gas.
transfer: **3191240**, 曲线**1191750**, 其余**1999490**

# 16bit

1. transfer, gas:**5894975**(曲线消耗**2914000**, 其余消耗**2980975**)

使用分叉后: 
mul: 6000 gas; add: 150 gas.
gas:**2162394**(曲线消耗**442200**, 其余消耗**1720194**)


# 最新对比

zether transfer(32bit):
* gascost: 7188k.
* eccost: 6455k.
* tx: 1472bytes.


pgc transfer(32bit);
* gascost: 9183k.
* eccost: 7871k.
* tx: 3268bytes(未对点进行压缩优化).


// 