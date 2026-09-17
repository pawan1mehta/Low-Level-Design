package main

import "errors"

var (
	ErrWarehouseNotFound    = errors.New("warehouse not found")
	ErrInsufficientStock    = errors.New("insufficient stock")
	ErrWarehouseFull        = errors.New("warehouse capacity exceeded")
	ErrNoAvailableWarehouse = errors.New("no available warehouse found")
)
