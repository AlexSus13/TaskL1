package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//Контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())
	//Контекст с отменой по ТаймАуту
	ctxT, _ := context.WithTimeout(context.Background(), 1 * time.Second)
	//Канал отмены
	cancelChanel := make(chan struct{})

	go worker1(ctx)
	go worker2(cancelChanel)
	go worker3(ctxT)

	time.Sleep(1 * time.Second)
	cancel() //Отменяем выполнение горутины в которую передан ctx
	cancelChanel <- struct{}{} //Передаем в канал отмены пустую структуру, 
				   //чтобы сигнализировать об отмене выполнения
	close(cancelChanel) //Закрываем канал + это условие отмены работы worker4

	fmt.Println("Main Stop")
}

func worker3(ctx context.Context) {

	for {
		select {
		case <-ctx.Done(): //Метод Done() возвращает канал , который получает пустой struct{}
				   //каждый раз, когда контекст получает событие отмены(ТаймАут).
			fmt.Println("worker3 Stop")
			return
		default:
			fmt.Println("worker3 Ok")
		}
	}
}

func worker2(ch chan struct{}) {

	for {
		select {
		case <-ch: //Можно прочитать из канала, после того как в main`e cancelChanel <- struct{}{}
			fmt.Println("worker2 Stop")
			return
		default:
			fmt.Println("worker2 Ok")
		}
	}
}

func worker1(ctx context.Context) {

	for {
		select {
		case <-ctx.Done(): //Метод Done() возвращает канал , который получает пустой struct{}
				   //каждый раз, когда контекст получает событие отмены(вызов cancel()).
			fmt.Println("worker1 Stop")
			return
		default:
			fmt.Println("worker1 Ok")
		}
	}
}
