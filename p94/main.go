package main

import (
	"fmt"
	"time"
)

func main() {
	// 慣習的に `done` という名前が使われる
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			// 前節で説明した for-select
			for {
				select {
				case s := <-strings:
					// おもろい処理
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		// 1秒後似操作をキャンセルする
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork groutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}
