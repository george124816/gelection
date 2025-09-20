package migrate

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	engine "github.com/george124816/gelection/internal/db"
)

func Migrate() error {
	postgresUrl := engine.GetEnvOrDefault(
		"POSTGRES_URL", "postgres://postgres:postgres@localhost:5555/postgres?sslmode=disable",
	)
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///bin/db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != migrate.ErrNoChange {
		return err
	}

	return nil
}
