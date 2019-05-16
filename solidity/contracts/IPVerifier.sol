pragma solidity >= 0.5.0 < 0.6.0;
pragma experimental ABIEncoderV2;

import "./library/BN128.sol";

interface IPParams {
  function getBitSize() external view returns(uint);
  function getN() external view returns(uint);
  function getGVector() external view returns(BN128.G1Point[16] memory);
  function getHVector() external view returns(BN128.G1Point[16] memory);
  function getU() external view returns(BN128.G1Point memory);
}

contract IPVerifier {
  using BN128 for BN128.G1Point;
  using BN128 for uint;

  // bit size of value.
  uint public constant bitSize = 16;
  // number of inner product points.
  uint public constant n = 4;

  // u fix point used in inner product.
  BN128.G1Point public u;

  // public params contract.
  IPParams public ipParams;

  // IPProof contains proof to verify inner product.
  struct IPProof {
    BN128.G1Point[n] l;
    BN128.G1Point[n] r;
    uint a;
    uint b;
  }

  // for tmp calculation.
  struct Board {
    BN128.G1Point[bitSize] hv;
    BN128.G1Point[bitSize] gv;
  }

  constructor(address params) public {
    ipParams = IPParams(params);
    require(bitSize == ipParams.getBitSize(), "bis size not equal");
    require(n == ipParams.getN(), "number of l,r not equal");
    BN128.G1Point memory tmpU = ipParams.getU();
    u.X = tmpU.X;
    u.Y = tmpU.Y;
  }

  /*
   * @dev verify inner product proof.
   * p[0-1]: p point.
   * scalar[0]: a.
   * scalar[1]: b.
   * scalar[2]: c.
   */
  function verifyIPProof(uint[2] memory p, uint[n*2] memory l, uint[n*2] memory r, uint[3] memory scalar) public view returns(bool) {
    IPProof memory proof;
    for (uint i = 0; i < n; i++) {
      proof.l[i] = BN128.G1Point(l[i*2], l[i*2+1]);
      proof.r[i] = BN128.G1Point(r[i*2], r[i*2+1]);
    }
    proof.a = scalar[0];
    proof.b = scalar[1];

    return verifyIPProofStep1(ipParams.getGVector(), ipParams.getHVector(), BN128.G1Point(p[0], p[1]), scalar[2], proof);
  }


  function verifyIPProofStep1(BN128.G1Point[bitSize] memory gv, BN128.G1Point[bitSize] memory hv, BN128.G1Point memory p, uint c, IPProof memory proof) internal view returns(bool) {
    // compute challenge e.
    uint e = computeChallengeStep1(p.X, p.Y, c);

    BN128.G1Point memory ue = u.mul(e);

    return verifyIPProofStep2(gv, hv, ue, ue.mul(c).add(p), proof);
  }

  //
  function verifyIPProofStep2(BN128.G1Point[bitSize] memory gv, BN128.G1Point[bitSize] memory hv, BN128.G1Point memory newU, BN128.G1Point memory p, IPProof memory proof) internal view returns(bool) {

    Board memory b;
    uint step = 2;
    b.gv = gv;
    b.hv = hv;

    for (uint i = 0; i < n; i++) {
      uint e = computeChallengeStep2(proof.l[i].X, proof.l[i].Y, proof.r[i].X, proof.r[i].Y);
      uint eInverse = e.inv();

      for (uint j = 0; j < bitSize/step; j++) {
        // compute gv prime.
        b.gv[j] = b.gv[j].mul(eInverse).add(b.gv[bitSize/step+j].mul(e));
        // compute hv prime.
        b.hv[j] = b.hv[j].mul(e).add(b.hv[bitSize/step+j].mul(eInverse));
      }

      // compute p points.
      // p' = l*x^2 + r^*xInv^2 + p.
      p = proof.l[i].mul(e.mul(e).mod()).add(proof.r[i].mul(eInverse.mul(eInverse).mod())).add(p);

      step = step * 2;
    }

    // c = a * b;
    uint c = proof.a.mul(proof.b).mod();

    // want = gv[0]*a + hv[0]*b + u*c.
    BN128.G1Point memory want = newU.mul(c).add(b.gv[0].mul(proof.a)).add(b.hv[0].mul(proof.b));
    return p.X == want.X && p.Y == want.Y;
  }

  function computeChallengeStep1(uint a, uint b, uint c) internal view returns(uint) {
    return uint(keccak256(abi.encodePacked(a, b, c))).mod();
  }

  function computeChallengeStep2(uint a, uint b, uint c, uint d) internal view returns(uint) {
    return uint(keccak256(abi.encodePacked(a, b, c, d))).mod();
  }
}
