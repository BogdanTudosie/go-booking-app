package models

import "github.com/google/uuid"

type Booking struct {
	Number  string
	Tickets []Ticket
}

// Generates a UUID for the booking
func (b *Booking) GenerateNumber() {
	b.Number = uuid.NewString()
}

// Adds a Ticket to a booking
func (b *Booking) AddTicket(ticket Ticket) {
	b.Tickets = append(b.Tickets, ticket)
}

func (b *Booking) RemoveTicket(withUUID string) {
	for index, value := range b.Tickets {
		if value.TicketNumber == withUUID {
			b.Tickets = append(b.Tickets[:index], b.Tickets[index+1:]...)
		}
	}
}
