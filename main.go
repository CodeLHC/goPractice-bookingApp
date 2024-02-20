package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
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
		OrderID string `json:"orderId"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		Email string `json:"email"`
		NumberOfTickets uint `json:"numberOfTickets"`
	}

	// var wg = sync.WaitGroup{}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/conference", getConferenceDetails).Methods("GET")
	r.HandleFunc("/bookings", getBookings).Methods("GET")
	r.HandleFunc("/order", postTicketsOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))

	
		// firstName, lastName, email, userTickets := getUserInput()
		// isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// if isValidName && isValidEmail && isValidTicketNumber{
		// 	bookings, remainingTickets = bookTickets(remainingTickets, userTickets, firstName, lastName, email, orderId)
		// 	wg.Add(1)
		// 	// go sendTicket(userTickets, firstName, lastName, email)


		// if remainingTickets == 0 {
		// 	fmt.Println("Our conference is booked out. Come back next year")
		// }
		// } else {
		// 	if !isValidName{
		// 		fmt.Print("First name or last name you entered is too short\n")
		// 	}
		// 	if !isValidEmail{
		// 		fmt.Print("Email address you entered does not contain an @ sign\n")
		// 	}
		// 	if !isValidTicketNumber{
		// 		fmt.Print("The number of tickets you've entered is invalid\n")
		// 	}
		// }
		// wg.Wait()
	}



func bookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, orderID string) ([]UserData, uint){
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		OrderID: orderID,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		NumberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
return bookings, remainingTickets
}

func getConferenceDetails( w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	response:= ConferenceInfo{
		ConferenceName: conferenceName,
		TotalTickets: conferenceTickets,
		TicketsLeft: remainingTickets,
	}
json.NewEncoder(w).Encode(response)
}

func getBookings( w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func postTicketsOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var order UserData
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderID:= uuid.New().String()
	order.OrderID = orderID
	bookings, remainingTickets = bookTickets(remainingTickets, order.NumberOfTickets, order.FirstName, order.LastName, order.Email, order.OrderID)
	json.NewEncoder(w).Encode(bookings)
}

// func getUserInput()(string, string, string, uint){
// 	var firstName string
// 		var lastName string
// 		var email string
// 		var userTickets uint

// 		fmt.Println("Enter your first name: ")
// 		fmt.Scan(&firstName)

// 		fmt.Println("Enter your last name: ")
// 		fmt.Scan(&lastName)

// 		fmt.Println("Enter your email address: ")
// 		fmt.Scan(&email)

// 		fmt.Println("Enter amount of tickets: ")
// 		fmt.Scan(&userTickets)

// 		return firstName, lastName, email, userTickets
// }

// func sendTicket(userTickets uint, firstName string, lastName string, email string){
// 	time.Sleep(10 * time.Second)
// 	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
// 	fmt.Println("#################")
// 	fmt.Printf("Sending tickets:\n %v \nto email address %v", ticket, email)
// 	fmt.Println("#################")
// 	wg.Done()
// }

