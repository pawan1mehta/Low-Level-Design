package state

import "fmt"

type IdleState struct {
	atm *ATM
}

func NewIdleState(atm *ATM) *IdleState {
	return &IdleState{atm: atm}
}

func (idleState *IdleState) InsertCard(card string) {
	fmt.Print("Inserting atm card...")
	idleState.atm.SetATMCard(card)
	idleState.atm.SetCurrentState(NewHasCardState(idleState.atm))
}

func (idleState *IdleState) SelectOperation(operationType string) {
	fmt.Print("Please insert card first!")
}

func (idleState *IdleState) CashWithdrawal(withdrawalAmount int) {
	fmt.Print("Please insert card first!")
}

func (idleState *IdleState) DisplayBalance() {
	fmt.Print("Please insert card first!")
}

func (idleState *IdleState) TransferMoney(accountNumber string, transferAmount int) {
	fmt.Print("Please insert card first!")
}

func (idleState *IdleState) ReturnCard() {
	fmt.Print("No card to return!")
}

func (idleState *IdleState) Exit() {
	idleState.atm.SetCurrentState(NewIdleState(idleState.atm))
}
