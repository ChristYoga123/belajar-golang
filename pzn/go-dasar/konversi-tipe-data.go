package main

import "fmt"

func main() {
	//	Konversi Tipe Data
	/*Integer*/
	//	Widening Conversion
	var a int = 10
	var b int64 = int64(a) // Konversi dari int ke int64
	fmt.Println(b)
	//	Narrowing Conversion
	var c int64 = 129
	var d = int8(c) // Konversi dari int64 ke int
	fmt.Println(d)  // Output: -127 (karena int8 hanya bisa menyimpan nilai dari -128 sampai 127)
}
