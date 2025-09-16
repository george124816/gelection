package repository

import (
	"context"
	"errors"

	engine "github.com/george124816/gelection/internal/db"
)

func DeleteCandidate(id uint64) error {
	sqlStatement := `DELETE FROM candidates WHERE id = $1`

	result, err := engine.Db.Exec(context.Background(), sqlStatement, id)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 1 {
		return nil
	} else {
		return errors.New("failed to delete")

	}

}
