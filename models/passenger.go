package models

type Passenger struct {
	FirstName string
	LastName  string
	Email     string
	Telephone string
	Booking   []Booking
}
