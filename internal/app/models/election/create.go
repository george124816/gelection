package election

import (
	"context"
	"github.com/george124816/gelection/internal/db"
)

func Create(name string) error {
	_, err := db.Db.Exec(context.Background(), "INSERT INTO election (name) VALUES ($1)", name)
	if err != nil {
		return err
	}

	return nil
}
