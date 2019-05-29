pragma solidity >= 0.5.0 < 0.6.0;

import "./library/BN128.sol";

interface IPParams {
  function getBitSize() external view returns(uint);
  function getN() external view returns(uint);
  function getGVector() external view returns(uint[32] memory);
  function getHVector() external view returns(uint[32] memory);
  function getU() external view returns(uint[2] memory);
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

    uint[n] challenges;
    uint[n] challengesInverse;
  }

  constructor(address params) public {
    ipParams = IPParams(params);
    require(bitSize == ipParams.getBitSize(), "bis size not equal");
    require(n == ipParams.getN(), "number of l,r not equal");
    uint[2] memory tmpU = ipParams.getU();
    u.X = tmpU[0];
    u.Y = tmpU[1];
  }


  /*
   * @dev verify inner product proof.
   * p[0-1]: p point.
   * scalar[0]: a.
   * scalar[1]: b.
   * scalar[2]: c.
   */
  function verifyIPProof(uint[2] memory p, uint[n*2] memory l, uint[n*2] memory r, uint[3] memory scalar) public view returns(bool) {
    return verifyIPProofWithCustomParams(ipParams.getGVector(), ipParams.getHVector(), p, scalar[2], l, r, scalar[0], scalar[1]);
  }

  /*
   * @dev call by range proof.
   * @dev Warning: hv isn't the public h vector generator.
   */
  function verifyIPProofWithCustomParams(uint[bitSize*2] memory gv, uint[bitSize*2] memory hv, uint[2] memory p, uint c, uint[n*2] memory l, uint[n*2] memory r, uint a, uint b) public view returns(bool) {
    BN128.G1Point[bitSize] memory gvPoints;
    BN128.G1Point[bitSize] memory hvPoints;
    for (uint i = 0; i < bitSize; i++) {
      gvPoints[i] = BN128.G1Point(gv[2*i], gv[2*i+1]);
      hvPoints[i] = BN128.G1Point(hv[2*i], hv[2*i+1]);
    }
    IPProof memory proof;
    for (uint i = 0; i < n; i++) {
      proof.l[i] = BN128.G1Point(l[2*i], l[2*i+1]);
      proof.r[i] = BN128.G1Point(r[2*i], r[2*i+1]);
    }

    proof.a = a;
    proof.b = b;
    return verifyIPProofStep1(gvPoints, hvPoints, BN128.G1Point(p[0], p[1]), c, proof);
  }


  function verifyIPProofStep1(BN128.G1Point[bitSize] memory gv, BN128.G1Point[bitSize] memory hv, BN128.G1Point memory p, uint c, IPProof memory proof) internal view  returns(bool) {
    // compute challenge e.
    uint e = computeChallengeStep1(c);

    BN128.G1Point memory ue = u.mul(e);

    return verifyIPProofMulti(gv, hv, ue, ue.mul(c).add(p), proof);
  }


  function verifyIPProofMulti(BN128.G1Point[bitSize] memory gv, BN128.G1Point[bitSize] memory hv, BN128.G1Point memory newU, BN128.G1Point memory p, IPProof memory proof) internal view  returns(bool) {
    Board memory b;
    // compute formula on the right.
    // compute p + li * xi^2 + ri * xi^-2.
    for (uint i = 0; i < n; i++) {
      uint x = computeChallengeStep2(proof.l[i].X, proof.l[i].Y, proof.r[i].X, proof.r[i].Y);
      uint xInverse = x.inv();
      b.challenges[i] = x;
      b.challengesInverse[i] = xInverse;
      p = p.add(proof.l[i].mul(x.mul(x))).add(proof.r[i].mul(xInverse.mul(xInverse)));
    }

    // compute formula on the left.
    // compute g*s*a + h*s^-1*b + u*a*b.

    // compute s.
    uint[bitSize] memory s;
    for (uint i = 0; i < bitSize; i++) {
      for (uint j = 0; j < n; j++) {
        uint tmp;

        if (smallParseBinary(i, j, n)) {
          tmp = b.challenges[j];
        } else {
          tmp = b.challengesInverse[j];
        }

        if (j == 0) {
          s[i] = tmp;
        } else {
          s[i] = s[i].mul(tmp).mod();
        }
      }

    }

    BN128.G1Point memory left;
    left = multiExp(gv, s).mul(proof.a).add(multiExpInverse(hv, s).mul(proof.b)).add(newU.mul(proof.a.mul(proof.b)));

    return left.X == p.X && left.Y == p.Y;
  }

  //
  function verifyIPProofStep2(BN128.G1Point[bitSize] memory gv, BN128.G1Point[bitSize] memory hv, BN128.G1Point memory newU, BN128.G1Point memory p, IPProof memory proof) internal view  returns(bool) {

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

  function computeChallengeStep1(uint a) internal pure returns(uint) {
    return uint(keccak256(abi.encodePacked(a))).mod();
  }

  function computeChallengeStep2(uint a, uint b, uint c, uint d) internal pure returns(uint) {
    return uint(keccak256(abi.encodePacked(a, b, c, d))).mod();
  }

  function smallParseBinary(uint t, uint j, uint size) internal pure returns(bool) {
    uint w = 1 << (size - 1);

    for (uint i=0; i < j; i++) {
      w = w >> 1;
    }

    if ((t&w) != 0) {
      return true;
    }

    return false;
  }

  function multiExp(BN128.G1Point[bitSize] memory base, uint[bitSize] memory exp) internal view returns(BN128.G1Point memory) {
    BN128.G1Point memory res;
    res = base[0].mul(exp[0]);
    for (uint i = 1; i < bitSize; i++) {
      res = res.add(base[i].mul(exp[i]));
    }

    return res;
  }

  function multiExpInverse(BN128.G1Point[bitSize] memory base, uint[bitSize] memory exp) internal view returns(BN128.G1Point memory) {
    uint[bitSize] memory expInverse;
    for (uint i = 0; i < bitSize; i++) {
      expInverse[i] = exp[i].inv();
    }

    return multiExp(base, expInverse);
  }
}
