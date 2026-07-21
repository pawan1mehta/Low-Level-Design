package state

import "fmt"

type HasCardState struct {
	atm *ATM
}

func NewHasCardState(atm *ATM) *HasCardState {
	return &HasCardState{atm: atm}
}

func (hasCardState *HasCardState) InsertCard(card string) {
	fmt.Print("Card is already inserted!")
}

func (hasCardState *HasCardState) SelectOperation(operationType string) {
	fmt.Println("Operation selected: ", operationType)
}

func (hasCardState *HasCardState) CashWithdrawal(withdrawalAmount int) {
	fmt.Print("Please authenticate pin first!")
}

func (hasCardState *HasCardState) DisplayBalance() {
	fmt.Print("Please authenticate pin first!")
}

func (hasCardState *HasCardState) TransferMoney(accountNumber string, transferAmount int) {
	fmt.Print("Please authenticate pin first!")
}

func (hasCardState *HasCardState) ReturnCard() {
	fmt.Println("Card return successfully!")
	hasCardState.atm.SetCurrentState(NewIdleState(hasCardState.atm))
}

func (hasCardState *HasCardState) Exit() {
	fmt.Println("Machine exited.")
	hasCardState.atm.SetCurrentState(NewIdleState(hasCardState.atm))
}
