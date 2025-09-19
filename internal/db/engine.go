package engine

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Engine *pgxpool.Pool

func Connect() {
	postgresUrl := GetEnvOrDefault("POSTGRES_URL", "postgres://postgres:postgres@localhost:5555/postgres")
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
