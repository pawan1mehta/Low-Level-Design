package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Locker struct {
	compartments   []*Compartment
	accessTokenMap map[string]*AccessToken
}

func NewLocker(compartments []*Compartment) *Locker {
	return &Locker{
		compartments:   compartments,
		accessTokenMap: make(map[string]*AccessToken),
	}
}

func (l *Locker) DepositPackage(size Size) (string, error) {
	compartment := l.findUnOccupiedCompartmentBySize(size)
	if compartment == nil {
		return "", fmt.Errorf("no available compartment of size %s", size)
	}

	compartment.Open()
	compartment.MarkOccupied()
	accessToken := l.generateAccessToken(compartment)
	l.accessTokenMap[accessToken.GetCode()] = accessToken

	return accessToken.GetCode(), nil
}

func (l *Locker) Pickup(accessCode string) error {
	accessToken, ok := l.accessTokenMap[accessCode]
	if !ok {
		return fmt.Errorf("invalid access token code")
	}
	if accessToken.IsExpired() {
		return fmt.Errorf("access token has expired")
	}

	compartment := accessToken.GetCompartment()
	compartment.Open()
	l.clearDeposit(accessToken)

	return nil
}

func (l *Locker) OpenExpiredCompartment() {
	for _, accessToken := range l.accessTokenMap {
		if accessToken.IsExpired() {
			accessToken.GetCompartment().Open()
		}
	}
}

func (l *Locker) findUnOccupiedCompartmentBySize(size Size) *Compartment {
	for _, compartment := range l.compartments {
		if !compartment.IsOccupied() && compartment.GetSize() == size {
			return compartment
		}
	}
	return nil
}

func (l *Locker) generateAccessToken(compartment *Compartment) *AccessToken {
	code := fmt.Sprintf("%06d", rand.IntN(1000000))
	expiration := time.Now().Add(7 * 24 * time.Hour)
	return NewAccessToken(code, expiration, compartment)
}

func (l *Locker) clearDeposit(token *AccessToken) {
	compartment := token.GetCompartment()
	compartment.MarkFree()
	delete(l.accessTokenMap, token.GetCode())
}
