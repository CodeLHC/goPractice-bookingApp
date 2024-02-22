package main

import (
	"encoding/json"
	"fmt"
	"gopractice-bookingapp/helper"
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

	type BookingError struct {
		Message string `json:"message"`
	}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/conference", getConferenceDetails).Methods("GET")
	r.HandleFunc("/bookings", getBookings).Methods("GET")
	r.HandleFunc("/order", postTicketsOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))		
	}

func bookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, orderID string) ([]UserData, uint, string){

	var userData = UserData {
		OrderID: orderID,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		NumberOfTickets: userTickets,
	}
	
isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

var bookingErrorMessage string

    if !isValidName {
        bookingErrorMessage += "First name or last name you entered is too short"
    }
    if !isValidEmail {
        bookingErrorMessage += "Email address you entered does not contain an @ sign"
    }
    if !isValidTicketNumber {
        bookingErrorMessage += "The number of tickets you've entered is invalid"
    }

    if !isValidName || !isValidEmail || !isValidTicketNumber {
        return nil, remainingTickets, bookingErrorMessage
    }
	
	 bookings = append(bookings, userData)

	remainingTickets = remainingTickets - userTickets
return bookings, remainingTickets, bookingErrorMessage
}

func postTicketsOrder(w http.ResponseWriter, r *http.Request){
	fmt.Println("request method:", r.Method)
	(w).Header().Set("Content-Type", "application/json")
	 w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var order UserData
	var bookingErrorMessage string 
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderID:= uuid.New().String()
	order.OrderID = orderID
	bookings, remainingTickets, bookingErrorMessage = bookTickets(remainingTickets, order.NumberOfTickets, order.FirstName, order.LastName, order.Email, order.OrderID)
	
	 var response interface{}

    if bookingErrorMessage != "" {
        response = struct {
            Error *BookingError `json:"error"`
        }{
            Error: &BookingError{Message: bookingErrorMessage},
        }
    } else {
        response = struct {
            UserData        UserData `json:"order"`
            RemainingTickets uint     `json:"remainingTickets"`
        }{
            UserData:        order,
            RemainingTickets: remainingTickets,
        }
    }

	
	json.NewEncoder(w).Encode(response)
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