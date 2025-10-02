package repository

import (
	"context"
	"testing"

	"github.com/george124816/gelection/internal/election/model"
	"github.com/pashagolub/pgxmock/v4"
)

func TestGetElection(t *testing.T) {
	t.Run("should return election successfully", func(t *testing.T) {
		expectedElection := model.Election{Id: 1, Name: "Presidential Election"}
		mock, err := pgxmock.NewPool()

		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}

		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name"}).AddRow(int64(1), "Presidential Election")

		mock.ExpectQuery("SELECT \\* FROM elections WHERE id = \\$1").WithArgs(1).WillReturnRows(rows)

		election, err := GetElection(context.Background(), mock, 1)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if expectedElection.Name != election.Name {
			t.Errorf("Expected election name %s, got %s", expectedElection.Name, election.Name)
		}
		if expectedElection.Id != election.Id {
			t.Errorf("Expected election id %d, got %d", expectedElection.Id, election.Id)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should return error when election doesn't exist", func(t *testing.T) {
		mock, err := pgxmock.NewPool()

		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}

		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name"})

		mock.ExpectQuery("SELECT \\* FROM elections WHERE id = \\$1").WithArgs(99).WillReturnRows(rows)

		_, err = GetElection(context.Background(), mock, 99)

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

func TestGetAllElections(t *testing.T) {
	t.Run("should return all elections successfully", func(t *testing.T) {
		expectedElections := []model.Election{
			{Id: 1, Name: "Presidential Election"},
			{Id: 2, Name: "Senate Election"},
			{Id: 3, Name: "Local Election"},
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name"}).
			AddRow(int64(1), "Presidential Election").
			AddRow(int64(2), "Senate Election").
			AddRow(int64(3), "Local Election")

		mock.ExpectQuery("SELECT \\* FROM elections").WillReturnRows(rows)

		elections, err := GetAllElections(context.Background(), mock)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(elections) != len(expectedElections) {
			t.Errorf("Expected %d elections, got %d", len(expectedElections), len(elections))
		}

		for i, election := range elections {
			if election.Id != expectedElections[i].Id {
				t.Errorf("Expected election[%d].Id %d, got %d", i, expectedElections[i].Id, election.Id)
			}
			if election.Name != expectedElections[i].Name {
				t.Errorf("Expected election[%d].Name %s, got %s", i, expectedElections[i].Name, election.Name)
			}
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should return empty slice when no elections exist", func(t *testing.T) {
		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		rows := mock.NewRows([]string{"id", "name"})
		mock.ExpectQuery("SELECT \\* FROM elections").WillReturnRows(rows)

		elections, err := GetAllElections(context.Background(), mock)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(elections) != 0 {
			t.Errorf("Expected empty slice, got %d elections", len(elections))
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})
}