package main

import "fmt"

func add(a int, b int) int {
	sum := a + b
	return sum
}
func main() {
	sum := add(2, 3)
	fmt.Print(sum)
}
