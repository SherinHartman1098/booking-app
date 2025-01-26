package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

// package level variables
var conferenceName = "Go Conference" // alternative syntax=> conferenceName :="GoConference"
const conferenceTickets = 50

var remainingTickets = 50
var bookings = make([]map[string]string, 0) //creating empty list of maps with initial size as 0

func main() { //entry point

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isvalidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isvalidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names of booking are: %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked ouot. Please try next year.")
				break
			}
		} else {
			if !isvalidName {
				fmt.Printf("First name or last name entered is too short.\n")
			}
			if !isValidEmail {
				fmt.Printf("Email address you entered doesn't contain@sign.\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of ticket you entered is invalid.\n")
			}

		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string { //in this line, the inside the bracket []string are input parameters and outside are output parameters
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames

}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter your number of tickets needed: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = make(map[string]string) // we cannot have mixed data types
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // converting integer to string

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are reminaing for %v\n", remainingTickets, conferenceName)

}
