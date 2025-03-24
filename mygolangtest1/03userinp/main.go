package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMsg := "Welcome to useInp program"
	fmt.Println(welcomeMsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza")

	// comma ok syntax

	input, err := reader.ReadString('\n')

	if err == nil {
		fmt.Println("Thanks for reading", input)
		fmt.Printf("The rating type is %T\n", input)
		fmt.Printf("The rating type is %T", err)
	}

}
