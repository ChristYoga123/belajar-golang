package main

import "fmt"

func printValue(v interface{}) {
	fmt.Println("Value:", v) // mencetak nilai dari interface kosong
}
func main() {
	printValue(42)             // Output: Value: 42
	printValue("Hello")        // Output: Value: Hello
	printValue(3.14)           // Output: Value: 3.14
	printValue([]int{1, 2, 3}) // Output: Value: [1 2 3]
}
