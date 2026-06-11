package state

type ATMState interface {
	InsertCard(card string)
	SelectOperation(operationType string)
	CashWithdrawal(withdrawalAmount int)
	DisplayBalance()
	TransferMoney(accountNumber string, transferAmount int)
	ReturnCard()
	Exit()
}
