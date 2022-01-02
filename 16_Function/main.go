package main

import "fmt"

func Greeter(name string) {
	fmt.Printf("Welcome %v ", name)

}

func sum(a int, b int) int {
	return a + b
}

func prosum(values ...int) (int, int) {
	total := 0
	count := 0
	for _, val := range values {
		total += val
		count++
	}
	return total, count
}

func main() {
	fmt.Println("Welcome to functions in go lang ")
	// Main is the entry point for program
	// We dont need to call main unlike any other func

	// call greeter
	Greeter("Shubh")
	// we cannot declare a function inside a function

	sum := sum(4, 34)
	fmt.Println("The sum is ", sum)

	// can return more than one value in go lang using a function 
	prosum, count := prosum(3, 52, 53, 23, 564) 
	fmt.Println("The sum is ", prosum)
	fmt.Println("NO. of parameters is ", count)

}
