package basics

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

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehicle
	Seats int
}

func DisplayInfoPtr(c *Car) {
	fmt.Printf("Car(make=%s, model=%s, year=%d, seats=%d)\n", c.Make, c.Model, c.Year, c.Seats)
}

func (c *Car) DisplayInfoVal() {
	fmt.Printf("Car(make=%s, model=%s, year=%d, seats=%d)\n", c.Make, c.Model, c.Year, c.Seats)
}
