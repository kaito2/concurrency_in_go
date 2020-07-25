package main

import "fmt"

func main() {
	chanOwner := func() <-chan int { // 戻り値が読み込み専用のチャンネル
		// results チャンネルのスコープが
		// chanOwner に閉じているので外部から書き込みができない
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // 引数が読み込み専用チャンネル
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	// resultStream は読み込み専用なので
	// このスコープで悪いことができない
	resultStream := chanOwner()
	consumer(resultStream)
}
