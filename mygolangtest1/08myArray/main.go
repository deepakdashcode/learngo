package main

import "fmt"

func main() {

	var fruitList [4]string
	fruitList[0] = "Tomato"
	fruitList[1] = "Apple"

	fruitList[3] = "Mango"

	fmt.Println(fruitList)
	fmt.Println("Number of elements = ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "Mushroom"}
	fmt.Println(vegList)
}
