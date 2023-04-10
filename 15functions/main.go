package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to function")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	result2, message := proAdder(1, 2, 3, 4, 5)
	fmt.Println("proAdder Result is: ", result2, message)
}

func adder(val1, val2 int) int { // also can write val1 int, val2 int
	return val1 + val2
}

func proAdder(values ...int) (int, string) { //here argument is slice
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "hi"
}

//you are not allowed to write func inside function

func greeter() {
	fmt.Println("greeter called")
}
