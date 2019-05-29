pragma solidity >= 0.5.0 < 0.6.0;

import "./library/BN128.sol";

interface SigmaProofParams {
  function getG() external view returns(uint[2] memory);
  function getH() external view returns(uint[2] memory);
}

contract SigmaVerifier {
  using BN128 for BN128.G1Point;
  using BN128 for uint;

  BN128.G1Point public g;
  BN128.G1Point public h;

  constructor(address params) public {
    // init g and h point.
    SigmaProofParams sigmaParams = SigmaProofParams(params);
    uint[2] memory tmpG = sigmaParams.getG();
    uint[2] memory tmpH = sigmaParams.getH();
    g.X = tmpG[0];
    g.Y = tmpG[1];

    h.X = tmpH[0];
    h.Y = tmpH[1];
  }

  /*
   * @dev verify sigma proof.
   * points[0-1]: pk1
   * points[2-3]: ct1 X
   * points[4-5]: ct1 Y
   * points[6-7]: pk2
   * points[8-9]: ct2 X
   * points[10-11]: ct2 Y
   * points[12-13]: A1
   * points[14-15]: A2
   * points[16-17]: B1
   * points[18-19]: B2
   *
   */
  function verifySigmaProof(uint[20] memory points, uint z1, uint z2, uint z3) public view returns(bool) {
    uint e;
    // compute challenge e.
    e = computeChallenge(points[12], points[13], points[14], points[15], points[16], points[17], points[18], points[19]);

    // check pk1 * z1 == A1 + X1 * e.
    BN128.G1Point[3] memory point1;
    point1[0] = BN128.G1Point(points[0], points[1]);
    point1[1] = BN128.G1Point(points[12], points[13]);
    point1[2] = BN128.G1Point(points[2], points[3]);
    if (!verifyStep1(point1, z1, e)) {
      return false;
    }

    // check pk2 * z2 == A2 + X2 * e.
    BN128.G1Point[3] memory point2;
    point2[0] = BN128.G1Point(points[6], points[7]);
    point2[1] = BN128.G1Point(points[14], points[15]);
    point2[2] = BN128.G1Point(points[8], points[9]);
    if (!verifyStep1(point2, z2, e)) {
      return false;
    }

    // check h * z1 + g * z3 == B1 + Y1 * e.
    BN128.G1Point[2] memory point3;
    point3[0] = BN128.G1Point(points[16], points[17]);
    point3[1] = BN128.G1Point(points[4], points[5]);
    if (!verifyStep2(point3, z1, z3, e)) {
      return false;
    }

    // check h*z2 + g*z3 == B2 + Y2*e.
    BN128.G1Point[2] memory point4;
    point4[0] = BN128.G1Point(points[18], points[19]);
    point4[1] = BN128.G1Point(points[10], points[11]);
    if (!verifyStep2(point4, z2, z3, e)) {
      return false;
    }

    return true;
  }

  /*
   * @dev check pk * z == A + X * e.
   * points[0]: pk
   * points[1]: A
   * points[2]: X
   */
  function verifyStep1(BN128.G1Point[3] memory points, uint z, uint e) internal view returns(bool) {
    // compute x*e + A.
    BN128.G1Point memory want;
    want = points[2].mul(e).add(points[1]);

    BN128.G1Point memory actual;
    actual = points[0].mul(z);

    return want.X == actual.X && want.Y == actual.Y;
  }

  /*
   * @dev check h*za + g*zb == B + Y*e.
   * points[0]: B
   * points[1]: Y
   */
  function verifyStep2(BN128.G1Point[2] memory points, uint za, uint zb, uint e) internal view returns(bool) {
    // comute g * za + h * zb.
    BN128.G1Point memory want = h.mul(za).add(g.mul(zb));
    BN128.G1Point memory actual = points[0].add(points[1].mul(e));

    return want.X == actual.X && want.Y == actual.Y;
  }

  // compute challenge for proof.
  function computeChallenge(uint a, uint b, uint c, uint d, uint e, uint f, uint i, uint j) internal pure returns(uint) {
    return uint(keccak256(abi.encodePacked(a, b, c, d, e, f, i, j))).mod();
  }
}
