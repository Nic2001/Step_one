// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func main() { // по советам многих экспертов из тг канала "WB Стажеры Go 3" будем использовать map[string]struct{},
	// т.к. ключ будет уникальным, а большего нам и не надо

	s := []string{"cat", "cat", "dog", "cat", "tree"}

	sm := make(map[string]struct{})
	for _, i := range s {
		sm[i] = struct{}{}
	}

	for i := range sm {
		fmt.Print(i, " ")
	}
}
