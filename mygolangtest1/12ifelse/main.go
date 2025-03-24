package main

import "fmt"

func main() {
	fmt.Println("If else test")
	loginCount := 7
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "10 counts bro"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Number is greater than 10 or equal")
	}
}
