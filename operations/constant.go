package main

import "fmt"

func main() {
	// const firstName string = "Himena"
	// const lastName string = "Tsukimiya"

	//error
	// firstName = "Unchangeable"
	// lastName = "Unchangeable"

	const (
		firstName string = "Himena"
		lastName         = "Tsukimiya"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}