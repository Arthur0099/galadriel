pragma solidity >= 0.5.0 < 0.6.0;
pragma experimental ABIEncoderV2;

import "./library/BN128.sol";
import "./PublicParams.sol";

// response for verifying all proofs.
contract Verifier {
    using BN128 for BN128.G1Point;
    using BN128 for uint;

    // bitSize of balance value.
    uint public constant bitSize = 64;
    uint public constant n = 6;
    // agg size.
    uint constant aggSize = 2;
    // 2^step = aggSize
    uint constant step = 1;
    // vector size.
    uint constant vectorSize = bitSize*aggSize;
    // proof l, r size.
    uint constant lrSize = n+step;

    // public h point.
    BN128.G1Point public h;

    // public g point.
    BN128.G1Point public g;

    // uBase point.
    BN128.G1Point public uBase;

    // g vector generator used in range proof
    BN128.G1Point[vectorSize] public gVector;
    // h vector generator used in range proof
    BN128.G1Point[vectorSize] public hVector;

    // params.
    PublicParams public params;

    // encrypted ct.(balance of account)
    struct CT {
        BN128.G1Point X;
        BN128.G1Point Y;
    }

    struct MRCT {
        BN128.G1Point X1;
        BN128.G1Point X2;
        BN128.G1Point Y;
    }

    struct Board {
        // for tmp calculation.
        uint y;
        uint z;
        uint zNeg;
        uint zSquare;
        uint zCubed;
        uint x;
        uint x2;
        uint dleta;

        uint[vectorSize] yn;
        uint[vectorSize] ynInverse;
        uint[bitSize] n2;
        uint[lrSize] challenges;
        uint[lrSize] challengesInverse;
        uint[lrSize] challengesSquare;
        uint[lrSize] challengesSquareInverse;
        uint[vectorSize] l;
        uint[vectorSize] tl;
        uint[vectorSize] r;
        uint[vectorSize] tr;

        BN128.G1Point expect;
        BN128.G1Point actual;
        BN128.G1Point tmp;
    }

    struct TransferStatement {
        BN128.G1Point pk1;
        BN128.G1Point pk2;
        MRCT mrct;
        CT updated;
        CT refresh;
        uint[] custom;
    }

    struct TransferProof {
        PTEProof pteProof;
        DLEProof dleProof;
        CTVProof ctvProof;
        AggProof aggProof;
    }

    struct PTEProof {
        BN128.G1Point A1;
        BN128.G1Point A2;
        BN128.G1Point B;
        uint z1;
        uint z2;
    }

    struct DLEProof {
        BN128.G1Point A1;
        BN128.G1Point A2;
        uint z;
    }

    struct CTVProof {
        BN128.G1Point A;
        BN128.G1Point B;
        uint z1;
        uint z2;
    }

    struct AggProof {
        BN128.G1Point A;
        BN128.G1Point S;
        BN128.G1Point T1;
        BN128.G1Point T2;
        uint t;
        uint txx;
        uint u;

        // inner product proof
        BN128.G1Point[] l;
        BN128.G1Point[] r;
        uint ap;
        uint bp;
    }

    constructor(address params_) public {
        params = PublicParams(params_);

        uint[2] memory tmpH = params.getH();
        uint[2] memory tmpG = params.getG();
        uint[2] memory tmpU = params.getU();

        h.X = tmpH[0];
        h.Y = tmpH[1];
        g.X = tmpG[0];
        g.Y = tmpG[1];
        uBase.X = tmpU[0];
        uBase.Y = tmpU[1];
        require(bitSize == params.getBitSize(), "bitsize not equal");

        // 分开初始化, 第一次初始化32个.
        init(32);
    }

    function init(uint step) public {
        uint i = 0;
        for (; i < vectorSize; i++) {
            if (gVector[i].X == 0) {
                break;
            }
        }

        if (i < gVector.length) {
            initVector(i, i+step);
        }
        
    }

    // init public gv/hv of protocol.
    function initVector(uint start, uint end) internal {
        uint scalar;
        BN128.G1Point memory gBase = BN128.G1Point(1, 2);
        
        uint gvb = uint(keccak256(abi.encodePacked("gvs"))).mod();
        for (uint i = start; i < end && i < vectorSize; i++) {
            scalar = uint(keccak256(abi.encodePacked(gvb+i))).mod();
            gVector[i] = gBase.mul(scalar);
        }
        
        uint hvb = uint(keccak256(abi.encodePacked("hvs"))).mod();
        for (uint j = start; j < end && j < vectorSize; j++) {
            scalar = uint(keccak256(abi.encodePacked(j+hvb))).mod();
            hVector[j] = gBase.mul(scalar);
        }
    }

    function verify(
        uint[36] memory points,
        uint[10] memory scalar, 
        uint[2*lrSize] memory l, 
        uint[2*lrSize] memory r,
        uint[4] memory ct,
        uint nonce,
        uint token
    )
        public
        view 
        returns(bool) 
    {
        TransferProof memory proof;
        TransferStatement memory state;
        state.pk1 = BN128.G1Point(points[0], points[1]);
        state.pk2 = BN128.G1Point(points[2], points[3]);
        state.updated.X = BN128.G1Point(ct[0], ct[1]);
        state.updated.Y = BN128.G1Point(ct[2], ct[3]);
        state.refresh.X = BN128.G1Point(points[16], points[17]);
        state.refresh.Y = BN128.G1Point(points[18], points[19]);
        state.mrct.X1 = BN128.G1Point(points[4], points[5]);
        state.mrct.X2 = BN128.G1Point(points[6], points[7]);
        state.mrct.Y = BN128.G1Point(points[8], points[9]);
        state.custom = new uint[](12);
        state.custom[0] = nonce;
        state.custom[1] = token;
        for (uint i = 0; i < 10; i++) {
            state.custom[2+i] = points[i];
        }

        proof.pteProof.A1 = BN128.G1Point(points[10], points[11]);
        proof.pteProof.A2 = BN128.G1Point(points[12], points[13]);
        proof.pteProof.B = BN128.G1Point(points[14], points[15]);
        proof.pteProof.z1 = scalar[0];
        proof.pteProof.z2 = scalar[1];

        proof.dleProof.A1 = BN128.G1Point(points[24], points[25]);
        proof.dleProof.A2 = BN128.G1Point(points[26], points[27]);
        proof.dleProof.z = scalar[4];

        proof.ctvProof.A = BN128.G1Point(points[20], points[21]);
        proof.ctvProof.B = BN128.G1Point(points[22], points[23]);
        proof.ctvProof.z1 = scalar[2];
        proof.ctvProof.z2 = scalar[3];

        proof.aggProof.A = BN128.G1Point(points[28], points[29]);
        proof.aggProof.S = BN128.G1Point(points[30], points[31]);
        proof.aggProof.T1 = BN128.G1Point(points[32], points[33]);
        proof.aggProof.T2 = BN128.G1Point(points[34], points[35]);
        proof.aggProof.t = scalar[5];
        proof.aggProof.txx = scalar[6];
        proof.aggProof.u = scalar[7];
        proof.aggProof.ap = scalar[8];
        proof.aggProof.bp = scalar[9];

        proof.aggProof.l = new BN128.G1Point[](lrSize);
        proof.aggProof.r = new BN128.G1Point[](lrSize);
        for (uint i = 0; i < lrSize; i++) {
            proof.aggProof.l[i] = BN128.G1Point(l[i*2], l[2*i+1]);
            proof.aggProof.r[i] = BN128.G1Point(r[i*2], r[2*i+1]);
        }

        return verifyAggTransfer(proof, state);
    }

    function verifyAggTransfer(TransferProof memory proof, TransferStatement memory state) public view returns(bool) {
        
        // verify pte proof.
        // pk1, pk2, ct enc. pte proof.
        // 7 mul, 4 add.
        require(verifyPTEProof(state.pk1, state.pk2, state.mrct, proof.pteProof), "pte equal proof invalid");
        
        // 
        require(verifyDLEProof(state.updated, state.refresh, state.pk1, proof.dleProof, state.custom), "dle sigma proof failed");

        // 
        require(verifyCTValidProof(state.pk1, state.refresh, proof.ctvProof), "ct valid proof invalid");

        // verify agg range proof.
        BN128.G1Point[2] memory commit;
        commit[0] = state.mrct.Y;
        commit[1] = state.refresh.Y;
        require(verifyAggRangeProof(commit, proof.aggProof), "aggrate range proof invalid");

        return true;
    }

    function verifyBurn(uint amount, uint[2] memory publicKey, uint[4] memory proof, uint z, uint[4] memory ub, uint[] memory input) public view returns(bool) {
        CT memory userBalance;
        userBalance.X.X = ub[0];
        userBalance.X.Y = ub[1];
        userBalance.Y.X = ub[2];
        userBalance.Y.Y = ub[3];

        DLEProof memory proofins;
        proofins.A1 = BN128.G1Point(proof[0], proof[1]);
        proofins.A2 = BN128.G1Point(proof[2], proof[3]);
        proofins.z = z;

        // Board memory board;
        require(verifyEqualProof(amount, userBalance, BN128.G1Point(publicKey[0], publicKey[1]), proofins, input), "equal proof verify failed");

        return true;
    }

    function verifyEqualProof(uint amount, CT memory ct, BN128.G1Point memory pk, DLEProof memory proof,uint[] memory input) internal view returns(bool) {
        BN128.G1Point memory g1 = ct.Y.add(g.mul(amount).neg());
        BN128.G1Point memory h1 = ct.X;
        BN128.G1Point memory g2 = h;
        BN128.G1Point memory h2 = pk;

        return checkDLEProof(g1, h1, g2, h2, proof, input);
    }

    function verifyAggRangeProof(BN128.G1Point[2] memory v, AggProof memory proof) internal view returns(bool) {
        BN128.G1Point[vectorSize] memory gv = gVector;
        BN128.G1Point[vectorSize] memory hv = hVector;
        Board memory board;

        // compute
        board.y = uint(keccak256(abi.encodePacked(proof.A.X, proof.A.Y, proof.S.X, proof.S.Y))).mod();
        board.z = uint(keccak256(abi.encodePacked(proof.S.X, proof.S.Y, proof.A.X, proof.A.Y))).mod();
        board.yn = powers(board.y);
        board.ynInverse = powers(board.y.inv());
        board.zNeg = board.z.neg();
        board.zSquare = board.z.mul(board.z).mod();
        board.zCubed = board.zSquare.mul(board.z).mod();
        board.n2 = powersBitSize(2);
        // 9 mul, 6 add.
        board.x = uint(keccak256(abi.encodePacked(proof.T1.X, proof.T1.Y, proof.T2.X, proof.T2.Y))).mod();
        board.x2 = board.x.mul(board.x);

        // check g*tx + h*t ?= v*(z^2 * z^m) + g*dleta + T1*x + T2*x^2. (z^m is a vector)
        // check g*(tx-dleta) + h*t ?= v*(z^2 * z^m) + T1*x + T2*x^2.
        // 6 mul. 4 add.
        board.expect = v[0].mul(board.zSquare).add(v[1].mul(board.zCubed)).add(proof.T1.mul(board.x)).add(proof.T2.mul(board.x2));
        // delta = (z - z^2) * <1^mn, y^mn> - z^(j+2) * <1^n, 2^n>. (j is [1, m])
        board.dleta = board.z.sub(board.zSquare).mul(sum(board.yn)).sub(board.zCubed.mul(sumBitSize(board.n2)));
        board.dleta = board.dleta.sub(board.zCubed.mul(board.z).mul(sumBitSize(board.n2)));
        board.actual = g.mul(proof.t.sub(board.dleta)).add(h.mul(proof.txx));
        if (!board.expect.eq(board.actual)) {
            return false;
        }

        // 1 add, 1 mul.
        BN128.G1Point memory p;
        p = proof.A.add(proof.S.mul(board.x));

        // compute formula on the right.
        // compute p + li * xi^2 + ri * xi^-2.
        // lrSize*2 mul, lrSize*2 add.
        for (uint i = 0; i < proof.l.length; i++) {
            uint x = uint(keccak256(abi.encodePacked(
                proof.l[i].X, proof.l[i].Y,
                proof.r[i].X, proof.r[i].Y
            ))).mod();
            board.challenges[i] = x;
            board.challengesSquare[i] = x.mul(x).mod();
            board.challengesSquareInverse[i] = board.challengesSquare[i].inv();

            board.tmp = proof.l[i];
            p = p.add(board.tmp.mul(board.challengesSquare[i]));
            board.tmp = proof.r[i];
            p = p.add(board.tmp.mul(board.challengesSquareInverse[i]));
        }

        // scalar mul, add.
        for (uint i = 0; i < gv.length; i++) {
            if (i == 0) {
                for (uint j = 0; j < proof.l.length; j++) {
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
                uint k = getBiggestPos(i, proof.l.length);

                // tl, tr should not be changed.
                board.tl[i] = board.tl[i - pow(k - 1)].mul(board.challengesSquare[proof.l.length - k]).mod();
                board.tr[i] = board.tr[i - pow(k - 1)].mul(board.challengesSquareInverse[proof.l.length - k]).mod();
            }

            board.l[i] = board.tl[i];

            // set si and si^-1.
            board.r[i] = board.tr[i];

            board.l[i] = board.l[i].mul(proof.ap).add(board.z);
            board.r[i] = board.r[i].mul(proof.bp);
            if (i < bitSize) {
                board.r[i] = board.r[i].sub(board.zSquare.mul(board.n2[i]));
            } else {
                board.r[i] = board.r[i].sub(board.zCubed.mul(board.n2[i % bitSize]));
            }

            board.r[i] = board.r[i].mul(board.ynInverse[i]).sub(board.z);
        }

        uint xu = uint(keccak256(abi.encodePacked(proof.t))).mod();

        // commit.
        // a: commit: 2*(vectorSize mul, vectorSize - 1 add).
        // b: normal: 3 add, 2 mul.
        board.actual = commit(gv, board.l).add(commit(hv, board.r)).add(uBase.mul(xu.mul(proof.ap.mul(proof.bp).sub(proof.t)))).add(h.mul(proof.u));

        return board.actual.X == p.X && board.actual.Y == p.Y;
    }

    // 5 mul, 3 add.
    function verifyCTValidProof(
        BN128.G1Point memory pk,
        CT memory refresh,
        CTVProof memory proof
    )
        internal
        view
        returns(bool) 
    {
        uint e = uint(keccak256(abi.encodePacked(proof.A.X, proof.A.Y, proof.B.X, proof.B.Y))).mod();

        // check pk*z1 = A + X*e.
        // 2 mul, 1 add.
        BN128.G1Point memory tmp = pk.mul(proof.z1);
        BN128.G1Point memory actual = proof.A.add(refresh.X.mul(e));
        if (!tmp.eq(actual)) {
            return false;
        }

        // check g*z2 + h*z1 = B + Y*e.
        // 3 mul, 2 add.
        tmp = g.mul(proof.z2).add(h.mul(proof.z1));
        actual = proof.B.add(refresh.Y.mul(e));

        return tmp.eq(actual);
    }

    // 7 mul, 4 add.
    function verifyPTEProof(
        BN128.G1Point memory pk1,
        BN128.G1Point memory pk2,
        MRCT memory ct,
        PTEProof memory proof
    )
        internal
        view
        returns(bool) 
    {
        uint e = uint(keccak256(abi.encodePacked(proof.A1.X, proof.A1.Y, proof.A2.X, proof.A2.Y, proof.B.X, proof.B.Y))).mod();

        // 2 mul, 1 add.
        // check pk1 * z1 == A1 + X1 * e.
        BN128.G1Point memory tmp = pk1.mul(proof.z1);
        BN128.G1Point memory actual = ct.X1.mul(e).add(proof.A1);
        if (!tmp.eq(actual)) {
            return false;
        }

        // 2 mul, 1 add.
        // check pk2 * z1 == A2 + X2 * e.
        tmp = pk2.mul(proof.z1);
        actual = ct.X2.mul(e).add(proof.A2);
        if (!tmp.eq(actual)) {
            return false;
        }

        // 3 mul, 2 add.
        // Check h * z1 + g * z2 == B + Y * e.
        tmp = h.mul(proof.z1).add(g.mul(proof.z2));
        actual = proof.B.add(ct.Y.mul(e));

        return tmp.eq(actual);
    }

    function verifyDLEProof(
        CT memory ori,
        CT memory refresh,
        BN128.G1Point memory pk,
        DLEProof memory proof,
        uint[] memory custom
    )
        internal
        view
        returns(bool) 
    {
        BN128.G1Point memory g1 = refresh.Y.add(ori.Y.neg());
        BN128.G1Point memory h1 = refresh.X.add(ori.X.neg());

        return checkDLEProof(g1, h1, h, pk, proof, custom);
    }

    function checkDLEProof(
        BN128.G1Point memory g1,
        BN128.G1Point memory h1,
        BN128.G1Point memory g2,
        BN128.G1Point memory h2,
        DLEProof memory proof,
        uint[] memory custom
    )
        internal
        view
        returns(bool)
    {
        uint eprime = uint(keccak256(abi.encodePacked(proof.A1.X, proof.A1.Y, proof.A2.X, proof.A2.Y))).mod();
        uint hcustom = uint(keccak256(abi.encodePacked(custom))).mod();
        uint e = uint(keccak256(abi.encodePacked(eprime, hcustom))).mod();

        // check g1 * z == A1 + h1 * e.
        if (!checkDLEStep(g1, h1, proof.A1, proof.z, e)) {
            return false;
        }
            // check g2 * z == A2 + h2 * e.
        return checkDLEStep(g2, h2, proof.A2, proof.z, e);
    }

    /*
     * @dev 2 mul, 1 add.
     */
    // check g^z == h*e+A
    function checkDLEStep(
        BN128.G1Point memory gt,
        BN128.G1Point memory ht,
        BN128.G1Point memory A,
        uint z,
        uint e
    )
        internal
        view
        returns(bool)
    {
        gt = gt.mul(z);
        ht = ht.mul(e).add(A);

        return gt.eq(ht);
    }

    function powersBitSize(uint256 base) internal pure returns(uint256[bitSize] memory powersRes) {
        powersRes[0] = 1;
        powersRes[1] = base;
        for (uint256 i = 2; i < bitSize; i++) {
            powersRes[i] = powersRes[i - 1].mul(base).mod();
        }
    }

    /*
     * @dev 
     */
    function sumBitSize(uint256[bitSize] memory data) internal pure returns(uint) {
        uint res = data[0];
        for (uint i = 1; i < bitSize; i++) {
            res = res.add(data[i]);
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
        return 1 + s - calTimes;
    }

    function commit(
        BN128.G1Point[vectorSize] memory vector, 
        uint[vectorSize] memory scalar
    ) 
        internal
        view 
        returns(BN128.G1Point memory) 
    {
        BN128.G1Point memory res = vector[0].mul(scalar[0]);
        for (uint i = 1; i < vector.length; i++) {
            res = res.add(vector[i].mul(scalar[i]));
        }

        return res;
    }

    function powers(uint256 base) internal pure returns (uint256[vectorSize] memory powersRes) {
        powersRes[0] = 1;
        powersRes[1] = base;
        for (uint256 i = 2; i < vectorSize; i++) {
          powersRes[i] = powersRes[i-1].mul(base).mod();
        }
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
    
    function sum(uint256[vectorSize] memory data) internal pure returns(uint) {
        uint res = data[0];
        for (uint i = 1; i < vectorSize; i++) {
            res = res.add(data[i]);
        }
        
        return res;
    }
}