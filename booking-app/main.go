package main

import (
	"fmt"
	"strings"
)

func main() {

	var ConferenceName = "Go Conference" // conferenceName := "go conference"
	dec := "- - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string //var bookings := []string{}

	// fmt.Printf("conferenceTickets is %T, remaining is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, ConferenceName)
	fmt.Println(dec)
	fmt.Printf("Welcome to %v booking application!\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
	fmt.Println(dec)

	for { 											//remainingTickets > 0 && len(bookings) < 50
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

		if userTickets < int(remainingTickets) {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, userFirstName+" "+userLastName)

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("the first value: %v\n", bookings[0])
			// fmt.Printf("slice type: %T\n", bookings)
			// fmt.Printf("slice length: %v\n", len(bookings))

			fmt.Println(dec)
			fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive your confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, ConferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("These first names of bookings are: %v\n", firstNames)
			fmt.Println(dec)
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our Conference is booked out. Come back next year.")
				break
			}
		} else if userTickets == int(remainingTickets) {
			// fmt.Printf()
		} else {
			fmt.Printf("we only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
		
	}

}
