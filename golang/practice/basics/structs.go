package basics

import (
	"fmt"
)

func Main_structs() {
	// 1-6, 12, 13
	p1 := newPerson("Eshmat", 12, "Tashkent", "Uzbekistan")
	p1.greeting()
	// 7: Compare 2 structs
	p2 := newPerson("Eshmat", 13, "Tashkent", "Uzbekistan")
	fmt.Println(p1 == p2)
	// 8: Add JSON tags to Person fields
	// 9-10: Marshal and Unmarshal a Person struct to JSON
	// 11: Create anonymous struct and use
	// 14: Add methods to embedded struct
}

// 1-6
type Person struct {
	Name    string  `json:"name"`
	Age     uint8   `json:"age"`
	Address Address `json:"address"`
}

type Address struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func newPerson(name string, age uint8, city string, country string) Person {
	p := Person{Name: name, Age: age, Address: Address{City: city, Country: country}}

	return p
}

func (p Person) greeting() {
	fmt.Printf("Hello %s! You are %d years old and live in %s, %s.\n", p.Name, p.Age, p.Address.City, p.Address.Country)
}
