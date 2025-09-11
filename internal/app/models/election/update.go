package election

import (
	"context"
	"github.com/george124816/gelection/internal/db"
)

func Update(id int, name string) error {
	_, err := db.Db.Exec(context.Background(), "UPDATE election SET name = $1 WHERE id = $2", name, id)
	if err != nil {
		return err
	}

	return nil
}
