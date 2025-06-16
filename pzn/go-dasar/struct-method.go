package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height // method untuk menghitung luas
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height) // method untuk menghitung keliling
}
func main() {
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println("Area:", rect.Area())           // Output: Area: 50
	fmt.Println("Perimeter:", rect.Perimeter()) // Output: Perimeter: 30
}
