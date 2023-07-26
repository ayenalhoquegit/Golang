package main

import "fmt"

func main() {

	name := "Jon"
	age := 30
	city := "New York"
	fmt.Printf("This is %s age : %d and he is from %s", name, age, city)
	const firstName = "Jon"
	const lastName = "Doe"
	const fullName = firstName + " " + lastName
	fmt.Print("\n", fullName)

	ID, roll := 30, 1
	fmt.Print("\n", ID, roll)

	const number = 10
	if number < 11 {
		fmt.Print("\n number is less than 11")
	}
}
