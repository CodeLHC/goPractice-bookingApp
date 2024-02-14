package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("Only %v tickets left out of %v spots\n", remainingTickets, conferenceTickets )
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter amount of tickets: ")
		fmt.Scan(&userTickets)

		isValidName:= len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber{
			remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " "+ lastName)

		


		fmt.Printf("Thank you %v %v, for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for the %v", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings{
			 var names = strings.Fields(booking)
			 firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)


		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}
		} else {
			if !isValidName{
				fmt.Print("First name or last name you entered is too short\n")
			}
			if !isValidEmail{
				fmt.Print("Email address you entered does not contain an @ sign\n")
			}
			if !isValidTicketNumber{
				fmt.Print("The number of tickets you've entered is invalid\n")
			}
		
			
		}
		
		
	}
}