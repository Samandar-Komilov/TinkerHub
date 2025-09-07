package basics

import (
	"fmt"
	"io"
	"math"
	"os"
)

// Type and interface definitions

type Shape interface {
	Area() float64
	Perimeter() float64
}

type CustomReader struct {
	io.Reader
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func Main_interfaces() {
	r1 := Rectangle{Width: 5, Height: 5}
	c1 := Circle{Radius: 4.5}
	PrintShapeInfo(c1)
	PrintShapeInfo(r1)

	// 7
	var slice []any
	slice = append(slice, 10, "hello", true)
	fmt.Println(slice)

	// 8
	var i any = "hello"
	value, ok := i.(int)
	if ok {
		fmt.Println("Value is", value)
	} else {
		fmt.Println("Type assertion failed.")
	}

	// 11
	myReader := CustomReader{Reader: os.Stdin}
	buf := make([]byte, 1024)
	n, err := myReader.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

// Struct methods

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Utils

func PrintShapeInfo(s Shape) {
	fmt.Println("Area:", s.Area(), "Perimeter:", s.Perimeter())
}
