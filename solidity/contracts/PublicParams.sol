pragma solidity >= 0.5.0 < 0.6.0;
import "./library/BN128.sol";

contract PublicParams {
  using BN128 for BN128.G1Point;
  using BN128 for uint;

  // g, h generator for sigma proof.
  // Warning: g is not the g base point of curve.
  BN128.G1Point public g;
  BN128.G1Point public h;

  // gBase represents curve base point on curve.
  BN128.G1Point public gBase;

  // bit size for vector.
  // value range [0, 2^bitSize-1]
  uint public constant bitSize = 64;
  // 2^n=bitSize; n is the number of ecpoints in innerproduct.
  uint public constant n = 6;
  // aggSize for agg range proof. must be pow of 2.
  uint public constant aggSize = 2;

  // fix point used in inner product.
  BN128.G1Point public u;

  // g vector generator used in range proof.
  BN128.G1Point[bitSize*aggSize] public gVector;
  // h vector generator used in range proof.
  BN128.G1Point[bitSize*aggSize] public hVector;

  // 在实际使用时, 应该在创建合约时, 由外部系统直接传入. 保证隐私
  constructor() public {
    // set h generator.
    // this should be the first step to set which will be used to
    // generate generator point.
    gBase.X = 1;
    gBase.Y = 2;

    g = generatePointByString("g generator of twisted elg");
    h = generatePointByString("h generator of twisted elg");
    u = generatePointByString("u generator of innerproduct");

    initVector();
  }

  // init public gv/hv of protocol.
  function initVector() internal {
    uint scalar;

    uint gvb = uint(keccak256(abi.encodePacked("gvs"))).mod();
    for (uint i = 0; i < bitSize*aggSize; i++) {
      scalar = uint(keccak256(abi.encodePacked(gvb+i))).mod();
      gVector[i] = gBase.mul(scalar);
    }

    uint hvb = uint(keccak256(abi.encodePacked("hvs"))).mod();
    for (uint j = 0; j < bitSize*aggSize; j++) {
      scalar = uint(keccak256(abi.encodePacked(j+hvb))).mod();
      hVector[j] = gBase.mul(scalar);
    }
  }

  // return g generator.
  function getG() public view returns(uint[2] memory) {
    uint[2] memory res;
    res[0] = g.X;
    res[1] = g.Y;
    return res;
  }

  // return h generator.
  function getH() public view returns(uint[2] memory) {
    uint[2] memory res;
    res[0] = h.X;
    res[1] = h.Y;
    return res;
  }

  // return u fix point.
  function getU() public view returns(uint[2] memory) {
    uint[2] memory res;
    res[0] = u.X;
    res[1] = u.Y;
    return res;
  }

  // return aggGV.
  function getAggGVector() public view returns(uint[bitSize*aggSize*2] memory) {
    uint[bitSize*aggSize*2] memory res;
    for (uint i = 0; i < bitSize*aggSize; i++) {
      res[2*i] = gVector[i].X;
      res[2*i+1] = gVector[i].Y;
    }
    return res;
  }

  // return aggHV.
  function getAggHVector() public view returns(uint[bitSize*aggSize*2] memory) {
    uint[bitSize*aggSize*2] memory res;
    for (uint i = 0; i < bitSize*aggSize; i++) {
      res[2*i] = hVector[i].X;
      res[2*i+1] = hVector[i].Y;
    }
    return res;
  }

  // return g vector generator.
  function getGVector() public view returns(uint[bitSize*2] memory) {
    uint[bitSize*2] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[2*i] = gVector[i].X;
      res[2*i+1] = gVector[i].Y;
    }
    return res;
  }

  // return h vector generator.
  function getHVector() public view returns(uint[2*bitSize] memory) {
    uint[bitSize*2] memory res;
    for (uint i = 0; i < bitSize; i++) {
      res[2*i] = hVector[i].X;
      res[2*i+1] = hVector[i].Y;
    }
    return res;
  }

  // return bit size for value.
  function getBitSize() public pure returns(uint) {
    return bitSize;
  }

  // return number of ecpoints in innerproduct.
  function getN() public pure returns(uint) {
    return n;
  }

  // generatePointByString
  function generatePointByString(string memory s) internal view returns(BN128.G1Point memory) {
    uint scalar = uint(keccak256(abi.encodePacked(s))).mod();
    return gBase.mul(scalar);
  }
}
