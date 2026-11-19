package main

import "fmt"

func main() {

	// Tasking 1
	a := 5
	b := 10

	a, b = b, a

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}
