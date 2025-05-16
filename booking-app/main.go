package main

import (
	"booking-app/helper"
	"fmt"
	
)

const conferenceTickets int = 50

var ConferenceName = "Go Conference" // conferenceName := "go conference"
var dec = "- - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
var remainingTickets uint = 50
var bookings []string //var bookings := []string{}
func main() {

	helper.GreetUsers(dec, ConferenceName, conferenceTickets, remainingTickets)

	for {

		userFirstName, userLastName, userEmail, userTickets := helper.GetUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userFirstName, userLastName, userEmail, userTickets, remainingTickets)

		if isValidTicketNumber && isValidName && isValidEmail {

			helper.BookTicket(userTickets, userFirstName, userLastName, dec, userEmail, remainingTickets, bookings, ConferenceName)

			firstNames := helper.GetFirstNames(bookings)

			fmt.Printf("The first names of bookins are: %v\n", firstNames)
			fmt.Println(dec)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our Conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("oh! Invalid First or Last name, too short.\n")
			}
			if !isValidEmail {
				fmt.Printf("oh! Invalid Email entered. Address you entered doesn't contain @. \n")
			}
			if !isValidTicketNumber {
				fmt.Printf("oh! Number of tickets entered is invalid.\n")
			}
		}

	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	// some code for new york tickets
	// case "London", "Hong Kong":
	// 	// some more code
	// default:
	//print statement
	// }

}








