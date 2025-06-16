package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	// Menyimpan fungsi dalam variabel
	sum := add                     // sum sekarang adalah variabel yang menyimpan fungsi add. Add tidak diberi tanda kurung karena kita hanya menyimpan referensinya, bukan memanggilnya.
	result := sum(3, 4)            // memanggil fungsi melalui variabel
	fmt.Println("Result:", result) // Output: Result: 7

	// Menggunakan fungsi sebagai argumen
	applyFunction := func(f func(int, int) int, x, y int) int {
		return f(x, y)
	}
	fmt.Println("Apply Function Result:", applyFunction(add, 5, 6)) // Output: Apply Function Result: 11
}
