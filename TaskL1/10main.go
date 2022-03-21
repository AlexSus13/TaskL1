package main

import (
	"fmt"
)

func main() {

	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)

	for _, v := range temp {
		i := (int(v) / 10) * 10 //int(v) преобразует число float64 в int, 
				        //например: 32,9 будет 32 
		m[i] = append(m[i], v)
	}
	fmt.Println(m)
}
