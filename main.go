package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()
	defer DB.Close()

	r := gin.Default()
	r.GET("/transactions", ReadTransactions)
	r.POST("/transactions", CreateTransaction)
	r.PUT("/transactions/:id", UpdateTransaction)
	r.DELETE("/transactions/:id", DeleteTransaction)

	r.Run(":8080")
}
