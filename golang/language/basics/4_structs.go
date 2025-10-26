package basics

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

/*
IGNORED: 15, 16, 25, 30, 31, 34, 35, 37, 38, 39, 40
*/

func Main_structs() {
	// 1-6, 12, 13
	p1 := newPerson("Eshmat", 12, "Tashkent", "Uzbekistan")
	p1.greeting()
	// 7: Compare 2 structs
	p2 := newPerson("Toshmat", 13, "Samarkand", "Uzbekistan")
	fmt.Println(p1 == p2)
	// 8: Add JSON tags to Person fields
	// Done

	// 9-10: Marshal and Unmarshal a Person struct to JSON
	jsonBytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Marshalled JSON: ", string(jsonBytes))

	var p3 Person
	err = json.Unmarshal(jsonBytes, &p3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Unmarshalled Person: ", p3)

	// 11: Create anonymous struct and use
	ast1 := struct {
		Id    int     `json:"id"`
		Grade float32 `json:"grade"`
	}{
		Id:    1,
		Grade: 3.45,
	}
	fmt.Println("Anonymous struct: ", ast1)

	// 14: Add methods to embedded struct
	p1.Address.print_address()

	// 17: Create a slice of structs
	var slc1 []Person
	slc1 = append(slc1, p1)
	slc1 = append(slc1, p2)
	fmt.Println("Slice array", slc1)

	// 18: Sort slices based on various fields
	// slices.SortFunc(slc1, compareStructs())

	// 19: Create a map of structs
	mp1 := make(map[int]Person)
	mp1[1] = p1
	mp1[2] = p2
	fmt.Println("Map of structs:", mp1)

	// 20-21: Change embedded struct field with pointer received method
	p3.Address.changeCity("Bukhara")
	fmt.Println("Struct 3:", p3)

	// 22-23: Working with XML
	xmlBytes, err := xml.Marshal(p3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(xmlBytes))

	// 26-27: Use struct as map key and create zero value struct
	var zerostc Person
	mp2 := make(map[Person]int)
	mp2[zerostc] = 0
	mp2[p1] = 1
	fmt.Println(mp2)

	// 28: Add validations to struct fields
	// Using go-playground/validator, soon in Network Programming we'll use

	// 36: Create a struct with a function field (similar to function pointers in C)
	mystruct1 := MyStruct{
		field1: func() string {
			return "Hello World!"
		},
	}
	fmt.Println(mystruct1, mystruct1.field1())

}

// 1-6
type Person struct {
	Name    string  `json:"name" xml:"name"`
	Age     uint8   `json:"age" xml:"age"`
	Address Address `json:"address" xml:"address"`
}

type Address struct {
	Country string `json:"country" xml:"country"`
	City    string `json:"city" xml:"city"`
}

type MyStruct struct {
	field1 func() string
}

func newPerson(name string, age uint8, city string, country string) Person {
	p := Person{Name: name, Age: age, Address: Address{City: city, Country: country}}

	return p
}

func (p Person) greeting() {
	fmt.Printf("Hello %s! You are %d years old and live in %s, %s.\n", p.Name, p.Age, p.Address.City, p.Address.Country)
}

func (a Address) print_address() {
	fmt.Printf("Location: City=%s Country=%s.\n", a.City, a.Country)
}

func (a *Address) changeCity(val string) {
	if val == "" || strings.TrimSpace(val) == "" {
		fmt.Println("Empty or whitespace-only string value!")
		return
	}
	a.City = val
}
