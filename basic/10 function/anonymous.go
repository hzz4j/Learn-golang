package main

import "fmt"

func main() {
	test2()
}

func test1() {
	// 1. 定义
	add := func(num1, num2 int) int {
		return num1 + num2
	}
	fmt.Println(add(1, 2))
}

func test2() {
	// 2. 立即执行
	func(name, msg string) {
		fmt.Println(name + " say: " + msg)
	}("静默", "Hello World")
}
