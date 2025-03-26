package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=10erea"

func main() {
	fmt.Println("Welcome to handling url")
	fmt.Println(myUrl)

	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query is %T\n", qparams)
	fmt.Println(qparams)

	fmt.Println(qparams["coursename"])

	for key, val := range qparams {
		fmt.Println("Key is ", key)
		fmt.Println("Value is ", val)
	}

	// Building an URL
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
