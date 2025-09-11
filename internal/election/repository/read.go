package election

import (
	"context"
	"github.com/george124816/gelection/internal/db"
)

func Read(id int) (Election, error) {
	var election Election
	sqlStatement := `
SELECT name FROM election WHERE id = $1
`

	err := db.Engine.QueryRow(context.Background(), sqlStatement, id).Scan(&election.name)
	if err != nil {
		return "", err
	}

	return election, nil
}
