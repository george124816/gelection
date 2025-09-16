package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/george124816/gelection/internal/candidate/model"
	engine "github.com/george124816/gelection/internal/db"
	"github.com/jackc/pgx/v5"
)

type DBQueries interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func GetCandidate(ctx context.Context, db DBQueries, id uint64) (model.Candidate, error) {
	var candidate model.Candidate
	sqlStatement := `SELECT * FROM candidates WHERE id = $1`

	err := db.QueryRow(ctx, sqlStatement, id).Scan(&candidate.Id, &candidate.Name, &candidate.ElectionId)

	if err != nil {
		log.Println(err)
		return candidate, err
	}

	fmt.Println(candidate)

	return candidate, nil

}

func GetAll() ([]model.Candidate, error) {

	var candidates []model.Candidate

	sqlStatement := `SELECT * FROM candidates`

	result, err := engine.Db.Query(context.Background(), sqlStatement)

	if err != nil {
		return nil, err
	}

	for result.Next() {
		var c model.Candidate
		// scan into fields
		if err := result.Scan(&c.Id, &c.Name, &c.ElectionId); err != nil {
			log.Fatal(err)
		}
		candidates = append(candidates, c)
	}

	return candidates, nil
}
