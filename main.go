package main

import (
	"fmt"
	"golang/helper"
	// "strconv"
	// "strings"
)

var totalticket uint= 20 
// var cars = make([]map[string]string,0)
var cars = make([]userdata,0)

type userdata struct{
	firstname string
	lastname string
	email string
	numberoftickets uint
	
}


func main() {
	
	// aks the username for their name
	for {
		username, surname, email, userticket := getuserinput()
            
        // create the map for the user
	    //  var userdata = make(map[string]string)
		//  userdata["firstname"]= username
		//  userdata["lastname"]= surname
		//  userdata["email"]= email
		// userdata["numberoftickets"]=strconv.FormatUint(uint64(userticket),10)


	   var userdata =userdata{
		   firstname:username,
		   lastname:surname,
		   email:email,
		   numberoftickets:userticket,
	   }

		// cars = append(cars, username+" "+surname)
		cars = append(cars, userdata)// using the map functionality
		// callling the functions
		greetusers(username) // functions to greet
  
		//calling the functions to print the first name

		// firstname := printfirstname(cars)
		// fmt.Printf("the first name of the user is %v\n", firstname)

		isvalidname, isvlaliedemail := helper.IsvalidUseInput(username, surname, email)

		if isvalidname && isvlaliedemail {
 
			


			totalticket -= userticket
			fmt.Printf("ther whole booklking array is %v\n", cars)
			fmt.Printf("the 1st vlaue is %v\n", cars[0])
			fmt.Printf("the type of the array is %T\n", cars)
			fmt.Printf("the length of the slice is %v\n", len(cars))
			fmt.Printf("the remaning tickets are  %v\n", totalticket)

			// _ blank indetifier it is used for the unused variabels in th GO

			// fmt.Printf("the all the booking are %v\n",cars)

			var noticket bool = totalticket <= 0
			if noticket {
				fmt.Printf("Srry the all the tickets are booked see you next time\n")
				break
			}
		} else {
			if !isvlaliedemail {
				println("the email you entered is invalid pls enter the proper eamil address \n")
			}
			fmt.Printf("the enter input values are invalid\n")
		}

	}
	city := "London"

	switch city {

	case "new york":
		// execute the code for the booking
	case "Singapore":
		// code
	case "London", "Hongkon":
		//code
	default:
		// when none of the above city matches then the default statement get executed
	}

}

func greetusers(username string) {
	fmt.Printf("welcome to our conference and have the happy life %v as you want be aware fo what youn want to do\n", username)
}

// func printfirstname(cars []string) []string {
// 	firstnames := []string{}
// 	for _, booking := range cars {
// 		var names = strings.Fields(booking)
// 		firstnames = append(firstnames, names[0])
// 		firstnames = append(firstnames, booking["firstname"])
// 	}
// 	return firstnames
// }

func getuserinput() (string, string, string, uint) {
	var username string
	var surname string
	var email string
	var userticket uint
	fmt.Println("Enter the username")
	fmt.Scan(&username)
	fmt.Println("Enter the surname")
	fmt.Scan(&surname)
	fmt.Println("Enter the email")
	fmt.Scan(&email)
	fmt.Println("Enter the numbers of the ticket")
	fmt.Scan(&userticket)

	return username, surname, email, userticket
}

func sendticket(userticket uint,firstname string, lastname string,email string ){

  var ticket = fmt.Sprintf("%v ticketf for %v  %v",userticket,firstname,lastname)
  fmt.Println("##########")
  fmt.Printf("send ticket to :\n%v\nto email address %v \n",ticket,email)
  fmt.Println("##########")

}