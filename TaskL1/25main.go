package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Start")

	start1 := time.Now()
	Sleep1(3 * time.Second)
	fmt.Println("Sleep1", time.Since(start1))

	start2 := time.Now()
	Sleep2(3 * time.Second)
	fmt.Println("Sleep2", time.Since(start2))
}

func Sleep1(t time.Duration) {
	timer := time.NewTimer(t)
	<-timer.C //канал таймера
}

func Sleep2(t time.Duration) {
	<-time.After(t)
}

//time.After мы остановить не можем, и пока он не выполнится, даже если мы завершили функцию, он не освободит ресурсы.
//Канал таймера нужно освобождать, если есть шанс, что таймер не успеет сработать:
//if !timer.Stop() {
//	<-timer.C//освобождаем канал таймера
//}
