/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

const (
	ProductDocument = "PRODUCT_DOCUMENT"
	UserDocument = "USER_DOCUMENT"
)

	
type Product struct {
	DocumentType string `json:"documenttype"`
	Name string `json:"name"`
	Category string `json:"category"`
	Price string `json:"price"`
	Brand string `json:"brand"`
}

type User struct {
	DocumentType string `json:"documenttype"`
	UserId string `json:"userId"`
	UserName string `json:"username"`
	UserEmail string `json:"useremail"`
	UserType string `json:"usertype"`
}

// QueryResult structure used for handling result of query
type QueryResultProduct struct {
	Key    string `json:"Key"`
	Record *Product
}

type QueryResultUser struct {
	Key    string `json:"Key"`
	Record *User
}


// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	

	products := []Product{
		Product{DocumentType: ProductDocument, Name: "Redmi Note 10", Category: "Electronics", Price: "12000", Brand:"Xiaomi"},
		Product{DocumentType: ProductDocument, Name: "Dell Vostro", Category: "Computers", Price: "42000", Brand:"Dell"},
		Product{DocumentType: ProductDocument, Name: "I9-9700K", Category: "Computers", Price: "68000", Brand:"Intel"},
		Product{DocumentType: ProductDocument, Name: "RTX 2090", Category: "Computers", Price: "24000", Brand:"Nvidia"},
	}

	users := []User {
		User{DocumentType: UserDocument,UserEmail:"djoynes0@bing.com" ,UserId:"653a890afc13ae192fc76854",UserName:"Rachèle",UserType:"Premium"},
		User{DocumentType: UserDocument,UserEmail:"bskryne1@apple.com" ,UserId:"653a890afc13ae192fc76856",UserName:"Clémentine",UserType:"Premium"},
		User{DocumentType: UserDocument,UserEmail:"wcrunkhorn2@t-online.de" ,UserId:"653a890afc13ae192fc76855",UserName:"Tán",UserType:"Premium"},
		User{DocumentType: UserDocument,UserEmail:"biwanowski3@time.com" ,UserId:"653a890afc13ae192fc76857",UserName:"Mahélie",UserType:"Standard"},
		User{DocumentType: UserDocument,UserEmail:"aaukland4@geocities.jp" ,UserId:"653a890afc13ae192fc76858",UserName:"Märta",UserType:"Standard"},
	}
	
		
	for i, user := range users {
		userAsBytes, _ := json.Marshal(user)
		err := ctx.GetStub().PutState("USER"+strconv.Itoa(i), userAsBytes)
		
		if err != nil {
			return fmt.Errorf("Failed to put to user world state. %s", err.Error())
		}
	}

	for j, product := range products {
		prodAsBytes, _ := json.Marshal(product)
		err := ctx.GetStub().PutState("PRODUCT"+strconv.Itoa(j), prodAsBytes)
		
		if err != nil {
			return fmt.Errorf("Failed to put to product world state. %s", err.Error())
		}
	}

	return nil
}

func (s *SmartContract) QueryAllProduct(ctx contractapi.TransactionContextInterface) ([]QueryResultProduct, error){

	queryString := fmt.Sprintf(`{"selector":{"documenttype":"%v"}}`, ProductDocument)
// process, err := ctx.GetStub().GetPrivateDataQueryResult("collectionContract",queryString)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("no contract found")
	}
	defer resultsIterator.Close()
	results := []QueryResultProduct{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		product := new(Product)
		_ = json.Unmarshal(queryResponse.Value, product)
		queryResult := QueryResultProduct{Key: queryResponse.Key, Record: product}
		results = append(results, queryResult)
	}
	return results, nil
}

func (s *SmartContract) QueryAllUser (ctx contractapi.TransactionContextInterface) ([]QueryResultUser, error) {
	queryString := fmt.Sprintf(`{"selector":{"documenttype":"%v"}}`, UserDocument)
// process, err := ctx.GetStub().GetPrivateDataQueryResult("collectionContract",queryString)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("no contract found")
	}
	defer resultsIterator.Close()
	results := []QueryResultUser{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		user := new(User)
		_ = json.Unmarshal(queryResponse.Value, user)
		queryResult := QueryResultUser{Key: queryResponse.Key, Record: user}
		results = append(results, queryResult)
	}
	return results, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}