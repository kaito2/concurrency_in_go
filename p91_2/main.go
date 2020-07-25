package main

import (
	"fmt"
	"time"
)

func printUntilDone(done <-chan interface{}) {
	for {
		select {
		case <-done:
			return
		default: // Do nothing.
		}
		// なんかの処理(割り込みはできない)
		fmt.Println("Sleeping...")
		time.Sleep(500 * time.Millisecond)
	}
}

func printUntilDone2(done <-chan interface{}) {
	for {
		select {
		case <-done:
			return
		default:
			// なんかの処理(割り込みはできない)
			// ...
		}
	}
}

func main() {
	done := make(chan interface{})
	go printUntilDone(done)
	time.Sleep(3 * time.Second)
	close(done)
}
