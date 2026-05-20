package main

import (
	"github.com/filipbabicdev/finance-tracker-api/internal/config"
	"github.com/filipbabicdev/finance-tracker-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	r := gin.Default()
	r.GET("/transactions", handler.ReadTransactions)
	r.POST("/transactions", handler.CreateTransaction)
	r.PUT("/transactions/:id", handler.UpdateTransaction)
	r.DELETE("/transactions/:id", handler.DeleteTransaction)

	r.Run(":8090")
}
