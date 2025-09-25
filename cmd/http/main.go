package http

import (
	"log"
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

	router.HandleFunc("/elections", electionHandler.ElectionHandler)
	router.HandleFunc("/election/{id}", electionHandler.ElectionHandler)
	router.HandleFunc("/election", electionHandler.ElectionHandler)

	router.HandleFunc("/candidates", handler.CandidateListCreateHandler)
	router.HandleFunc("/candidates/{id}", handler.CandidateRetrieveUpdateDestroyHandler)

	log.Println("starting server on port", config.Port)

	http.ListenAndServe(config.GetStringPort(), router)

}
