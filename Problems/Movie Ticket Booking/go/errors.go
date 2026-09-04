package main

import "errors"

var (
	ErrShowNotFound    = errors.New("show not found")
	ErrSeatReserved    = errors.New("seat already reserved")
	ErrBookingNotFound = errors.New("booking not found")
	ErrSeatNotFound    = errors.New("seat not found")
)
