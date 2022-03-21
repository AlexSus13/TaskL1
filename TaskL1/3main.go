package main

import (
        "fmt"
)

func main() {

	nums := []int{2, 4, 6, 8, 10}

	sum := SumSqr(nums)

	fmt.Println("Sum = ", sum)
}

func SumSqr(nums []int) int {

	if len(nums) < 1 {
		return 0
	}

	var sum int

	ch := make(chan int, len(nums))

	go func(ch chan int) {

		for _, v := range nums {
			ch <- v * v
		}
		close(ch)
	}(ch)

	for {
		sqr, ok := <-ch    //Проверка на закрытие канала
		if ok {            //ok = true канал открыт
			sum += sqr
		} else {           //канал закрыт
			break
		}
	}

	return sum
}
