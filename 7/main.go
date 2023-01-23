// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

// здесь придётся использовать блокировку записи с помощью мьютекса
func main() {
	var (
		m     = make(map[int]int)
		mutex sync.Mutex
		wg    sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)        // увеличиваем счётчик группы ожидания
		go func(i int) { // в анонимной горутине блокируем доступ к map и записываем в неё значение
			mutex.Lock()
			m[i] = i
			defer mutex.Unlock() // затем разблокирываем
			defer wg.Done()      // уменьшаем счётчик группы ожидания
		}(i)
	}
	wg.Wait() // блокировка потока
	fmt.Println(m)
}
