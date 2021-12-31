package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices ")

	// slice
	var fruitlist = []string{"Mango", "Coconut", "Mongodb"}
	fmt.Printf("fruit list is of data type : %T", fruitlist)

	fmt.Println("Fruit list is : ", fruitlist)

	// to add elements to the end
	fruitlist = append(fruitlist, "Coco", "Banana")
	fmt.Println("Fruit list is : ", fruitlist)

	// to trim the slice
	fruitlist = append(fruitlist[:3], "hello") // 3 is not included
	fmt.Println("Fruit list is : ", fruitlist)

	// make and new
	highscores := make([]int, 4)  // make (data type , no.of vars)

	highscores[0] = 43
	highscores[1] = 84
	highscores[2] = 65
	highscores[3] = 72

	fmt.Println(highscores)
	highscores = append(highscores, 434, 453, 564)
	fmt.Println(highscores)
	
	// append actually reallocates the memory so we can add 
	// element to slice which only had memory for 4 elements 
	
	// fmt.Print(highscores[5]) 
	
	// sort is a package used to sort dlices 
	fmt.Println("SLice is sorted : ", sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println("SLice is sorted : ", sort.IntsAreSorted(highscores))
	
}
