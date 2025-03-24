package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Mango", "Tomato"}

	fruitList = append(fruitList, "Orange", "Cherry")

	fmt.Println(fruitList)

	fruitList = append(fruitList[:3], "Test")
	fmt.Println(fruitList)

	fruitList = append(fruitList, fruitList...)
	fmt.Println(fruitList)

	highScores := make([]int, 5, 10)
	highScores[0] = 234
	highScores[1] = 214
	highScores[2] = 244
	highScores[3] = 254
	fmt.Println(&highScores[0])
	// highScores[4] = 213 panic
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(&highScores[0])
	fmt.Println(highScores)
	sort.Ints(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
}
