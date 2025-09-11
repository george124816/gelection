package election

import (
	"context"
	"github.com/george124816/gelection/internal/db"
)

func Delete(id int) error {
	sqlStatement := `
DELETE FROM election WHERE id = $1
`

	_, err := db.Engine.Exec(context.Background(), sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}
