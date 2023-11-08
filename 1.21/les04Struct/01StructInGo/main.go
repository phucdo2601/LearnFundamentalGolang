package main

import "fmt"

type Student struct {
	name        string
	age         int
	address     string
	phoneNumber int
	email       string
	gender      bool
	nickname    string
}

func main() {
	/**
	Declaring a variable of a `struct` type
	All the struct fields are initialized
	with their zero value
	*/

	var newStuA Student
	fmt.Println(newStuA)

	// Naming fields while
	// initializing a struct
	newStuB := Student{
		name:        "Phuc Do Ngoc",
		age:         23,
		address:     "Ho Chi Minh City, Viet Nam",
		phoneNumber: 42434323,
		email:       "testEmail01@gmail.com",
		gender:      true,
		nickname:    "phucdn",
	}

	fmt.Println("Student B: ", newStuB)

	/**
	Declaring and initializing a
	struct using a struct literal
	*/
	newStuC := Student{
		"testName02",
		22,
		"Ha Noi, Viet Nam",
		242343,
		"testEmail02@gmail.com",
		false,
		"testNick02",
	}

	fmt.Println("Student C: ", newStuC)

	// Uninitialized fields are set to
	// their corresponding zero-value
	newStudD := Student{name: "Delhi"}
	fmt.Println("Address3: ", newStudD)

	// Accessing struct fields
	// using the dot operator
	fmt.Println("Name of Student B: ", newStuB.name)

	// Assigning a new value
	// to a struct field
	newStuB.name = "Billy Stone"
	// Displaying the result
	fmt.Println("Update name of student B: ", newStuB)
}
