package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Ticket struct {
	TicketNumber string
	Class        string
	Cost         int
	Steamer      Steamer
}

func (t *Ticket) BookTicket() {
	if t.Steamer.TicketsLeft == 0 {
		fmt.Println("No more tickets left, please try another vessel to become available")
		return
	}
	t.Steamer.TicketsLeft -= 1
}

func (t *Ticket) GenerateTicketNumber() {
	t.TicketNumber = uuid.NewString()
}
