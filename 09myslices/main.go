package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// two key word for memory management new and make
	highScore := make([]int, 4)

	highScore[0] = 245
	highScore[1] = 246
	highScore[2] = 247
	highScore[3] = 248
	// highScore[4] = 249 //error saying accessing out of boundary

	highScore = append(highScore, 4454, 4545, 4546)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slices based on indexes
	var courses = []string{"react", "js", "swift", "python", "ruby"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
