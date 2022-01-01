package main

import "fmt"

func main() {
	fmt.Println("Welcome to Control structure ")

	// If else
	var result string
	age := 21

	if age > 18 {
		result = "Welcome to the party"
	} else {
		result = "You are not allowed in here"
	}

	fmt.Println(result)

	if num := 5; num%2 == 0 {
		fmt.Println("Number is even ")
	} else {
		fmt.Println("Number is odd ")
	}

}
