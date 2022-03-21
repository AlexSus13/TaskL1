package main

import (
        "os/signal"
        "math/rand"
        "context"
        "syscall"
        "time"
        "sync"
        "fmt"
)

func main() {

        var n int
        fmt.Print("Введите кол-во Воркеров: ")
        fmt.Scan(&n)

        dataChanel := make(chan interface{})
        defer close(dataChanel)

	ctxN, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

        ctx, cancel := context.WithCancel(context.Background())

        wg := &sync.WaitGroup{}

        for i := 1; i <= n; i++ {
                wg.Add(1)
                go worker(ctx, wg, dataChanel, i)
        }

        generator := rand.New(rand.NewSource(time.Now().UnixNano()))

        for {
                select {
                case <-ctxN.Done():
                        cancel()
                        wg.Wait()
                        fmt.Println("Stop")
			stop()
                        return
                default:
                        dataChanel <- generator.Int63()
                }
        }
}

func worker(ctx context.Context, wg *sync.WaitGroup, dataChanel chan interface{}, workerNum int) {

        defer wg.Done()

        for {
                select {
                case <-ctx.Done():
                        fmt.Printf("worker%d STOP\n", workerNum)
                        return
                case result, ok := <-dataChanel:
                        if ok {
                                fmt.Printf("worker%d: %v\n", workerNum, result)
                                time.Sleep(150 * time.Millisecond)
                        } else {
                                return
                        }
                }
        }
}

