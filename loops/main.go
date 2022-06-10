package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	// normal for loop type 1
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// normal for loop type 2
	// only day
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// normal for loop type 3
	// index - day
	// for index, day := range days {
	// 	fmt.Println(index, day)
	// }

	rougeValue := 1

	for rougeValue < 10 {
		fmt.Println("Value is", rougeValue)
		rougeValue++
	}
}
