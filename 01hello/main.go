package main

import (
    "fmt"
	"hello/helper"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	helper.Count();

	var MAP = make(map[string]string)
	MAP["name"] = "vasu"
	// fmt.Printf("%T", conferenceName);

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Printf("We have total of %v tickets and %v are still available \n", 40, remainingTickets)
	fmt.Println("get your tickets here to attend")

	//variables either need a initialization or type like ts
	// var firstName string
	// var lastName string
	// var email string
	// var tickets int
	// var userTickets int

	// // taking input
	// fmt.Println("Enter your first name:")
	// fmt.Scan(&firstName);
	// fmt.Println("Enter your last name:")
	// fmt.Scan(&lastName);
	// fmt.Println("Enter your email:")
	// fmt.Scan(&email);
	// fmt.Println("Enter number of tickets you want:")
	// fmt.Scan(&tickets);

	// fmt.Printf("Thank You %v %v for the %v tickets. You will receive a confirmation email at %v soon\n", firstName, lastName, tickets, email);

	 
	// for {
	// 	var name string;

	// 	fmt.Scan(&name);
	// 	fmt.Println(name);
	// }

	// var arry = []int{1,2,3,4,5}

	// for index, value := range arry{
	// 	fmt.Println(value, " at index ", index);
	// }


// var p = fmt.Println

//     p("Contains:  ", s.Contains("test", "es"))
//     p("Count:     ", s.Count("test", "t"))
//     p("HasPrefix: ", s.HasPrefix("test", "te"))
//     p("HasSuffix: ", s.HasSuffix("test", "st"))
//     p("Index:     ", s.Index("test", "e"))
//     p("Join:      ", s.Join([]string{"a", "b"}, "-"))
//     p("Repeat:    ", s.Repeat("a", 5))
//     p("Replace:   ", s.Replace("foo", "o", "0", -1))
//     p("Replace:   ", s.Replace("foo", "o", "0", 1))
//     p("Split:     ", s.Split("a-b-c-d-e", "-"))
//     p("ToLower:   ", s.ToLower("TEST"))
//     p("ToUpper:   ", s.ToUpper("test"))
    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

