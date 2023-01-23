//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func() {
		for i := 0; i < len(array); i++ {
			ch1 <- array[i]
			ch2 <- array[i] * array[i]
		}
	}()
	go func() {
		for i := 0; i < len(array); i++ {
			<-ch1
			fmt.Println(<-ch2)
		}
	}()

	time.Sleep(1 * time.Second) // чуть чуть держим время, чтоб горутины успели выполниться

	// можно сделать вариант с группой ожиания

	var wg sync.WaitGroup
	wg.Add(len(array)) // заполняем счётчик горутин

	go func() {
		for i := 0; i < len(array); i++ {
			ch1 <- array[i]
			ch2 <- array[i] * array[i]
		}
	}()
	go func() {
		for i := 0; i < len(array); i++ {
			<-ch1
			fmt.Println(<-ch2)
			wg.Done() // уменьшаем счётчик горутин
		}
	}()

	wg.Wait() // блокировка потока
}
