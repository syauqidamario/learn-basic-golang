package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	dan := Man{"Dan"}
	dan.Married()
	fmt.Println(dan.Name)
}