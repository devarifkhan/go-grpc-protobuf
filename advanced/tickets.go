package main

import (
	"fmt"
	"sync"
	"time"
)

// Ticket represents a support ticket
type Ticket struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	Status    TicketStatus
}

// TicketStatus represents the status of a ticket
type TicketStatus int

const (
	Open TicketStatus = iota
	InProgress
	Resolved
	Closed
)

// TicketManager manages tickets
type TicketManager struct {
	tickets map[int]*Ticket
	nextID  int
	mu      sync.Mutex
}

// NewTicketManager creates a new ticket manager
func NewTicketManager() *TicketManager {
	return &TicketManager{
		tickets: make(map[int]*Ticket),
		nextID:  1,
	}
}

// CreateTicket adds a new ticket to the system
func (tm *TicketManager) CreateTicket(title, content string) *Ticket {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	ticket := &Ticket{
		ID:        tm.nextID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		Status:    Open,
	}

	tm.tickets[ticket.ID] = ticket
	tm.nextID++
	return ticket
}

// GetTicket retrieves a ticket by ID
func (tm *TicketManager) GetTicket(id int) (*Ticket, bool) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	ticket, exists := tm.tickets[id]
	return ticket, exists
}

// UpdateStatus changes the status of a ticket
func (tm *TicketManager) UpdateStatus(id int, status TicketStatus) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	ticket, exists := tm.tickets[id]
	if !exists {
		return false
	}

	ticket.Status = status
	return true
}

func main() {
	tm := NewTicketManager()

	// Create some tickets
	ticket1 := tm.CreateTicket("Server Down", "The main server is not responding")
	ticket2 := tm.CreateTicket("Database Error", "Getting timeout on DB connections")

	fmt.Printf("Created ticket: #%d - %s\n", ticket1.ID, ticket1.Title)
	fmt.Printf("Created ticket: #%d - %s\n", ticket2.ID, ticket2.Title)

	// Update status
	tm.UpdateStatus(1, InProgress)

	// Get a ticket
	if t, exists := tm.GetTicket(1); exists {
		fmt.Printf("Ticket #%d status: %v\n", t.ID, t.Status)
	}
}
