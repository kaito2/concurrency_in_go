package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// おもろい処理
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	// 何かしらの処理
	fmt.Println("Done.")
}
