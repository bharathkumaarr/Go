package main

import (
	validate "booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var ConferenceName = "Go Conference" // conferenceName := "go conference"
var dec = "- - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
var remainingTickets uint = 50
var bookings []string //var bookings := []string{}
func main() {

	greetUsers()

	for {

		userFirstName, userLastName, userEmail, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validate.ValidateUserInput(userFirstName, userLastName, userEmail, userTickets, remainingTickets)

		if isValidTicketNumber && isValidName && isValidEmail {

			bookTicket(userTickets, userFirstName, userLastName, dec, userEmail)

			firstNames := getFirstNames()

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

func greetUsers() {
	// fmt.Printf("conferenceTickets is %T, remaining is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, ConferenceName)
	fmt.Println(dec)
	fmt.Printf("Welcome to %v booking application!\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
	fmt.Println(dec)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames

}

func getUserInput() (string, string, string, int) {
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

	return userFirstName, userLastName, userEmail, userTickets
}

func bookTicket(userTickets int, userFirstName string, userLastName string, dec string, userEmail string) []string {
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, userFirstName+" "+userLastName)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("the first value: %v\n", bookings[0])
	// fmt.Printf("slice type: %T\n", bookings)
	// fmt.Printf("slice length: %v\n", len(bookings))

	fmt.Println(dec)
	fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive your confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, ConferenceName)
	return bookings
}
