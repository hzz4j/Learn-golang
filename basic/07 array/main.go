package main

import (
	"fmt"
	"reflect"
)

func main() {
	// declareArray1()
	// declareArray2()
	// iterateArray()
	// arrayPointer()

	a := [...]int{1, 2, 3}
	transferArray(a)
	fmt.Println(&a, a)
}

func transferArray(a [3]int) {
	a[0] = 3
	fmt.Println(&a[0])
}

func declareArray1() {
	var a1 [5]int
	var a2 [3]string
	fmt.Println(a1) // [0 0 0 0 0]
	fmt.Println(a2) // [  ]
}

func declareArray2() {
	a1 := [3]int{1, 2, 3}
	// 如果将元素个数指定为特殊符号...，则表示通过初始化时的给定的值个数来推断数组长度
	a2 := [...]int{1, 2, 3, 4, 5}
	// 如果声明数组时，只想给其中某几个元素初始化赋值，则使用索引号
	a3 := [...]int{0: 1, 7: 7}
	fmt.Println(a1) // [1 2 3]
	fmt.Println(a2) // [1 2 3 4 5]
	fmt.Println(a3) // [1 0 0 0 0 0 0 7]
}

// 数组类型
func arrayType() {
	var (
		a [3]int
		b [5]int
	)

	fmt.Println(reflect.TypeOf(a)) // [3]int
	fmt.Println(reflect.TypeOf(b)) // [5]int
}

// 遍历数组
func iterateArray() {
	a := [...]int{3, 2, 5, 6}
	for i, v := range a {
		fmt.Println(i, v)
	}
}

// 指针数组
func arrayPointer() {
	a := [3]*int{0: new(int), 2: new(int)}
	fmt.Println(a) // [0xc0000180a8 <nil> 0xc0000180c0]
}
