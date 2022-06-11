package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=free"

func main() {
	fmt.Println("Welcome to handling url in golang")
	fmt.Println(myurl)

	//parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("The type of query params are : %T", queryParams)

	fmt.Println("The course name is : ", queryParams.Get("coursename"))

	for _, val := range queryParams {
		fmt.Println("Params is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev:3000",
		Path:     "/tutcss",
		RawQuery: "user=nimish",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println("The url is:", anotherURL)
}
