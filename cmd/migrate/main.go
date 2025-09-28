package migrate

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/george124816/gelection/internal/configs"
)

func Migrate() error {
	databaseConfig := configs.GetPostgresConfig()
	postgresUrl := databaseConfig.String()

	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		slog.Error("failed to open database connection", err)
		os.Exit(1)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.Error("failed to configure drive", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		slog.Error("failed to migrate", err)
	}

	slog.Info("running migrate")

	err = m.Up()
	if err != migrate.ErrNoChange {
		return err
	}

	return nil
}
