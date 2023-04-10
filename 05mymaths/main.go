package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)


func main(){
	fmt.Println("Maths in go lang")

	// var num1 int = 2
	// var num2 float64 = 4.5

	// fmt.Println("The sum is: ", num1 + int(num2))

	//random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5)+1)

	//random from crypto

	myRandomNum, _ := rand.Int(rand.Reader,big.NewInt(5))
	fmt.Println(myRandomNum)
}