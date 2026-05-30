package main

func main() {
	bankAccount := NewRealBankAccount()

	proxyBankAccount := NewProxyBankAccount("Pawan", "admin", bankAccount)

	proxyBankAccount.Deposit(456.0)
	proxyBankAccount.Deposit(456.0)
	proxyBankAccount.Deposit(456.0)

	proxyBankAccount.Withdraw(400.0)

	proxyBankAccount.PrintLogs()
}
