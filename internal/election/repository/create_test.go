package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/george124816/gelection/internal/election/model"
	"github.com/pashagolub/pgxmock/v4"
)

func TestCreate(t *testing.T) {
	t.Run("should create election successfully", func(t *testing.T) {
		election := model.Election{
			Name: "Municipal Election",
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("INSERT INTO elections \\(name\\) VALUES \\(\\$1\\)").
			WithArgs(election.Name).
			WillReturnResult(pgxmock.NewResult("INSERT", 1))

		err = Create(context.Background(), mock, election)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should return error when database insert fails", func(t *testing.T) {
		election := model.Election{
			Name: "Failed Election",
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("INSERT INTO elections \\(name\\) VALUES \\(\\$1\\)").
			WithArgs(election.Name).
			WillReturnError(errors.New("database error"))

		err = Create(context.Background(), mock, election)

		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should handle empty election name", func(t *testing.T) {
		election := model.Election{
			Name: "",
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("INSERT INTO elections \\(name\\) VALUES \\(\\$1\\)").
			WithArgs(election.Name).
			WillReturnResult(pgxmock.NewResult("INSERT", 1))

		err = Create(context.Background(), mock, election)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})
}