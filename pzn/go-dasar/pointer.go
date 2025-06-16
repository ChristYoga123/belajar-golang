package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x // mendeklarasikan pointer yang menyimpan alamat x

	fmt.Println("Nilai dari x:", x)                        // Output: Value of x: 10
	fmt.Println("Alamat dari x:", &x)                      // Output: Address of x: <alamat memori>
	fmt.Println("Pointer p (merupakan alamat dari x):", p) // Output: Pointer p: <alamat memori>
	fmt.Println("Nilai yang ditunjuk oleh pointer p:", *p) // Output: Value pointed by p: 10

	*p = 20                                                               // mengubah nilai yang ditunjuk oleh pointer
	fmt.Println("Nilai baru dari x setelah diubah melalui pointer p:", x) // Output: New value of x after change through pointer p: 20
}
