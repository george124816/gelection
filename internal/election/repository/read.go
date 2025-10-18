package repository

import (
	"context"
	"log/slog"

	"github.com/george124816/gelection/internal/election/model"
)

func (d DefaultAdapter) GetElection(ctx context.Context, db DBQueries, id int) (*model.Election, error) {
	var election model.Election
	sqlStatement := `SELECT * FROM elections WHERE id = $1`

	err := db.QueryRow(ctx, sqlStatement, id).Scan(&election.Id, &election.Name)
	if err != nil {
		return nil, err
	}

	return &election, nil
}

func (d DefaultAdapter) GetAllElections(ctx context.Context, db DBQueries) ([]model.Election, error) {
	var elections []model.Election

	sqlStatement := `SELECT * FROM elections`

	result, err := db.Query(ctx, sqlStatement)

	if err != nil {
		return nil, err
	}

	for result.Next() {
		var e model.Election
		if err := result.Scan(&e.Id, &e.Name); err != nil {
			slog.Error("failed to scan result from database", "error", err)
		}
		elections = append(elections, e)
	}

	return elections, nil
}
