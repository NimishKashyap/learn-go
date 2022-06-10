package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs course")
	// no inheritence in golang
	// no super or parent

	nimish := User{"Nimish", "nimish@go.dev", true, 20}

	fmt.Println(nimish)
	fmt.Printf("Nimish Details: %+v\n", nimish)
	fmt.Printf("Name is %v and email is %v", nimish.Name, nimish.Email)

	nimish.GetStatus()
	nimish.NewMail()

	fmt.Printf("Name is %v and email is %v", nimish.Name, nimish.Email)
}

// First Capital letter is very very important
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u *User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
