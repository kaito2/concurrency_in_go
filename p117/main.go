package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main () {
	// fan in fan out してない ver
	randFn := func() interface{} { return rand.Intn(50000000)}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, randFn))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}

func main2() {
	// fan in fan out した ver

	// *** fan in (マルチプレキシング) ***
	// 複数のストリームを単一のストリームにまとめる
	fanIn := func(
	// done チャンネルは嗜み
		done <-chan interface{},
		channels ...<-chan interface{},
	) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		// すべてのチャンネルが close するまで待ち合わせる
		wg.Add(len(channels))
		// すべてのチャンネルから select する
		for _, c := range channels {
			go multiplex(c)
		}

		// すべてのチャンネルが close したら返すチャネルを close する。（後始末）
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	randFn := func() interface{} { return rand.Intn(50000000)}
	done := make(chan interface{})
	defer close(done)

	start := time.Now()
	randIntStream := toInt(done, repeatFn(done, randFn))

	// *** fan out ***
	// 驚くほど簡単(らしい)
	// => 同じチャンネルを複数の goroutine に渡す
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	fmt.Println("Primes:")
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}
