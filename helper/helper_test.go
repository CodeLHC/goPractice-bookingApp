package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserInput(t *testing.T) {

	testcases:= []struct{
		desc string
		firstName string
		lastName string
		email string
		userTickets uint
		remainingTickets uint
		expectedValidNameOutput bool
		expectedValidEmailOutput bool
		expectedValidTicketNumber bool
	}{{
		desc: "all returns true with valid inputs",
		firstName: "Laura",
		lastName: "Crawford",
		email: "laura.crawford@test.co.uk",
		userTickets: 2,
		remainingTickets: 50,
		expectedValidNameOutput: true,
		expectedValidEmailOutput: true,
		expectedValidTicketNumber: true,
	},
{
		desc: "valid name is false if first name is less than 2 chars",
		firstName: "L",
		lastName: "Crawford",
		email: "laura.crawford@test.co.uk",
		userTickets: 2,
		remainingTickets: 50,
		expectedValidNameOutput: false,
		expectedValidEmailOutput: true,
		expectedValidTicketNumber: true,
},
{
		desc: "valid name is false if last name is less than 2 chars",
		firstName: "La",
		lastName: "C",
		email: "laura.crawford@test.co.uk",
		userTickets: 2,
		remainingTickets: 50,
		expectedValidNameOutput: false,
		expectedValidEmailOutput: true,
		expectedValidTicketNumber: true,
},
{
		desc: "valid email is false if no @ symbol",
		firstName: "Laura",
		lastName: "Crawford",
		email: "test.co.uk",
		userTickets: 2,
		remainingTickets: 50,
		expectedValidNameOutput: true,
		expectedValidEmailOutput: false,
		expectedValidTicketNumber: true,
	},
{
		desc: "valid ticket number is false if more user tickets than remaining tickets",
		firstName: "Laura",
		lastName: "Crawford",
		email: "laura.crawford@test.co.uk",
		userTickets: 51,
		remainingTickets: 50,
		expectedValidNameOutput: true,
		expectedValidEmailOutput: true,
		expectedValidTicketNumber: false,
	},
}
for _,tc := range testcases{
	t.Run(tc.desc, func(t *testing.T){
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(tc.firstName, tc.lastName, tc.email, tc.userTickets, tc.remainingTickets)
		assert.Equal(t, tc.expectedValidNameOutput, isValidName)
		assert.Equal(t, tc.expectedValidEmailOutput, isValidEmail)
		assert.Equal(t, tc.expectedValidTicketNumber, isValidTicketNumber)
	})
}
}