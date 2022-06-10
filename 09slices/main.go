package main

import "fmt"

func main() {
	// fmt.Println("Welcome to slices")

	// var fruitList = []string{"Apple", "Banana", "Grape", "Orange"}

	// fmt.Printf("Type of fruitlist: %T\n", fruitList)

	// fruitList = append(fruitList, "Pineapple")

	// fmt.Println("The fruit list is:", fruitList)

	// // Slice slicing LOL
	// fruitList = append(fruitList[1:3])

	// fmt.Println("The fruit list is:", fruitList)

	// highScores := make([]int, 4)

	// highScores[0] = 100
	// highScores[1] = 200
	// highScores[2] = 300
	// highScores[3] = 400

	// highScores = append(highScores, 555, 666, 777)

	// fmt.Println("The high scores are:", len(highScores))

	// sort.Ints(highScores)

	// fmt.Println("Sorted high scores are:", highScores)

	// How to remove value from slice based on index

	var courses = []string{"Reactjs", "Javascript", "Switft", "python", "C#"}
	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
