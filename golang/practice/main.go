package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-practice/bank"
	"github.com/go-practice/basics"
	"github.com/go-practice/car"
	"github.com/go-practice/shapes"
)

func main() {
	rect := shapes.Rectangle{Width: 10, Height: 5}
	fmt.Println("Rectangle Area:", rect.Area())

	circle := shapes.Circle{Radius: 7}
	fmt.Println("Circle Area:", circle.Area())

	car1 := car.Car{
		Vehicle: car.Vehicle{Make: "BMW", Model: "X5", Year: 2020},
		Seats:   7,
	}

	car1.DisplayInfoVal()

	account1 := bank.BankAccount{AccountNumber: 1, Balance: 0}

	account1.Deposit(300)
	account1.WithDraw(100)

	account1.CheckBalance()
	account1.WithDraw(400)

	// Variadic function usage
	shapes.Sum(1, 5, 2, 5, 1)

	// Anonymous functions
	basics.Anonym()

	// Closures 1
	basics.Closure1()

	// Closures 2: HTTP middleware
	http.Handle("/", basics.Logger(http.HandlerFunc(basics.MainHandler)))
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
