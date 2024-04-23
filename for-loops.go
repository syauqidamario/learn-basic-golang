package main

import "fmt"

func main(){
	counter := 1

	for counter <=100{
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter := 2; counter <= 100; counter++{
		fmt.Println("The Tortured Poets Department", counter)
	}

	//for loops with range
	names := []string{"Taylor", "Jeremy", "Dylan"}
	for index, name := range names{
		fmt.Println("index", index, "=", name)
	}
}