// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
// с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var sum1, sum2, sum3 int

	// можно использовать предыдущую программу, добавив суммирование
	// с помощью группы ожидания
	var wg sync.WaitGroup
	wg.Add(len(numbers)) // заполняем счётчик горутин

	for _, number := range numbers {
		go func(n int) { // анонимная горутина
			defer wg.Done() // уменьшаем счётчик горутин
			sum1 += Square(n)
		}(number)
	}
	wg.Wait() // блокировка потока

	fmt.Println(sum1)

	// простой вариант с задержкой временем
	for i := 0; i < 5; i++ {
		go func(i int) {
			sum2 += Square(numbers[i])
		}(i)
	}
	time.Sleep(5 * time.Millisecond)

	fmt.Println(sum2)

	// а можно использовать блокировку горутин каналами

	ch := make(chan int)
	defer close(ch)

	for _, number := range numbers {
		go func(n int) {
			ch <- n * n

		}(number)
	}

	for i := 0; i < 5; i++ {
		sum3 += <-ch
	}

	fmt.Println(sum3)

}

func Square(number int) int { // функция расчёта и вывода квадрата
	return number * number
}
