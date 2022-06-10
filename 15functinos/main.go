package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()

	fmt.Println(add(10, 20))
	fmt.Println(proAdder(1, 2, 3, 4, 5))

}
func greeterTwo() {
	fmt.Println("Anoter method")
}
func greeter() {
	fmt.Println("Namaste from Golang")
}
func add(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}
	return total
}
