package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to tim estudy golang")

	presentTime := time.Now()

	fmt.Println("The present time is:", presentTime)

	fmt.Println("The formatted time is:", presentTime.Format("2006-01-02 Monday 15:04:05"))

	createdDate := time.Date(2020, time.April, 20, 15, 04, 05, 0, time.UTC)
	fmt.Println("The created date is:", createdDate.Format("2006-01-02 Monday 15:04:05"))
}
