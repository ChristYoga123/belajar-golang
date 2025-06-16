package main

import "fmt"

func main() {
	//	Deklarasi variabel
	var a int
	a = 10
	fmt.Println(a)

	//	Deklarasi variabel dan inisialisasi
	var b int = 20
	fmt.Println(b)

	//	Deklarasi variabel dengan tipe data yang di infer
	var c = 30
	fmt.Println(c)

	//	Deklarasi variabel dengan tipe data yang di infer menggunakan :=
	d := 40
	fmt.Println(d)

	//	Deklarasi Multiple Variabel
	var e, f int = 50, 60
	fmt.Println(e, f)

	var (
		g = 70
		h = 80
	)
	fmt.Println(g, h)
}
