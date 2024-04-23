package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHi(hasName HasName) {
	fmt.Println("Hello there", hasName.GetName())
}

//Implementing interface
//If the data type was the same as the interface's, it will automatically considered that interface
//This means it doesn't need to be manually implemented

type Person struct {
	Name string
}

func (person Person) GetName() string { 
	return person.Name 
}

type Animal struct{
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
func main(){
	animal := Animal{Name: "Cat"}
	person := Person{Name: "Mirai" }
	SayHi(person)
	SayHi(animal)
}