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
- Cuncurrency is the important need for deploying servers, only languages like c++, java provides this
- Go comes handy as its easy to write like javascript and static, concurrent like c++
- The answer is no: Go is neither a functional nor an object-oriented language. Rather, it is a procedural language similar to C or even Perl (before it got an object system).

### LEXER
responsible for handling the semi-colons and other syntax stuff

### ```go mod tidy```
go mod tidy ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.

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

### defer
- putting a ```defer``` keyword before the statement, makes it run at the near end of program, reverse order (LIFO) in case of multiple defers
```
defer fmt.Println("World")
fmt.Println("hello")

=> hello World
```
```
defer fmt.Println("One")
defer fmt.Println("Two")
defer fmt.Println("Three")

=> Three Two One
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

### Maps
```
var MAP = make(map[string]string)
var MAP1 map[string]string
MAP["name"] = "vasu"

//list of maps

var mapList []map[string]string
var mapListOne = make([]map[string]string, 0) //, 0 is for initialization, not required
```
```
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
```

### Structs
```
type UserData struct {
name string
age uint
email string
}

var userOne = UserData {
name: "vasu",
age: 20,
email: "vasu4arodev@gmail.com"
}

fmt.Println(userOne.name)
```
- printing struct data
  1) %v
     ```
     {hitesh, 20, hitesh@gmail.com}
     ```
  2) %+v
     ```
     {name: hitesh, age: 20, email: hitesh@gmail.com}
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

## Multiple package structure
- while we use anything from a package it needs to be imported first that too varies depending type of package
- create a folder say 'goPrograms' with main.go and mode file with a package name say 'hello'
- for new package, create a new folder say 'helper', create files in it.
- while using a package exports in another package, always import as
```
import {
   "fmt",
   "hello/helper"
}
```
- for export variables or functions write them in capitalized form like ```GetData()```
- for import ```helper.GetData()```

### http calls
- get request
```
import (
   "fmt"
   "net/http"
   "io/ioutil"
)

response, err := http.Get("http://hello.com")
if(err != nil)
   panic(err)

// here content is of type []uint8 that is converted into string

content, err := ioutil.ReadAll(response.Body)
fmt.Print(string(content))
fmt.Printf("%T", content)

```
- url decoding
```
const url = "http://local_mania.org"
func main(){
	result, _ := url.Parse(url)
	fmt.Println(url.Scheme)
	fmt.Println(url.Host)
	fmt.Println(url.RawQuery)
	
	var qParams = url.query() //gives the map of queries 
}
```
- url contruction
```
var url = &url.URL{
   Scheme: "https",
   Host: "",
   Path: ""
}

anotherURL = url.String();
```

- POST req
```
const URL = "http://fakejson.com"

requestBody := strings.NewReader(`
    {
    	"courseName": "Btech",
	"duration": "4 yrs"
    }
`)

response, _ := http.Post(URL, "application/json", requestBody);
```
- JSON formatting

```
type course struct {
	Name string `json:"website"` //following json script used as key
	Price int
	Password string `json:"-"` // "-" is left in main json
	Tags []string `json:"tags,omitempty"` //omitmpty removes nil fields
}

func main() {

courses := []course{
	{"ReactJS", 299, "abc123", []string{"mern stack", "new course"}},
	{"JavaScript", 229, "def123", []string{"frontend", "Month ago"}},
}

finalJson, err := json.MarshalIndent(courses, "", "\t")

// bytes format json
fmt.Print(finalJson)

fmt.Print(string(finaJson))
}
```
- json.parse, storing fetched json data using structs
```
jsonData := []byte(`
	{
		"name": "vasu",
		"Price": 288
	}
`)

var newData course
checkValid := json.Valid(jsonData)

if checkValid {
	json.Unmarshal(jsonData, &newData)
} else {
	fmt.Print("not valid")
}

fmt.Printf("%#v\n", newData)
```

## Concurrency
- This is like creating a asynchronous path for long duartion work like fetching data from api or sending automated mails
- Lets say we have a function ```sendTicket``` that takes a long 10 sec duration, when multiple users are on the application this could block the go program thread flow and block other users from using the app concurrently
- Go provides a way to separate this blocking task in a separate thread and get its results when its completed, deleting the thread as well
- This way if user ```A``` enters the app, fills details and then reaches this blocking task, this would be separated in separate thread and another user ```B``` can use the application
- just use ```go``` keyword to make this happen
```
go getData();

//blocking code
func getData(){

}
```
### WaitGroups
- concurrency could lead to exit of application without waiting for all threads to complete
- thereby we use wautgroups that keeps app running until unless all added concurrent threads are done'
- https://gobyexample.com/waitgroups

## Gorilla mux and CRUD operations using GoLang
- The name mux stands for "HTTP request multiplexer". Like the standard http.ServeMux, mux.Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions. 
- https://github.com/gorilla/mux

## RESOURCES
- https://go.dev/doc/ - official docs
- https://www.youtube.com/watch?v=yyUHQIec83I&t=2291s&ab_channel=TechWorldwithNana - YouTube, Thanks Nana
- https://mj-go.in/golang/crud-rest-api-with-gorilla-mux - CRUD operations in go using mux
- https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa - YouTube, Thanks Hitesh
