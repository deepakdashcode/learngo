package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice val is 1, you are ready to open")

	case 2:
		fmt.Println("Dice val is 2")
	case 3:
		fmt.Println("Dice val is 3")
	case 4:
		fmt.Println("Dice val is 4")
		fallthrough
	case 5:
		fmt.Println("Dice val is 5")
	case 6:
		fmt.Println("Dice val is 6, roll again")

	default:
		fmt.Println("What bro ?")
	}
}
