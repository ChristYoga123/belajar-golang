package main

import "fmt"

func increment(x *int) {
	*x++ // mengubah nilai yang ditunjuk oleh pointer. Fungsi ini akan menambah nilai yang ditunjuk oleh x sebesar 1.
}
func main() {
	var num int = 10
	fmt.Println("Value before increment:", num) // Output: Value before increment: 10

	increment(&num)                            // meneruskan alamat memori dari num
	fmt.Println("Value after increment:", num) // Output: Value after increment: 11

	// Menggunakan pointer untuk mengalokasikan memori
	p := new(int) // mengalokasikan memori untuk int dan mengembalikan pointer. Nilai awal dari p adalah 0 (nilai default untuk int).
	*p = 20
	fmt.Println("Value pointed by p:", *p) // Output: Value pointed by p: 20
}
