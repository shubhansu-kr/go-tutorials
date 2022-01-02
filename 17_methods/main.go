package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Name     string
	Email    string
	Nickname string
}

func (u User) setuser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name ")
	Name, _ := reader.ReadString('\n')
	u.Name = Name

	fmt.Println("Enter Email ")
	Email, _ := reader.ReadString('\n')
	u.Email = Email

	fmt.Println("Enter nickname ")
	NiName, _ := reader.ReadString('\n')
	u.Nickname = NiName

	u.display()
}

func (u User) display() {
	fmt.Println("Name - ", u.Name)
	fmt.Println("Email - ", u.Email)
	fmt.Println("Nickname - ", u.Nickname)
}

func main() {
	fmt.Println("Welcome to methods ")

	// Shubh := User {"Shubhansu", "shubh@go.dev", "Shubh"}
	// Shahi := User {"Raushan", "shahi@gmail.com", "sobhu"}
	// Vivek := User {"Vivek", "Vivek@gmail.com", "thorki"}

	Shubh := User{}
	Shubh.Name = "Shubhansu"
	Shubh.Email = "shubh@go.dev"
	Shubh.Nickname = "Shubh"

	Vivek := User{}
	// Shahi := User{}

	Shubh.display()
	Vivek.setuser()
	// Shahi.setuser()
	// Vivek.display() -> empty

}
