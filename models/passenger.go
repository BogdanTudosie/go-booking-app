package models

type Passenger struct {
	FirstName string
	LastName  string
	Email     string
	Telephone string
	Bookings  []Booking // history of bookings created by this passenger
}

func (p *Passenger) RemoveBooking(withUUID string) {
	for index, value := range p.Bookings {
		if value.Number == withUUID {
			p.Bookings = append(p.Bookings[:index], p.Bookings[index+1:]...)
		}
	}
}
