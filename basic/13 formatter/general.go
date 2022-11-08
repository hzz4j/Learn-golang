package main

import "fmt"

type Book struct {
	name  string
	price float64
}

func main() {
	book := Book{"Learning Go programming", 100}
	fmt.Printf("%v\n", book)
	fmt.Printf("%+v\n", book) // 推荐
	fmt.Printf("%#v\n", book)
	fmt.Printf("%T\n", book)
	fmt.Printf("%%\n")
}

/**
{Learning Go programming 100}
{name:Learning Go programming price:100}
main.Book{name:"Learning Go programming", price:100}
main.Book
%
*/
