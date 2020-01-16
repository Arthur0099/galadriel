pragma solidity >= 0.5.0 < 0.6.0;

import "./library/Secp256k1.sol";
import "./library/BN128.sol";

// verify dle sigmaproof.
contract DLESigmaVerifier {

  using Secp256k1 for *;

  using BN128 for BN128.G1Point;
  using BN128 for *;

  /**
   * @dev verify dle sigma proof. 4 mul, 2 add.
   * points[0-1]: A1 point.
   * points[2-3]: A2 point.
   * points[4-5]: g1 point.
   * points[6-7]: h1 point.
   * points[8-9]: g2 point.
   * points[10-11]: h2 point.
   * z = a + e*w.
   */
  function verifyDLESigmaProof(uint[12] memory points, uint z) public view returns(bool) {
    uint e = computeChallenge(points[0], points[1], points[2], points[3]);
    uint[6] memory check1;
    check1[0] = points[4];
    check1[1] = points[5];
    check1[2] = points[0];
    check1[3] = points[1];
    check1[4] = points[6];
    check1[5] = points[7];
    // g1^z == h1^e+A1
    if (!checkDLESigmaProofBn128(check1, z, e)) {
        return false;
    }

    uint[6] memory check2;
    check2[0] = points[8];
    check2[1] = points[9];
    check2[2] = points[2];
    check2[3] = points[3];
    check2[4] = points[10];
    check2[5] = points[11];
    // g2^z == h2^e+A2
    if (!checkDLESigmaProofBn128(check2, z, e)) {
        return false;
    }

    return true;
  }

  /*
   * @dev check g * z == A + h * e.
   * points[0-1]: g point.
   * points[2-3]: A point.
   * points[4-5]: h point.
   */
  function checkDLESigmaProof(uint[6] memory points, uint z, uint e) internal pure returns(bool) {
    uint gzX;
    uint gzY;
    uint heX;
    uint heY;
    uint checkX;
    uint checkY;
    // g * z.
    (gzX, gzY) = points[0].ecmul(points[1], z);
    // h * e.
    (heX, heY) = points[4].ecmul(points[5], e);
    // a + h * e.
    (checkX, checkY) = points[2].ecadd(points[3], heX, heY);

    return (gzX == checkX && gzY == checkY);
  }

  /*
   * @dev 2 mul, 1 add.
   */
  // check g^z == h*e+A.
  // points[0-1]: g
  // points[2-3]: A
  // points[4-5]: h
  function checkDLESigmaProofBn128(uint[6] memory points, uint z, uint e) internal view returns(bool) {
    BN128.G1Point memory g = BN128.G1Point(points[0], points[1]);
    BN128.G1Point memory A = BN128.G1Point(points[2], points[3]);
    BN128.G1Point memory h = BN128.G1Point(points[4], points[5]);

    g = g.mul(z);
    h = h.mul(e).add(A);

    return g.X == h.X && g.Y == h.Y;
  }

  function computeChallenge(uint a, uint b, uint c, uint d) internal pure returns(uint) {
    return uint(keccak256(abi.encodePacked(a, b, c, d))).mod();
  }

  //
}
