package main

import "fmt"

func main() {
	var i int
	typeAssert(i)
}

func typeAssert(i interface{}) {
	if v, ok := i.(int); ok {
		fmt.Printf("i is int -> %d\n", v)
	} else if v, ok := i.(float32); ok {
		fmt.Printf("i is float32 -> %f\n", v)
	} else {
		fmt.Printf("i is not int or float32\n")
	}
}
