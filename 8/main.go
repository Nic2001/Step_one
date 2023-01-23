// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, i, ch int64

	fmt.Scan(&a, &i, &ch)

	// Конвертация двоичного значения в шестнадцатеричное
	b := strconv.FormatInt(a, 2)

	c := []byte(b)

	if ch == 1 {
		c[i] = 49
	} else {
		c[i] = 48
	}
	fmt.Println("было ", b, " стало ", string(c))
}
