package main

import (
	"fmt"
)

// method is same as function but when it is related to class we term it as methos

// there is no class concept in go lang instead there is struct
func main() {
	fmt.Println("welcome to struct in go lang")

	//  there is no inheritance in golang; no super or parent
	jyoti := User{"Jyoti", "abc@gmail.com", true, 21}

	fmt.Println(jyoti)
	fmt.Printf("Jyoti details are %+v\n", jyoti)
	fmt.Printf("Jyoti name id %v\n", jyoti.Name)
	jyoti.GetStatus()
	jyoti.NewEmail()
	fmt.Println(jyoti)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//oneAge int //if first letter is not capital then is is not accessible outside
}

// you can define function below and above the struct
// if func is related to struct then it is called method
func (u User) GetStatus() {
	fmt.Println("Is user active?", u.Status)
}

func (u User) NewEmail() { // here user passes as a copy so original object won't get change that's why pointer is there
	u.Email = "newmail@gmail.com"
	fmt.Println("user new email is: ", u.Email)
}
