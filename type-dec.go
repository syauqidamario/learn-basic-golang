package main

import "fmt"

func main(){
	type NoKTP string

	var ktpSyauqi NoKTP = "13579013"
	fmt.Println(ktpSyauqi)
	fmt.Println(NoKTP("3333333"))
}