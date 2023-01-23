// Реализовать бинарный поиск встроенными методами языка.
package main

import "fmt"

func binarySearch(a int, array []int) {

	low := 0 // первый и последний индексы массива
	high := len(array) - 1

	for low <= high { // пока массив существует
		mid := (low + high) / 2 // находим середину

		if array[mid] < a { // сравниваем срединный элемент с искомым и по результатам выбираем какую половину дальше исследуем
			low = mid + 1 // вот поэтому массив должен быть отсортирован
		} else {
			high = mid - 1
		}
	}

	if low == len(array) || array[low] != a {
		fmt.Println("Элемент отсутствует")
	}

	fmt.Println("Элемент присутствует, индекс элемента: ", low)
}

func main() {
	array := []int{1, 2, 3, 4, 6, 10, 20, 30, 40, 50} // ОТСОРТИРОВАННЫЙ массив целых чисел
	binarySearch(6, array)
}

//много взял отсюда https://dzen.ru/a/YkQvRcEV3CBQCAs-, https://www.golangprograms.com/golang-program-for-implementation-of-binary-search.html
