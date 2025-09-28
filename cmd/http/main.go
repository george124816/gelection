package http

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/george124816/gelection/internal/candidate/handler"
	"github.com/george124816/gelection/internal/configs"
	engine "github.com/george124816/gelection/internal/db"
	electionHandler "github.com/george124816/gelection/internal/election/handler"
)

func Start() {
	config := configs.HttpConfig{Port: 4000}

	engine.Connect()

	router := http.NewServeMux()

	router.HandleFunc("/elections", electionHandler.ElectionListCreateHandler)
	router.HandleFunc("/elections/{id}", electionHandler.ElectionRetrieveHandler)

	router.HandleFunc("/candidates", handler.CandidateListCreateHandler)
	router.HandleFunc("/candidates/{id}", handler.CandidateRetrieveUpdateDestroyHandler)

	slog.Info(fmt.Sprintf("starting server on port %d", config.Port))

	http.ListenAndServe(config.GetStringPort(), router)

}
