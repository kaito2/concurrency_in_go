package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// ちゃんと goroutine の外で Add する
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All goroutine complete.")
}

/* output
$ go run p48/main.go
1st goroutine sleeping...
2nd goroutine sleeping...
All goroutine complete.
*/