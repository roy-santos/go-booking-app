package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application!")

	if remainingTickets == 0 {
		fmt.Println("The conference is sold out")
	} else {
		fmt.Printf("We have a total of %v tickets and %v tickets  are"+
			"still available.\n", conferenceTickets, remainingTickets)

	}

	fmt.Println("Get your tickets here to attend.")
}
