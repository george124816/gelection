package election

import (
	"context"
	"github.com/george124816/gelection/internal/db"
)

func Create(election Election) error {
	sqlStatement := `
INSERT INTO election (name) VALUES ($1)
`

	_, err := db.Engine.Exec(context.Background(), sqlStatement, election.name)
	if err != nil {
		return err
	}

	return nil
}
