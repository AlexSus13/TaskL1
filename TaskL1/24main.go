package main

import (
	"fmt"
	"math"
)

type CoordinatePlane struct {
	p1 Point
	p2 Point
}

type Point struct {
	x float64
	y float64
}

func main() {

	p1 := Point{x: 2, y: 2}
	p2 := Point{x: 4, y: 3}

	CoordP := &CoordinatePlane {
			p1: p1,
			p2: p2,
		}
	CoordP.distance()
}


func (cp *CoordinatePlane) distance() {

	distX := math.Abs(cp.p1.x - cp.p2.x) //math.Abs возвращает абсолютное значение т.е. модуль

	distY := math.Abs(cp.p1.y - cp.p2.y)

	distance := math.Sqrt(distX * distX + distY * distY) //math.Sqrt возвращает корень числа

	fmt.Println("distance = ", distance)
}
