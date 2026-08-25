package main

import (
	"log"

	"github.com/filipbabicdev/finance-tracker-api/internal/config"
	"github.com/filipbabicdev/finance-tracker-api/internal/database"
	"github.com/filipbabicdev/finance-tracker-api/internal/handler"
	"github.com/filipbabicdev/finance-tracker-api/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	db, err := database.New(database.Config{
		DSN:             cfg.DatabaseURL,
		MaxOpenConns:    cfg.DBMaxOpenConns,
		MaxIdleConns:    cfg.DBMaxIdleConns,
		ConnMaxLifetime: cfg.DBConnMaxLifetime,
	})

	if err != nil {
		log.Fatalf("database error: %v", err)
	}
	defer db.Close()

	transactionRepo := repository.NewTransactionRepo(db)
	transactionHandler := handler.NewTransactionHandler(transactionRepo)

	r := gin.Default()
	_ = r.SetTrustedProxies(nil)
	handler.SetupRoutes(r, transactionHandler)

	r.GET("/", handler.RootHandler(r, cfg.Env))

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
