package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("This will be printed at the end") // akan dieksekusi setelah main selesai

	fmt.Println("Starting main function")

	// Contoh panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("Something went wrong!")                // akan menghentikan eksekusi program
	fmt.Println("This line will not be executed") // tidak akan dieksekusi
}
