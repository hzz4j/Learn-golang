package main

import "fmt"

const name = "静默"

var greeting = "Hello"

func main() {
	fmt.Println(greeting, name)

	//name = "Q10Viking" 常量不能改变
	greeting = "Hello again"
	fmt.Println(greeting, name)
}

/**
Hello 静默
Hello again 静默
*/
