package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")

	user := User{"Deepak", "deepak@gmail.com", true, 21}
	fmt.Println(user)
	fmt.Printf("Deatils are %+v\n", user)
	fmt.Printf("Name is %v\n", user.Name)
	fmt.Printf("%v\n", user)
	user.GetStatus()
	user.NewMail()

	fmt.Println(user)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int // this var cant be exported
}

func (user User) GetStatus() {
	fmt.Println("Is user active", user.Status)
}

func (user User) NewMail() {
	user.Email = "test@goemail.com"
	fmt.Println("Email of user", user.Name, "is", user.Email)
}
