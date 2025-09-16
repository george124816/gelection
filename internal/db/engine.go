package engine

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var Engine *pgxpool.Pool

func Connect() {
	poolConfig, err := pgxpool.ParseConfig("postgres://postgres:postgres@localhost:5555/postgres")
	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}

	Engine, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool", err)
	}
}
