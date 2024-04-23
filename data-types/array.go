package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "Dhanu"
	names[1] = "Bimo"
	names[2] = "Gibran"
	names[3] = "Syauqi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	var values = [...] int{
		66,
		67,
		71,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}