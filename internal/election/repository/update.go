package election

import (
	"context"
	"github.com/george124816/gelection/internal/db"
)

func Update(election Election) error {
	sqlStatement := `
UPDATE election SET name = $1 WHERE id = $2
`

	_, err := db.Engine.Exec(context.Background(), sqlStatement, election.name, election.id)
	if err != nil {
		return err
	}

	return nil
}
