package repository

import (
	"context"
	"log"
	"testing"

	"github.com/george124816/gelection/internal/candidate/model"
	"github.com/pashagolub/pgxmock/v4"
)

func TestGetCandidate(t *testing.T) {
	t.Run("should return candidate successfully", func(t *testing.T) {
		expectedCandidate := model.Candidate{Name: "Ernesto", Id: 3, ElectionId: 22}
		mock, err := pgxmock.NewPool()

		if err != nil {
			log.Fatal(err)
		}

		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name", "election_id"}).AddRow(uint(3), "Ernesto", 22)

		mock.ExpectQuery("SELECT \\* FROM candidates WHERE id = \\$1").WithArgs(uint64(3)).WillReturnRows(rows)

		candidate, err := GetCandidate(context.Background(), mock, 3)

		if err != nil {
			log.Fatal(err)
		}

		if expectedCandidate.Name != candidate.Name {
			log.Fatal("the name doesn't match")
		}
		if expectedCandidate.ElectionId != candidate.ElectionId {
			log.Fatal("the election_id doesn't match")
		}
		if expectedCandidate.Id != candidate.Id {
			log.Fatal("the id doesn't match")
		}

		if 1 != 2 {
			log.Fatal("should error")
		}

	})
	t.Run("should return not found when candidate doesn't exists", func(t *testing.T) {
		mock, err := pgxmock.NewPool()

		if err != nil {
			log.Fatal(err)
		}

		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name", "election_id"})

		mock.ExpectQuery("SELECT \\* FROM candidates WHERE id = \\$1").WithArgs(uint64(3)).WillReturnRows(rows)

		_, err = GetCandidate(context.Background(), mock, 3)

		if err.Error() != "no rows in result set" {
			t.Error(err)
		}

	})

}
