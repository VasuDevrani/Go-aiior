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
- The answer is no: Go is neither a functional nor an object-oriented language. Rather, it is a procedural language similar to C or even Perl (before it got an object system).

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

### Arrays and Slices in go
- arrays

```
	var array = [50]float32{1, 2, 4};
	// intializes the array with 0 or empty strings in case of string array
	var array2 [20]int;

	array2[0] = 10;
	array[5] = 30

	fmt.Println(array);
	fmt.Println(len(array2));
```
- slices

```
	var slices = []int{1, 2, 4};
	var slices2 []int;

	slices = append(array, 20);
	slices2 = append(array2, 30);

	fmt.Println(slices);
	fmt.Println(len(slices2));
```

### loops and conditionals
- for loop

```
for {
	var name string;

	fmt.Scan(&name);
	fmt.Println(name);
}
var arry = []int{1,2,3,4,5}

for index, value := range arry{
	fmt.Println(value, " at index ", index);
}

for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            i + j
        }
    }
```
- if else
	- uses && || operators
```
if len(name) == 0{
	break
}else if name == "ram"{
	continue
}else{
	var ind = 0;
	for ind < len(name){
		fmt.Println(name[ind])
		ind++
	}
}
```
- switch

```
switch expression {
case x:
   // code block
case y:
   // code block
case z:
...
default:
   // code block
}
```

## ⭐ for keeping a variable not in use
```
// no problem as index is not used here
for _, value := range array{
	fmt.Println(value);
}
```

### string methods
```
import (
	"fmt",
	s "strings"
)
var p = fmt.Println

p("Contains:  ", s.Contains("test", "es"))
p("Count:     ", s.Count("test", "t"))
p("HasPrefix: ", s.HasPrefix("test", "te"))
p("HasSuffix: ", s.HasSuffix("test", "st"))
p("Index:     ", s.Index("test", "e"))
p("Join:      ", s.Join([]string{"a", "b"}, "-"))
p("Repeat:    ", s.Repeat("a", 5))
p("Replace:   ", s.Replace("foo", "o", "0", -1))
p("Replace:   ", s.Replace("foo", "o", "0", 1))
p("Split:     ", s.Split("a-b-c-d-e", "-"))
p("ToLower:   ", s.ToLower("TEST"))
p("ToUpper:   ", s.ToUpper("test"))
```

### functions
```
func myMessage() {
  fmt.Println("I just got executed!")
}
func newMessage(name string){

}

//define data type when function returns something
func moreFunnction(name string, age uint, books []string) []string{
}

func main() {
  myMessage()
  newMessage("golang")
}
```
- multiple returns from a function
```
name, age, address = getUserDetails();

func getUserDetails() (string, uint, string){
//take inputs
return useName, userAge, userAddress
}
```

## Multiple file structure
- go files within a same package or folder are interconnected and should work under a single package
- Example inside folder 'hello';
  - ```
  	go.mod
	helper.go
	main.go
	utils.go
    ```
- the variables defined outside functions inside these files are open to other files as well
- a function present in helper.go can be used in main.go without any sort of import statements
- to run file ```go run main.go helper.go``` or ```go run .```


## RESOURCES
- https://go.dev/doc/ - official docs
- https://www.youtube.com/watch?v=yyUHQIec83I&t=2291s&ab_channel=TechWorldwithNana - YouTube, Thanks Nana
