package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["go"] = "Go"
	languages["py"] = "python"

	fmt.Println("List of all languages", languages)
	fmt.Println(languages["js"])

	delete(languages, "js")
	fmt.Println("List of all languages", languages)

	for key, val := range languages {
		fmt.Println("KEY =", key, "Value is =", val)
	}

}
