package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/transactions", ReadTransactions)
	r.POST("/transactions", CreateTransaction)
	r.PUT("/transactions/:id", UpdateTransaction)
	r.DELETE("/transactions/:id", DeleteTransaction)
}
