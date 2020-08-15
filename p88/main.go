package main

import (
	"fmt"
)

func main() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- i
		}
	}

	handleData := make(chan int, 1)
	// ...
	// XXX: loopData の外でも handleData に値を書き込める
	// MEMO: handleData のキャパシティが1以上である必要あり
	// handleData <- 10
	// ...
	go loopData(handleData)

	// ...
	// XXX: 後から data  を書き換えられる
	// data = make([]int, 100)
	// ...

	for num := range handleData {
		fmt.Println(num)
	}
}
