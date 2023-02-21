package main

import (
	"log"

	"github.com/amidgo/amidledger/chaincode/domain/contract"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&contract.ShopContract{})
	if err != nil {
		log.Panic(err)
	}
	if err := assetChaincode.Start(); err != nil {
		log.Panic(err)
	}
}
