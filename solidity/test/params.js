const PublicParams = artifacts.require("PublicParams");
import {ReadJsonFile} from "./file.js";

contract("public params", function (accounts){
  before(async function(){
    this.ContractIns = await PublicParams.new();
  });

  it("check public params", async function () {
    const params = ReadJsonFile("../params/params");
    const g = await this.ContractIns.getG();
    assert.equal(g[0], params.g.X, "g.x not equal");
    assert.equal(g[1], params.g.Y, "g.y not equal");

    const h = await this.ContractIns.getH();
    assert.equal(h[0], params.h.X, "h.x not equal");
    assert.equal(h[1], params.h.Y, "h.y not equal");

    // check fix u point.
    const u = await this.ContractIns.getU();
    assert.equal(u[0], params.u.X, "u.x not equal");
    assert.equal(u[1], params.u.Y, "u.y not equal");

    // check g vector/h vector.
    const contractGV = await this.ContractIns.getGVector();
    const contractHV = await this.ContractIns.getHVector();
    assert.equal(contractGV.length, params.gv.length, "gv length not equal");
    assert.equal(contractHV.length, params.hv.length, "hv length not equal");
    assert.equal(contractGV.length, contractHV.length, "gv length not equal hv length");

    for (let i = 0; i < contractGV.length; i++) {
      let conGVector = contractGV[i];
      let conHVector = contractHV[i];
      let goGVector = params.gv[i];
      let goHVector = params.hv[i];
      assert.equal(conGVector[0], goGVector.X, i+" g vector x not equal");
      assert.equal(conGVector[1], goGVector.Y, i+" g vector y not equal");
      assert.equal(conHVector[0], goHVector.X, i+" h vector x not equal");
      assert.equal(conHVector[1], goHVector.Y, i+" h vector y not equal");
    }
  });
});
