package main

import "fmt"

func main() {
	var i interface{} = "Hello, World!" // interface kosong yang berisi string

	// Type assertion untuk mengonversi ke string
	str, ok := i.(string)
	if ok {
		fmt.Println("String value:", str) // Output: String value: Hello, World!
	} else {
		fmt.Println("Type assertion failed")
	}

	// Type assertion untuk mengonversi ke int (akan gagal)
	num, ok := i.(int)
	if ok {
		fmt.Println("Integer value:", num)
	} else {
		fmt.Println("Type assertion failed for int") // Output: Type assertion failed for int
	}

	// Type Assertion dengan switch
	/*
			var i interface{} = "Hello, World!" // interface kosong yang berisi string

		    switch v := i.(type) {
		    case string:
		        fmt.Println("String value:", v) // Output: String value: Hello, World!
		    case int:
		        fmt.Println("Integer value:", v)
		    default:
		        fmt.Println("Unknown type")
		    }
	*/

	//	Contoh yang salah
	/*
		var i interface{} = "Hello, World!" // interface kosong yang berisi string

		// Type assertion yang salah (akan panic jika tidak menggunakan ok)
		str := i.(string) // ini akan berhasil
		fmt.Println("String value:", str)

		// Type assertion yang salah (akan panic)
		num := i.(int)                     // ini akan panic karena i bukan int
		fmt.Println("Integer value:", num) // tidak akan dieksekusi
	*/

}
