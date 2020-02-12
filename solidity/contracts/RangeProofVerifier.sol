pragma solidity >= 0.5.0 < 0.6.0;

import "./library/BN128.sol";
import "./PublicParams.sol";
import "./IPVerifier.sol";

contract RangeProofVerifier {
  using BN128 for BN128.G1Point;
  using BN128 for uint;

   // bit size of value.
  uint public constant bitSize = 32;
  // number of inner product points.
  uint public constant n = 5;

  // public params contract.
  PublicParams public params;

  // inner product verifier.
  IPVerifier public ipVerifier;

  // g point.
  BN128.G1Point public g;
  // h point.
  BN128.G1Point public h;
  // uBase point.
  BN128.G1Point public uBase;
  // g vector.
  BN128.G1Point[bitSize] public gvBase;
  // h vector.
  BN128.G1Point[bitSize] public hvBase;

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
    uint[2*n] l;
    uint[2*n] r;
    uint[] ll;
    uint[] rr;
    uint a;
    uint b;
  }

  // struct for tmp calculation.
  struct Board {
    // to call ip verifier.
    uint[] gvs;
    uint[] hvs;
    uint[2*n] ls;
    uint[2*n] rs;
    BN128.G1Point tmp;
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
    uint[n] challengesSquare;
    uint[n] challengesSquareInverse;
    uint[bitSize] l;
    uint[bitSize] tl;
    uint[bitSize] r;
    uint[bitSize] tr;
  }

  constructor(address params_, address ip) public {
    params = PublicParams(params_);
    ipVerifier = IPVerifier(ip);
    require(bitSize == params.getBitSize(), "bis size not equal");
    require(n == params.getN(), "number of l,r not equal");

    // set public params.
    uint[2] memory tmpG = params.getG();
    uint[2] memory tmpH = params.getH();
    uint[2] memory tmpU = params.getU();
    uint[2*bitSize] memory gv = params.getGVector();
    uint[2*bitSize] memory hv = params.getHVector();
    g.X = tmpG[0];
    g.Y = tmpG[1];
    h.X = tmpH[0];
    h.Y = tmpH[1];
    uBase.X = tmpU[0];
    uBase.Y = tmpU[1];

    for (uint i = 0; i < bitSize; i++) {
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
   * points[8-9]: V(commitment).
   * scalar[0]: t.
   * scalar[1]: tx.
   * scalar[2]: u.
   * scalar[3]: a.
   * scalar[4]: b.
   * l[0-1]: l.x,l.y.
   * r[0-1]: r.x,r.y.
   */
  function optimizedVerifyRangeProof(uint[10] memory points, uint[5] memory scalar, uint[2*n] memory l, uint[2*n] memory r) public view returns(bool) {
    RangeProof memory rangeProof;
    rangeProof.A = BN128.G1Point(points[0], points[1]);
    rangeProof.S = BN128.G1Point(points[2], points[3]);
    rangeProof.T1 = BN128.G1Point(points[4], points[5]);
    rangeProof.T2 = BN128.G1Point(points[6], points[7]);
    rangeProof.t = scalar[0];
    rangeProof.txx = scalar[1];
    rangeProof.u = scalar[2];

    rangeProof.ipProof.l = l;
    rangeProof.ipProof.r = r;
    rangeProof.ipProof.a = scalar[3];
    rangeProof.ipProof.b = scalar[4];

    return optimizedVerify(BN128.G1Point(points[8], points[9]), rangeProof);
  }


  function verifyRangeProof(uint[10] memory points, uint[5] memory scalar, uint[] memory l, uint[] memory r) public view returns(bool) {
    RangeProof memory rangeProof;
    rangeProof.A = BN128.G1Point(points[0], points[1]);
    rangeProof.S = BN128.G1Point(points[2], points[3]);
    rangeProof.T1 = BN128.G1Point(points[4], points[5]);
    rangeProof.T2 = BN128.G1Point(points[6], points[7]);
    rangeProof.t = scalar[0];
    rangeProof.txx = scalar[1];
    rangeProof.u = scalar[2];

    rangeProof.ipProof.ll = l;
    rangeProof.ipProof.rr = r;
    rangeProof.ipProof.a = scalar[3];
    rangeProof.ipProof.b = scalar[4];

    return verify(BN128.G1Point(points[8], points[9]), rangeProof);
  }

  // function getRangeProof()

  /*
   *
   */
  function verify(BN128.G1Point memory v, RangeProof memory rangeProof) internal view returns(bool) {
    BN128.G1Point[bitSize] memory gv = getGV();
    BN128.G1Point[bitSize] memory hv = getHV();
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


  /*
   * @dev 2*(bitSize+n)+14 mul, 2*(bitSize+n)+9 add.
   */
  function optimizedVerify(BN128.G1Point memory v, RangeProof memory rangeProof) internal view returns(bool) {
    BN128.G1Point[bitSize] memory gv = getGV();
    BN128.G1Point[bitSize] memory hv = getHV();
    Board memory board;
    // compute
    board.y = computeChallenge(rangeProof.A.X, rangeProof.A.Y, rangeProof.S.X, rangeProof.S.Y);
    board.z = computeChallenge2(board.y);
    board.yn = powers(board.y);
    board.ynInverse = powers(board.y.inv());
    board.zNeg = board.z.neg();
    board.zSquare = board.z.mul(board.z).mod();
    board.n2 = powers(2);
    // 9 mul, 6 add.
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

    // 1 add, 1 mul.
    BN128.G1Point memory p;
    p = rangeProof.A.add(rangeProof.S.mul(board.x));

    // compute formula on the right.
    // compute p + li * xi^2 + ri * xi^-2.
    // n*2 mul, n*2 add.
    for (uint i = 0; i < n; i++) {
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
    for (uint i = 0; i < bitSize; i++) {

      if (i == 0) {
        for (uint j = 0; j < n; j++) {
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
        uint k = getBiggestPos(i, n);

        // tl, tr should not be changed.
        board.tl[i] = board.tl[i-pow(k-1)].mul(board.challengesSquare[n-k]).mod();
        board.tr[i] = board.tr[i-pow(k-1)].mul(board.challengesSquareInverse[n-k]).mod();
      }

      board.l[i] = board.tl[i];

      // set si and si^-1.
      board.r[i] = board.tr[i];

      board.l[i] = board.l[i].mul(rangeProof.ipProof.a).add(board.z);
      board.r[i] = board.r[i].mul(rangeProof.ipProof.b);
      board.r[i] = board.r[i].sub(board.zSquare.mul(board.n2[i]));
      board.r[i] = board.r[i].mul(board.ynInverse[i]).sub(board.z);
    }

    uint xu = uint(keccak256(abi.encodePacked(rangeProof.t))).mod();

    // commit.
    // 2*bitSize+4 mul, 2*bitSize+2 add.
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

  function verifyIPInternal(BN128.G1Point[bitSize] memory gv, BN128.G1Point[bitSize] memory hv, BN128.G1Point memory p, RangeProof memory rangeProof) internal view returns(bool) {
    Board memory b;
    b.gvs = toUintArray(gv);
    b.hvs = toUintArray(hv);
    uint[2] memory pp;
    pp[0] = p.X;
    pp[1] = p.Y;
    uint[2] memory u;
    u[0] = uBase.X;
    u[1] = uBase.Y;

    return ipVerifier.optimizedVerifyIPProof(b.gvs, b.hvs, pp, u, rangeProof.t, rangeProof.ipProof.ll, rangeProof.ipProof.rr, rangeProof.ipProof.a, rangeProof.ipProof.b);
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
  function powers(uint256 base) internal pure returns (uint256[bitSize] memory powersRes) {
        powersRes[0] = 1;
        powersRes[1] = base;
        for (uint256 i = 2; i < bitSize; i++) {
          powersRes[i] = powersRes[i-1].mul(base).mod();
        }
    }

  /*
   * @dev sum []
   */
  function sum(uint256[bitSize] memory data) internal pure returns(uint) {
    uint res = data[0];
    for (uint i = 1; i < bitSize; i++) {
      res = res.add(data[i]);
    }

    return res;
  }

  /*
   * @dev modInverse return (a1.inv, a2.inv, ..., an.inv)
   */
  function modInverse(uint[bitSize] memory a) internal pure returns(uint[bitSize] memory) {
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
   * @dev bitSize mul, bitSize-1 add.
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
  function times(uint[bitSize] memory m, uint scalar) internal pure returns(uint[bitSize] memory) {
    uint[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = m[i].mul(scalar);
    }

    return res;
  }


  /*
   * @dev add field vector(a1, ..., an) + (b1, ..., bn) == (a1+b1, ..., an+bn)
   */
  function addFieldVector(uint[bitSize] memory a, uint[bitSize] memory b) internal pure returns(uint[bitSize] memory) {
    uint[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = a[i].add(b[i]);
    }

    return res;
  }

  /*
   *
   */
  function subFieldVector(uint[bitSize] memory a, uint[bitSize] memory b) internal pure returns(uint[bitSize] memory) {
    uint[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = a[i].add(b[i].neg());
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
  function multFieldVector(uint[bitSize] memory a, uint[bitSize] memory b) internal pure returns(uint[bitSize] memory) {
    uint[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i] = a[i].mul(b[i]);
    }

    return res;
  }

  function smallParseBinary(uint t, uint j, uint size) internal pure returns(bool) {
    uint w = 1 << (size - 1);

    for (uint i = 0; i < j; i++) {
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

  /*
   *
   */
  function getGV() internal view returns(BN128.G1Point[bitSize] memory) {
    BN128.G1Point[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i].X = gvBase[i].X;
      res[i].Y = gvBase[i].Y;
    }

    return res;
  }

  /*
   *
   */
  function getHV() internal view returns(BN128.G1Point[bitSize] memory) {
    BN128.G1Point[bitSize] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[i].X = hvBase[i].X;
      res[i].Y = hvBase[i].Y;
    }

    return res;
  }

  function toUintArray(BN128.G1Point[bitSize] memory points) internal pure returns(uint[] memory) {
    uint[] memory res = new uint[](2*bitSize);
    for (uint i = 0; i < bitSize; i++) {
      res[2*i] = points[i].X;
      res[2*i+1] = points[i].Y;
    }

    return res;
  }
}
