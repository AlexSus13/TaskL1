package main

import (
        "fmt"
        "context"
        "time"
	"math/rand"
	"sync"
)

func main() {
	//Контекст отменяется после дедлайна или вызова функции отмены.
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1 * time.Second))

        wg := &sync.WaitGroup{}

        dataChanel := make(chan int)
        defer close(dataChanel)

        start := time.Now()
        wg.Add(1)
        go worker(ctx, dataChanel, wg, start)

        for {
                select {
                case <-ctx.Done():
                        cancel() //func отмены контекста
                        wg.Wait()//ждем завершения горутин
                        duration := time.Since(start)
                        fmt.Println("TimeWorkM", duration) //время работы main
                        fmt.Println("StopMain")
                        return
                case dataChanel <- rand.Int():
                }
        }
}


/*
func main() {
	//После срабатывания таймера сработает канал таймера
	timer := time.NewTimer(1 * time.Second)
	//Контекст отменяется после вызова функции отмены.
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}

        dataChanel := make(chan int)
	defer close(dataChanel)

	start := time.Now()
	wg.Add(1)
        go worker(ctx, dataChanel, wg, start)

        for {
                select {
                case <- timer.C: //канал таймера
			cancel() //func отмены контекста
			wg.Wait()//ждем завершения горутин
			duration := time.Since(start)
			fmt.Println("TimeWorkM", duration) //время работы main
			fmt.Println("StopMain")
			return
                case dataChanel <- rand.Int():
                }
        }
}
*/
func worker(ctx context.Context, dataChanel chan int, wg *sync.WaitGroup, start time.Time) {

	defer wg.Done()

        for {
                select {
                case <-ctx.Done(): // сработает после завершения контекста или вызова cancel()
			duration := time.Since(start)
			fmt.Println("TimeWorkG", duration)
                        return
                case result, ok := <-dataChanel: //читаем из канала данные и проверяем закрыт или нет
			if ok { //ok = true открыт
				fmt.Printf("worker: %d\n", result)
			} else {
				duration := time.Since(start)
				fmt.Println("TimeWorkG", duration)
				return
			}
                }
        }
}

