package repository

import (
	"context"
	"log/slog"
	"testing"

	"github.com/george124816/gelection/internal/candidate/model"
	"github.com/pashagolub/pgxmock/v4"
)

var adapter = DefaultAdapter{}

func TestGetCandidate(t *testing.T) {
	t.Run("should return candidate successfully", func(t *testing.T) {
		expectedCandidate := model.Candidate{Name: "Ernesto", Id: 3, ElectionId: 22}
		mock, err := pgxmock.NewPool()

		if err != nil {
			slog.Error(err.Error())
		}

		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name", "election_id"}).AddRow(uint(3), "Ernesto", 22)

		mock.ExpectQuery("SELECT \\* FROM candidates WHERE id = \\$1").WithArgs(uint64(3)).WillReturnRows(rows)

		candidate, err := adapter.GetCandidate(context.Background(), mock, 3)

		if err != nil {
			slog.Error(err.Error())
		}

		if expectedCandidate.Name != candidate.Name {
			slog.Error(err.Error())
		}
		if expectedCandidate.ElectionId != candidate.ElectionId {
			slog.Error(err.Error())
		}
		if expectedCandidate.Id != candidate.Id {
			slog.Error(err.Error())
		}

	})
	t.Run("should return not found when candidate doesn't exists", func(t *testing.T) {
		mock, err := pgxmock.NewPool()

		if err != nil {
			slog.Error(err.Error())
		}

		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name", "election_id"})

		mock.ExpectQuery("SELECT \\* FROM candidates WHERE id = \\$1").WithArgs(uint64(3)).WillReturnRows(rows)

		_, err = adapter.GetCandidate(context.Background(), mock, 3)

		if err.Error() != "no rows in result set" {
			t.Error(err)
		}

	})

}
