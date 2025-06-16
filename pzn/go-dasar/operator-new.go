package main

import "fmt"

func main() {
	p := new(int)                           // mengalokasikan memori untuk int dan mengembalikan pointer
	fmt.Println("Value of p (pointer):", p) // Output: Value of p (pointer): <alamat memori>
	fmt.Println("Value pointed by p:", *p)  // Output: Value pointed by p: 0 (nilai default int)

	*p = 42                                        // mengubah nilai yang ditunjuk oleh pointer
	fmt.Println("Updated value pointed by p:", *p) // Output: Updated value pointed by p: 42
}
