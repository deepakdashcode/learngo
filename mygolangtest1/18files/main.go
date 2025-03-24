package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")
	content := "This is the content of file"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("The length is ", length)

	defer file.Close()

	readFile("myfile.txt")
	readNew("myfile.txt")
}

func readFile(filename string) {
	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(byteData)
	fmt.Println(string(byteData))

}

func readNew(filename string) {
	data, err := os.ReadFile(filename)
	fmt.Println("DATA OF FILE")
	fmt.Println(string(data))
	fmt.Println(err)

}
