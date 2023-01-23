// Реализовать все возможные способы остановки выполнения горутины.
package main

import (
	"context"
	"fmt"
)

func main() {

	// завершение работы горутины при закрытии канала
	c := make(chan int)
	go func() { // запускаем анонимную горутину
		for i := range c {
			fmt.Println(i) // в цикле читаем данные из канала
		}
	}()
	for i := 0; i < 5; i++ {
		c <- i // в цикле закидываем данные в канал
	}
	close(c) // закрываем канал, и горутина завершается

	// завершение с помощью контекста
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		<-ctx.Done()
	}(ctx)
	cancel()

}
