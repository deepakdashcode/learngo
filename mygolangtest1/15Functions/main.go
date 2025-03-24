package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in go")
	greeter("Deepak")
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Addition is ", result)

	proRes, msg, num := proAdder(3, 4, 5)
	fmt.Println("Addition is ", proRes, msg, num)

}

func adder(val1 int, val2 int) int {
	return val1 + val2
}
func greeterTwo() {
	fmt.Println("Another method")
}
func greeter(name string) {
	fmt.Printf("Hello %v\n", name)
}

func proAdder(values ...int) (int, string, int) {
	sum := 0
	for _, val := range values {
		fmt.Println("val is ", val)
		sum += val
	}
	return sum, "hii function", 0
}
