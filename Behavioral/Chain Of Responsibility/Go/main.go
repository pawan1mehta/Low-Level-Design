package main

import "fmt"

func main() {
	basic := NewBasicSupportHandler()
	technical := NewTechnicalSupportHandler()
	manager := NewManagerSupportHandler()

	basic.SetNextHandler(technical).SetNextHandler(manager)

	requests := []string{
		"General Inquiry",
		"Technical Issue",
		"Serious Complaint",
		"Billing Issue",
	}

	for _, request := range requests {
		fmt.Printf("Incoming Request: %s \n", request)
		basic.Handle(request)
	}
}
