package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to talk to local server")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/post"

	//form data

	data := url.Values{}
	data.Add("firstname", "Jyoti")
	data.Add("lastname", "Sharma")
	data.Add("email", "jojo@gmai.com")

	response, err := http.PostForm(myUrl, data)
	checkNilErr(err)

	defer response.Body.Close()

	dataByte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	content := string(dataByte)
	fmt.Println(content)
}

func PerformPostRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"GoLearn",
			"price":0,
			"platform":"LCO"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	checkNilErr(err)

	defer response.Body.Close()

	dataByte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	content := string(dataByte)
	fmt.Println(content)

}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	checkNilErr(err)

	defer response.Body.Close()
	fmt.Println("Status code is: ", response.Status)
	fmt.Println("Content length is: ", response.ContentLength)

	dataByte, err := ioutil.ReadAll(response.Body) //content in dataByte is in byte format
	checkNilErr(err)

	// content := string(dataByte)
	// fmt.Println(content)

	//Another way to simplyfy response using string builder library
	var responseString strings.Builder
	byteCount, _ := responseString.Write(dataByte)
	fmt.Println("Byte count in response string is: ", byteCount)
	fmt.Println("content is ", responseString.String())
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
