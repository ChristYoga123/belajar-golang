package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++ // mengubah nilai field count
}
func (c *Counter) GetCount() int {
	return c.count // mengembalikan nilai field count
}
func main() {
	c := &Counter{count: 0}                     // membuat pointer ke struct Counter.
	fmt.Println("Initial count:", c.GetCount()) // Output: Initial count: 0

	c.Increment()                                       // memanggil method Increment
	fmt.Println("Count after increment:", c.GetCount()) // Output: Count after increment: 1

	c.Increment()                                               // memanggil method Increment lagi
	fmt.Println("Count after another increment:", c.GetCount()) // Output: Count after another increment: 2
}
