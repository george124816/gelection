package election

import (
	"context"
	"github.com/george124816/gelection/internal/db"
)

func Delete(id int) error {
	_, err := db.Db.Exec(context.Background(), "DELETE FROM election WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
