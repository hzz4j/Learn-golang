package main

import "fmt"

// 指针参数
func argf(a, b *int) {
	*a = *a + *b
	*b = 666
}

func main() {
	var (
		a = 3
		b = 2
	)

	argf(&a, &b)
	fmt.Println(a, b)
}
