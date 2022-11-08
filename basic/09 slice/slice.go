package main

import "fmt"

func main() {
	// expanSlice()
	copySlice()
}

func create() {
	//创建一个长度为3，容量为4，int类型的切片
	s := make([]int, 3, 4)
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 4

	// 这种方法和创建数组类似，只是不需要指定[]运算符里的值。初始的长度和容量会基于初始化时提供的元素的个数确定
	// slice := []int{1,2,3}

	// 和数组一样也可以通过指定索引初始化, 比如index 4 值为100
	slice := []int{3: 100}
	fmt.Println(slice) // [0 0 0 100]
}

func sliceTest() {
	s1 := []int{10, 20, 30, 40}
	s2 := s1[1:3]
	fmt.Println(s1, s2) // [10 20 30 40] [20 30]
	s2[0] = 100
	fmt.Println(s1, s2) // [10 100 30 40] [100 30]
}

func expanSlice() {
	s1 := []int{10, 20, 30, 40}
	s2 := s1[1:3:3]
	fmt.Println(len(s2), cap(s2)) // 2 2
	fmt.Println(s1, s2)           // [10 20 30 40] [20 30]

	s2 = append(s2, 300) // s2 扩容
	s1[1] = 200          // 修改s1
	fmt.Println(s1, s2)  // s1修改并不会影响s2 [10 200 30 40] [20 30 300]
}

func copySlice() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, len(s1))
	copy(s2, s1)
	s2[0] = 100
	fmt.Println(s1, s2) // [1 2 3] [100 2 3]
}
