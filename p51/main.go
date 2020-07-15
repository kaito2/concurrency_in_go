package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwmutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwmutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}

/* ouptut
$ go run p51/main.go                                                                                                                                                              (git)-[master]
Readers  RWMutex       Mutex
1        34.869µs      7.045µs
2        11.24µs       32.771µs
4        7.641µs       5.481µs
8        9.661µs       7.223µs
16       35µs          26.44µs
32       34.244µs      47.911µs
64       88.778µs      35.348µs
128      52.257µs      53.719µs
256      123.802µs     164.445µs
512      224.542µs     232.432µs
1024     490.117µs     342.715µs
2048     700.828µs     740.074µs
4096     1.365087ms    1.422ms
8192     2.652015ms    2.618513ms
16384    4.988283ms    4.943878ms
32768    9.130666ms    8.611712ms
65536    17.108441ms   14.675342ms
131072   33.937963ms   38.177573ms
262144   65.707289ms   59.761357ms
524288   133.580727ms  119.810864ms
*/
