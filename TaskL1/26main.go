package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("result:", checkUniq("aAbcd"))
}

func checkUniq(str string) bool {

	m := make(map[string]int)

	s := strings.ToLower(str) //Т.к. функция проверки регистронезависима то,
				  //делаем все символы строчными. 

	for _, r := range s {
		m[string(r)]++
	}

	for _, v := range m {
		if v != 1 {
			return false
		}
	}

	return true
}
