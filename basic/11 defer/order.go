package main

import "fmt"

func main() {
	fmt.Println("Starting")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	fmt.Println("Ended")
}

/**
Starting
Ended
4
3
2
1
*/
