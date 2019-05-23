pragma solidity >= 0.5.0 < 0.6.0;

import "./library/BN128.sol";

contract PGC {

  using BN128 for BN128.G1Point;
  using BN128 for uint;

  // .
  //  mapping(string => BN128.G1Point) balance;
  //  mapping(string => BN128.G1Point) pk;
  //  mapping(string => bool) used;

  //
  function initialBalance(string memory name, uint[2] memory publicKey, uint[2] memory balanceCT) public returns(bool) {
    //    if (used[name]) {
    //      return false;
    //    }

    //    used[name] = true;
    //    balance[name] = BN128.G1Point(balanceCT[0], balanceCT[1]);
    //    pk[name] = BN128.G1Point(publicKey[0], publicKey[1]);

    return true;
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
