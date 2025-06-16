package main

import (
	"fmt"
	packageinitialization "go-dasar/package-initialization"
)

func main() {
	// Menggunakan package-initialization untuk mendapatkan koneksi database
	connection := packageinitialization.GetDatabase()
	fmt.Println(connection)
}
