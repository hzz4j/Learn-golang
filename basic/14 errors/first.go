package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {
	if _, err := divide(2, 0); err != nil {
		fmt.Println(err.Error())
	}
}
