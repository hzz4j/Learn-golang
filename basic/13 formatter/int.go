package main

import "fmt"

func main() {
	a := 255
	fmt.Printf("二进制: %b\n", a)    // 11111111
	fmt.Printf("八进制: %o\n", a)    // 377
	fmt.Printf("十进制: %d\n", a)    // 255
	fmt.Printf("十六进制: %x\n", a)   //ff
	fmt.Printf("大写十六进制: %X\n", a) //FF
}
