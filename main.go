package main

import "fmt"

func main() {
	var yourName = "Bruce Ismay"
	ticketsLeft := 2200
	ticketClass := "First"
	ticketValue := 200

	fmt.Println("Welcome to White Star Line Cruises")
	fmt.Println("You can book a ticket for your transatlantic trip here!")
	fmt.Printf("Welcome aboard %s\n", yourName)
	fmt.Printf("Your ticket is a %s class ticket and it costs %d pounds", ticketClass, ticketValue)
	fmt.Printf("There are %d tickets left", ticketsLeft-1)
}
