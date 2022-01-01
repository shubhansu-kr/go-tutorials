package main

import "fmt"

type User struct {
	// capital u in user defines the function to be public
	Name   string
	Age    int
	Email  string
	Status bool
	// all members are public since start with capital letter
}

func main() {
	fmt.Println("welcome to structs ")
	// NO inheritance in go lang, no super , no child

	// declare struct obj
	Shubh := User{"Shubhansu", 21, "shubhansu2021@gmail.com", true}
	fmt.Println(Shubh)
	fmt.Printf("User details : %+v\n", Shubh)
	fmt.Printf("User Name is : %+v\n", Shubh.Name)
	fmt.Printf("User Name is : %+v\n", Shubh.Email)
	fmt.Printf("User Name is : %+v\n", Shubh.Age)

}
