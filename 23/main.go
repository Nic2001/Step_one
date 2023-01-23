// Удалить i-ый элемент из слайса.
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 0

	fmt.Scan(&i)

	arr = append(arr[0:i], arr[i+1:]...) //при помощи append вроде проще всего

	fmt.Println(arr)
}
