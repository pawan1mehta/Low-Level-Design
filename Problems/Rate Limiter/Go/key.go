package main

type Key struct {
	ClientID string
	Endpoint string
}

func NewKey(clientID string, endpoint string) Key {
	return Key{
		ClientID: clientID,
		Endpoint: endpoint,
	}
}
