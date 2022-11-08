package main

import "fmt"

type Car struct {
	Color string
	Brand string
	Model string
	price int
}

func changeToBlue(car *Car) {
	car.Color = "blue"
}

func main() {
	car := &Car{
		Color: "red",
		Brand: "BMW",
		Model: "5",
		price: 100,
	}
	changeToBlue(car)
	fmt.Printf("%+v\n", *car) // {Color:red Brand:BMW Model:5 price:100}
}
