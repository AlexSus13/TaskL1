package main
import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)//map в Go  это конкурентно небезопасный тип данных.

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{} //Мьютексы позволяют разграничить доступ к некоторым общим ресурсам, 
			    //только одна горутина имеет к ним доступ в определенный момент времени.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(counters map[int]int, i int, mu *sync.Mutex, wg *sync.WaitGroup) {
			for j := 0; j < 5; j++ {
				mu.Lock() //Метод для блокирования доступа к общему разделяемому ресурсу.
				counters[i * 10 + j]++
				mu.Unlock() //Метод для разблокирования доступа к общему разделяемому ресурсу. 
			}
			defer wg.Done()
		}(m, i, mu, wg)
	}
	wg.Wait()
	fmt.Println("result", m) //Может произойти ситуация, когда мы пытаемся вычитать то, 
				 //что другая горутина сейчас меняет. Но т.к. WaitGroup не блокируем.
				 //Используя -race при запуске, можно отловит состояние гонки
				 //go run -race 7main.go
}

