package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type BookingSystem struct {
	movies   map[string]Movie
	theaters map[string]Theater
	shows    map[string]Show
	bookings map[string]Booking
}

func NewBookingSystem() *BookingSystem {
	return &BookingSystem{
		movies:   map[string]Movie{},
		theaters: map[string]Theater{},
		shows:    map[string]Show{},
		bookings: map[string]Booking{},
	}
}

func (bs *BookingSystem) Search(title string) []Movie {
	q := strings.ToLower(strings.TrimSpace(title))
	var movieList []Movie
	for _, movie := range bs.movies {
		if strings.Contains(strings.ToLower(movie.Title()), q) {
			movieList = append(movieList, movie)
		}
	}
	return movieList
}

func (bs *BookingSystem) GetShows(movieID string) []Show {
	var shows []Show
	for _, show := range bs.shows {
		if show.MovieID() == movieID {
			shows = append(shows, show)
		}
	}
	return shows
}

func (bs *BookingSystem) GetAvailableSeats(showID string) ([]Seat, error) {
	show, ok := bs.shows[showID]
	if !ok {
		return nil, ErrShowNotFound
	}
	return show.GetAvailableSeats(), nil
}

func (bs *BookingSystem) Booking(user User, showID string, seatIDs []int) (string, error) {
	show, ok := bs.shows[showID]
	if !ok {
		return "", ErrShowNotFound
	}

	err := show.ReserveSeats(seatIDs)
	if err != nil {
		return "", err
	}

	bookingId := fmt.Sprintf("%06d", rand.IntN(1000000))

	var tickets []Ticket
	for _, seatId := range seatIDs {
		ticketID := fmt.Sprintf("%03d", rand.IntN(1000))
		tickets = append(tickets, NewTicket(ticketID, bookingId, showID, seatId))
	}

	bs.bookings[bookingId] = NewBooking(bookingId, user.ID(), showID, tickets, BOOKED)

	return bookingId, nil
}

func (bs *BookingSystem) Cancel(bookingID string) error {
	booking, ok := bs.bookings[bookingID]
	if !ok {
		return ErrBookingNotFound
	}

	for _, ticket := range booking.Tickets() {
		show, ok := bs.shows[ticket.ShowID()]
		if !ok {
			continue
		}
		show.VacantSeat(ticket.SeatID())
	}

	booking.UpdateState(CANCELED)

	delete(bs.bookings, bookingID)

	return nil
}
