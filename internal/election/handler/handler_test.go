package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/george124816/gelection/internal/election/model"
	"github.com/george124816/gelection/internal/election/repository"
)

type MockAdapter struct{}

func (d MockAdapter) GetAllElections(ctx context.Context, db repository.DBQueries) ([]model.Election, error) {
	elections := []model.Election{
		model.Election{
			Id:   1,
			Name: "Zé do Caixão",
		},
	}

	return elections, nil
}

func (d MockAdapter) Create(ctx context.Context, db repository.DBQueries, election model.Election) error {
	return nil
}

func (d MockAdapter) GetElection(ctx context.Context, db repository.DBQueries, id int) (model.Election, error) {
	return model.Election{}, nil
}

func TestElectionListCreateHandler(t *testing.T) {
	adapter = MockAdapter{}

	t.Run("election list handler", func(t *testing.T) {
		request, err := http.NewRequest("GET", "/elections", nil)
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		ElectionListCreateHandler(response, request)

		expectedElection := []model.Election{
			model.Election{
				Id:   1,
				Name: "Zé do Caixão",
			},
		}

		var returnedBody []model.Election

		body, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}

		err = json.Unmarshal(body, &returnedBody)

		if !slices.Equal(expectedElection, returnedBody) {
			t.Fatalf("expected %v, returned %v", expectedElection, returnedBody)
		}
	})

	t.Run("election create handler", func(t *testing.T) {})
}
