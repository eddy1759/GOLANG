// package main

// import "fmt"

// func main() {
// 	var conferenceName = "Go Conference"
// 	const conferenceTickets = 50
// 	var remainingTickets uint = 50

// 	fmt.Printf("Welcome to %v booking application\n", conferenceName)
// 	fmt.Printf("We have total of %v tickets available and %v tickets remaining\n", conferenceTickets, remainingTickets)
// 	fmt.Println("Get your ticket here to attend")

// 	bookings := []string{}

// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint

// 	fmt.Println("Enter your first name: ")
// 	fmt.Scan(&firstName)

// 	fmt.Println("Enter your last name: ")
// 	fmt.Scan(&lastName)

// 	fmt.Println("Enter your email address: ")
// 	fmt.Scan(&email)

// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scan(&userTickets)

// 	remainingTickets = remainingTickets - userTickets
// 	bookings = append(bookings, firstName + " " + lastName)

// 	fmt.Printf("The whole Slice: %v\n", bookings)
// 	fmt.Printf("The first value: %v\n", bookings[0])
// 	fmt.Printf("Slice type: %T\n", bookings)
// 	fmt.Printf("Slice length: %v\n", len(bookings))

// 	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
// 	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)


//}