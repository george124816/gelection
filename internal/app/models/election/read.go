package election

import (
	"context"
	"github.com/george124816/gelection/internal/db"
)

func Read(id int) (string, error) {
	var name string
	err := db.Db.QueryRow(context.Background(), "SELECT name FROM election WHERE id = $1", id).Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}
