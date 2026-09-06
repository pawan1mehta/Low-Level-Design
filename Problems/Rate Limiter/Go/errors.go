package main

import "errors"

var (
	ErrEndpointConfigAlreadyExists = errors.New("endpoint config already exists")
	ErrConfigMismatch              = errors.New("config mismatch")
	ErrUnknowConfigType            = errors.New("unknow config type")
)
