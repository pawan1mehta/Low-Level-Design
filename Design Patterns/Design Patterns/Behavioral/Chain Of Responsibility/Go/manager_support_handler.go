package main

import "fmt"

type ManagerSupportHandler struct {
	*BaseHandler
}

func NewManagerSupportHandler() *ManagerSupportHandler {
	return &ManagerSupportHandler{
		BaseHandler: &BaseHandler{},
	}
}

func (ms *ManagerSupportHandler) Handle(request string) {
	if request == "Serious Complaint" {
		fmt.Println("Manager Support: Handling serious complaint.")
	} else {
		fmt.Printf("Can not handle request: '%s' - escalating...\n", request)
		ms.BaseHandler.Handle(request)
	}
}
