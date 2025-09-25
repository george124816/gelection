package repository

import (
	"context"
	"log"

	"github.com/george124816/gelection/internal/election/model"
)

func Create(ctx context.Context, db DBQueries, election model.Election) error {
	sqlStatement := `
	INSERT INTO elections (name) VALUES ($1)
	`

	_, err := db.Exec(ctx, sqlStatement, election.Name)

	if err != nil {
		log.Println(err)

		return err
	}

	return nil

}
