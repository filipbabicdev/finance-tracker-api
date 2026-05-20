package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *TransactionHandler) {
	r.POST("/transactions", h.Create)
	r.GET("/transactions", h.ReadAll)
	r.PUT("/transactions/:id", h.Update)
	r.DELETE("/transactions/:id", h.Delete)
}
