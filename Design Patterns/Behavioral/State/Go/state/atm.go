package state

type ATM struct {
	currentState ATMState
	atmCard      string
}

func NewATM() *ATM {
	atm := &ATM{}
	atm.currentState = NewIdleState(atm)
	return atm
}

func (atm *ATM) SetCurrentState(state ATMState) {
	atm.currentState = state
}

func (atm *ATM) SetATMCard(card string) {
	atm.atmCard = card
}

func (atm *ATM) InsertCard(card string) {
	atm.currentState.InsertCard(card)
}

func (atm *ATM) SelectOperation(operationType string) {
	atm.currentState.SelectOperation(operationType)
}

func (atm *ATM) CashWithdrawal(withdrawalAmount int) {
	atm.currentState.CashWithdrawal(withdrawalAmount)
}

func (atm *ATM) DisplayBalance() {
	atm.currentState.DisplayBalance()
}

func (atm *ATM) TransferMoney(accountNumber string, transferAmount int) {
	atm.currentState.TransferMoney(accountNumber, transferAmount)
}

func (atm *ATM) ReturnCard() {
	atm.currentState.ReturnCard()
}

func (atm *ATM) Exit() {
	atm.currentState.Exit()
}
