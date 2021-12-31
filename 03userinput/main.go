package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println ("Enter the rating for go :")

	// commma ok || comma error syntax 
	input, _ := reader.ReadString('\n') 
	// stores value in input if executed properly
	// else stores value in underscore 
	
	fmt.Println("Thanks for rating - ", input)

	fmt.Printf ("Type of input is %T", input) // string type 
}