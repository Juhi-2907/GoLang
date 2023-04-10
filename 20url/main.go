package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&payment_id=j54cevkr"

func main() {
	fmt.Println("welcome to handling go url")
	fmt.Println("Url is: ", myurl)

	result, err := url.Parse(myurl)

	checkNilErr(err)

	fmt.Println("Parts of url:")
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params is: %T\n", qparams) //key value pairs
	fmt.Println(qparams["coursename"])

	for i, val := range qparams {
		fmt.Println("val is: ", i, val)
	}

	urlFormation := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=jyoti",
	}

	myUrl2 := urlFormation.String()
	fmt.Println(myUrl2)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
