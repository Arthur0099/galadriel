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
  });
});
