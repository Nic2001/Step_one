//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.
package main

import (
	"fmt"
	"sync"
)

type Increment struct { // структура со счётчиком
	count int
	lock  sync.Mutex
}

func (i *Increment) Increase() { // инкрементирующий метод
	i.lock.Lock()
	i.count++
	defer i.lock.Unlock()
}

func main() {
	increment := Increment{count: 0} // экземпляр структуры
	var wg sync.WaitGroup            // группа ожидания
	n := 1000
	wg.Add(n)
	for i := 0; i < n; i++ { // в цикле инкрементируем счётчик
		go func() {
			for i := 0; i < 10; i++ {
				increment.Increase()
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(increment.count) // смотрим итоговое количество, должно быть 1000
}
