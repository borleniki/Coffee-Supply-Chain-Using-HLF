package main

import (
	"coffeesupply/contracts"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	farmerContract := new(contracts.FarmerContract)
	processContract := new(contracts.ProcessContract)

	chaincode, err := contractapi.NewChaincode(farmerContract, processContract)

	if err != nil {
		log.Panicf("Could not create chaincode : %v", err)
	}

	err = chaincode.Start()

	if err != nil {
		log.Panicf("Failed to start chaincode : %v", err)
	}
}
