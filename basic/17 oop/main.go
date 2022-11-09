package main

import "fmt"

// go run main.go pop.go
func main() {
	WashProcedure(10, 5)
	fmt.Println()
	washer := &Washer{State: false, Powder: 10}
	// washer.Powder = 0
	closes := []*Closes{{Clean: false}, {Clean: true}}
	err := washer.wash(closes)
	washer.check(err, closes)
}
