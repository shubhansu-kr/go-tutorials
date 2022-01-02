package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go lang")

	content := "Follow me on linked in "

	// os package is used to create a txt file in the pwd
	file, err := os.Create("./Follow.txt")

	if err != nil {
		panic(err)
	}

	// io is used for input putput
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("The length of file is ", length)

	defer file.Close()

	readfile("./Follow.txt")
}

func readfile(file string) {
	databyte, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	// text file is read in data bytes
	fmt.Println("Text inside the file is : \n", string(databyte))

}
