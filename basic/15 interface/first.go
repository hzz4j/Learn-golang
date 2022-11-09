package main

import "fmt"

type Transporter interface { // 定义接口，通常以er结尾
	move(src string, dst string) (int, error)
	// 变量名也可以省略
	whistle(int) int
}

type Car struct { // 定义结构体时无需指定要实现什么类型的接口
	price int
}

// 只要结构体实现了接口声明里的方法，就称为该结构体实现了该接口

func (c Car) move(src string, dst string) (int, error) {
	return c.price, nil
}

func (c Car) whistle(price int) int {
	return c.price
}

func main() {
	var t Transporter
	t = Car{
		price: 100,
	}

	t.whistle(3)
	fmt.Printf("%+v\n", t)
}
