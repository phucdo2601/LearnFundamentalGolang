package main

import "fmt"

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio, ", firstName)
}

func getNames() (firstName string, lastName string) {
	return "John", "Doe"
}
