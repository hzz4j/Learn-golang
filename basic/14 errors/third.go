package main

import "fmt"

func soo() {
	fmt.Println("enter soo function")
	defer func() {
		if err := recover(); err != nil { // 拦截处理异常
			fmt.Printf("soo function error-> %s", err)
		}
	}()

	fmt.Println("注册defer结束")

	defer fmt.Println("AAA")

	defer func() {
		n := 0
		_ = 3 / n
		fmt.Println("BBB")
		defer fmt.Println("CCC")
	}()
}

func main() {
	soo()
}

/**
enter soo function
注册defer结束
AAA
soo function error-> runtime error: integer divide by zer
*/
