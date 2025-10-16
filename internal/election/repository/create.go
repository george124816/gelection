package repository

import (
	"context"
	"log/slog"

	"github.com/george124816/gelection/internal/election/model"
)

func (d DefaultAdapter) Create(ctx context.Context, db DBQueries, election model.Election) error {
	sqlStatement := `
	INSERT INTO elections (name) VALUES ($1)
	`

	_, err := db.Exec(ctx, sqlStatement, election.Name)

	if err != nil {
		slog.Error("failed to create election", "error", err)

		return err
	}

	return nil

}
