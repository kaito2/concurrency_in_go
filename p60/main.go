package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get() // この呼出では instance は作成されない
}

/* output
$ go run p60/main.go
Creating new instance.
Creating new instance.
*/
