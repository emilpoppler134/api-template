package main

import (
	"github.com/emilpoppler134/api-template/internal/config"
	"github.com/emilpoppler134/api-template/internal/db"
	"github.com/emilpoppler134/api-template/internal/handlers"
	"github.com/emilpoppler134/api-template/internal/http"
)

func main() {
	// Load configuration
	configuration := config.Load()

	// Connect to database and Initialize handlers
	database := db.Connect(configuration.DatabaseDSN)
	handlers := handlers.Init(database)

	// Initialize, Register and Run server
	server := http.Init(configuration.Port)
	server.Register(handlers)
	server.Run()
}
