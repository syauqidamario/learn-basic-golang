package main

import "fmt"

func main() {
	name := "Bimo"

	switch name {
	case "Bimo":
		fmt.Println("Hello Bimo")
	case "Dhanu":
		fmt.Println("Hello Dhanu")
	default:
		fmt.Println("Hello there")
	}

	switch length := len(name); length > 10{
	case true:
		fmt.Println("Name too long")
	case false:
		fmt.Println("Name is correct")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name too long")
	case length > 5:
		fmt.Println("Name is a bit long")
	default:
		fmt.Println("Name is correct")
	}
}