package main

import "fmt"

type Shape interface {
	Area() float64      // mendefinisikan method Area
	Perimeter() float64 // mendefinisikan method Perimeter
}
type SegiEmpat struct {
	Width  float64
	Height float64
}

func (r SegiEmpat) Area() float64 {
	return r.Width * r.Height // implementasi method Area
}
func (r SegiEmpat) Perimeter() float64 {
	return 2 * (r.Width + r.Height) // implementasi method Perimeter
}
func printShapeInfo(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
func main() {
	rect := SegiEmpat{Width: 5, Height: 10}
	printShapeInfo(rect) // memanggil fungsi dengan tipe Shape
}
