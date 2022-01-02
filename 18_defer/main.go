package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer ")

	// defer -> marks the line for execution to last
	// last in first out
	// in case of multile defer statements
	// the line which is at last will be ecexuted first
	// that is in reverse order

	defer fmt.Println("First ")
	defer fmt.Println("Second ")
	defer fmt.Println("Third ")

	mydefer()
	fmt.Printf ("\n")
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
