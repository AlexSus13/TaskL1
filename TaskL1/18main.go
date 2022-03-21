package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int32 = 0

func main() {

	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter(wg)
	}

	wg.Wait()
	fmt.Println("count = ", count)
}

func counter(wg *sync.WaitGroup) {
	atomic.AddInt32(&count, 1) //AddInt32 атомарно добавляет delta=1 к *count и возвращает новое значение.
				  //atomic помогает избегать состояние гонки.
	defer wg.Done()
}
