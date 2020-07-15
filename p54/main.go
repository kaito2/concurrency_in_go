package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        //1
		queue = queue[1:] //9
		fmt.Println("Removed from queue")
		c.L.Unlock() //10
		c.Signal()   //11
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()            //3
		for len(queue) == 2 { //4
			fmt.Println("waiting...")
			// Wait => c.L.Unlock() して Signal() を受け取るまで sleep ...
			c.Wait() //5
			// time.Sleep(100 * time.Millisecond) // remove
		}
		// c.L.Lock() // remove
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) //6
		c.L.Unlock()                        //7
	}
}

/* output
$ go run p54/main.go
Adding to queue
Adding to queue
waiting...
Removed from queue
Adding to queue
waiting...
Removed from queue
Adding to queue
waiting...
Removed from queue
Adding to queue
waiting...
Removed from queue
Adding to queue
waiting...
Removed from queue
Removed from queue
Adding to queue
Adding to queue
waiting...
Removed from queue
Adding to queue
waiting...
Removed from queue
Adding to queue
*/
