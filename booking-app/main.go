package main

import (
	"fmt"
)

func main() {
	var ConferenceName = "Go Conference"
	// conferenceName := "go conference"

	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remaining is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, ConferenceName)


	fmt.Printf("Welcome to %v booking application!\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")




	var userName string
	var userTickets int
	//ask user for name

	userName = "bharath"
	userTickets = 2
	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)


}


