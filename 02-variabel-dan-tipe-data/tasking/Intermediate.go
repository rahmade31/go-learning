package main

import "fmt"

func main() {

	// Tasking 1
	a := 5
	b := 10

	a, b = b, a

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// Tasking 2

	celsius := 30.0
	fahrenheit := celsius*9/5 + 32
	fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)
}
