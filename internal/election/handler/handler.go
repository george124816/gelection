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

var adapter repository.Adapter = election.DefaultAdapter{}

func ElectionListCreateHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		elections, err := adapter.GetAllElections(context.Background(), engine.Engine)
		if err != nil {
			fmt.Fprintln(w, err)
			return
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

		err = adapter.Create(context.Background(), engine.Engine, election)

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
			return
		}

		election, err := adapter.GetElection(context.Background(), engine.Engine, inputId)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "election_not_found")
			return
		}

		resultJson, err := json.Marshal(election)

		fmt.Println(string(resultJson))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(resultJson))
	}
}
