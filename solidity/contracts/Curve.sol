pragma solidity >= 0.5.0 < 0.6.0;

import "./library/BN128.sol";
import "./library/Secp256k1.sol";

contract Curve {

  // 0 represent secp256k1.
  // 1 represent bn128.
  uint public curveType;

  function add(uint x1, uint y1, uint x2, uint y2) public view returns(uint x3, uint y3) {
    return (0, 0);
  }

  function scalarMult(uint x, uint y, uint scalar) public view returns(uint x1, uint y1) {
    return (0, 0);
  }

  // add point(x1, y1) with point(x2, y2) under bn128 curve.
  function bn128Add(uint x1, uint y1, uint x2, uint y2) public view returns(uint, uint) {
    return (0, 0);
  }

  // add point(x1, y1) with point(x2, y2) under secp256k1 curve.
  function secpAdd(uint x1, uint y1, uint x2, uint y2) public view returns(uint, uint) {
    return Secp256k1.ecadd(x1, y1, x2, y2);
  }

  // mult point(x1, y1) with scalar under secp256k1.
  function secpScalarMult(uint x, uint y, uint scalar) public view returns(uint, uint) {
    return Secp256k1.ecmul(x, y, scalar);
  }
}
