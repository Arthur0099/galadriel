pragma solidity >= 0.5.0 < 0.6.0;
pragma experimental ABIEncoderV2;

import "./library/BN128.sol";
import "./DLESigmaVerifier.sol";
import "./PublicParams.sol";

contract PGC {

  using BN128 for BN128.G1Point;
  using BN128 for uint;

  // encrypted ct.(balance of account)
  struct CT {
    BN128.G1Point X;
    BN128.G1Point Y;
  }

  // pk.x => pk.y => encrypt balance.
  mapping(uint => mapping(uint => CT)) balance;

  // bitSize of balance value.
  uint public constant bitSize = 16;
  uint public constant n = 4;

  // public h point.
  BN128.G1Point public h;

  // public g point.
  BN128.G1Point public g;

  // dle sigma verifier.
  DLESigmaVerifier public dleSigmaVerifier;

  // params.
  PublicParams public params;

  //
  constructor(address params_, address dleSigmaVerifier_) public {
    params = PublicParams(params_);
    dleSigmaVerifier = DLESigmaVerifier(dleSigmaVerifier_);

    BN128.G1Point memory tmpH = params.getH();
    BN128.G1Point memory tmpG = params.getG();

    h.X = tmpH.X;
    h.Y = tmpH.Y;
    g.X = tmpG.X;
    g.Y = tmpG.Y;
    require(bitSize == params.getBitSize(), "bitsize not equal");
  }

  /*
   * @dev deposit eth to an account.
   */
  function depositAccount(uint[2] memory publicKey) public payable returns(bool) {
    // check eth account.
    require(msg.value > 1 ether, "eth deposited less than 1 eth");
    require(msg.value/1 ether*1 ether == msg.value, "eth amount not an integer");

    // add amount to user account.
    uint amount = msg.value/1 ether;

    CT storage userBalance = balance[publicKey[0]][publicKey[1]];
    CT memory deposited = encrypt(amount, BN128.G1Point(publicKey[0], publicKey[1]));
    userBalance.X = userBalance.X.add(deposited.X);
    userBalance.Y = userBalance.Y.add(deposited.Y);

    return true;
  }

  event T(uint x);
  /*
   * @dev burn withdraw all eth back to eth account.
   * proof[0-1]: A1 point.
   * proof[2-3]: A2 point.
   *
   */
  function burn(address payable receiver, uint amount, uint[2] memory publicKey, uint[4] memory proof, uint z) public returns(bool) {
    // do nothing
    if (amount < 1) {
      return false;
    }
    // compute y' = Y - g*m.
    CT storage userBalance = balance[publicKey[0]][publicKey[1]];
    // todo: check not zero.

    BN128.G1Point memory y = userBalance.Y.add(g.mul(amount).neg());
    uint[12] memory points;
    points[0] = proof[0];
    points[1] = proof[1];
    points[2] = proof[2];
    points[3] = proof[3];
    points[4] = y.X;
    points[5] = y.Y;
    points[6] = userBalance.X.X;
    points[7] = userBalance.X.Y;
    points[8] = h.X;
    points[9] = h.Y;
    points[10] = publicKey[0];
    points[11] = publicKey[1];

    // do nothing when error.
    if (!dleSigmaVerifier.verifyDLESigmaProof(points, z)) {
      return false;
    }

    // update user encrypted balance.
    CT memory updatedBalance = encrypt(0, BN128.G1Point(publicKey[0], publicKey[1]));
    userBalance.X = updatedBalance.X;
    userBalance.Y = updatedBalance.Y;

    // transfer eth back to user.
    receiver.transfer(amount*1 ether);
  }

  /*
   * @dev encrypt for despoit eth account. the randomness is set to 0.
   * @dev it's ok since the eth amount is public know to everyone.
   * @dev but the ctx's randomness is not know.
   */
  function encrypt(uint amount, BN128.G1Point memory pk) internal view returns(CT memory) {
    require(amount < 2**bitSize, "amount out of range");

    // no sense.
    uint r = 0;

    CT memory ct;
    // X = pk * r.
    ct.X = pk.mul(r);
    // Y = g*m + h*r.
    ct.Y = g.mul(amount).add(h.mul(r));

    return ct;
  }

  /*
   * @dev transfer from one to another.
   * pks[0-1]: sender
   * pks[2-3]: receiver
   *****************************************************
   * rangeProof1: prove transfer amount v in right range.
   * points[0-1]: A.
   * points[2-3]: S.
   * points[4-5]: T1.
   * points[6-7]: T2.
   * points[8-9]: V(commitment).
   * scalar[0]: t.
   * scalar[1]: tx.
   * scalar[2]: u.
   * scalar[3]: a.
   * scalar[4]: b.
   * l[0-1]: l0.x, l0.y
   ...
   * l[6-7]: l3.x, l3.y
   * r[0-1]: r0.x, r0.y
   ...
   * r[6-7]: r3.x, r3.y
   *****************************************************
   * rangeProof2: prove sender's updated balance is in right range.
   * points[10-11]: A.
   * points[12-13]: S.
   * points[14-15]: T1.
   * points[16-17]: T2.
   * points[18-19]: V.
   * scalar[5]: t.
   * scalar[6]: tx.
   * scalar[7]: u.
   * scalar[8]: a.
   * scalar[9]: b.
   * l[8-9]: l0.x, l0.y
   ...
   * l[14-15]: l3.x, l3.y
   ...
   * r[8-9]: r0.x, r0.y
   ...
   * r[14-15]: r3.x, r4.y
   *****************************************************
   * dle sigma proof prove two encrypted data was same data under
   * same public key.
   * points[20-21]: A1.
   * points[22-23]: A2.
   * points[24-25]: g1.
   * points[26-27]: h1.
   * points[28-29]: g2.
   * points[30-31]: h2.
   * scalar[10]: z.
   ******************************************************
   * sigma proof proving two encrypted was same value under two
   * public key.
   * points[32-33]:
   */
  function transfer() public returns(bool) {
    //
  }
}
