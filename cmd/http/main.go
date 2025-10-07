package http

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/george124816/gelection/internal/candidate/handler"
	"github.com/george124816/gelection/internal/configs"
	engine "github.com/george124816/gelection/internal/db"
	electionHandler "github.com/george124816/gelection/internal/election/handler"
	healthHandler "github.com/george124816/gelection/internal/health/handler"
	voteHandler "github.com/george124816/gelection/internal/vote/handler"
)

func Start() (*http.Server, error) {
	config := configs.HttpConfig{Port: 4000}

	engine.Connect()

	router := http.NewServeMux()

	router.HandleFunc("/elections", electionHandler.ElectionListCreateHandler)
	router.HandleFunc("/elections/{id}", electionHandler.ElectionRetrieveHandler)

	router.HandleFunc("/candidates", handler.CandidateListCreateHandler)
	router.HandleFunc("/candidates/{id}", handler.CandidateRetrieveUpdateDestroyHandler)

	router.HandleFunc("/health", healthHandler.HealthCheckHandler)

	router.HandleFunc("/votes", voteHandler.VoteListCreateHandler)

	slog.Info(fmt.Sprintf("starting server on port %d", config.Port))

	server := &http.Server{
		Addr:    config.GetStringPort(),
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			slog.Error(err.Error())
		}
	}()

	return server, nil

}
