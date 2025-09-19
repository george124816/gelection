package http

import (
	"log"
	"net/http"

	candidateHandler "github.com/george124816/gelection/internal/candidate/handler"
	"github.com/george124816/gelection/internal/configs"
	engine "github.com/george124816/gelection/internal/db"
	electionHandler "github.com/george124816/gelection/internal/election/handler"
)

func Start() {
	config := configs.HttpConfig{Port: 4000}

	engine.Connect()

	router := http.NewServeMux()

	// Election Routes
	router.HandleFunc("/elections", electionHandler.ElectionHandler)
	router.HandleFunc("/election/{id}", electionHandler.ElectionHandler)
	router.HandleFunc("/election", electionHandler.ElectionHandler)

	// Candidate Routes
	router.HandleFunc("/candidate", candidateHandler.CandidateHandler)
	router.HandleFunc("/candidates", candidateHandler.CandidateHandler)
	router.HandleFunc("/candidate/{id}", candidateHandler.CandidateHandler)

	log.Println("starting server on port", config.Port)

	http.ListenAndServe(config.GetStringPort(), router)

}
