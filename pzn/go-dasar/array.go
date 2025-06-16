package main

import "fmt"

func main() {
	// Deklarasi array
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// Inisialisasi array langsung
	fruits := [3]string{"Apple", "Banana", "Cherry"}

	// Inisialisasi array dengan panjang otomatis
	vegetables := [...]string{"Carrot", "Potato", "Tomato"}

	// Array multidimensi
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// Mengakses elemen array
	fmt.Println("First number:", numbers[0])
	fmt.Println("Second fruit:", fruits[1])
	fmt.Println("Third vegetable:", vegetables[2])
	fmt.Println("Matrix element:", matrix[1][2]) // Output: 6

	// Menggunakan loop untuk iterasi
	for i, v := range numbers {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}
}
