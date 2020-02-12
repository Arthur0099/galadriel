import { italic } from "ansi-colors";

const Test = artifacts.require("Test");

contract("test", function(account){
        before(async function(){
                this.TestIns = await Test.new();
        });

        it("", async function(){
                let arr = [];
                arr.push(1);
                arr.push(1);
                arr.push(1);
                arr.push(1);
                let tx = await this.TestIns.set(arr);
                console.log(tx.receipt.gasUsed);
                tx = await this.TestIns.set2(arr);
                console.log(tx.receipt.gasUsed);        
        })
})