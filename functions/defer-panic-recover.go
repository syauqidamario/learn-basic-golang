package main

import "fmt"

//Defer is a function that has a scheduled execution after another function was executed

func logging(){
	fmt.Println("Done calling the function")
}

func endApp(){
	fmt.Println("End App")
	message := recover()
	fmt.Println("Error occurs", message)
}

func runApp(error bool){
	defer endApp()
		if error{
			panic("ERROR")
		}
		message := recover()
		fmt.Println("Error happens", message)
}
