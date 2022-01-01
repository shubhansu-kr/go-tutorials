package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch case in go lang ")

	// rand/math - random number
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1 // 6 is exclusive
	fmt.Println("Value of dice number is : ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Dice value is 1, You can open")
	case 2:
		fmt.Println("You can move 2 blocks")
	case 3:
		fmt.Println("You can move 3 blocks")
	case 4:
		fmt.Println("You can move 4 blocks")
	case 5:
		fmt.Println("You can move 5 blocks")
	case 6:
		fmt.Println("You can move 6 block and roll the dice again")
		fallthrough
	default:
		fmt.Println("What was that !!")
	}

	// fallthrough is used to jump to next case too 
}
