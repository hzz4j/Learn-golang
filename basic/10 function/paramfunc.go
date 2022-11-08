package main

import "fmt"

type addFunc func(num1, num2 int) int

func funcAsArg(f addFunc) int {
	return f(1, 2)
}

func main() {
	var a addFunc = func(num1, num2 int) int {
		return num1 + num2
	}

	fmt.Println(funcAsArg(a))
}
