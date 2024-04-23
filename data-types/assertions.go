//Type assertions is the ability to transform data types into the desired data type

package main

import "fmt"
func random() interface{}{
	return "OK"
}

func main(){
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)
	// resultInt := result.(int)
	// fmt.Println(resultInt)
	switch value := result.(type){
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}