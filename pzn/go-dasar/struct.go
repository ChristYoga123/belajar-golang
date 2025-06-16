package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// Inisialisasi struct
	p := Person{Name: "Alice", Age: 30, Address: "123 Main St"}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Address:", p.Address)

	// Mengakses field struct
	p.Age = 31 // mengubah nilai field Age
	fmt.Println("Updated Age:", p.Age)
}
