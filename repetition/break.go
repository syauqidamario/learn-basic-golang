package main

import "fmt"

func main(){
	for i := 0; i < 20; i++{
		if i == 10{
			break
		}
		fmt.Println("Queen", i)
	}

	for i := 0; i < 30; i++{
		if i%2 == 0 {
			continue
		}
		fmt.Println("Even Number Loop", i)
	}
}