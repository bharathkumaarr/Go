package helper

import (
	"fmt"
	"strings"
)

func ValidateUserInput(userFirstName string, userLastName string, userEmail string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}

func GreetUsers(dec string, ConferenceName string, conferenceTickets int, remainingTickets uint) {
	// fmt.Printf("conferenceTickets is %T, remaining is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, ConferenceName)
	fmt.Println(dec)
	fmt.Printf("Welcome to %v booking application!\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
	fmt.Println(dec)
}

func GetFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames

}

func GetUserInput() (string, string, string, int) {
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

func BookTicket(userTickets int, userFirstName string, userLastName string, dec string, userEmail string, remainingTickets uint, bookings []string, ConferenceName string) []string {
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
