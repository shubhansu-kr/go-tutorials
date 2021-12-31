package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to golang ")

	// time is the package

	presenttime := time.Now()
	fmt.Println(presenttime) // prints the current time in default format 
	// to format the output format package is used 
	fmt.Println(presenttime.Format("01-02-2006 Monday"))

	// createdate 
	birthdate := time.Date(2003, time.July, 21 , 11, 35, 44, 32, time.Local)
	fmt.Println (birthdate)

}
