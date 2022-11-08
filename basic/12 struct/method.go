package main

import "fmt"

type Person struct {
	name, msg string
}

func (p Person) sayHello() {
	fmt.Printf("Hello, %s! %s\n", p.name, p.msg)
}

func main() {
	p := Person{
		name: "静默",
		msg:  "Learning Go programming",
	}
	p.sayHello() // Hello, 静默! Learning Go programming
}
