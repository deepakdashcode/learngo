package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNum := 23
	var ptr = &myNum
	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)

	*ptr *= 2
	fmt.Println("Value of pointer is", *ptr)
}
