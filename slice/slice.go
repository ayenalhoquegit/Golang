package main

import "fmt"

func main() {
	array := [4]int{1, 2, 3, 4}
	slice := array[0:3]
	fmt.Println("Array ", array)
	fmt.Println("Slice ", slice)

	newSlice := []int{1, 2, 3}
	newSlice = append(newSlice, 4, 5)
	fmt.Println("New slice", newSlice)
}
