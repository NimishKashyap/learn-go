package main

import "fmt"

func main() {

	defer fmt.Println("Hello")
	defer fmt.Println("World")
	defer fmt.Println("One")
	fmt.Println("Two")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
