/*
-khi muon chay mot du an go:
+ Phai create a new moudle bang lenh go mod init ten project
+ Cau truc cua mot file .go la phai nam trong mot package, cau dau tien trong mot file .go phai la package
+ cac cau lenh xu ly cua trong go phai dc nhet trong mot func: func funcName (){}
+ Phai import cac thu vien ma go ho tro de thuc hien cac thao tac lenh
+ Phai import thu vien fmt de dung dc lenh print: st1: import "fmt"; st2: fmt.Print("")
+ De chay mot file go: go run file name, ex: go run main.go
*/
package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

/**
* Const in go: khi dat bien const co gia tri xac di thi khong thay doi dc gia tri do
 */
const conferenceTickets int = 50

// var conferenceName = "Go Conference"
var conferenceName = "Go Conference"

var remainingTickets uint = 50

// slice demo
// var bookings []string
// map demo
// var bookings = []string{}
// struct demo
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	/**
	* Variable in go: la mot bien trong go, co the thay doi dc gia tri
		bien khai bao trong go thi phai su dung, khong su dung thi se bao loi
		 + cach khai bao bien trong go: var varName = "gia tri", ex: var conferenceName = "Go Conference"
		 + cach khai bao khac:
			"Syntactic Sugar": varName := "gia tri", ex: conferenceName := "Go Conference"
	*/

	/**
	* print in go lang using library fmt and fmt.Print
		su dung fmt.Printf de nhet bien khi in trong go, cau lenh nay se goi tuan tu cac bien ma minh khai bao sau dau phay
		 + %v: la in ra gia tri cua bien
		 + %T: la in ra kieu du lieu cua bien
		\n: la de xuong dong
		- Khi in ra gia tri, them dau & truoc ten bien dc in, thi go se cho ta ket qua vi tri cua bien do trong memory
		- "string.Fields()": (phai import lib strings truoc khi dung)
			+ Spits the string with white space as separator
			+ And returns a slice with the split elements

		- len():
			+ Build in function returning length of variable, according to its type
			+ Arrays & Slices: Size of the list (number of elements)
	*/
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceNames is %T\n", conferenceTickets, remainingTickets, conferenceName)

	greetUsers()

	/**
	Logical Operators:
	 - &&: logical AND operator, both conditions must be true, for the whole expression to be true
	 - ||: logical OR operator, any of conditions must be true, for the whole expression to be true
	*/

	/**
	Array in Go:
	Fixed Size -- Phai fix cung mot size khi khai bao mang
	Array in Go is Static
	Same data type
	- Cach khai bao mang trong go: var varName [size]variable_type
	- Cach khai bao mang co gia tri khoi tao trong go: var bookings = [50]string{"Nana", "Billy", "Peter"}
	*/

	/**
	Slices in Go:
	 - Slice is abstraction of an Array (la mot thanh phan con trong mang)
	 - Slices are more flexible and powerful: variable-length or get an sub-array of its own
	 - Slices are also indexed-based and have a size, but is resized when needed
	 - Khai bao slice:
	 	+ C1: var varName []variable_type
		+ C2: var varName = []variable_type{} // khai bao mot slice rong
		+ C3: varName := []string{} // tuong tu cach khai bao "Syntactic Sugar", phai khai bao mot mang rong
	*/

	/**
	Loop in Go: (Vong lap) -- for, for-each
	Cu phap vong for each: for index, objName := range varName {}
	- Blank Identifier: (_)
		+ To ignore a variable you dont want to use
		+ So with Go, you need to make unused variables explicit
	- Infinite Loop (never ends):
		+ A loop that repeats endlessly, as the condition is always true
	- Loop with conditionals: ending when excution conditions is done

	- "break" statement:
		+ Terminate the loop for
		+ And continues with  the code right after the for loop (in our case we dont have any code)

	- "continue" statement:
		+ Cause loop to skip the remainer of its body
		+ Add immediately resetting its conditions (in our infinite loop our condition is always true)

	- conditionals in Loop:
		+ End of a loop:
			A Loop continues as long as a condition is true
	*/

	/**
	Rannge:
		- Range iterates over elements for different data structures (so not only arrays and slices)
		- For arrays and slices, range provides the index and value for each element
	*/

	/**
	if ... else:
	- Cu phap if condition:
		if condition {
			// code to be executed if condition is true
		}

	- Cu phap if else condition:
		if condition {
			// code to be executed if condition is true
		} else {
			// code to be executed if condition is false
		}

	- Cu phap else if condition:
		if condition {
			// code to be executed if condition is true
		} else if condition {
			// code to be executed if condition is true
		} else {
			// code to be executed if both conditions are false
		}
	*/

	/**
	User input validation:
	*/

	/**
	Switch Statement:
		- Allows a variable to be tested for equality against a list of values
		- Default handles the case, if no match is found
	*/

	/**
	Functions
		- Function is only executed when called
		- You can call a function as many times you want
		- So a function is also used to reduced a code duplication
		- Returning multiple values: (Function trong go co the tra ve nhieu gia tri theo kieu va thu tu ma minh da quy uoc san)
			+ A Go function can return multiple values
	*/

	/**
	Package  Level Variables
		- Defined at the outside all functions
		- They can be accessed inside any of the functions
		- And in all files, which are in the same package

	Local variables
		- Defined at the inside a functions or a block
		- They can be accessed only inside that function or block of code
	*/

	/**
	Date types in go: tham khao https://www.geeksforgeeks.org/data-types-in-go/
	khi muon khai bao bien trong go ma khong can khai bao gia tri khoi tao ban dau, thi chi can the dataType dang sau ten bien
	Cau truc var varName dataType; ex: var address string
	+ Cac kieu du lieu so trong go: int, uint, byte, rune, float64, float32
		int la so binh thuong,
		float64, float32: la so co phan thap phan
	+ Phai cung kieu du lieu thi moi thuc hien dc cac phep toan logic
	*/

	// for remainingTickets > 0 && len(bookings) < 50 {

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := printFirstNamesFunc()
		fmt.Printf("There first name of bookings are: %v\n", firstNames)

		// var noTicketsRemaining bool = remainingTickets == 0
		// noTicketsRemaining := remainingTickets == 0

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or Last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign.")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid.")
		}
	}
	wg.Wait()
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking applications\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

}

func printFirstNamesFunc() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name

	/**
	User input: fmt.Scan() -- (Input da tu ban phim)
	Den nhan mot input tu ban phim trong go thi chung ta su dung fmt.Scan(&varName):
		Trong do varName la bien chung ta can lay gia tri input dua vao
	*/
	// fmt.Scan(userName)
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter numbers of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	// var userData = make(map[string]string)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookings = append(bookings, firstName+" "+lastName)
	bookings = append(bookings, userData)

	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Array Type: %T\n", bookings)
	// fmt.Printf("Array Length: %v\n", len(bookings))

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice Type: %T\n", bookings)
	// fmt.Printf("Slice Length: %v\n", len(bookings))

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")
	wg.Done()
}
