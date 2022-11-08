package main

import "fmt"

func main() {
	str := "Hello world!"
	fmt.Printf("%s\n", str) // Hello world!
	fmt.Printf("%q\n", str) // "Hello world!"
	fmt.Printf("%x\n", str) // 48656c6c6f20776f726c6421
	fmt.Printf("%X\n", str) // 48656C6C6F20776F726C6421
}
