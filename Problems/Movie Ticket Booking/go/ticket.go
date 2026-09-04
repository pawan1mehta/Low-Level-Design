package main

type Ticket struct {
	id        string
	bookingID string
	showID    string
	seatID    int
}

func NewTicket(id, bookingId, showID string, seatId int) Ticket {
	return Ticket{
		id:        id,
		bookingID: bookingId,
		showID:    showID,
		seatID:    seatId,
	}
}

func (t Ticket) ID() string {
	return t.id
}

func (t Ticket) ShowID() string {
	return t.showID
}

func (t Ticket) SeatID() int {
	return t.seatID
}
