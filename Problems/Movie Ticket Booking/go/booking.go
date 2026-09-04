package main

type BookingState int

const (
	PENDING BookingState = iota
	BOOKED
	CANCELED
)

type Booking struct {
	id      string
	userID  string
	showID  string
	tickets []Ticket
	state   BookingState
}

func NewBooking(id, userID, showID string, ticketIDs []Ticket, state BookingState) Booking {
	return Booking{
		id:      id,
		userID:  userID,
		showID:  showID,
		tickets: ticketIDs,
		state:   state,
	}
}

func (b Booking) ID() string {
	return b.id
}

func (b Booking) Tickets() []Ticket {
	return b.tickets
}

func (b Booking) UpdateState(canceled BookingState) {
	b.state = canceled
}
