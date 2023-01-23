// Реализовать пересечение двух неупорядоченных множеств.
package main

import (
	"fmt"
	"math"
)

// если бы все элементы были бы уникальными - можно было бы просто перебрать 2 массива, но...
func intersect(firstArr []int, secondArr []int) []int {

	m1 := make(map[int]int) // создаём 2 мапы для вычисления количества уникальны элементов в каждом массиве
	m2 := make(map[int]int)
	resArr := []int{}

	for _, i := range firstArr { // заполняем мапы
		m1[i] += 1
	}
	for _, j := range secondArr {
		m2[j] += 1
	}

	for i := range m1 {
		for j := range m2 { // ищем совпадения элементов
			if i == j {
				for k := 1.0; k <= math.Min(float64(m1[i]), float64(m2[i])); k++ { // заполняем результирующий массив минимальным количеством
					resArr = append(resArr, i) // одинаковых элементов в каждом массиве
				}
			}
		}
	}

	return resArr
}

func main() {

	firstArr := []int{1, 3, 4, 5, 5}
	secondArr := []int{3, 4, 5, 6, 7, 8}

	fmt.Println("Пересечение множеств: ", intersect(firstArr, secondArr))

	firstArr = []int{1, 3, 4, 5, 5}
	secondArr = []int{3, 4, 5, 6, 7, 8, 5}

	fmt.Println("Пересечение множеств: ", intersect(firstArr, secondArr))
}
