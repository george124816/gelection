package http

import (
	"log"
	"net/http"

	otel "github.com/george124816/gelection/internal"
	"github.com/george124816/gelection/internal/candidate/handler"
	"github.com/george124816/gelection/internal/configs"
	engine "github.com/george124816/gelection/internal/db"
	electionHandler "github.com/george124816/gelection/internal/election/handler"
)

func Start() {
	config := configs.HttpConfig{Port: 4000}

	engine.Connect()

	err := otel.StartExporter()

	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

	router.HandleFunc("/elections", electionHandler.ElectionListCreateHandler)
	router.HandleFunc("/elections/{id}", electionHandler.ElectionRetrieveHandler)

	router.HandleFunc("/candidates", handler.CandidateListCreateHandler)
	router.HandleFunc("/candidates/{id}", handler.CandidateRetrieveUpdateDestroyHandler)

	log.Println("starting server on port", config.Port)

	http.ListenAndServe(config.GetStringPort(), router)

}
