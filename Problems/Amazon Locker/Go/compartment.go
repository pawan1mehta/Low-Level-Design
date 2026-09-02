package main

import "time"

type AccessToken struct {
	code        string
	expiration  time.Time
	compartment *Compartment
}

func NewAccessToken(code string, expiration time.Time, compartment *Compartment) *AccessToken {
	return &AccessToken{
		code:        code,
		expiration:  expiration,
		compartment: compartment,
	}
}

func (a *AccessToken) IsExpired() bool {
	if time.Now().After(a.expiration) {
		return true
	}
	return false
}

func (a *AccessToken) GetCompartment() *Compartment {
	return a.compartment
}

func (a *AccessToken) GetCode() string {
	return a.code
}
