package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	// passing the address of struct variable
	// emp8 is a pointer to the Employee struct
	emp8 := &Employee{
		"PhucDn", "Do Ngoc", 23, 5000,
	}
	// (*emp8).firstName is the syntax to access
	// the firstName field of the emp8 struct
	// fmt.Println("First Name: ", (*emp8).firstName)
	// fmt.Println("Age: ", (*emp8).age)

	// emp8.firstName is used to access
	// the field firstName
	fmt.Println("First Name: ", emp8.firstName)
	fmt.Println("Age: ", emp8.age)
}
