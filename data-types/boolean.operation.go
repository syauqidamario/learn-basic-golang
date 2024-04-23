package main

import "fmt"

func main() {
	var finalGrade = 98
	var attendance = 95

	var finalPassGrade bool = finalGrade > 98
	var attendancePass bool = attendance > 95

	var pass bool = finalPassGrade && attendancePass

	fmt.Println(pass)
}