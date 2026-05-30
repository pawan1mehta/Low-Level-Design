package main

type SupportHandler interface {
	SetNextHandler(handler SupportHandler) SupportHandler
	Handle(request string)
}
