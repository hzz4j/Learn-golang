package main

import "fmt"

func main() {
	// pointerVsVariable()
	// declarePointer()
	multiPointer()
}

func pointerVsVariable() {
	a := "Hello, world!"
	fmt.Println(&a) // 0xc000054250
}

// 声明指针
func declarePointer() {
	// panic: runtime error: invalid memory address or nil pointer dereference
	// var a *int
	// *a = 10
	// fmt.Println(a)

	var a *int = new(int)
	*a = 10
	fmt.Println(a, *a, &a) // 0xc0000180a8 10 0xc000006028
}

// 多维指针
func multiPointer() {
	var a ***int
	v := 3
	p1 := &v  // *int
	p2 := &p1 // **int
	p3 := &p2 // ***int
	a = p3
	fmt.Println(***a)
}
