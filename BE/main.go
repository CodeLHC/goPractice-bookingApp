package main

import (
	"encoding/json"
	"fmt"
	"gopractice-bookingapp/helper"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

	const conferenceTickets = 50
	var conferenceName = "Go Conference"
	var remainingTickets uint = 50 
	var bookings = make([]UserData, 0)


	type ConferenceInfo struct {
		ConferenceName string
		TotalTickets uint
		TicketsLeft uint
	}

	type UserData struct {
		firstName string
		lastName string
		email string
		numberOfTickets uint
	}

	var wg = sync.WaitGroup{}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/conference", getConferenceDetails).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

	greetUsers()

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber{
			bookings, remainingTickets = bookTickets(remainingTickets, userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

		firstNames:= getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year")
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
		wg.Wait()
	}

func bookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string, email string) ([]UserData, uint){
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}


	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v, for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v", remainingTickets, conferenceName)
return bookings, remainingTickets
}

func getConferenceDetails( w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response:= ConferenceInfo{
		ConferenceName: conferenceName,
		TotalTickets: conferenceTickets,
		TicketsLeft: remainingTickets,
	}
json.NewEncoder(w).Encode(response)
}

func greetUsers() string {
	ticketInfo := fmt.Sprintf("Welcome to the %v booking application\n", conferenceName)
	ticketInfo += fmt.Sprintf("Only %v tickets left out of %v spots\n", remainingTickets, conferenceTickets)
	ticketInfo += "Get your tickets here to attend\n"
	return ticketInfo
}

func getFirstNames()[]string {
	firstNames := []string{}
		for _, booking := range bookings{
			 firstNames = append(firstNames, booking.firstName)
		}
		return firstNames
		
}

func getUserInput()(string, string, string, uint){
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

		return firstName, lastName, email, userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending tickets:\n %v \nto email address %v", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
