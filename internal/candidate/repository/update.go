package repository

import (
	"context"
	"errors"

	"github.com/george124816/gelection/internal/candidate/model"
	engine "github.com/george124816/gelection/internal/db"
)

func Update(id int, candidate model.Candidate) error {
	sqlStatement := `UPDATE candidates SET name = $1 WHERE id = $2`

	result, err := engine.Engine.Exec(context.Background(), sqlStatement, candidate.Name, id)

	if err != nil {
		return err
	}

	if result.RowsAffected() != 1 {
		return errors.New("failed to update")
	}

	return nil
}
