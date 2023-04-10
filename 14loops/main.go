package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println(fruitList)

	// for d := 0; d < len(fruitList); d++ {
	// 	fmt.Println(fruitList[d])
	// }

	// for index := range fruitList {
	// 	fmt.Println(fruitList[index])
	// }

	// for index, fruit := range fruitList {
	// 	fmt.Printf("index is %v and value is %v\n", index, fruit)
	// }

	for _, fruit := range fruitList {
		fmt.Printf("value is %v\n", fruit)
	}
	index := 1
	for index < 10 {
		fmt.Println(index)
		index++
		if index == 4 {
			goto loc
		}
	}
loc:
	fmt.Println("got statement")
}
