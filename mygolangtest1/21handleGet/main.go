package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web req")
	// PerformGetRequest("http://localhost:8000/get")
	// PerformPostJSONRequest()
	PerformPostFormDataRequest()
}

func PerformGetRequest(myUrl string) {
	// const myUrl = "http://localhost:8000"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Response Body:", string(body))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(body)
	fmt.Println("Bytecount is", byteCount)
	fmt.Println("Response is ", responseString.String())

}
func PerformPostJSONRequest() {
	myUrl := "http://localhost:8000/post"

	// fake json payLoad

	requestBody := strings.NewReader(`
		{
			"CourseName": "go lang",
			"Price": 0,
			"Platform": "LCO"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("The response is ", string(content))
}

func PerformPostFormDataRequest() {
	myUrl := "http://localhost:8000/postform"

	// form Data

	data := url.Values{}

	data.Add("First Name", "Deepak")
	data.Add("Last Name", "Dash")
	data.Add("email", "dk@gmail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Content is ", string(content))
}
