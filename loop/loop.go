package main

import "fmt"

func main() {
	// As simple for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop as while loop
	i := 0
	for i < 3 {
		i += 2
	}
	fmt.Print(i)

	// sample range in for loop
	rvariable := []string{"one", "two", "three", "four", "five"}
	for i, v := range rvariable {
		fmt.Println(i, v)
	}

	// using for loop for strings
	for i, j := range "abcd" {
		fmt.Println(i, j)
	}

}
