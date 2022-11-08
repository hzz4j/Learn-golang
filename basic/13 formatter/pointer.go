package main

import "fmt"

func main() {
	name := new(string)
	fmt.Printf("%p", name) // 0xc000054250
}
