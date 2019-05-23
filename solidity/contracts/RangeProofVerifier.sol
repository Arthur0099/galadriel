pragma solidity >= 0.5.0 < 0.6.0;
pragma experimental ABIEncoderV2;

import "./library/BN128.sol";

interface RangeParams {
  function getBitSize() external view returns(uint);
  function getN() external view returns(uint);
  function getGVector() external view returns(BN128.G1Point[16] memory);
  function getHVector() external view returns(BN128.G1Point[16] memory);
  function getG() external view returns(BN128.G1Point memory);
  function getH() external view returns(BN128.G1Point memory);
  function getU() external view returns(BN128.G1Point memory);
}

interface IPVerifierInterface {
  function verifyIPProofWithCustomParams(BN128.G1Point[16] calldata gv, BN128.G1Point[16] calldata hv, BN128.G1Point calldata p, uint c, BN128.G1Point[4] calldata l, BN128.G1Point[4] calldata r, uint a, uint b) external returns(bool);
}

contract RangeProofVerifier {
  using BN128 for BN128.G1Point;
  using BN128 for uint;

   // bit size of value.
  uint public constant bitSize = 16;
  // number of inner product points.
  uint public constant n = 4;

  // public params contract.
  RangeParams public rpParams;

  // inner product verifier.
  IPVerifierInterface public ipVerifier;

  // g point.
  BN128.G1Point public g;
  // h point.
  BN128.G1Point public h;
  // uBase point.
  BN128.G1Point public uBase;

  // range proof.
  struct RangeProof {
    BN128.G1Point A;
    BN128.G1Point S;
    BN128.G1Point T1;
    BN128.G1Point T2;
    uint t;
    uint txx;
    uint u;
    IPProof ipProof;
  }

  // ip proof.
  struct IPProof {
    BN128.G1Point[n] l;
    BN128.G1Point[n] r;
    uint a;
    uint b;
  }

  // struct for tmp calculation.
  struct Board {
    uint x;
    uint x2;
    uint y;
    uint[bitSize] yn;
    uint[bitSize] ynInverse;
    uint z;
    uint zSquare;
    uint zNeg;
    uint[bitSize] n2;
    BN128.G1Point[bitSize] hPrime;
    BN128.G1Point expect;
    BN128.G1Point actual;
    uint dleta;
    uint[bitSize] exp;
    uint[n] challenges;
    uint[n] challengesInverse;
    uint[bitSize] l;
    uint[bitSize] r;
  }

  constructor(address params, address ip) public {
    rpParams = RangeParams(params);
    ipVerifier = IPVerifierInterface(ip);
    require(bitSize == rpParams.getBitSize(), "bis size not equal");
    require(n == rpParams.getN(), "number of l,r not equal");

    // set public params.
    BN128.G1Point memory tmpG = rpParams.getG();
    BN128.G1Point memory tmpH = rpParams.getH();
    BN128.G1Point memory tmpU = rpParams.getU();
    g.X = tmpG.X;
    g.Y = tmpG.Y;
    h.X = tmpH.X;
    h.Y = tmpH.Y;
    uBase.X = tmpU.X;
    uBase.Y = tmpU.Y;
  }

  /*
   * @dev verify range proof valid or not.
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
   * l[0-1]: l.x,l.y.
   * r[0-1]: r.x,r.y.
   */
  function verifyRangeProof(uint[10] memory points, uint[5] memory scalar, uint[2*n] memory l, uint[2*n] memory r) public returns(bool) {
    RangeProof memory rangeProof;
    rangeProof.A = BN128.G1Point(points[0], points[1]);
    rangeProof.S = BN128.G1Point(points[2], points[3]);
    rangeProof.T1 = BN128.G1Point(points[4], points[5]);
    rangeProof.T2 = BN128.G1Point(points[6], points[7]);
    rangeProof.t = scalar[0];
    rangeProof.txx = scalar[1];
    rangeProof.u = scalar[2];

    for (uint i = 0; i < n; i++) {
      rangeProof.ipProof.l[i] = BN128.G1Point(l[i*2], l[i*2+1]);
      rangeProof.ipProof.r[i] = BN128.G1Point(r[i*2], r[i*2+1]);
    }
    rangeProof.ipProof.a = scalar[3];
    rangeProof.ipProof.b = scalar[4];

    return optimizedVerify(BN128.G1Point(points[8], points[9]), rangeProof);
  }


  /*
   *
   */
  function verify(BN128.G1Point memory v, RangeProof memory rangeProof) public  returns(bool) {
    BN128.G1Point[bitSize] memory gv = rpParams.getGVector();
    BN128.G1Point[bitSize] memory hv = rpParams.getHVector();
    Board memory board;
    // compute challenge.
    board.y = computeChallenge(rangeProof.A.X, rangeProof.A.Y, rangeProof.S.X, rangeProof.S.Y);

    board.z = computeChallenge2(board.y);
    board.yn = powers(board.y);
    board.zNeg = board.z.neg();
    board.zSquare = board.z.mul(board.z).mod();
    board.n2 = powers(2);
    board.hPrime = hadamard(hv, powers(board.y.inv()));

    board.x = computeChallenge(rangeProof.T1.X, rangeProof.T1.Y, rangeProof.T2.X, rangeProof.T2.Y);
    board.x2 = board.x.mul(board.x);

    // check g*tx + h*t ?= v*z^2 + g*dleta + T1*x + T2*x^2.
    board.expect = v.mul(board.zSquare).add(rangeProof.T1.mul(board.x)).add(rangeProof.T2.mul(board.x2));
    // delta = (z - z^2) * <1^n, y^n> - z^3 * <1^n, 2^n>.
    board.dleta = board.z.sub(board.zSquare).mul(sum(board.yn)).sub(board.zSquare.mul(board.z).mul(sum(board.n2)));
    board.expect = board.expect.add(g.mul(board.dleta));
    board.actual = g.mul(rangeProof.t).add(h.mul(rangeProof.txx));
    if (board.expect.X != board.actual.X || board.expect.Y != board.actual.Y) {
      return false;
    }

    // compute p point. p = A + S*x + gv*-z + h'*(z*y^n + z^2 * 2^n).
    BN128.G1Point memory p = rangeProof.A.add(rangeProof.S.mul(board.x));
    p = p.add(sumVector(gv).mul(board.zNeg));
    board.exp = addFieldVector(times(board.yn, board.z), times(board.n2, board.zSquare));
    p = p.add(commit(board.hPrime, board.exp));
    // compute p'. p' = p - h*u. == g*l + h'*r.(this could be apply on inner product).
    p = p.add(h.mul(rangeProof.u).neg());

    return verifyIPInternal(gv, board.hPrime, p, rangeProof);
  }


  event Test(uint x, uint y);
  event SI(uint x, uint i);
  /*
   *
   */
  function optimizedVerify(BN128.G1Point memory v, RangeProof memory rangeProof) public returns(bool) {
    //
    BN128.G1Point[bitSize] memory gv = rpParams.getGVector();
    BN128.G1Point[bitSize] memory hv = rpParams.getHVector();
    Board memory board;
    // compute
    board.y = computeChallenge(rangeProof.A.X, rangeProof.A.Y, rangeProof.S.X, rangeProof.S.Y);
    board.z = computeChallenge2(board.y);
    board.yn = powers(board.y);
    board.zNeg = board.z.neg();
    board.zSquare = board.z.mul(board.z).mod();
    board.n2 = powers(2);

    board.x = computeChallenge(rangeProof.T1.X, rangeProof.T1.Y, rangeProof.T2.X, rangeProof.T2.Y);
    board.x2 = board.x.mul(board.x);

    // check g*tx + h*t ?= v*z^2 + g*dleta + T1*x + T2*x^2.
    // check g*(tx-dleta) + h*t ?= v*z^2 + T1*x + T2*x^2.
    board.expect = v.mul(board.zSquare).add(rangeProof.T1.mul(board.x)).add(rangeProof.T2.mul(board.x2));
    // delta = (z - z^2) * <1^n, y^n> - z^3 * <1^n, 2^n>.
    board.dleta = board.z.sub(board.zSquare).mul(sum(board.yn)).sub(board.zSquare.mul(board.z).mul(sum(board.n2)));
    board.actual = g.mul(rangeProof.t.sub(board.dleta)).add(h.mul(rangeProof.txx));
    if (board.expect.X != board.actual.X || board.expect.Y != board.actual.Y) {
      return false;
    }

    BN128.G1Point memory p;
    p = rangeProof.A.add(rangeProof.S.mul(board.x));

    // compute formula on the right.
    // compute p + li * xi^2 + ri * xi^-2.
    for (uint i = 0; i < n; i++) {
      uint x = computeChallenge(rangeProof.ipProof.l[i].X, rangeProof.ipProof.l[i].Y, rangeProof.ipProof.r[i].X, rangeProof.ipProof.r[i].Y);
      uint xInverse = x.inv();
      board.challenges[i] = x;
      board.challengesInverse[i] = xInverse;
      p = p.add(rangeProof.ipProof.l[i].mul(x.mul(x))).add(rangeProof.ipProof.r[i].mul(xInverse.mul(xInverse)));
    }

    for (uint i = 0; i < bitSize; i++) {
      for (uint j = 0; j < n; j++) {
        uint tmp;

         if (smallParseBinary(i, j, n)) {
          tmp = board.challenges[j];
        } else {
          tmp = board.challengesInverse[j];
        }

        if (j == 0) {
          board.l[i] = tmp;
        } else {
          board.l[i] = board.l[i].mul(tmp).mod();
        }
      }
      //
      board.r[i] = board.l[i].inv();
    }

    // compute a * s + z.
    board.l = addScalar(times(board.l, rangeProof.ipProof.a), board.z);
    // compute y^-n * (b*s^-1 - (z*y^n+z^2*2^n)).
    board.r = times(board.r, rangeProof.ipProof.b);
    board.exp = addFieldVector(times(board.yn, board.z), times(board.n2, board.zSquare));
    board.r = subFieldVector(board.r, board.exp);
    board.r = multFieldVector(powers(board.y.neg()), board.r);

    // xu=19285133356664502155487922239640858641706256654660833335164497227122088314101
    // Test(x: 19285133356664502155487922239640858641706256654660833335164497227122088314101 (uint256), y: 0 (uint256))
    uint xu = uint(keccak256(abi.encodePacked(rangeProof.t))).mod();
    emit Test(xu, 0);


    // commit.
    board.actual = commit(gv, board.l).add(commit(hv, board.r)).add(uBase.mul(xu.mul(rangeProof.ipProof.a.mul(rangeProof.ipProof.b).sub(rangeProof.t)))).add(h.mul(rangeProof.u));
    emit Test(board.actual.X, board.actual.Y);
    emit Test(p.X, p.Y);

    return board.actual.X == p.X && board.actual.Y == p.Y;
  }

  function verifyIPInternal(BN128.G1Point[bitSize] memory gv, BN128.G1Point[bitSize] memory hv, BN128.G1Point memory p, RangeProof memory rangeProof) internal returns(bool) {
    return ipVerifier.verifyIPProofWithCustomParams(gv, hv, p, rangeProof.t, rangeProof.ipProof.l, rangeProof.ipProof.r, rangeProof.ipProof.a, rangeProof.ipProof.b);
  }

  function computeChallenge(uint a, uint b, uint c, uint d) internal pure returns(uint) {
    return uint(keccak256(abi.encodePacked(a, b, c, d))).mod();
  }

  function computeChallenge2(uint a) internal pure returns(uint) {
    return uint(keccak256(abi.encodePacked(a))).mod();
  }

  /*
   * @dev compute [1, base, base^2, ... , base^(bitSize-1)]
   */
  function powers(uint256 base) internal pure returns (uint256[bitSize] memory powers) {
        powers[0] = 1;
        powers[1] = base;
        for (uint256 i = 2; i < bitSize; i++) {
          powers[i] = powers[i-1].mul(base).mod();
        }
    }

  /*
   * @dev sum []
   */
  function sum(uint256[bitSize] memory data) internal view returns(uint) {
    uint res = data[0];
    for (uint i = 1; i < bitSize; i++) {
      res = res.add(data[i]);
    }

    return res;
  }

  /*
   * @dev modInverse return (a1.inv, a2.inv, ..., an.inv)
   */
  function modInverse(uint[bitSize] memory a) internal view returns(uint[bitSize] memory) {
    uint[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = a[i].inv();
    }

    return res;
  }

  /*
   * @dev hadamard compute (h1, h2, ..., hn) * (a1, a2, ..., an) = (h1*a1, h2*a2, ..., hn*an)
   */
  function hadamard(BN128.G1Point[bitSize] memory m, uint[bitSize] memory a) internal view returns(BN128.G1Point[bitSize] memory) {
    BN128.G1Point[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = m[i].mul(a[i]);
    }

    return res;
  }

  /*
   * @dev sum vector.
   */
  function sumVector(BN128.G1Point[bitSize] memory v) internal view returns(BN128.G1Point memory) {
    BN128.G1Point memory res = v[0];
    for (uint i = 1; i < bitSize; i++) {
      res = res.add(v[i]);
    }

    return res;
  }

  /*
   * @dev commit compute (h1, h2, ..., hn) * (a1, a2, ..., an) == h1*a1 + h2*a2 + ... + hn*an.
   */
  function commit(BN128.G1Point[bitSize] memory vector, uint[bitSize] memory scalar) internal view returns(BN128.G1Point memory) {
    BN128.G1Point memory res = vector[0].mul(scalar[0]);
    for (uint i = 1; i < bitSize; i++) {
      res = res.add(vector[i].mul(scalar[i]));
    }

    return res;
  }

  /*
   * @dev (m1, m2, ..., mn) * scalar == (m1*scalar, m2*scalar, ..., mn*scalar)
   */
  function times(uint[bitSize] memory m, uint scalar) internal view returns(uint[bitSize] memory) {
    uint[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = m[i].mul(scalar);
    }

    return res;
  }


  /**
   * @dev (m1, m2, ..., mn) + scalar = (m1+z, m2+z, ..., mn+z)
   */
  function addScalar(uint[bitSize] memory m, uint scalar) internal view returns(uint[bitSize] memory) {
    uint[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = res[i].add(scalar).mod();
    }

    return res;
  }

  /*
   * @dev add field vector(a1, ..., an) + (b1, ..., bn) == (a1+b1, ..., an+bn)
   */
  function addFieldVector(uint[bitSize] memory a, uint[bitSize] memory b) internal view returns(uint[bitSize] memory) {
    uint[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = a[i].add(b[i]);
    }

    return res;
  }

  /*
   *
   */
  function subFieldVector(uint[bitSize] memory a, uint[bitSize] memory b) internal view returns(uint[bitSize] memory) {
    uint[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = a[i].add(b[i].neg());
    }

    return res;
  }

  /*
   *
   */
  function multFieldVector(uint[bitSize] memory a, uint[bitSize] memory b) internal view returns(uint[bitSize] memory) {
    uint[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = a[i].mul(b[i]);
    }

    return res;
  }

  function smallParseBinary(uint n, uint j, uint size) internal view returns(bool) {
    uint w = 1 << (size - 1);

    for (uint i=0; i < j; i++) {
      w = w >> 1;
    }

    if ((n&w) != 0) {
      return true;
    }

    return false;
  }
}
