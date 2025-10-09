package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	engine "github.com/george124816/gelection/internal/db"
	"github.com/george124816/gelection/internal/election/model"
	"github.com/george124816/gelection/internal/election/repository"
	election "github.com/george124816/gelection/internal/election/repository"
)

func ElectionListCreateHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		elections, err := election.GetAllElections(context.Background(), engine.Engine)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		resultJson, err := json.Marshal(elections)

		if err != nil {
			fmt.Fprintln(w, err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(resultJson))
	case r.Method == "POST":
		var election model.Election
		bodyRequest, err := io.ReadAll(r.Body)

		if err != nil {
			fmt.Fprintln(w, err)
		}

		err = json.Unmarshal(bodyRequest, &election)

		if err != nil {
			fmt.Fprintln(w, err)
		}

		err = repository.Create(context.Background(), engine.Engine, election)

		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "created")
	}

}

func ElectionRetrieveHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		inputId, err := strconv.Atoi(r.PathValue("id"))

		if err != nil {
			fmt.Fprintln(w, err)
		}

		election, err := election.GetElection(context.Background(), engine.Engine, inputId)

		if err != nil {
			fmt.Fprintln(w, err)
		}

		resultJson, err := json.Marshal(election)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(resultJson))
	}
}
