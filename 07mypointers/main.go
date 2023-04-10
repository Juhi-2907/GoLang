package main

import "fmt"

func main() {
	fmt.Println("welcome to pointer")

	var ptr *int
	fmt.Println("value to unassigned ", ptr)

	myNumber := 4

	var myptr = &myNumber
	fmt.Println("value to assigned ", myptr)
	fmt.Println("value in pointer is ", *myptr)

	*myptr = *myptr + 2
	fmt.Println("new value of myNumber ", myNumber)

}
