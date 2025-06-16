package main

import "fmt"

func divide(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return // mengembalikan named return value
	}
	result = a / b
	return // mengembalikan named return value
}
func main() {
	res, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res) // Output: Result: 5
	}

	res, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: division by zero
	} else {
		fmt.Println("Result:", res)
	}
}
