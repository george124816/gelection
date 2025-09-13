package repository

import (
	"context"
	"log"

	"github.com/george124816/gelection/internal/candidate/model"
	engine "github.com/george124816/gelection/internal/db"
)

func Create(candidate model.Candidate) error {

	sqlStatement := `
	INSERT INTO candidates (name, election_id) VALUES ($1, $2)
	`

	_, err := engine.Db.Exec(context.Background(), sqlStatement, candidate.Name, candidate.ElectionId)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
