pragma solidity >= 0.5.0 < 0.6.0;

import "./library/BN128.sol";
import "./PublicParams.sol";

contract AggRangeProofVerifier {
  using BN128 for BN128.G1Point;
  using BN128 for uint;

   // bit size of value.
  uint constant bitSize = 64;
  // number of inner product points.
  uint constant n = 6;
  // agg size.
  uint constant aggSize = 2;
  // 2^step = aggSize
  uint constant step = 1;
  // vector size.
  uint constant vectorSize = bitSize*aggSize;
  // proof l, r size.
  uint constant lrSize = n+step;

  // public params contract.
  PublicParams public params;

  // g point.
  BN128.G1Point public g;
  // h point.
  BN128.G1Point public h;
  // uBase point.
  BN128.G1Point public uBase;
  // g vector.
  BN128.G1Point[vectorSize] public gvBase;
  // h vector.
  BN128.G1Point[vectorSize] public hvBase;

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
    uint[lrSize*2] l;
    uint[lrSize*2] r;
    uint a;
    uint b;
  }

  // struct for tmp calculation.
  struct Board {
    // to call ip verifier.
    uint[2*vectorSize] gvs;
    uint[2*vectorSize] hvs;
    BN128.G1Point tmp;
    uint x;
    uint x2;
    uint y;
    uint[vectorSize] yn;
    uint[vectorSize] ynInverse;
    uint z;
    uint zSquare;
    uint zCubed;
    uint zNeg;
    uint[bitSize] n2;
    BN128.G1Point[bitSize] hPrime;
    BN128.G1Point expect;
    BN128.G1Point actual;
    uint dleta;
    uint[lrSize] challenges;
    uint[lrSize] challengesInverse;
    uint[lrSize] challengesSquare;
    uint[lrSize] challengesSquareInverse;
    uint[vectorSize] l;
    uint[vectorSize] tl;
    uint[vectorSize] r;
    uint[vectorSize] tr;
  }

  // 64位，该步骤会超出8000000限制
  constructor(address params_) public {
    // revert("test");
    params = PublicParams(params_);
    

    require(bitSize == params.getBitSize(), "bis size not equal");
    

    // set public params.
    uint[2] memory tmpG = params.getG();
    uint[2] memory tmpH = params.getH();
    uint[2] memory tmpU = params.getU();
    
    uint[2*vectorSize] memory gv = params.getAggGVector();
    uint[2*vectorSize] memory hv = params.getAggHVector();
    g.X = tmpG[0];
    g.Y = tmpG[1];
    h.X = tmpH[0];
    h.Y = tmpH[1];
    uBase.X = tmpU[0];
    uBase.Y = tmpU[1];

    for (uint i = 0; i < vectorSize; i++) {
      gvBase[i].X = gv[2*i];
      gvBase[i].Y = gv[2*i+1];
      hvBase[i].X = hv[2*i];
      hvBase[i].Y = hv[2*i+1];
    }
  }

  /*
   * @dev verify range proof valid or not.
   * points[0-1]: A.
   * points[2-3]: S.
   * points[4-5]: T1.
   * points[6-7]: T2.
   * points[8-9]: V(commitment, 1).
   * points[10-11]: V(commitment, 2).
   * scalar[0]: t.
   * scalar[1]: tx.
   * scalar[2]: u.
   * scalar[3]: a.
   * scalar[4]: b.
   * l[0-11]: l.x,l.y.
   * r[0-11]: r.x,r.y.
   */
  function aggVerifyRangeProof(uint[10+aggSize] memory points, uint[5] memory scalar, uint[2*lrSize] memory l, uint[2*lrSize] memory r) public view returns(bool) {
    RangeProof memory rangeProof;
    rangeProof.A = BN128.G1Point(points[0], points[1]);
    rangeProof.S = BN128.G1Point(points[2], points[3]);
    rangeProof.T1 = BN128.G1Point(points[4], points[5]);
    rangeProof.T2 = BN128.G1Point(points[6], points[7]);
    rangeProof.t = scalar[0];
    rangeProof.txx = scalar[1];
    rangeProof.u = scalar[2];

    BN128.G1Point[2] memory commits;
    commits[0] = BN128.G1Point(points[8], points[9]);
    commits[1] = BN128.G1Point(points[10], points[11]);

    rangeProof.ipProof.l = l;
    rangeProof.ipProof.r = r;
    rangeProof.ipProof.a = scalar[3];
    rangeProof.ipProof.b = scalar[4];

    return optimizedVerify(commits, rangeProof);
  }


  /*
   * @dev 
   */
  function optimizedVerify(BN128.G1Point[2] memory v, RangeProof memory rangeProof) internal view returns(bool) {
    BN128.G1Point[vectorSize] memory gv = getGV();
    BN128.G1Point[vectorSize] memory hv = getHV();
    Board memory board;
    // compute
    board.y = computeChallenge(rangeProof.A.X, rangeProof.A.Y, rangeProof.S.X, rangeProof.S.Y);
    board.z = computeChallenge(rangeProof.S.X, rangeProof.S.Y, rangeProof.A.X, rangeProof.A.Y);
    board.yn = powers(board.y);
    board.ynInverse = powers(board.y.inv());
    board.zNeg = board.z.neg();
    board.zSquare = board.z.mul(board.z).mod();
    board.zCubed = board.zSquare.mul(board.z).mod();
    board.n2 = powersBitSize(2);
    // 9 mul, 6 add.
    board.x = computeChallenge(rangeProof.T1.X, rangeProof.T1.Y, rangeProof.T2.X, rangeProof.T2.Y);
    board.x2 = board.x.mul(board.x);

    // check g*tx + h*t ?= v*(z^2 * z^m) + g*dleta + T1*x + T2*x^2. (z^m is a vector)
    // check g*(tx-dleta) + h*t ?= v*(z^2 * z^m) + T1*x + T2*x^2.
    // 6 mul. 4 add.
    board.expect = v[0].mul(board.zSquare).add(v[1].mul(board.zCubed)).add(rangeProof.T1.mul(board.x)).add(rangeProof.T2.mul(board.x2));
    // delta = (z - z^2) * <1^mn, y^mn> - z^(j+2) * <1^n, 2^n>. (j is [1, m])
    board.dleta = board.z.sub(board.zSquare).mul(sum(board.yn)).sub(board.zCubed.mul(sumBitSize(board.n2)));
    board.dleta = board.dleta.sub(board.zCubed.mul(board.z).mul(sumBitSize(board.n2)));
    board.actual = g.mul(rangeProof.t.sub(board.dleta)).add(h.mul(rangeProof.txx));
    if (board.expect.X != board.actual.X || board.expect.Y != board.actual.Y) {
      return false;
    }

    // 1 add, 1 mul.
    BN128.G1Point memory p;
    p = rangeProof.A.add(rangeProof.S.mul(board.x));

    // compute formula on the right.
    // compute p + li * xi^2 + ri * xi^-2.
    // lrSize*2 mul, lrSize*2 add.
    for (uint i = 0; i < lrSize; i++) {
      uint x = computeChallenge(rangeProof.ipProof.l[2*i], rangeProof.ipProof.l[2*i+1], rangeProof.ipProof.r[2*i], rangeProof.ipProof.r[2*i+1]);
      board.challenges[i] = x;
      board.challengesSquare[i] = x.mul(x).mod();
      board.challengesSquareInverse[i] = board.challengesSquare[i].inv();

      board.tmp = BN128.G1Point(rangeProof.ipProof.l[2*i], rangeProof.ipProof.l[2*i+1]);
      p = p.add(board.tmp.mul(board.challengesSquare[i]));
      board.tmp = BN128.G1Point(rangeProof.ipProof.r[2*i], rangeProof.ipProof.r[2*i+1]);
      p = p.add(board.tmp.mul(board.challengesSquareInverse[i]));
    }

    // scalar mul, add.
    for (uint i = 0; i < vectorSize; i++) {
      if (i == 0) {
        for (uint j = 0; j < lrSize; j++) {
          uint tmp = board.challenges[j];
          if (j == 0) {
            board.tl[i] = tmp;
          } else {
            board.tl[i] = board.tl[i].mul(tmp).mod();
          }
        }

        board.tr[i] = board.tl[i];
        board.tl[i] = board.tl[i].inv();
      } else {
        // i is start from 0.
        // 5 >= k >= 1.
        uint k = getBiggestPos(i, lrSize);

        // tl, tr should not be changed.
        board.tl[i] = board.tl[i-pow(k-1)].mul(board.challengesSquare[lrSize-k]).mod();
        board.tr[i] = board.tr[i-pow(k-1)].mul(board.challengesSquareInverse[lrSize-k]).mod();
      }

      board.l[i] = board.tl[i];

      // set si and si^-1.
      board.r[i] = board.tr[i];

      board.l[i] = board.l[i].mul(rangeProof.ipProof.a).add(board.z);
      board.r[i] = board.r[i].mul(rangeProof.ipProof.b);
      if (i < bitSize) {
        board.r[i] = board.r[i].sub(board.zSquare.mul(board.n2[i]));
      } else {
        board.r[i] = board.r[i].sub(board.zCubed.mul(board.n2[i%bitSize]));
      }
      board.r[i] = board.r[i].mul(board.ynInverse[i]).sub(board.z);
    }

    uint xu = uint(keccak256(abi.encodePacked(rangeProof.t))).mod();

    // commit.
    // a: commit: 2*(vectorSize mul, vectorSize - 1 add).
    // b: normal: 3 add, 2 mul.
    board.actual = commit(gv, board.l).add(commit(hv, board.r)).add(uBase.mul(xu.mul(rangeProof.ipProof.a.mul(rangeProof.ipProof.b).sub(rangeProof.t)))).add(h.mul(rangeProof.u));

    // return true;
    return board.actual.X == p.X && board.actual.Y == p.Y;
  }

  function pow(uint kk) internal pure returns(uint) {
    uint i = kk;
    if (i == 0) {
      return 1;
    }
    uint res = 1;
    while(i > 0) {
      res = res * 2;
      i--;
    }

    return res;
  }

  function computeChallenge(uint a, uint b, uint c, uint d) internal pure returns(uint) {
    return uint(keccak256(abi.encodePacked(a, b, c, d))).mod();
  }

  /*
   * @dev compute [1, base, base^2, ... , base^(bitSize-1)]
   */
  function powers(uint256 base) internal pure returns (uint256[vectorSize] memory powersRes) {
        powersRes[0] = 1;
        powersRes[1] = base;
        for (uint256 i = 2; i < vectorSize; i++) {
          powersRes[i] = powersRes[i-1].mul(base).mod();
        }
  }

  /*
   * @dev compute [1, base, base^2, ... , base^(bitSize-1)]
   */
  function powersBitSize(uint256 base) internal pure returns (uint256[bitSize] memory powersRes) {
        powersRes[0] = 1;
        powersRes[1] = base;
        for (uint256 i = 2; i < bitSize; i++) {
          powersRes[i] = powersRes[i-1].mul(base).mod();
        }
  }

  /*
   * @dev sum []
   */
  function sumBitSize(uint256[bitSize] memory data) internal pure returns(uint) {
    uint res = data[0];
    for (uint i = 1; i < bitSize; i++) {
      res = res.add(data[i]);
    }

    return res;
  }

  function sum(uint256[vectorSize] memory data) internal pure returns(uint) {
    uint res = data[0];
    for (uint i = 1; i < vectorSize; i++) {
      res = res.add(data[i]);
    }

    return res;
  }

  function commit(BN128.G1Point[vectorSize] memory vector, uint[vectorSize] memory scalar) internal view returns(BN128.G1Point memory) {
    BN128.G1Point memory res = vector[0].mul(scalar[0]);
    for (uint i = 1; i < vectorSize; i++) {
      res = res.add(vector[i].mul(scalar[i]));
    }

    return res;
  }

  /*
   * @dev commit compute (h1, h2, ..., hn) * (a1, a2, ..., an) == h1*a1 + h2*a2 + ... + hn*an.
   * @dev bitSize mul, bitSize-1 add.
   */
  function commit(BN128.G1Point[bitSize] memory vector, uint[bitSize] memory scalar) internal view returns(BN128.G1Point memory) {
    BN128.G1Point memory res = vector[0].mul(scalar[0]);
    for (uint i = 1; i < bitSize; i++) {
      res = res.add(vector[i].mul(scalar[i]));
    }

    return res;
  }

  function getBiggestPos(uint i, uint s) internal pure returns(uint) {
    uint l = 1 << s;
    uint calTimes;
    while (i < l && l > 0) {
      l = l >> 1;
      calTimes++;
    }
    return 1+s-calTimes;
  }
  /*
   *
   */
  function getGV() internal view returns(BN128.G1Point[vectorSize] memory) {
    BN128.G1Point[vectorSize] memory res;
    for (uint i = 0; i < vectorSize; i++) {
      res[i].X = gvBase[i].X;
      res[i].Y = gvBase[i].Y;
    }

    return res;
  }

  /*
   *
   */
  function getHV() internal view returns(BN128.G1Point[vectorSize] memory) {
    BN128.G1Point[vectorSize] memory res;
    for (uint i = 0; i < vectorSize; i++) {
      res[i].X = hvBase[i].X;
      res[i].Y = hvBase[i].Y;
    }

    return res;
  }

  function toUintArray(BN128.G1Point[bitSize] memory points) internal pure returns(uint[2*bitSize] memory) {
    uint[2*bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[2*i] = points[i].X;
      res[2*i+1] = points[i].Y;
    }

    return res;
  }
}
