package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	engine "github.com/george124816/gelection/internal/db"
	"github.com/george124816/gelection/internal/vote/model"
	"github.com/george124816/gelection/internal/vote/repository"
	vote "github.com/george124816/gelection/internal/vote/repository"
)

func VoteListCreateHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		votes, err := vote.GetAllVotes(context.Background(), engine.Engine)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		resultJson, err := json.Marshal(votes)

		if err != nil {
			fmt.Fprintln(w, err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(resultJson))
	case r.Method == "POST":
		var vote model.Vote
		bodyRequest, err := io.ReadAll(r.Body)

		if err != nil {
			fmt.Fprintln(w, err)
		}

		err = json.Unmarshal(bodyRequest, &vote)

		if err != nil {
			fmt.Fprintln(w, err)
		}

		err = repository.Create(context.Background(), engine.Engine, vote)

		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "created")
	}

}
