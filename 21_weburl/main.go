package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=32asd432ersg"

func main() {
	fmt.Println("Welcome to web url in go lang")
	fmt.Println(myurl)

	// parsing the url - extract some information by resolving into syntatic parts
	// url library
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}
	// fmt.Println(result.Scheme)       ->http
	// fmt.Println(result.Host)         ->lco.dev:3000
	// fmt.Println(result.Path)	        ->/learn
	// fmt.Println(result.Port())       ->3000
	// fmt.Println(result.RawQuery)

	qparam := result.Query()
	fmt.Printf("The type is %T \n", qparam) // url.values -> key pair format

	fmt.Println(qparam["coursename"])
	fmt.Println(qparam["paymentid"])

	for _, val := range qparam {
		fmt.Println("Parameter is : ", val)
	}

	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherurl := partsofurl.String()
	fmt.Println(anotherurl)
}
