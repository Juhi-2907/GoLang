package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to file in golang")

	content := "This needs to go in file - google.com"

	file, err := os.Create("./mygolang.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length of file is : ", length)
	defer file.Close()
	ReadFile("./mygolang.txt")
}

func ReadFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)

	checkNilErr(err)
	fmt.Println("Text data in file is \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
