package main

import "fmt"

type BankAccount interface {
	Deposit(amount float32)
	Withdraw(amount float32)
	Balance() float32
}

type RealBankAccount struct {
	balance float32
}

func NewRealBankAccount() *RealBankAccount {
	return &RealBankAccount{
		balance: 0.0,
	}
}

func (rb *RealBankAccount) Deposit(amount float32) {
	rb.balance = rb.balance + amount
	fmt.Printf("Deposit %g , new balance is %g \n", amount, rb.balance)
}

func (rb *RealBankAccount) Withdraw(amount float32) {
	if rb.balance >= amount {
		rb.balance = rb.balance - amount
		fmt.Printf("Withdrew %g, new balance is %g \n", amount, rb.balance)
	} else {
		fmt.Println("Insufficient funds, withdrawal denied.")
	}
}

func (rb *RealBankAccount) Balance() float32 {
	return rb.balance
}
