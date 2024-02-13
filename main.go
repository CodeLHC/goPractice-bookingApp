package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50 

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("Only %v tickets left out of %v spots\n", remainingTickets, conferenceTickets )
	fmt.Println("Get your tickets here to attend")


	var userName string
	var userTickets int

	userName = "Laura"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
	
}