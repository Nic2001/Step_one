//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.
package main

import "fmt"

func main() {

	s := "" //создадим пустую строку для задания переменной

	fmt.Scan(&s)    //заполним её
	b1 := []rune(s) //сделаем два одинаковых массива рун, т.к. символы могут состоять из двух байтов
	b2 := []rune(s) //чтоб было удобней, сразу сделаем их путём конвертирования строки

	for i := 0; i < len(b1); i++ { //перевернём один массив
		b2[len(b1)-i-1] = b1[i]
	}

	fmt.Println(string(b2)) //и выведем его
}
