package main

type SeatType int

const (
	NORMAL SeatType = iota
	GOLD
	PREMIUM
)

type SeatState int

const (
	OCCUPIED SeatState = iota
	UNOCCUPIED
)

type Seat struct {
	id       int
	seatType SeatType
	price    float64
}

func (s Seat) ID() int {
	return s.id
}

type Screen struct {
	id    string
	seats map[int]*Seat
}

func (s *Screen) ID() string {
	return s.id
}

func (s *Screen) Seat(id int) (Seat, error) {
	seat, ok := s.seats[id]
	if !ok {
		return Seat{}, ErrSeatNotFound
	}
	return *seat, nil
}

type Theater struct {
	id      string
	name    string
	screens []Screen
}

func (t Theater) ID() string {
	return t.id
}
