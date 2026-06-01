package main

import "state/state"

func main() {
	atm := state.NewATM()

	atm.InsertCard("atm_card")
	atm.SelectOperation("withdraw")
	atm.TransferMoney("account_number", 0)
}
