package main

import (
	"fmt"
)

// -----------------------
// Struct to represent an Account
// -----------------------
type Account struct {
	OwnerName string
	Balance   float64
	LoanAmount float64
}

// Method to deposit money into the account
func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.Balance += amount
		fmt.Println(a.OwnerName, "deposited", amount)
	} else {
		fmt.Println("Invalid deposit amount")
	}
}

// Method to withdraw money from the account
func (a *Account) Withdraw(amount float64) {
	if amount > 0 && amount <= a.Balance {
		a.Balance -= amount
		fmt.Println(a.OwnerName, "withdrew", amount)
	} else {
		fmt.Println("Invalid withdraw attempt")
	}
}

// Method to take a loan
func (a *Account) Loan(amount float64) {
	if amount > 0 {
		a.LoanAmount += amount
		a.Balance += amount
		fmt.Println(a.OwnerName, "took a loan of", amount)
	} else {
		fmt.Println("Invalid loan amount")
	}
}

// Method to pay back a loan
func (a *Account) PayLoan(amount float64) {
	if amount > 0 && amount <= a.Balance && amount <= a.LoanAmount {
		a.Balance -= amount
		a.LoanAmount -= amount
		fmt.Println(a.OwnerName, "paid", amount, "towards loan")
	} else {
		fmt.Println("Invalid loan payment attempt")
	}
}

// -----------------------
// Struct to represent a Bank
// -----------------------
type Bank struct {
	Name     string
	Accounts []Account
}

// Method to add a new account to the bank
func (b *Bank) AddAccount(ownerName string, initialDeposit float64) {
	account := Account{
		OwnerName: ownerName,
		Balance:   initialDeposit,
	}
	b.Accounts = append(b.Accounts, account)
	fmt.Println("Account created for", ownerName)
}

// Method to find an account by owner name
func (b *Bank) FindAccount(ownerName string) *Account {
	for i := range b.Accounts {
		if b.Accounts[i].OwnerName == ownerName {
			return &b.Accounts[i]
		}
	}
	return nil
}

// -----------------------
// Main function - Simulation
// -----------------------
func main() {
	// Create a new bank
	myBank := Bank{Name: "GoLang National Bank"}
	
	// Add accounts
	myBank.AddAccount("Alice", 1000)
	myBank.AddAccount("Bob", 500)

	// Find accounts
	aliceAccount := myBank.FindAccount("Alice")
	bobAccount := myBank.FindAccount("Bob")

	// Perform operations for Alice
	if aliceAccount != nil {
		aliceAccount.Deposit(500)
		aliceAccount.Withdraw(200)
		aliceAccount.Loan(1000)
		aliceAccount.PayLoan(300)
		fmt.Println("Final Balance:", aliceAccount.Balance, "Loan Remaining:", aliceAccount.LoanAmount)
	}

	// Perform operations for Bob
	if bobAccount != nil {
		bobAccount.Deposit(300)
		bobAccount.Withdraw(100)
		bobAccount.Loan(500)
		bobAccount.PayLoan(200)
		fmt.Println("Final Balance:", bobAccount.Balance, "Loan Remaining:", bobAccount.LoanAmount)
	}
}
