package main

import "fmt"

func main() {
	//创建一个长度为3，容量为4，int类型的切片
	s := make([]int, 3, 4)
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 4

	// 这种方法和创建数组类似，只是不需要指定[]运算符里的值。初始的长度和容量会基于初始化时提供的元素的个数确定
	// slice := []int{1,2,3}

	// 和数组一样也可以通过指定索引初始化, 比如index 4 值为100
	slice := []int{3: 100}
	fmt.Println(slice) // [0 0 0 100]
}
