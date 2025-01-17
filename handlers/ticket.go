package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fabrianivan-id/ticket-api/models"
	db "github.com/fabrianivan-id/ticket-api/utils"
)

func CreateTicketHandler(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket

	// Parse the incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %s", err), http.StatusBadRequest)
		return
	}

	// Validation
	if ticket.TicketTitle == "" || len(ticket.TicketTitle) < 10 || len(ticket.TicketTitle) > 100 {
		http.Error(w, "ticket_title must be between 10 and 100 characters", http.StatusBadRequest)
		return
	}

	if ticket.TicketMsg == "" || len(ticket.TicketMsg) < 100 {
		http.Error(w, "ticket_msg must be at least 100 characters", http.StatusBadRequest)
		return
	}

	if ticket.UserID <= 0 {
		http.Error(w, "user_id must be a positive integer", http.StatusBadRequest)
		return
	}

	ticket.Status = "opn" // Default status to 'Open'

	// Insert into the database
	query := `INSERT INTO tickets (ticket_title, ticket_msg, user_id, status, created_at) 
              VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id`

	var id int
	err := db.DB.QueryRow(query, ticket.TicketTitle, ticket.TicketMsg, ticket.UserID, ticket.Status).Scan(&id)
	if err != nil {
		log.Println("Error inserting ticket:", err)
		http.Error(w, "Unable to create ticket", http.StatusInternalServerError)
		return
	}

	ticket.ID = id
	ticket.CreatedAt = time.Now()

	// Respond with the newly created ticket
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}

func GetTicketListHandler(w http.ResponseWriter, r *http.Request) {
	createdAt := r.URL.Query().Get("created_at")
	sort := r.URL.Query().Get("sort")
	pageSize := r.URL.Query().Get("page_size")
	pageSizeInt, _ := strconv.Atoi(pageSize)

	page := r.URL.Query().Get("page")
	pageInt, _ := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	if pageSizeInt < 10 {
		pageSizeInt = 10
	} else if pageSizeInt > 50 {
		pageSizeInt = 50
	}

	query := "SELECT id, ticket_title, ticket_msg, user_id, status, created_at FROM tickets WHERE status = 'opn'"

	if createdAt != "" {
		query += " AND created_at < $1"
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSizeInt, (pageInt-1)*pageSizeInt)

	if sort == "asc" {
		query += " ORDER BY created_at ASC"
	} else if sort == "desc" {
		query += " ORDER BY created_at DESC"
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSizeInt, (page-1)*pageSizeInt)

	var tickets []models.Ticket
	err := db.DB.Select(&tickets, query)
	if err != nil {
		http.Error(w, "Error fetching tickets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}