package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	// Loops are interesting in Golang
	for _, value := range languages {
		fmt.Println("Value:", value)
	}
}
