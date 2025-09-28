package repository

import (
	"context"
	"log/slog"

	"github.com/george124816/gelection/internal/candidate/model"
	engine "github.com/george124816/gelection/internal/db"
)

func Create(candidate model.Candidate) error {

	sqlStatement := `
	INSERT INTO candidates (name, election_id) VALUES ($1, $2)
	`

	_, err := engine.Engine.Exec(context.Background(), sqlStatement, candidate.Name, candidate.ElectionId)

	if err != nil {
		slog.Error("failed to create candidate", err)
		return err
	}

	return nil

}
