package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Monday", "Tuesday", "Thursday", "Friday", "Sunday"}
	fmt.Println(days)

	// loop syntax

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// range is used to set the conditions for loop
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v, Day is %v \n", index, day)
	// }

	// while loop
	roguevalue := 2
	for roguevalue < 10 {

		if roguevalue == 5 {
			roguevalue++
			continue
		}
		if roguevalue == 7 {
			goto lco
		}

		fmt.Println(roguevalue)
		roguevalue++
	}
	// go lang dont have the concept of preincrement operator -> only post

	// goto in golang is similar to c/c++
	// declare label and jump between them

lco:
	fmt.Println("used goto")
}
