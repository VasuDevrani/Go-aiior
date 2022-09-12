# All information in one repo - GoLang
<img src="https://cdn.thenewstack.io/media/2022/05/57bb2a1f-golang.png" height="250px"/>

## Installation and HELLO WORLD
- install latest go version from https://go.dev/learn/
- use a proper code editor like VS Code
- create a folder dedicated to goLang
- create ```main.go```
- install recommended packages
- using terminal create go mod file as a good go practice, ```go mod init module_name```
- create first go file main.go, run ```go run file_name```

- ```
    package main
    import "fmt"

    func main() {
      fmt.Println("hello world");
    }
  ```
 ## GENERAL INFO
 - Go is a statically typed language

### LEXER
responsible for handling the semi-colons and other syntax stuff

### Basic Program
```
package main

import "fmt" //hover over to see different % verbs

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// fmt.Printf("%T", conferenceName);

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Printf("We have total of %v tickets and %v are still available \n", 40, remainingTickets)
	fmt.Println("get your tickets here to attend")

	//variables either need a initialization or type like ts
	var firstName string
	var lastName string
	var email string
	var tickets int
	// var userTickets int

	// taking input
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName);
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName);
	fmt.Println("Enter your email:")
	fmt.Scan(&email);
	fmt.Println("Enter number of tickets you want:")
	fmt.Scan(&tickets);

	fmt.Printf("Thank You %v %v for the %v tickets. You will receive a confirmation email at %v soon\n", firstName, lastName, tickets, email); 
}
```

## RESOURCES
- https://go.dev/doc/ - official docs
- https://www.youtube.com/watch?v=yyUHQIec83I&t=2291s&ab_channel=TechWorldwithNana - YouTube, Thanks Nana
