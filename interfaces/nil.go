package main

import "fmt"

//A Nil is an empty data type
//A Nil is only usable in certain data types, such as interface, function, map, slice, pointer, and channel

func NewMap(name string) map[string]string{
	if name == "" {
		return nil
	} else {
		return map[string] string{
			"name": name,
		}
	}
}

func main(){
	data := NewMap(" ")
	if data == nil{
		fmt.Println("Empty Data")
	}else {
		fmt.Println(data)
	}
}