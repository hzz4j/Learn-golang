package main

import "fmt"

func typeAssert(t interface{}) {
	switch v := t.(type) {
	case int:
		fmt.Printf("i is int -> %d\n", v)
	case float32:
		fmt.Printf("i is float32 -> %f\n", v)
	default:
		fmt.Printf("i is %T -> %v\n", v, v)
	}
}

func main() {
	typeAssert(float64(1)) // i is float64 -> 1
}
