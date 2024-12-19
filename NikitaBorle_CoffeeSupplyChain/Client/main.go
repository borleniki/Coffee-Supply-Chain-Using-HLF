package main

import "fmt"

func main() {
	// result := submitTxnFn(
	// 	"farmers",
	// 	"coffeechannel",
	// 	"Coffeesupply",
	// 	"FarmerContract",
	// 	"invoke",
	// 	make(map[string][]byte),
	// 	"RegisterFarmer",
	// 	"Nikita", "nikku", "06/03/2002", "Female", "nikku@gmail.com", "1234567891", "Hyderabad", "Arabica", "Brown Black", "InFarm",
	// )

	privateData := map[string][]byte{
		"quality":     []byte("High"),
		"color":       []byte("Brown"),
		"name":        []byte("Amjad"),
		"typeofBeans": []byte("Arabica"),
		"grindSize":   []byte("small"),
		"brewTime":    []byte("5 min"),
	}

	result := submitTxnFn("processors", "coffeechannel", "Coffeesupply", "ProcessContract", "private", privateData, "CreateProcessOrder", "01")

	// result := submitTxnFn("processors", "coffeechannel", "Coffeesupply", "ProcessContract", "query", make(map[string][]byte), "ReadProcessOrder", "01")

	// result := submitTxnFn("manufacturer", "coffeechannel", "Coffeesupply", "FarmerContract", "query", make(map[string][]byte), "GetAllFarmers")

	// result := submitTxnFn("manufacturer", "coffeechannel", "Coffeesupply", "ProcessContract", "query", make(map[string][]byte), "GetAllOrders")

	// result := submitTxnFn("manufacturer", "coffeechannel", "Coffeesupply", "FarmerContract", "query", make(map[string][]byte), "GetMatchingOrders", "02")

	// result := submitTxnFn("manufacturer", "coffeechannel", "Coffeesupply", "FarmerContract", "invoke", make(map[string][]byte), "MatchOrder", "02", "01")

	// result := submitTxnFn("manufacturer", "coffeechannel", "Coffeesupply", "FarmerContract", "query", make(map[string][]byte), "ReadFarmerDetails", "01")

	fmt.Println(result)

}
