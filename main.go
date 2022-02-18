package main

import (
	"fmt"
	"strings"
)

func main() {

	var firstName string
	var lastName string
	var email string
	var ticketCount int
	var bookings [10]string
	var bookingsSlice []string

	for {
		fmt.Print("Please enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Please enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Please enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Please enter your ticket count: ")
		fmt.Scan(&ticketCount)
		conferenceName := "Go Conference"
		const conferenceTickets int = 50

		
		var remainingTickets = conferenceTickets - ticketCount

		if remainingTickets >= 0 {
			fmt.Printf("Welcome to %v booking\n", conferenceName)
			fmt.Printf("We have total of %v tickets and %v are still availabe ...\n", conferenceTickets, remainingTickets)
	
			fmt.Printf("My name is %v %v I signed out with this email %v and I have %v ticket\n", firstName, lastName, email, ticketCount)
	
			// Array ...
			bookings[0] = firstName + " " + lastName
	
			fmt.Printf("whole values of booking array is %v \n", bookings)
			fmt.Printf("first value of booking array is %v \n", bookings[0])
			fmt.Printf("size of booking array is %v \n", len(bookings))
	
			// Slice ...
	
			bookingsSlice = append(bookingsSlice, firstName+" "+lastName)
	
			fmt.Printf("whole values of booking array is %v \n", bookingsSlice)
			fmt.Printf("first value of booking array is %v \n", bookingsSlice[0])
			fmt.Printf("size of booking array is %v \n", len(bookingsSlice))
	
			var firstNames []string
			for _, element := range bookingsSlice {
				name := strings.Fields(element)
				// fmt.Println()
				firstNames = append(firstNames, name[0])
			}
			fmt.Println(firstNames)
		} else{
			fmt.Println("Please enter valid ticket count ...")
			continue
		}
		
	}

}
