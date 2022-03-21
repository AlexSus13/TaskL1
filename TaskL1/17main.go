package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	//для бинарного поиска только отсортированные данные
	fmt.Println(binarySearch(slice, 15))
}

func binarySearch(slice []int, item int) int {

	start := 0        //позиция 1-го элемента
	end := len(slice) //позиция последнего элемента
	var middle int    //позиция среднего элемента
	var position int  //позиция найденного элемента
	found := false    //флаг, нашли или нет

	for found == false && start < end {

		middle = ((start + end) / 2) //Результат деления округляется в меньшую сторону.

		if slice[middle] == item { //Если искомый элемент центральный, то...
			found = true //true => нашли
			position = middle //позиция равна позиции среднего элемента
			return position
		}
		if item < slice[middle] {  //Если искомый элемент меньше центрального, то...
			end = middle - 1   //дальше рассматриваем левую от середины часть массива 
		} else {		   //Если искомый элемент больше центрального, то...
			start = middle + 1 //дальше рассматриваем правую от середины часть массива
		}
		fmt.Println("start = ", start)
		fmt.Println("middle = ", middle)
		fmt.Println("end = ", end)
		fmt.Println()
	}//крутимся в цикле пока не найдем элемент или массив не будет размерности 1

	if found == false {
		fmt.Println("Not found")
		return -1
	} else {
		return position
	}
}

