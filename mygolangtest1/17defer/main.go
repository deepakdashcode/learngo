package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer myDefer()
	fmt.Println("Hello")
}

// world, one, two
// Hello

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, " ")
	}
}
