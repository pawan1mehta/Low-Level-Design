package main

import (
	"server"
	"log"
	"time"
)

func main() {
	svr := server.New(
		server.WithTimeout(time.Minute),
		server.WithPort(8080),
		server.WithHost("localhost"),
	)

	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
