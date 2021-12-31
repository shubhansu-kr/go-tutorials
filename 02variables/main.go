package main

import "fmt"


// constant declaration
// const logintoken string = "sahuera kda" // private 
const Logintoken string = "sahuera kda" // public 

func main() {
	// syntax to declare new variable
	// var <variable name> <vartiable type> = value ;

	var user string = "Shubhansu"
	
	// fp is the shortcut for printline
	fmt.Println(user)
	
	// capital is the format specifier for variable type 
	fmt.Printf ("The type of variable is : %T \n", user)
	
	// bool data type - true or false 
	var isloggedin bool = false 
	fmt.Println(isloggedin)
	fmt.Printf ("The type of variable is : %T \n", isloggedin)
	
	var score int = 54 
	fmt.Println(score)
	fmt.Printf ("The type of variable is : %T \n", score)
	
	var smallfloat float32 = 432.76534563 
	fmt.Println(smallfloat)
	fmt.Printf ("The type of variable is : %T \n", smallfloat)

	//default value of variables if not initialised 
	var roll int 
	fmt.Println(roll) // 0 is the default value 
	fmt.Printf ("The type of variable is : %T \n", roll) 

	var name string 
	fmt.Println(name) // null is the default value 
	fmt.Printf ("The type of variable is : %T \n", name)

	// implicit way of declaring variables 
	var portfolio = "github/shubhansu-kr"
	fmt.Println(portfolio)
	fmt.Printf ("The type of variable is : %T \n", portfolio)

	// no var style - implicit declaration - use walrus operator
	// this method is only allowed inside a method or function 
	
	nousers := 3200  
	fmt.Println(nousers)
	fmt.Printf ("The type of variable is : %T \n", nousers)

	// using constant 
	fmt.Println(Logintoken)
	fmt.Printf ("The type of variable is : %T \n", Logintoken)
}
