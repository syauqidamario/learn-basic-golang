package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Subaru",
		"address": "Kyoto",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	series := make(map[string]string)
	series["title"] = "Kamen Rider Gotchard"
	series["producer"] = "Yousuke Minato"
	series["writer"] = "Keiichi Hasegawa"
	series["wrong"] = "Oops"

	delete(series, "wrong")
	fmt.Println(series)
}