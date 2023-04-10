package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("welcome to web request")
	response, err := http.Get(url)

	checkNilErr(err)
	fmt.Printf("Type of the response is:%T\n", response)

	defer response.Body.Close() //it caller responsibility to close all connetion

	dataByte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	content := string(dataByte)
	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
