package main

import (
	"fmt"
	"strings"
) 

var conferenceName = "Go Conference"
const conferenceTickets uint= 50
var remainingTickets uint= 50
var bookings = []string{}


func main() {


	greetUsers()
	
	for {

		firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
	
		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTickets(userTickets,firstName, lastName, email)

			//call printFirstNames
			firstNames := getFirstNames();

			fmt.Printf("The first name of bookings are %v\n", firstNames);

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break;
			}	
		} else {
			if !isValidName {
				fmt.Println("First name or Last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email does not contain @. Invalid email")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}		
	}
	

}

func greetUsers() {
	fmt.Printf("Welcome to %v.\n",conferenceName)
	fmt.Printf("We have %v tickets and %v tickets are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Book your tickets to attend.\n")
}

func getFirstNames() []string {
	firstNames := []string{} 
	for _,booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames;
}


func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets required: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)	
	fmt.Printf("User %v %v has booked %v tickets. Email confirmation will be sent at %v \n",firstName, lastName, userTickets, email) 
	fmt.Printf("%v remaining tickets\n", remainingTickets);
}