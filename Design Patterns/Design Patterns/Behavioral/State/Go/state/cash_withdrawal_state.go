package state

import "fmt"

type CashWithdrawalState struct {
	atm *ATM
}

func NewCashWithdrawalState(atm *ATM) *CashWithdrawalState {
	return &CashWithdrawalState{atm: atm}
}

func (cws *CashWithdrawalState) InsertCard(card string) {
	fmt.Print("Card is alreadyt inserted")
}

func (cws *CashWithdrawalState) SelectOperation(transactionType string) {
	fmt.Print("Option is already selected")
}

func (cws *CashWithdrawalState) CashWithdrawal(withdrawalAmount int) {
	fmt.Print("Cash withdrawning...")
	fmt.Print("Cash withdrawn successfully")
}

func (cws *CashWithdrawalState) DisplayBalance() {
	fmt.Print("Unsupported Operation")
}

func (cws *CashWithdrawalState) TransferMoney(accountNumber string, transferAmmount int) {
	fmt.Print("Unsupported Operation")
}

func (cws *CashWithdrawalState) ReturnCard() {
	fmt.Print("Card is alreadyt inserted")
}

func (cws *CashWithdrawalState) Exit() {
	cws.atm.SetCurrentState(NewIdleState(cws.atm))
}
