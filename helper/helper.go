package helper

import "strings"

var MYGlobalVariableTest = "Test"

func ValidateUserInput(firstName string, lastName string, email string, ticketCount int64, conferenceTickets int64) (bool, bool, bool) {
	isValidName := len(firstName) > 4 && len(lastName) > 4
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := ticketCount < conferenceTickets

	return isValidName, isValidEmail, isValidTicketCount

}
