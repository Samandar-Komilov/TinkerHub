package basics

import (
	"fmt"
	"reflect"
)

const (
	Low = iota
	Medium
	High
)
const (
	Read = 1 << iota
	Write
	Execute
	Crash
)
const pi = 3.14

type Celsius float64
type Farenheit float64

func Swap_without_tmp(a *int, b *int) {
	// Q4
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func Complex1() {
	// 7
	var c1 complex128 = 1 + 2i

	fmt.Println(real(c1), imag(c1))
}

func Rune1() {
	// 10
	var r1 rune = '💀'

	fmt.Printf("%c, %U\n", r1, r1)
}

func Struct1() {
	type person struct {
		name string
	}
	var fred person

	fmt.Println(fred, fred.name)
}

func Main_vardt() {
	var a, b = 5, 6

	Swap_without_tmp(&a, &b)

	fmt.Println(a, b)

	b1 := true
	// 5
	fmt.Println(!b1)
	// 6
	fmt.Println(Low, Medium, High)
	// 7
	Complex1()
	// 8
	fmt.Printf("%.2f\n", float64(a))
	// 10
	Rune1()
	// 11
	fmt.Println(Read, Write, Execute, Crash)
	// 13
	fmt.Println(pi)
	// 15
	Struct1()
	// 18
	var t1 Celsius = 34.4
	fmt.Println(Farenheit(t1))
	// 19
	fmt.Println(reflect.TypeOf(a))
	// 20
	const x = 10 / 2
	fmt.Println(x)
}
