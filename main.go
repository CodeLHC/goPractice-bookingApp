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

		if userTickets <= remainingTickets{
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

			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n",remainingTickets, userTickets)
			
		}
		
		
	}
}