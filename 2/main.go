// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	number := []int{2, 4, 6, 8, 10}

	// с помощью группы ожидания
	var wg sync.WaitGroup
	wg.Add(len(number)) // заполняем счётчик горутин

	for _, number := range number { //Вопрос
		go func(number int) { // анонимная горутина
			defer wg.Done() // уменьшаем счётчик горутин
			Square(number)
		}(number)
	}
	wg.Wait() // блокировка потока

	// простой вариант с задержкой временем
	for i := 0; i < 5; i++ {
		go Square(number[i])
	}
	time.Sleep(5 * time.Millisecond)

}

func Square(number int) { // функция расчёта и вывода квадрата
	fmt.Println(number * number)
}
