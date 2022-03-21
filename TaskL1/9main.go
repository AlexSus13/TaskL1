package main

import (
	"fmt"
	"sync"
)

type job func(wg *sync.WaitGroup, in, out chan int)

func main() {

	jobs := []job{ //slice функций обработчиков
		job(data),
		job(nums),
		job(sqrNums),
	}
	pipeline(jobs...) //запускаем конвеер
}

func pipeline(jobs ...job) {

	wg := &sync.WaitGroup{}

	chNum := make(chan int) //канал для передачи чисел
	chSqrNum := make(chan int) //канал для передачи квадратов чисел

	for i := 0; i < len(jobs); i++ {
		wg.Add(1)
		go jobs[i](wg, chNum, chSqrNum)//итерируемся по slice функций обработчиков
					       //и запускаем их в горутине
	}
	wg.Wait()//ждем выполнения всех горутин
}

func data(wg *sync.WaitGroup, chNum, chSqrNum chan int) {

	data := []int{2, 4, 6, 8, 10, 2, 4, 6, 8, 10, 2, 4, 6, 8, 10, 2, 4, 6, 8, 10}
	for _, n := range data {
		chNum <- n //итерируемся по slice с данными
	}
	defer close(chNum) //отправитель закрывае канал для передачи чисел
	defer wg.Done()
}

func nums(wg *sync.WaitGroup, chNum, chSqrNum chan int) {

	for n := range chNum { //range получает значения из канала до тех пор, пока он не закрыт
		chSqrNum <- n * n
	}
	defer close(chSqrNum) //отправитель закрывае канал для передачи квадратов чисел
	defer wg.Done()
}

func sqrNums(wg *sync.WaitGroup, chNum, chSqrNum chan int) {

	for result := range chSqrNum {
		fmt.Println("SqrResult = ", result)
	}
	defer wg.Done()
}

