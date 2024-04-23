package main

import "fmt"

//An empty interface lacks a method declaration, this means every single data type will be its implementation

func Ups()interface{}{
	return "Oops"
}

func main(){
	empty := Ups()
	fmt.Println(empty)
}