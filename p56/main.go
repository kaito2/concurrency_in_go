package main

import (
	"fmt"
	"sync"
)

type Button struct { //1
	Clicked *sync.Cond
}

func main() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) { //2
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			// ここで goroutine が呼び出されたかを確認している
			// => 本当は c.Wait() が呼び出されたかを確認したい
			goroutineRunning.Done()
			fmt.Println("goroutine start!")
			c.L.Lock()
			defer c.L.Unlock()
			// c.Wait() が呼ばれる前に button.Clicked.Broadcast()
			// が呼ばれると c.Wait() が永遠に待ってデッドロックする
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup //3
	clickRegistered.Add(3)

	subscribe(button.Clicked, func() { //4
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() { //5
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() { //6
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast() //7
	clickRegistered.Wait()
}
