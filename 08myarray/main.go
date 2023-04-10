package main

import "fmt"

func main() {
	fmt.Println("welcome to array")

	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "banana"
	fruitlist[3] = "pears"

	fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("Length of fruit list is: ", len(fruitlist))

	var veglist = [4]string{"tomato", "brinjal", "potato"}
	fmt.Println("veg list is: ", veglist)
}
