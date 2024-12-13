package utils

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func DbConnection() (*sql.DB, func()) {
	ctx := context.Background()

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(filepath.Join("..", "utils", "migrations.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(50000000*time.Second)),
	)

	if err != nil {
		panic(err)
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")

	db, err := sql.Open("pgx", connStr)

	cleanup := func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			fmt.Println("failed to terminate pgContainer: %s", err)
		}
		if err := db.Close(); err != nil {
			fmt.Printf("Failed to close database connection: %v\n", err)
		}
	}

	return db, cleanup

}
