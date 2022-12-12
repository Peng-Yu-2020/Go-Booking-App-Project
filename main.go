package main

import ("fmt"
		"strings"
		"booking-app/helper"
	)


	// var conferenceName = "Go Conference"
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint= 50
	var bookings = []string{}

func main()  {
	
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, and remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)

	greetUsers()

	// var bookings [50]string // definition for array
	// var bookings []string // definition for slice
	

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if  isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstName()
			fmt.Printf("The first names of bookings are: %v \n", firstNames)
			
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			} 
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
		
	}
	
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email address:")
	fmt.Scan(&email)
	fmt.Printf("Enter number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets
	fmt.Println("------------")

	// bookings[0] = firstName + " " + lastName // used in array case
	bookings = append(bookings, firstName + " " + lastName)
	// fmt.Printf("The whole slice: %v \n", bookings)
	// fmt.Printf("The first value: %v \n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

}