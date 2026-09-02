package main

func main() {
	var compartments []*Compartment
	for i := 0; i <= 10; i++ {
		compartments = append(compartments, NewCompartment(SMALL, false))
	}
	for i := 0; i <= 10; i++ {
		compartments = append(compartments, NewCompartment(MEDIUM, false))
	}
	for i := 0; i <= 10; i++ {
		compartments = append(compartments, NewCompartment(LARGE, false))
	}

	locker := NewLocker(compartments)

	accessTokenCode, _ := locker.DepositPackage(SMALL)
	accessTokenCode, _ = locker.DepositPackage(MEDIUM)
	accessTokenCode, _ = locker.DepositPackage(LARGE)

	locker.Pickup(accessTokenCode)

}
