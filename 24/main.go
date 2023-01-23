//Разработать программу нахождения расстояния между двумя точками, которые представлены
//в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package main

import (
	"fmt"
	"math"
)

// структура точки
type Point struct {
	x float64
	y float64
}

// конструктор, создаёт точку
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// возвращает расстояние между точками a, b
func Distance(a, b Point) float64 {
	x := a.x - b.x
	y := a.y - b.y
	return math.Sqrt(x*x + y*y)
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(1.414, 1.414)

	fmt.Println(Distance(a, b))

}
