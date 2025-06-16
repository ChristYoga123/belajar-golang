package main

import "fmt"

type Operation func(int, int) int // mendeklarasikan tipe fungsi Operation
func add(a, b int) int {
	return a + b
}
func main() {
	var op Operation               // mendeklarasikan variabel dengan tipe Operation
	op = add                       // mengassign fungsi add ke variabel op
	result := op(3, 4)             // memanggil fungsi melalui variabel
	fmt.Println("Result:", result) // Output: Result: 7

	// Menggunakan tipe fungsi sebagai parameter
	applyOperation := func(f Operation, x, y int) int {
		return f(x, y)
	}
	fmt.Println("Apply Operation Result:", applyOperation(op, 5, 6)) // Output: Apply Operation Result: 11
}
