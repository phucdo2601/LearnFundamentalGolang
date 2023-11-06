package main

import "fmt"

func main() {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	password := "vbcvbvcb"
	username = "Phucdn"

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username, password)

	if lengthPass := len(password); lengthPass > 4 {
		fmt.Printf("the length of pasword string is greater than 4")
	} else {
		fmt.Printf("the length of pasword string is so small")
	}
}
