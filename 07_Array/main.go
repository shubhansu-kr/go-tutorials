package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array ")

	// declaring array
	var fruits [4]string

	// inititalising
	fruits[0] = "Apple"
	fruits[1] = "Pineapple"
	fruits[3] = "Guava"

	// we can print the whole array with the array name
	fmt.Println("Fruit list is : ", fruits)
	// a blank space is the indication of empty elment

	// len() is used to print the length of array
	fmt.Println("The length of array is : ", len(fruits))

	// declare and initialise
	var veglist = [3]string{"mushroom", "Spinach", "cucumber"}
	fmt.Println("Veg list is ", veglist)

}
