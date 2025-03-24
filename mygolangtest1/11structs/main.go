package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")

	user := User{"Deepak", "deepak@gmail.com", true, 21}
	fmt.Println(user)
	fmt.Printf("Deatils are %+v", user)
	fmt.Printf("Name is %v", user.Name)
	fmt.Printf("%v", user)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
