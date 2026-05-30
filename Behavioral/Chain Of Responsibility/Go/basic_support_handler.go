package main

import "fmt"

type BasicSupportHandler struct {
	*BaseHandler
}

func NewBasicSupportHandler() *BasicSupportHandler {
	return &BasicSupportHandler{
		BaseHandler: &BaseHandler{},
	}
}

func (bs *BasicSupportHandler) Handle(request string) {
	if request == "General Inquiry" {
		fmt.Println("Basic Support: Handling general inquiry.")
	} else {
		fmt.Printf("Can not handle request: '%s' - escalating...\n", request)
		bs.BaseHandler.Handle(request)
	}
}
