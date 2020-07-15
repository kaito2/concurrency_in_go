package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	time.Sleep(1 * time.Second)
	fmt.Printf("Count is %d\n", count)
}

/* output
$ go run p57/main.go
Count is 1
*/
