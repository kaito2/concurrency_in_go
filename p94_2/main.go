package main

import (
	"fmt"
	"math/rand"
)

func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	// "newRandStream closure exited." が表示されないので、
	// newRandStream() で作られたゴルーチンは残り続ける。
}

/* output
$ go run p94_2/main.go
3 random ints:
1: 5577006791947779410
2: 8674665223082153551
3: 6129484611666145821
*/
