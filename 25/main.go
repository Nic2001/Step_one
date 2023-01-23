// Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d) // time.After ожидает истечения времени d, а затем отправляет текущее время по возвращенному каналу
	// но нам это не нужно, надо чтоб просто функция подождала
	fmt.Println("Конец")
}

func main() {
	fmt.Println("Начало") // засечём старт программы
	sleep(time.Second * 2)
}

// https://pkg.go.dev/time#After здесь про time.After
