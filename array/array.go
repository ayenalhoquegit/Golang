package main

import "fmt"

func main() {
	names := [3]string{"Jon", "Doe", "John"}
	fmt.Println(len(names))
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
