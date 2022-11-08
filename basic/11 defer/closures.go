package main

import "fmt"

func main() {
	x := 10
	defer func() {
		fmt.Printf("x = %d\n", x)
	}()
	x = 100
	fmt.Println("func end: ", x)
}

/**
func end:  100
x = 100
*/
