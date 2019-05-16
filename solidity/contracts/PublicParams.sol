pragma solidity >= 0.5.0 < 0.6.0;
pragma experimental ABIEncoderV2;

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
  uint public constant bitSize = 16;
  // 2^n=bitSize; n is the number of ecpoints in innerproduct.
  uint public constant n = 4;

  // fix point used in inner product.
  BN128.G1Point public u;

  // g vector generator used in range proof.
  BN128.G1Point[bitSize] public gVector;
  // h vector generator used in range proof.
  BN128.G1Point[bitSize] public hVector;

  constructor() public {
    // set alt_128 g base point to g.
    h.X = 1;
    h.Y = 2;
    // set u same with h.
    u.X = h.X;
    u.Y = h.Y;

    // set h generator.
    gBase.X = 1;
    gBase.Y = 2;

    g = generatePointByString("g generator of twisted elg");

    initVector();
  }

  // init public gv/hv of protocol.
  function initVector() internal {
    uint scalar;
    for (uint i = 0; i < bitSize; i++) {
      scalar = uint(keccak256(abi.encodePacked(i))).mod();
      gVector[i] = gBase.mul(scalar);
    }

    for (uint j = 0; j < bitSize; j++) {
      scalar = uint(keccak256(abi.encodePacked(j+bitSize))).mod();
      hVector[j] = gBase.mul(scalar);
    }
  }

  // return g generator.
  function getG() public view returns(BN128.G1Point memory) {
    return g;
  }

  // return h generator.
  function getH() public view returns(BN128.G1Point memory) {
    return h;
  }

  // return u fix point.
  function getU() public view returns(BN128.G1Point memory) {
    return u;
  }

  // return g vector generator.
  function getGVector() public view returns(BN128.G1Point[bitSize] memory) {
    return gVector;
  }

  // return h vector generator.
  function getHVector() public view returns(BN128.G1Point[bitSize] memory) {
    return hVector;
  }

  // return bit size for value.
  function getBitSize() public view returns(uint) {
    return bitSize;
  }

  // return number of ecpoints in innerproduct.
  function getN() public view returns(uint) {
    return n;
  }

  // generatePointByString
  function generatePointByString(string memory s) internal view returns(BN128.G1Point memory) {
    uint scalar = uint(keccak256(abi.encodePacked(s))).mod();
    return gBase.mul(scalar);
  }
}
