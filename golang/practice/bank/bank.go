package bank

import "fmt"

type BankAccount struct {
	AccountNumber int
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance = b.Balance + amount
}

func (b *BankAccount) WithDraw(amount float64) error {
	if b.Balance < amount {
		return fmt.Errorf("insufficient funds")
	}

	b.Balance = b.Balance - amount
	return nil
}

func (b *BankAccount) CheckBalance() {
	fmt.Printf("Current balance: %f", b.Balance)
}
