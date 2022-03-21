package main

import (
	"fmt"
)

func main() {

	str := "главрыба"

	method1(str)
}

func method1 (str string) {

        var newStr string //переменная для новой перевернутой строки
        var slStr []string //массив для символов строки

        for _, r := range str { //for-range для строки перебирает точки Unicode-(руны).
                slStr = append(slStr, string(r))
        }

        for i := len(slStr)-1; i >= 0; i-- { //проходимся по массиву символов в обратном порядке
                newStr += fmt.Sprintf("%s", slStr[i]) //конкатинция строк
        }

        fmt.Println("method1 = ", newStr)

}
