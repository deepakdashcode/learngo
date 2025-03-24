package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for ind, val := range days {
		fmt.Printf("Index is %v and Value is %v\n", ind, val)
	}

	rougueVal := 1

	for rougueVal < 10 {

		if rougueVal == 7 {
			goto location
			// break // also allowed
		}
		if rougueVal%3 == 0 {
			rougueVal++
			continue
		}
		fmt.Println("Value is : ", rougueVal)
		rougueVal++
	}

location:
	fmt.Println("Jumping at this location")
}
