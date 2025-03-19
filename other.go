// package main

// import "fmt"

// func main() {

// 	//    var conferencename ="Go Confernce"
// 	//    var conferencename string ="Go Confernce"
// 	conferencename := "Go Confernce"
// 	//    const conferenctickets = 50
// 	const conferenctickets int = 50
// 	var tickets int = 50

// 	fmt.Printf("conference ticketsis %T ,remainingtickets is %T, conferencename is  %T \n", conferenctickets, tickets, conferencename)

// 	fmt.Println("welcome to", conferencename, "the booking application")
// 	fmt.Printf("welcome to %v of the booking application \n", conferenctickets)
// 	//    fmt.Println("we have total of",conferenctickets,"and",remainingtickets,"are still avialble")
// 	fmt.Printf("we have total of %v and %v are still avialble\n", conferenctickets, tickets)
// 	fmt.Println("get your ticket to attend")

// 	var username string
// 	var surname string
// 	var remainingtickets int
// 	var userticket int
// 	// aks the username for their name
// 	fmt.Println("Enter the username")
// 	fmt.Scan(&username)
// 	fmt.Println("Enter the surname")
// 	fmt.Scan(&surname)
// 	fmt.Println("Enter the numbers of the ticket")
// 	fmt.Scan(&userticket)

// 	remainingtickets = tickets - userticket

// 	// username="tommy"
// 	// remainingtickets = 2
// 	fmt.Printf("use %v %vbook  %v this many ticktets\n", username, surname, remainingtickets)
// 	fmt.Printf("the remainingtickets are %v\n", remainingtickets)


// 	// Arrays and slice in GO

// 	var bookings [50]string
// 	bookings[0]=username+" "+surname

// 	fmt.Printf("ther whole booklking array is %v\n",bookings)
// 	fmt.Printf("the 1st vlaue is %v\n",bookings[0])
// 	fmt.Printf("the type of the array is %T\n",bookings)
// 	fmt.Printf("the size  of the array is %v\n",len(bookings))


// 	// slices in GO
// 	// var cars [] string
// 	// alternative decleration
// 	// var cars= [] string{}
// 	 cars := [] string{}
// 	cars=append(cars, username+" "+surname)
	
// 	fmt.Printf("ther whole booklking array is %v\n",cars)
// 	fmt.Printf("the 1st vlaue is %v\n",cars[0])
// 	fmt.Printf("the type of the array is %T\n",cars)
// 	fmt.Printf("the length of the slice is %v\n",len(cars))
    


// }

