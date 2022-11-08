package main

import "fmt"

type Author struct {
	Name   string
	Macket string
}

// 组合
type Book struct {
	Title  string
	Author Author
}

func main() {
	book := Book{
		Title:  "Go Programming",
		Author: Author{Name: "静默", Macket: "全站开发工程师"},
	}
	fmt.Printf("%+v\n", book) // {Title:Go Programming Author:{Name:静默 Macket:全站开发工程师}}

}
