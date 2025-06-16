package main

import (
	"fmt"
	"time"
)

func displayGoroutine(i int) {
	fmt.Println("goroutine", i)
}

func main() {
	for i := 0; i < 10000; i++ {
		go displayGoroutine(i)
	}

	time.Sleep(10 * time.Second) // Wait for goroutines to finish

}
