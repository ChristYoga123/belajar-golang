package main

import "fmt"

func main() {
	var a int = 5 // 0101 dalam biner
	var b int = 3 // 0011 dalam biner

	fmt.Println("a & b =", a&b)   // AND: 0001 (1)
	fmt.Println("a | b =", a|b)   // OR: 0111 (7)
	fmt.Println("a ^ b =", a^b)   // XOR: 0110 (6)
	fmt.Println("a << 1 =", a<<1) // Shift kiri: 1010 (10)
	fmt.Println("a >> 1 =", a>>1) // Shift kanan: 0010 (2)
}
