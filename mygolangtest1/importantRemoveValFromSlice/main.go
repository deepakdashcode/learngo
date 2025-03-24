package main

import "fmt"

func main() {
	fmt.Println("Test")

	var coursesList = []string{"react", "js", "Swift", "Python", "ruby"}
	fmt.Println(coursesList)

	var index int = 2
	coursesList = append(coursesList[:index], coursesList[index+1:]...)
	fmt.Println(coursesList)
}
