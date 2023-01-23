//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
package main

import (
	"fmt"
	"strings"
)

func main() {

	firstString := "abcd"
	secondString := "abCdefAaf"
	thirdString := "aabcd"

	unicStr(firstString)
	unicStr(secondString)
	unicStr(thirdString)
}

func unicStr(s string) {
	unic := 1
	s = strings.ToLower(s) //т.к регистр не важен, переводим всю строку в нижний регистр

	symSet := make(map[rune]struct{}) // на всякий случай используем руны, а не байты

	for _, r := range s {

		if _, ok := symSet[r]; ok { // Если такой символ уже есть во множестве, возвращаем false
			fmt.Println("false")
			unic = 0
			break
		}

		symSet[r] = struct{}{} // Если символ ранее не встречался, добавляем его во множество
	}

	if unic == 1 {
		fmt.Println("true")
	}
}
