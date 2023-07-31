package main

import "fmt"

// * Operator =  dereferencing operator
// & operator = returns the address of a variable

func modifyPointer(p *int) {
	*p = 20
}

func main() {
	x := 10
	p := &x
	fmt.Println("Value of x before modify:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value stored in pointer P ", p)
	fmt.Println("value store in x(*p)", *p)
	modifyPointer(p)
	fmt.Println("Value of x after pointer modify:", x)
}
