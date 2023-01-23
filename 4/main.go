//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func runWorkers(n int, c chan int) { // запускаем n горутин, читающих из канала c
	for i := 0; i < n; i++ {
		go func() { // анонимная горутина, читающая из канала, пока он не закрыт
			for v := range c {
				fmt.Println(v) // для проверки работы
			}
		}()
	}
}

func main() {
	c := make(chan int)

	osSign := make(chan os.Signal, 1) // канал для записи сигнала ОС

	signal.Notify(osSign, syscall.SIGINT) // в канал будет записано значение, если поступил SIGINT, генерируемый Ctrl+С

	n := 10
	runWorkers(n, c) // запускаем воркеры

	for { // в бесконечном цикле закидываем данные в канал
		select {
		case <-osSign: // если есть значение в канале для сигнала системы - закрываем канал, прекращаем цикл
			close(c)
			break // Читает из канала, пока он не закрыт
		default: // закидываем данные
			c <- rand.Int()
		}
	}
}

// про воркеры немного посмотрел здесь https://proglib.io/p/parallelizm-v-golang-i-workerpool-chast-1-2020-12-24
