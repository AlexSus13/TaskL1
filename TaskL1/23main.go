package main

import (
	"fmt"
)

func main() {

	var NumMethod int

	sl := []int{2, 4, 6, 8, 10}
	var i int

	fmt.Println(sl)

	begin := 0
	end := len(sl) - 1
	fmt.Printf("%v <= i <= %v\n", begin, end)
	fmt.Println("Введите индекс удаляемого элемента i = ")

	fmt.Scan(&i)

	if begin <= i && i <= end {
		NumMethod, sl = method1(sl, i) //Не зменяет исходный sl в main если нет return
		//NumMethod = method2(sl, i) //Изменяет исходный sl в main
		//NumMethod = method3(sl, i) //Изменяет исходный sl в main
		//NumMethod = method4(sl, i) //Изменяет исходный sl в main
		sl = sl[:len(sl)-1] //для method 2, 3, 4
	} else {
		fmt.Println("Нет элемента с таким индексом")
		return
	}
	fmt.Printf("method%d %v\n", NumMethod, sl)
	//Если вы просто присвоите существующий срез новой переменной, срез не будет дублироваться.
	//Обе переменные будут ссылаться на один и тот же слайс, поэтому любые изменения значения слайса
	//будут отражены во всех его ссылках. Это также верно при передаче слайса в функцию, поскольку в Go слайсы передаются по ссылке.
}

func method4(sl []int, i int) int { //method2 изменяет исходный sl в func main, т.к.
	//усеченный срез находитя в том же участке памяти!!!.
	//элементы в исходном sl заменятся на скопиррованные из sl[i+1:] начиная с i.
	copy(sl[i:], sl[i+1:]) //Копируем значения

	sl = sl[:len(sl)-1] //Усекаем срез(для правильного вывода в пределах func).
	return 4
}

func method3(sl []int, i int) int { //method3 изменяет исходный sl в func main, т.к.
	//method3 меняет порядок, но он быстрее
	sl[i] = sl[len(sl)-1] //Копируем последний элемент в индекс i.

	return 3
}

func method2(sl []int, i int) int { //method2 изменяет исходный sl в func main, т.к.
	//усеченный срез находитя в том же участке памяти!!!
	//элементы в исходном sl заменятся на добавленные начиная с i.
	//i := 2; было в main [2 4 6 8 10] стало [2 4 8 10 10]
	//Если при добавлении элементов привысили cap(sl), то
	//создастся новый срез, и он будет указат на дугую обл. амяти.
	sl = append(sl[:i], sl[i+1:]...)
	return 2
}

func method1(sl []int, i int) (int, []int) { //method1 не изменяет исходный sl в func main если нет return, т.к.
	//мы создали новый slice и добавили в него данные
	//при помощи append
	newSl := append([]int(nil), sl[:i]...) //Элементы sl добавляются к нулевому срезу, а результирующий срез присваивается newSl.
	newSl = append(newSl, sl[i+1:]...)     //Усеченный срез указывает на тот же участок памяти!!!

	return 1, newSl
}
