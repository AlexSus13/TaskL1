package main

import (
        "fmt"
	"strings"
)

func main() {

        str := "snow dog sun"

	method1(str)
}

func method1(str string) {

	var newStr string

	words := strings.Fields(str)//Разбивает исходную строку "по пробелам" на массив строк

	for i := len(words) - 1; i >= 0; i-- {
		if i != 0 {
			newStr += fmt.Sprintf("%s", words[i]) + " "
		} else {
			newStr += fmt.Sprintf("%s", words[i])
		}
	}
	fmt.Println(newStr)
}

