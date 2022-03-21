package main

import (
	"fmt"
)

func main() {
	var num int64 = 13
	var pos uint = 4
	fmt.Printf("num = %d in bytes = %08b\n", num, num)
	fmt.Println("bit position(1-8) = ", pos)

	result := bitReplace(num, pos)
	fmt.Printf("numAfter = %d in bytes = %08b\n", result, result)
}

func bitReplace(num int64, pos uint) int64 {

	pos-- //т.к. счет битов начинается с 0

	val := num & (1 << pos) //Сдвигаем бит у еденицы на выбранную позицию (1 << pos).
				//Оператор “И” & копирует бит в результат, если бит задан в обоих операндах.
				//Если на выбранной позиции у num - 0, то val = 0,
				//ели на выбранной позиции у num - 1, то val > 0.
	if val > 0 { //на месте выбранного бита в num 1
		num = num ^ (1 << pos) //Оператор “X-ИЛИ” ^ копирует бит в результат, 
					//если бит задан в одном из операндов, но не в обоих
	} else { //на месте выбранного бита в num 0
		num = num | (1 << pos) //Оператор “ИЛИ” | копирует бит в результат, 
				       //если бит задан в одном из операндов,
				       //следовательно, мы "копируем" этот бит в num
	}
	return num
}
