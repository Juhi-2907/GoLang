package main

import "fmt"

// there is no class concept in go lang instead there is struct
func main() {
	fmt.Println("welcome to struct in go lang")

	//  there is no inheritance in golang; no super or parent
	jyoti := User{"Jyoti", "abc@gmail.com", true, 21}

	fmt.Println(jyoti)
	fmt.Printf("Jyoti details are %+v\n", jyoti)
	fmt.Printf("Jyoti name id %v", jyoti.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
