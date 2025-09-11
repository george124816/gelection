package db

import (
	"context"
	"log"
	"os"
)

func Schema() {
	sql := structure()
	_, err := Db.Exec(context.Background(), sql)
	if err != nil {
		log.Fatalln("Error creating schema:", err)
	}
}

func structure() string {
	sql, err := os.ReadFile("internal/db/structure.sql")
	if err != nil {
		log.Fatalln(err)
	}

	return string(sql)
}
