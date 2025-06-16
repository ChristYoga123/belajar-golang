package main

import (
	"fmt"
	"time"
)

func goRoutine() {
	fmt.Println("go routine")
}

func main() {
	go goRoutine()
	fmt.Println("main function")
	// Wait for the goroutine to finish
	time.Sleep(1 * time.Second)
}
