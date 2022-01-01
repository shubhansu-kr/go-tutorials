package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")

	language := make(map[string]string)

	language["PY"] = "Python"
	language["JS"] = "Javascript"
	language["Rb"] = "Ruby"

	fmt.Println("List of all languages : ", language)

	// delete type is used to delete elements
	delete(language, "Rb")
	fmt.Println("List of all languages : ", language)

	// single element can be accessed using the index specified
	fmt.Println("JS stands for : ", language["JS"])

	// loops in go lang

	for key, value := range language {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}

	// _ (underscore is used to ignore in go lang )
}
