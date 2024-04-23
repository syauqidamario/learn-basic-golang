package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Tangerang"
	fmt.Println(address1) // change too
	fmt.Println(address2) // change to Tangerang
}