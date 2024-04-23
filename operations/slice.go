package main

import "fmt"

func main() {
	names := [...]string{"Dhanu", "Bimo", "Gibran", "Syauqi", "Kamio"}
	slice := names[2:3]
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	months :=[...]string{"January", "February", "March", "April", "May", "June"}
	monthSlice1 := months[5:]
	monthSlice1[0] = "New May"
	monthSlice1[1] = "New January"
	fmt.Println(months)

	monthSlice2 := append(monthSlice1, "New Holiday")
	monthSlice2[0] = "Oops"
	fmt.Println(monthSlice2)
	fmt.Println(months)

	newSlice := make([]string, 3, 5)
	newSlice[0] = "Spontan"
	newSlice[1] = "Uhuy"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := months[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
}