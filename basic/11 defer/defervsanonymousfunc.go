package main

import "fmt"

func main() {
	x := 2
	defer func(num int) {
		fmt.Printf("x = %d\n", num)
	}(x)
	x = 10
	fmt.Printf("x = %d\n", x)
}

/**
x = 10
x = 2
*/
