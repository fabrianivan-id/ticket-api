package main

import (
	"log"
	"net/http"

	"github.com/fabrianivan-id/ticket-api/handlers"
	"github.com/fabrianivan-id/ticket-api/utils"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	utils.InitDB()

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/tickets", handlers.CreateTicketHandler).Methods("POST")

	// Start the server
	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
