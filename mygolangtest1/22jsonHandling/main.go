package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json Handling")
	// encodeJson()
	DecodeJson()
}

func encodeJson() {

	lcoCourses := []course{
		{"JS", 100, "LCO", "abc123", []string{"web", "js"}},
		{"MERN", 200, "LCO", "bc123", []string{"web", "js", "full stack"}},
		{"JAVA", 100, "LCO", "ab123", nil},
	}
	fmt.Println(lcoCourses)
	fmt.Println()

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
            "courseName": "JS",
            "Price": 100,
            "website": "LCO",
            "tags": ["web","js"]
    }
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON WAS VALID")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json aint valid bro")
	}

	// some cases where we want only key value pair

	var onlineData map[string]any

	json.Unmarshal(jsonDataFromWeb, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	for key, val := range onlineData {
		fmt.Println("Key = ", key, "Value is = ", val)
		fmt.Printf("Type of data is %T\n", val)
	}

	var test any
	test = "hii"
	fmt.Println("test = ", test)
	fmt.Printf("test Type %T\n", test)
	test = 10
	fmt.Println("test = ", test)
	fmt.Printf("test Type %T\n", test)

}
