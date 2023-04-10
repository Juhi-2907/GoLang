package main

import "fmt"

func main() {
	fmt.Println("welcome to if else")

	var num = 10
	if num < 10 {
		fmt.Println("number is less than 10")
	} else if num > 10 {
		fmt.Println("number is greater than 10")
	} else {
		fmt.Println("number is exactly 10")
	}

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if num2 := 3; num2 < 10 {
		fmt.Println("number 2 is less than 10")
	} else {
		fmt.Println("number is greater than equal to 10")
	}
}
