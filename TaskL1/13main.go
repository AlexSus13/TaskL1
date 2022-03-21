package main

import (
	"fmt"
)

func main() {
	n1 := 1
	n2 := 2

	fmt.Println("before swap", n1, n2)

	n1, n2 = n2, n1 //swap

	fmt.Println("after swap", n1, n2)
}
