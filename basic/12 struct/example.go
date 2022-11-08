package main

import "fmt"

type Car struct {
	Color string
	Brand string
	Model string
	price int
}

func main() {
	var car Car
	fmt.Printf("%+v\n", car) // {Color: Brand: Model: price:0}
	car = Car{
		Color: "red",
		Brand: "BMW",
		Model: "5",
		price: 100,
	}
	fmt.Printf("%+v\n", car) // {Color:red Brand:BMW Model:5 price:100}
}
