package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int
	var one int = 1

	ptr = &one

	fmt.Println("The value of ptr is:", ptr)
	fmt.Println("The value of *ptr is:", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is: ", one)
}
