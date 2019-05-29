package pgc

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	// ContractsPath is the location of truffle contract build files.
	ContractsPath = "./solidity/build/contracts"

	// ABI is the special field in contract json files which specify
	// contract's abi.
	ABI = "abi"
	// BIN is the special filed in contract json files which specify
	// contract's initial binary code to be deployed.
	BIN = "bytecode"
)

var contractsName = []map[string]string{
	{
		"name": "DLESigmaVerifier",
	},
	{
		"name": "PublicParams",
	},
	{
		"name": "IPVerifier",
	},
	{
		"name": "PGC",
	},
	{
		"name": "RangeProofVerifier",
	},
	{
		"name": "SigmaVerifier",
	},
}

// GenerateGoCode automatically generates go files for contract defined in truffle contracts.
func GenerateGoCode(newFlag bool) {
	for _, v := range contractsName {
		path := filepath.Join(ContractsPath, v["name"]+".json")
		files := generateGoCode(path, v["name"], newFlag)
		for _, f := range files {
			Delete(f)
		}
	}
}

// GenerateJavaCode automatically generates java code to call/send contract tx.
func GenerateJavaCode(flag bool) {
	for _, v := range contractsName {
		path := filepath.Join(ContractsPath, v["name"]+".json")
		files := generateJavaCode(path, v["name"])
		if flag {
			for _, f := range files {
				Delete(f)
			}
		}
	}
}

func generateGoCode(path, name string, newFlag bool) []string {
	raw := Read(path)

	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		panic(err)
	}

	abiRaw, err := json.Marshal(data[ABI])
	binRaw, err := json.Marshal(data[BIN])
	if err != nil {
		panic(err)
	}

	// Trim ""
	trimBin := []byte(strings.Trim(string(binRaw), "\""))

	// Write abi and bin file
	name = strings.ToLower(name)
	abiName := name + ".abi"
	binName := name + ".bin"
	var outName string
	if newFlag {
		outName = filepath.Join("./", "new", name, name+".go")
	} else {
		outName = filepath.Join("./contracts", name+".go")
	}

	Write(abiName, abiRaw)
	Write(binName, trimBin)

	cmd := "./abigen"
	command := exec.Command(cmd, "--bin", binName, "--abi", abiName, "--pkg", "contracts", "--out", outName, "--type", name)
	fmt.Print(outName, "\n")
	err = command.Run()
	if err != nil {
		panic(err)
	}

	return []string{
		filepath.Join("./", abiName),
		filepath.Join("./", binName),
	}
}

func generateJavaCode(path, name string) []string {
	raw := Read(path)

	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		panic(err)
	}

	abiRaw, err := json.Marshal(data[ABI])
	binRaw, err := json.Marshal(data[BIN])
	if err != nil {
		panic(err)
	}

	// Trim ""
	trimBin := []byte(strings.Trim(string(binRaw), "\""))

	// Write abi and bin file
	abiName := name + ".abi"
	binName := name + ".bin"

	Write(abiName, abiRaw)
	Write(binName, trimBin)

	cmd := "/home/ubuntu/web3j-3.6.0/bin/web3j"
	command := exec.Command(cmd, "solidity", "generate", binName, abiName, "-p", "contracts", "-o", "./contracts/java")
	err = command.Run()
	if err != nil {
		panic(err)
	}

	return []string{
		filepath.Join("./", abiName),
		filepath.Join("./", binName),
	}
}
