package main

import "fmt"

type Address struct {
	Street, City, PostalCode string
}

type Person struct {
	FirstName, LastName string
	Age                 int
	Address             Address
}

func main() {
	p := Person{FirstName: "Jon", LastName: "Doe", Age: 30, Address: Address{"123 Main Street", "New York", "12345"}}
	fmt.Println(p.FirstName, p.LastName)
	fmt.Println("Address:", p.Address.Street, p.Address.City, p.Address.PostalCode)
	//fmt.Printf("Person name %s , city %s  and age %d ", p.name, p.city, p.age)
}
