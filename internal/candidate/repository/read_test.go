package repository

import (
	"context"
	"testing"

	"github.com/george124816/gelection/internal/candidate/model"
	"github.com/pashagolub/pgxmock/v4"
)

func TestGetCandidate(t *testing.T) {
	t.Run("should return candidate successfully", func(t *testing.T) {
		expectedCandidate := model.Candidate{Name: "Ernesto", Id: 3, ElectionId: 22}
		mock, err := pgxmock.NewPool()

		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}

		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name", "election_id"}).AddRow(uint(3), "Ernesto", 22)

		mock.ExpectQuery("SELECT \\* FROM candidates WHERE id = \\$1").WithArgs(uint64(3)).WillReturnRows(rows)

		candidate, err := GetCandidate(context.Background(), mock, 3)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if expectedCandidate.Name != candidate.Name {
			t.Errorf("Expected candidate name %s, got %s", expectedCandidate.Name, candidate.Name)
		}
		if expectedCandidate.ElectionId != candidate.ElectionId {
			t.Errorf("Expected candidate election_id %d, got %d", expectedCandidate.ElectionId, candidate.ElectionId)
		}
		if expectedCandidate.Id != candidate.Id {
			t.Errorf("Expected candidate id %d, got %d", expectedCandidate.Id, candidate.Id)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should return not found when candidate doesn't exists", func(t *testing.T) {
		mock, err := pgxmock.NewPool()

		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}

		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name", "election_id"})

		mock.ExpectQuery("SELECT \\* FROM candidates WHERE id = \\$1").WithArgs(uint64(3)).WillReturnRows(rows)

		_, err = GetCandidate(context.Background(), mock, 3)

		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err.Error() != "no rows in result set" {
			t.Errorf("Expected 'no rows in result set', got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})
}

func TestGetAllCandidates(t *testing.T) {
	t.Run("should return all candidates successfully", func(t *testing.T) {
		expectedCandidates := []model.Candidate{
			{Id: 1, Name: "Alice", ElectionId: 10},
			{Id: 2, Name: "Bob", ElectionId: 10},
			{Id: 3, Name: "Charlie", ElectionId: 20},
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name", "election_id"}).
			AddRow(int64(1), "Alice", int64(10)).
			AddRow(int64(2), "Bob", int64(10)).
			AddRow(int64(3), "Charlie", int64(20))

		mock.ExpectQuery("SELECT \\* FROM candidates").WillReturnRows(rows)

		candidates, err := GetAllCandidates(context.Background(), mock)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(candidates) != len(expectedCandidates) {
			t.Errorf("Expected %d candidates, got %d", len(expectedCandidates), len(candidates))
		}

		for i, candidate := range candidates {
			if candidate.Id != expectedCandidates[i].Id {
				t.Errorf("Expected candidate[%d].Id %d, got %d", i, expectedCandidates[i].Id, candidate.Id)
			}
			if candidate.Name != expectedCandidates[i].Name {
				t.Errorf("Expected candidate[%d].Name %s, got %s", i, expectedCandidates[i].Name, candidate.Name)
			}
			if candidate.ElectionId != expectedCandidates[i].ElectionId {
				t.Errorf("Expected candidate[%d].ElectionId %d, got %d", i, expectedCandidates[i].ElectionId, candidate.ElectionId)
			}
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should return empty slice when no candidates exist", func(t *testing.T) {
		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name", "election_id"})
		mock.ExpectQuery("SELECT \\* FROM candidates").WillReturnRows(rows)

		candidates, err := GetAllCandidates(context.Background(), mock)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(candidates) != 0 {
			t.Errorf("Expected empty slice, got %d candidates", len(candidates))
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})
}
