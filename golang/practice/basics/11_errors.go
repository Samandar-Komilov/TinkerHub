package basics

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// type definitions

type User struct {
	id   int
	name string
}

// package level error variables
var ErrNotFound = errors.New("not found")

func Main_errors() {
	// 1
	res, err := Divide(3, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// 3, 4
	users := []*User{
		{id: 1, name: "John"},
		{id: 2, name: "Jane"},
	}
	fmt.Println(users)
	user, err := FindUser(&users, 2)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("User not found")
		} else {
			fmt.Println("Error finding user:", err)
		}
	} else {
		fmt.Println("User found:", user)
	}

	// 5 - Error wrapping (need more research)

}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return -1, errors.New("b cannot be zero")
	}

	return a / b, nil
}

func FindUser(users *[]*User, id int) (*User, error) {
	for i, v := range *users {
		if v.id == id {
			return (*users)[i], nil
		}
	}

	return nil, ErrNotFound
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
