package engine

import (
	"context"
	"log"
	"os"

	"github.com/george124816/gelection/internal/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Engine *pgxpool.Pool

func Connect() {
	postgresUrl := configs.GetPostgressConfig().String()
	poolConfig, err := pgxpool.ParseConfig(postgresUrl)

	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}

	Engine, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool", err)
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
