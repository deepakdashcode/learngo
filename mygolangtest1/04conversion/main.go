package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	inp, err := reader.ReadString('\n')
	fmt.Print("Rating is ", inp)

	fmt.Printf("Error is %v\n", err)

	numRating, err1 := strconv.ParseFloat(strings.TrimSpace(inp), 64)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Added 1 to rating", numRating+1)
	}

}
