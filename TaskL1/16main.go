package main

import "fmt"

func main() {
	nums := []int{8, 1, 2, 2, 3, 7, 0, 4, 32, 17, 21, 5, -6}
	qSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
func qSort(nums []int, start int, end int) {
	if len(nums) <= 1 {
		return
	}
	left := start //Ставим левый "указатель"-(индекс элемента массива) на начало массива().
	right := end //Ставим правый "указатель" на конец массива.
	pivot := nums[(left+right)/2] //Берем опорную точку - (элемент из массива).
	for left <= right {
		for ; nums[left] < pivot; left++ { //Двигаем левый "указатель" вправо до тех пор, пока
		}				   //не найдем элемент больше опорного.
		for ; nums[right] > pivot; right-- { //Двигаем правый "указатель" влево до тех пор, пока
		}				     //не найдем элемент меньше опорного.
		if left <= right { //Если правый "указатель" справа от левого, то... 
			nums[left], nums[right] = nums[right], nums[left] //меняем элементы местами
			left++ //Сдвигаем левый "указатель" вправо
			right--//Сдвигаем правый "указатель" влево
		}
	}//продолжаем цикл, пока правый "указатель" не окажется слева от левого "указателя" 

	if left < end { //рекурсивно вызываем функцию
		qSort(nums, left, end)
	}
	if right > start { //рекурсивно вызываем функцию
		qSort(nums, start, right)
	}
	//
}
