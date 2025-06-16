package main

import "fmt"

func main() {
	// Deklarasi slice
	var numbers []int                     // slice kosong
	numbers = append(numbers, 10, 20, 30) // menambahkan elemen
	fmt.Println("Numbers:", numbers)
	// Inisialisasi slice dengan elemen
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruits:", fruits)
	// Slice dari array
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:4] // slice dari indeks 1 hingga 3
	fmt.Println("Slice from array:", s)
	// Menggunakan fungsi make
	s2 := make([]int, 5) // slice dengan panjang 5
	fmt.Println("Slice with make:", s2)
	// Menambahkan elemen ke slice
	s2 = append(s2, 6, 7)
	fmt.Println("Slice after append:", s2)
	// Menggunakan copy
	s3 := make([]int, len(s2))
	copy(s3, s2) // menyalin elemen dari s2 ke s3
	fmt.Println("Copied slice:", s3)
	// Menghapus elemen (tidak ada fungsi delete, menggunakan slicing)
	s2 = append(s2[:1], s2[2:]...) // menghapus elemen kedua
	fmt.Println("Slice after delete:", s2)
	// Menggunakan slice multidimensi
	matrix := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Matrix slice:", matrix)
}
