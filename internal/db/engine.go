package engine

import (
	"context"
	"log/slog"
	"os"

	"github.com/george124816/gelection/internal/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Engine *pgxpool.Pool

func Connect() {
	postgresUrl := configs.GetPostgresConfig().String()
	poolConfig, err := pgxpool.ParseConfig(postgresUrl)

	if err != nil {
		slog.Error("Unable to parse DATABASE_URL:", err)
		os.Exit(1)
	}

	Engine, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		slog.Error("Unable to create connection pool", err)
		os.Exit(1)
	}
}

func GetEnvOrDefault(envName string, defaultValue string) string {
	resultValue := os.Getenv(envName)
	if resultValue != "" {
		return resultValue
	} else {
		return defaultValue
	}
}
