package main

import (
	"fmt"
)

func multiVar() {
	var (
		hello string = "Hello "
		world string = "World"
	)
	fmt.Println(hello)
	fmt.Println(world)
}

func main() {
	// 自动类型推断
	var name = "静默"
	fmt.Println(name)

	// 声明多个变量并指定类型
	var a, b int = 3, 2
	fmt.Println(a, b)

	// var 简化写法
	f := true
	fmt.Println(f)

	// var tmp1  会报错：没有指定类型
	var tmp2 int
	fmt.Println(tmp2)

	multiVar()
}

/**output
静默
3 2
true
0
*/
