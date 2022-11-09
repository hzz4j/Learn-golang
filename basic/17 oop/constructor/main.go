package main

import "fmt"

func main() {
	u1 := GetUserInstance()
	u2 := GetUserInstance()
	fmt.Printf("u1 = %p, u2 = %p\n", u1, u2)
}
