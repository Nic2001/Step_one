// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	a := 5
	b := 3

	fmt.Println(a, b)

	a = a * b
	b = a / b
	a = a / b //это я сам придумал

	fmt.Println(a, b)

	a, b = b, a //кроме того, в go можно менят местами две перенменные не создавая третью

	fmt.Println(a, b)
}
