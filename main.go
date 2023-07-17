package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	welcome(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email name")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booked %v tickets. You will receive it at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets reamaining", remainingTickets)
			fmt.Printf("These are all our bookings: %v\n", bookings)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := []string{}
			for _, bookings := range bookings {
				var names = strings.Fields(bookings)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}
			fmt.Printf("These are all the first names: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("There no more tickets")
				break
			}
		} else {
			fmt.Printf("There no enougth tickets. Currently available %v", remainingTickets)
			continue
		}		
	}

	wg.Wait()
}

func welcome(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get here your tickets to attend")
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
  fmt.Println("##################")
	fmt.Printf("Sending tickets (%v) to %v %v => %v", userTickets, firstName, lastName, email)
	fmt.Println("##################")
	wg.Done()
}
