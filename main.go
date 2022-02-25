package main

import (
	"awesomeProject/helper"
	"fmt"
	"strconv"
)

func main() {

	bookingsSlice := []map[string]string{}
	const conferenceTickets int64 = 50

	for {
		fmt.Println("This is global variable ... ", helper.MYGlobalVariableTest)

		firstName, lastName, email, ticketCount := setUserInput()
		isValidName, isValidEmail, isValidTicketCount := helper.ValidateUserInput(firstName, lastName, email, ticketCount, conferenceTickets)
		conferenceName := "Go Conference"

		var remainingTickets = conferenceTickets - ticketCount

		if isValidName && isValidEmail && isValidTicketCount {
			WelcomeMsgForValidatedUsers(conferenceName, conferenceTickets, remainingTickets, firstName, lastName, email, ticketCount)
			userData := map[string]string{}
			userData["firstname"] = firstName
			userData["lastName"] = lastName
			userData["email"] = email
			userData["ticketCount"] = strconv.FormatInt(ticketCount, 10)
			bookingsSlice = append(bookingsSlice, userData)
			firstNames := getUserFirstName(bookingsSlice)
			fmt.Println(firstNames)

		} else {
			if !isValidName {
				fmt.Println("Please enter valid name ...")
			}
			if !isValidEmail {
				fmt.Println("Please enter valid email ...")
			}
			if !isValidTicketCount {
				fmt.Println("Please enter valid ticket count ...")
			}

			continue
		}

	}

}
func WelcomeMsgForValidatedUsers(conferenceName string, conferenceTickets int64, remainingTickets int64, firstName string, lastName string, email string, ticketCount int64) {
	fmt.Printf("Welcome to %v booking\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availabe ...\n", conferenceTickets, remainingTickets)

	fmt.Printf("My name is %v %v I signed up with this email %v and I have %v ticket\n", firstName, lastName, email, ticketCount)
}

func getUserFirstName(bookingsSlice map[string]string) []string {
	var firstNames []string
	for _, element := range bookingsSlice {
		fmt.Println(element)
		//firstNames = append(firstNames, element["firstName"])
	}
	return firstNames
}

func setUserInput() (string, string, string, int64) {
	var firstName string
	var lastName string
	var email string
	var ticketCount int64
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Please enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Please enter your ticket count: ")
	fmt.Scan(&ticketCount)
	return firstName, lastName, email, ticketCount
}
