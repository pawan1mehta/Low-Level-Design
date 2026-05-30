package main

import "fmt"

type BaseHandler struct {
	nextHandler SupportHandler
}

func (bs *BaseHandler) SetNextHandler(nextHandler SupportHandler) SupportHandler {
	bs.nextHandler = nextHandler
	return nextHandler
}

func (bs *BaseHandler) Handle(request string) {
	if bs.nextHandler != nil {
		bs.nextHandler.Handle(request)
	} else {
		fmt.Printf("No handler available for: '%s' - Request rejected!\n", request)
	}
}
