package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/filipbabicdev/finance-tracker-api/internal/database"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		log.Fatalf("repository test setup failed: %v", err)
	}
	os.Exit(code)
}

func run(m *testing.M) (int, error) {
	ctx := context.Background()

	// Start a PostgreSQL container
	container, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("finance_test"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		return 0, fmt.Errorf("start postgress container: %w", err)
	}
	defer func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			log.Printf("terminate container: %v", err)
		}
	}()

	dsn, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return 0, fmt.Errorf("connection string: %w", err)
	}

	testDB, err = database.New(database.Config{
		DSN:             dsn,
		MaxOpenConns:    5,
		MaxIdleConns:    2,
		ConnMaxLifetime: 5,
	})
	if err != nil {
		return 0, fmt.Errorf("connect and migrate: %w", err)
	}
	defer testDB.Close()

	return m.Run(), nil
}

func resetDB(t *testing.T) {
	t.Helper()
	if _, err := testDB.Exec(`
		TRUNCATE transactions, categories
		RESTART IDENTITY CASCADE;
	`); err != nil {
		t.Fatalf("reset db: %v", err)
	}
}
