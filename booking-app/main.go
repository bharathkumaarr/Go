package main

import (
	"fmt"
)

func main() {
	var ConferenceName = "Go Conference" // conferenceName := "go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// fmt.Printf("conferenceTickets is %T, remaining is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, ConferenceName)
	fmt.Println("- - - - - - - - - - - - - - - - - - -")
	fmt.Printf("Welcome to %v booking application!\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
	fmt.Println("- - - - - - - - - - - - - - - - - - -")

	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets int

	//ask user for name
	fmt.Print("Enter your First name: ")
	fmt.Scan(&userFirstName)

	fmt.Print("Enter your Last name: ")
	fmt.Scan(&userLastName)

	fmt.Print("Enter your Email: ")
	fmt.Scan(&userEmail)

	fmt.Print("Enter the Tickets you want to book: ")
	fmt.Scan(&userTickets)

	fmt.Println("- - - - - - - - - - - - - - - - - - -")

	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive your confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)
	fmt.Println("")
	fmt.Printf("%v tickets are remaining for %v\n",remainingTickets, ConferenceName)



}
