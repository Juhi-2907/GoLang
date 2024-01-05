package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "hey user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating of my course:")

	//comma ok//err ok syntax
	// eg: input, err := ? But if you don't need to care for input or err just put _ there
	input, _ := reader.ReadString('\n') //read string take a value until what you want to read
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of the input is %T", input) // this gave it is a string
}
