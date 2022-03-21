package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

var justString string

func createHugeString(size int) string {

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	var hugeStr string
	for i := 0; i <= size; i++ {
		if i%2 == 0 {
			//Возвращает случайный символ между 'a' и 'z'
			hugeStr += string('a' + rune(generator.Intn('z'-'a'+1)))
		} else {
			//Генерирует символ из произвольного набора, выберает случайный индекс из среза символов:
			chars := []rune("∑♥⌘")
			hugeStr += string(chars[rand.Intn(len(chars))])
		}

	}

	return hugeStr
}

func someFunc() {
	v := createHugeString(1 << 10)

	justString = v[:100]                                  //- получение подстроки, в байтах, не в символах.
	numOfjustString := utf8.RuneCountInString(justString) //кол-во символов
	fmt.Println("numOfjustString = ", numOfjustString)

	justString1 := string([]rune(v)[:100]) //Усеченный срез.
	numOfjustString1 := utf8.RuneCountInString(justString1)
	fmt.Println("numOfjustString1 = ", numOfjustString1)

	//Повторная нарезка среза не делает копию базового массива.
	//Полный массив будет храниться в памяти до тех пор,
	//пока на него больше не будут ссылаться.

	//==> Создадим новый слайс.
	justString3 := string(append([]rune{}, []rune(v)[:100]...))
	fmt.Println("justString3 = ", justString3)
}

func main() {
	someFunc()
}
