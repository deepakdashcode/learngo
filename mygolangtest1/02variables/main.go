package main

import "fmt"

const LoginToken = "ajsdhfguiasjiu"

func main() {

	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type %T \n", smallValue)

	var smallFloatValue float32 = 255.352348
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of type %T \n", smallFloatValue)

	var largeFloat float64 = 255.35234843253534
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type %T \n", largeFloat)

	// default values and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T \n", anotherVariable)

	// implicite type
	var website = "github.com"
	fmt.Println(website)

	// short style init
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
