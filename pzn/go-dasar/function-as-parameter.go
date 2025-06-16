package main

import "fmt"

func applyFunction(f func(int, int) int, x, y int) int {
	return f(x, y)
}
func add(a, b int) int {
	return a + b
}
func main() {
	result := applyFunction(add, 3, 5) // meneruskan fungsi add sebagai parameter
	fmt.Println("Result:", result)     // Output: Result: 8

	// Meneruskan fungsi anonim sebagai parameter
	multiply := func(a, b int) int {
		return a * b
	}
	result2 := applyFunction(multiply, 4, 6)
	fmt.Println("Multiply Result:", result2) // Output: Multiply Result: 24
}
