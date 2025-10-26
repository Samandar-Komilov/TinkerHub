package basics

import (
	"errors"
	"fmt"
	"io"
	"os"
)

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
	errr := InitApp()
	fmt.Println(errr)

	// 6 - errors.Is and errors.As
	err2 := Load()

	var pe ParseError
	if errors.Is(err2, ErrConfig) {
		fmt.Println("is -> config not found")
	}

	if errors.As(err2, &pe) {
		fmt.Println("as - file:", pe.File)
	}

	// 7 - wrapping errors with defer
	fmt.Println(doStuff())
}

// type definitions

type User struct {
	id   int
	name string
}

type ParseError struct {
	File string
}

// package level error variables
var ErrNotFound = errors.New("not found")
var ErrConfig = errors.New("config not found")

// errors as values
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

// Wrapping errors
func ReadConfig() error {
	return errors.New("file not found")
}

func InitApp() error {
	if err := ReadConfig(); err != nil {
		return fmt.Errorf("InitApp failed: %w", err)
	}
	return nil
}

func (e ParseError) Error() string {
	return "parse error in " + e.File
}

func Load() error {
	return fmt.Errorf("load failed: %w", ParseError{"config.toml"})
}

// wrapping errors with defer - wrapping while returning
func doStuff() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("doStuff failed: %w", err)
		}
	}()
	return errors.New("low level failure")
}

// Panic and Recover
