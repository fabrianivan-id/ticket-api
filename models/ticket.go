package models

import "time"

// Ticket represents the ticket structure
type Ticket struct {
	ID          int       `json:"id"`
	TicketTitle string    `json:"ticket_title"`
	TicketMsg   string    `json:"ticket_msg"`
	UserID      int       `json:"user_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
