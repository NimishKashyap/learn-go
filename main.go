package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// ask the user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter the Last: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter the number of tickets you want: ")
		fmt.Scan(&userTickets)

		if uint(userTickets) < remainingTickets {
			remainingTickets -= uint(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The whole Slice: %v\n", bookings)
			fmt.Printf("First value: %v\n", bookings[0])
			fmt.Printf("Slice Type: %T\n", bookings)
			fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("We have %v remaining tickets\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry we are sold out")
				break
			}

		} else if userTickets == int(remainingTickets) {
			//do someting else
		} else {
			fmt.Println("Sorry, we don't have enough tickets")
			fmt.Println("Please try again")
			continue
		}
	}

}
