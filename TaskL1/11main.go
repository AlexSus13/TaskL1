package main

import (
	"fmt"
)

func main() {

	set1 := []int{1, 2, 3, 4, 5, 6}
	set2 := []int{1, 20, 3, 40, 5, 60}

	fmt.Println(crossSet(set1, set2))
}

func crossSet(set1, set2 []int) []int {

	var crossSet []int

	m := make(map[int]int)

	for _, v := range set1 { //итерируемся по slice set1
		m[v]++		 //добавляем значения в мап,
				 //используя значения из slice как ключ
	}

	for _, v := range set2 { //итерируемся по slice set2
		if _, ok := m[v]; ok { //проверем существование ключа в мап
			crossSet = append(crossSet, v) //если уже есть добавляем в массив
		}
	}

	return crossSet
}
