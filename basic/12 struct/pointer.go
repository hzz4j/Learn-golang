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
	fmt.Printf("%p\n", car) //0xc0000d6000
}

func main() {
	car := &Car{
		Color: "red",
		Brand: "BMW",
		Model: "5",
		price: 100,
	}
	fmt.Printf("%p\n", car) //0xc0000d6000
	changeToBlue(car)
	fmt.Printf("%+v\n", *car) // {Color:blue Brand:BMW Model:5 price:100}

	fmt.Println(car.Color)    // blue
	fmt.Println((*car).Color) // blue

}
