package main

import "fmt"

const (
	RoleAdmin = "admin"
	RoleGuest = "gues"
)

type ProxyBankAccount struct {
	user            string
	role            string
	realBankAccount *RealBankAccount
	logs            []string
}

func NewProxyBankAccount(user, role string, bankAccount *RealBankAccount) *ProxyBankAccount {
	return &ProxyBankAccount{
		user:            user,
		role:            role,
		realBankAccount: bankAccount,
		logs:            []string{},
	}
}

func (p *ProxyBankAccount) Deposit(amount float32) {
	log := fmt.Sprintf("User: %-8s | Role: %-6s | Action: Deposit  | Amount: $%.2f", p.user, p.role, amount)
	p.logs = append(p.logs, log)

	p.realBankAccount.Deposit(amount)
}

func (p *ProxyBankAccount) Withdraw(amount float32) {
	log := fmt.Sprintf("User: %-8s | Role: %-6s | Action: Withdraw | Amount: $%.2f", p.user, p.role, amount)
	p.logs = append(p.logs, log)
	p.realBankAccount.Withdraw(amount)
}

func (p *ProxyBankAccount) Balance() float32 {
	return p.realBankAccount.Balance()
}

func (p *ProxyBankAccount) PrintLogs() {
	fmt.Println("\n Transaction Logs:")
	for _, log := range p.logs {
		fmt.Printf(log)
	}
}
