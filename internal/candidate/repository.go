package candidate

import (
	"context"
	"fmt"
	"log"

	"github.com/george124816/gelection/internal/db"
)

func CreateCandidate(candidate Candidate) error {
	sqlStatement := `
	INSERT INTO candidates (name, election_id) VALUES ($1, $2)
	`

	_, err := engine.Db.Exec(context.Background(), sqlStatement, candidate.Name, candidate.ElectionId)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func GetCandidate(id uint64) (Candidate, error) {
	var candidate Candidate
	sqlStatement := `SELECT * FROM candidates WHERE id = $1`

	err := engine.Db.QueryRow(context.Background(), sqlStatement, id).Scan(&candidate.Id, &candidate.Name, &candidate.ElectionId)

	if err != nil {
		log.Println(err)
		return candidate, err
	}

	fmt.Println(candidate)

	return candidate, nil

}

// func GetAll() ([]Candidate, error) {
// 	var candidates []Candidate
//
// 	sqlStatement := `SELECT * FROM candidates`
//
// 	db.Db.Query(context.Background(), sqlStatement)
// }
//
// func DeleteCandidate(id uint64) error {
// 	sqlStatement := `DELETE FROM candidates WHERE id = $1`
//
// 	_, err := db.Db.Exec(context.Background(), sqlStatement, id)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
