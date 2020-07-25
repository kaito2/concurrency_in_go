package main

import "fmt"

func getStringStream(done <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		defer fmt.Println("wrting to stringStream is done")
		// ↓ が本文中に記載されていたコード
		for _, s := range []string{"a", "b", "c"} {
			select {
			case <-done:
				return
			case stringStream <- s:
			}
		}
		// ↑
	}()
	return stringStream
}

func main() {
	done := make(chan interface{})
	// 使い所が難しいが、受け取るに停止の権限を与えたい場合に使うのかな
	// go close(done)
	for s := range getStringStream(done) {
		fmt.Println(s)
	}
}
