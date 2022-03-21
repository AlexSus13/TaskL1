package main

import (
	"fmt"
	"math/big"
)

func main() {
	//Создаем большое числа из string с 10 онованием.
	var a, _ = new(big.Int).SetString("100000000000000000000000000", 10)
	//Создаем большое число с помощью функции big.NewInt.
	b := big.NewInt(int64(1 << 37))

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	//Методы для *Int
	sum := new(big.Int).Add(a, b)
	fmt.Println("sum = ", sum)

	difference := new(big.Int).Sub(a, b)
	fmt.Println("difference = ", difference)

	composition := new(big.Int).Mul(a, b)
	fmt.Println("composition = ", composition)

	division := new(big.Int).Div(a, b)
	fmt.Println("division = ", division)
}
