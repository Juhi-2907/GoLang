package main

import "fmt"

// wherever defer keyword is written it will not execute that line
// else it will execute that  statement at last of this function
func main() {
	//it will work as LIFO
	defer fmt.Println("five")
	defer fmt.Println("six")
	fmt.Println("come to end")
	defDef()
}

func defDef() { // also can write val1 int, val2 int
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	defer fmt.Println("four")
}
