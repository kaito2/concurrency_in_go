package main

func main() {
	loop:
		for {
			select {
			case <-done:
				break loop
			case val, ok = <- valueStream:
				if ok == false {
					return  // or `break loop`
				}
				// val に対する処理
			}
		}

	orDone := func(done, c <-chan interface{}) <- chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <- done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}
}
