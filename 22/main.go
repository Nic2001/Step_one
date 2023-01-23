// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

// я так понял, здесь надо использовать big.Int
func main() {
	a := big.NewInt(123456789)
	b := big.NewInt(123456788)

	fmt.Println(new(big.Int).Add(a, b))
	fmt.Println(new(big.Int).Sub(a, b))
	fmt.Println(new(big.Int).Mul(a, b))
	fmt.Println(new(big.Int).Div(a, b))
}
