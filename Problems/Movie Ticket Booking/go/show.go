package main

import (
	"sync"
	"time"
)

type Show struct {
	id        string
	movieId   string
	screen    *Screen
	seating   *ShowSeating
	startTime time.Time
}

func (s *Show) MovieID() string {
	return s.movieId
}

func (s *Show) ReserveSeats(seatIds []int) error {
	return s.seating.ReserveSeats(seatIds)
}

func (s *Show) GetAvailableSeats() []Seat {
	return s.seating.Available(s.screen)
}

func (s *Show) VacantSeat(id int) error {
	return s.seating.VacantSeat(id)
}

type ShowSeating struct {
	status map[int]SeatState
	mu     sync.Mutex
}

func (s *ShowSeating) ReserveSeats(ids []int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var seats []int

	for _, id := range ids {
		st, ok := s.status[id]
		if !ok {
			return ErrSeatNotFound
		}
		if st == OCCUPIED {
			return ErrSeatReserved
		}
		seats = append(seats, id)
	}

	for _, id := range seats {
		s.status[id] = OCCUPIED
	}

	return nil
}

func (s *ShowSeating) VacantSeat(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.status[id]; !ok {
		return ErrSeatNotFound
	}
	s.status[id] = UNOCCUPIED
	return nil
}

func (s *ShowSeating) Available(screen *Screen) []Seat {
	s.mu.Lock()
	defer s.mu.Unlock()

	var free []Seat
	for id, status := range s.status {
		if status != UNOCCUPIED {
			continue
		}
		seat, err := screen.Seat(id)
		if err != nil {
			continue
		}
		free = append(free, seat)
	}

	return free
}
