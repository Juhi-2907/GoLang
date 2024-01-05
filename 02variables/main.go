package main

import "fmt"

// globally or outside of any method walrus operator is not allowed
//jwtToken := 34735756

// first capital letter of the variable name is making it public so you can use anywhere in other file
// int other language we need to specify by using public keyword
const LoggingToken = "hey"

// As js every variable is object here it is of types
func main() {
	var username string = "Jyoti Kumari"
	fmt.Println(username)
	//placeholder %T for type and \n for new line
	fmt.Printf("variable is of type: %T \n", username)
	//implicit var type
	var user = "jyoti"
	fmt.Printf(user)
	//walrus operator
	noUser := "Hemma"
	fmt.Printf(noUser)
}
