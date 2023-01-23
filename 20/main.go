// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	words := strings.Fields(str) // находим слова в строке

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 { // меняем местами слова в массиве
		words[i], words[j] = words[j], words[i]
	}

	fmt.Println(words)
}

// про функции пакета почитал здесь https://pkg.go.dev/strings#Fields
