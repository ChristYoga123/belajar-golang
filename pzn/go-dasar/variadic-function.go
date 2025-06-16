package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// Menggunakan parameter dengan jumlah variadic
	fmt.Println("Sum:", sum(1, 2, 3))       // Output: Sum: 6
	fmt.Println("Sum:", sum(4, 5, 6, 7, 8)) // Output: Sum: 30

	// Menggunakan slice sebagai argumen variadic function
	nums := []int{10, 20, 30}
	fmt.Println("Sum with slice:", sum(nums...)) // Output: Sum with slice: 60
}
