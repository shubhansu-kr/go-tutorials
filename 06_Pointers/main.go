package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers ")

	// pointer in go is similar to c++

	var ptr *int
	var mynumber int = 32
	ptr = &mynumber

	fmt.Println("The address of my number is ", &mynumber)
	fmt.Println("The address of my number is ", ptr)
	fmt.Println("Value of my number is ", mynumber)

	*ptr = *ptr + 32
	fmt.Println("Value of my number is ", mynumber)
}
