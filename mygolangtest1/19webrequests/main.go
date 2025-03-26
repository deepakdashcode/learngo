package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://example.com"

func main() {
	fmt.Println("Welcome")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type is %T\n", response)
	// fmt.Printf("Response type is %v\n", response)
	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	os.Create("test.html")
	os.WriteFile("test.html", dataBytes, os.ModePerm)
	fmt.Println(string(dataBytes))
}
