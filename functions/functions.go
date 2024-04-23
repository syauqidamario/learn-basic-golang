package main

import "fmt"

func sayHi(firstName string, lastName string){
	fmt.Println("Hello there", firstName, lastName)
}

func getHello(name string) string{
	return "Hello " + name
}

func getFullName()(string, string){
	return "Eko", "Khannedy"
}

func getCompleteName()(firstName, middleName, lastName string){
	firstName = "Kaguya"
	middleName = "Quartz"
	lastName = "Houou"
	return firstName, middleName, lastName
}

func sumAll(numbers ...int)int{
	total := 0
	for _, number := range numbers{
		total += number
	}
	return total
}

func getGoodBye(name string)string{
	return "Goob Bye " + name
}

//function as parameter

type Filter func(string) string
func sayHelloWithFilter(name string, filter func(string) string){
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string)string{
	if name == "You bitch"{
		return "..."
	}else{
		return name
	}
}

type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You're blocked", name)
	}else{
		fmt.Println("Welcome", name)
	}
}

//Recursive function
func factorialRecursive(value int) int{
	if value == 1{
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

//Closure is an ability for a function to interact with its surrounding data

func main(){
	counter := 0
	increment := func(){
		fmt.Println("Increment")
		counter++
	}
	increment()
	increment()
	fmt.Println(counter)
	sayHi("General", "Kenobi")
	result := getHello("General Kenobi")
	firstName, lastName := getFullName()
	fmt.Println(result)
	fmt.Println(firstName, lastName)
	// firstName, _ := getFullName()
	fmt.Println(firstName)
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
	total := sumAll(100, 200, 300, 400, 500, 600)
	fmt.Println(total)
	//slice parameter
	// numbers :=[]int{100, 200, 300, 400, 500, 600}
	// total = sumAll(numbers...)
	fmt.Println()

	goodbye := getGoodBye
	fmt.Println(goodbye("Syauqi"))
	//functions as parameter
	sayHelloWithFilter("Syauqi", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("You bitch", filter)
	//anonymous function
	blacklist := func(name string)bool{
		return name=="babi"
	}
	registerUser("Spontan", blacklist)
	registerUser("babi", func(name string)bool{
		return name == "babi"
	})
	//Recursive function example
	fmt.Println(factorialRecursive(100))
}