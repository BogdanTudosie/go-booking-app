package models

import "fmt"

type Ticket struct {
	Class   string
	Cost    int
	Steamer Steamer
}

func (t *Ticket) BookTicket() {
	if t.Steamer.TicketsLeft == 0 {
		fmt.Println("No more tickets left, please try another vessel to become available")
		return
	}
	t.Steamer.TicketsLeft -= 1
}
