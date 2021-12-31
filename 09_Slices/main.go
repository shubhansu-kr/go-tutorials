package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices part 2 ")

	// create a slice
	var marks = []int{78, 87, 67, 74, 55}
	fmt.Println(marks)

	// removing an element from slice
	// we gonna use append to remove elements
	var index int = 2
	marks = append(marks[:index], marks[index+1:]...)
	fmt.Println(marks)

}
