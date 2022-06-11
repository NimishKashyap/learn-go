package main

import (
	"encoding/json"
	"fmt"
)

// alias for course struct
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON Video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	nickCourse := []course{
		{"ReactJS bootcamp", 299, "learncode", "password123", []string{"react", "js", "bootcamp"}},
		{"Go bootcamp", 299, "learncode", "password456", []string{"go", "bootcamp"}},
		{"Python bootcamp", 299, "learncode", "password789", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(nickCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS bootcamp",
		"Price": 299,
		"website": "learncode",
		"tags": ["react", "js", "bootcamp"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON WAS VALID")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// Some cases when you just want to add key value pair

	// use interface when not sure about format of value
	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	// fmt.Println(myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("%s: %s -> type : %T\n", key, value, value)
	}
}
