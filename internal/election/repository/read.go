package repository

import (
	"context"
	"log"

	"github.com/george124816/gelection/internal/db"
	"github.com/george124816/gelection/internal/election/model"
)

func GetElection(ctx context.Context, db DBQueries, id int) (model.Election, error) {
	var election model.Election
	sqlStatement := `
SELECT name FROM election WHERE id = $1
`

	err := engine.Engine.QueryRow(ctx, sqlStatement, id).Scan(election.Name)
	if err != nil {
		return model.Election{}, err
	}

	return election, nil
}

func GetAll(ctx context.Context, db DBQueries) ([]model.Election, error) {
	var elections []model.Election

	sqlStatement := `SELECT * FROM elections`

	result, err := db.Query(ctx, sqlStatement)

	if err != nil {
		return nil, err
	}

	for result.Next() {
		var e model.Election
		if err := result.Scan(&e.Id, &e.Name); err != nil {
			log.Fatal(err)
		}
		elections = append(elections, e)
	}

	return elections, nil
}
