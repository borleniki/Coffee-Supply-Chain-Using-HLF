package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Farmer struct {
	Address     string `json:"address"`
	Color       string `json:"color"`
	ContactNo   string `json:"contactNo"`
	DateOfBirth string `json:"dob"`
	EmailId     string `json:"emailId"`
	Gender      string `json:"gender"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	TypeofBeans string `json:"typeofBeans"`
	UserName    string `json:"userName"`
}

type Process struct {
	ProcessId     string `json:"processId"`
	CoffeeQuality string `json:"quality"`
	Color         string `json:"color"`
	Name          string `json:"name"`
	TypeofBeans   string `json:"typeofBeans"`
	GrindSize     string `json:"grindSize"`
	BrewTime      string `json:"brewTime"`
}
type OrderData struct {
	ProcessId     string `json:"processId"`
	CoffeeQuality string `json:"quality"`
	Color         string `json:"color"`
	Name          string `json:"name"`
	TypeofBeans   string `json:"typeofBeans"`
	GrindSize     string `json:"grindSize"`
	BrewTime      string `json:"brewTime"`
}

func main() {
	router := gin.Default()

	router.Static("/public", "./public")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Farmer Dashboard",
		})

	})

	router.POST("/api/farmer", func(ctx *gin.Context) {

		var req Farmer
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}
		fmt.Println("request", req)
		result := submitTxnFn(
			"farmers",
			"coffeechannel",
			"Coffeesupply",
			"FarmerContract",
			"invoke",
			make(map[string][]byte),
			"RegisterFarmer",
			req.UserName,
			req.Name,
			req.Gender,
			req.Color,
			req.ContactNo,
			req.EmailId,
			req.Status,
			req.Address,
			req.DateOfBirth,
			req.TypeofBeans,
		)
		ctx.JSON(http.StatusOK, gin.H{"message": "Created new farmer with new coffee", "result": result})

	})

	router.GET("/api/farmer/:id", func(ctx *gin.Context) {
		userName := ctx.Param("id")

		result := submitTxnFn("farmers", "coffeechannel", "Coffeesupply", "FarmerContract", "query", make(map[string][]byte), "ReadFarmerDetails", userName)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.POST("/api/order", func(ctx *gin.Context) {
		var req Process
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		fmt.Printf("order  %s", req)

		privateData := map[string][]byte{
			"processId":   []byte(req.ProcessId),
			"quality":     []byte(req.CoffeeQuality),
			"color":       []byte(req.Color),
			"name":        []byte(req.Name),
			"typeofBeans": []byte(req.TypeofBeans),
			"grindSize":   []byte(req.GrindSize),
			"brewTime":    []byte(req.BrewTime),
		}

		submitTxnFn("processors", "coffeechannel", "Coffeesupply", "ProcessContract", "private", privateData, "CreateProcessOrder", req.ProcessId)

		ctx.JSON(http.StatusOK, req)
	})

	router.GET("/api/order/:id", func(ctx *gin.Context) {
		processId := ctx.Param("id")

		result := submitTxnFn("processors", "coffeechannel", "Coffeesupply", "ProcessContract", "query", make(map[string][]byte), "ReadProcessOrder", processId)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.GET("/api/order/all", func(ctx *gin.Context) {

		result := submitTxnFn("processors", "coffeechannel", "Coffeesupply", "ProcessContract", "query", make(map[string][]byte), "GetAllOrders")

		var orders []OrderData

		if len(result) > 0 {
			// Unmarshal the JSON array string into the orders slice
			if err := json.Unmarshal([]byte(result), &orders); err != nil {
				fmt.Println("Error:", err)
				return
			}

		}
		ctx.JSON(http.StatusOK, gin.H{"data": result})

	})

	router.Run(":3000")

}
