package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[2] = "Banana"
	fruitList[3] = "Grape"

	fmt.Println("The fruit list is:", fruitList)
	// always returns the size of array
	fmt.Println("Length of fruit list: ", len(fruitList))

	var vegLisst = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Veggies list is", vegLisst)
	fmt.Println("Length of Veggies list is", len(vegLisst))
}
