package main

import "fmt"

type TechnicalSupportHandler struct {
	*BaseHandler
}

func NewTechnicalSupportHandler() *TechnicalSupportHandler {
	return &TechnicalSupportHandler{
		BaseHandler: &BaseHandler{},
	}
}

func (ts *TechnicalSupportHandler) Handle(request string) {
	if request == "Technical Issue" {
		fmt.Println("Technical Support: Handling technical issue.")
	} else {
		fmt.Printf("Can not handle request: '%s' - escalating...\n", request)
		ts.BaseHandler.Handle(request)
	}
}
