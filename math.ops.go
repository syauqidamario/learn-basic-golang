package main

import "fmt"

func main(){
	var p = 100
	var q = 100
	var r = p + q
	fmt.Println(r)

	var s = 100
	s += 100
	s -= 100
	fmt.Println(s)

	var t = 999
	t++
	t++
	fmt.Println(t)
}