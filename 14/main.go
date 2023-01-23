//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
//из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a interface{} = 1
	var b interface{} = "str"
	var c interface{} = true
	var d interface{} = make(chan int)

	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)
	cType := reflect.TypeOf(c)
	dType := reflect.TypeOf(d)

	fmt.Println(aType)
	fmt.Println(bType)
	fmt.Println(cType)
	fmt.Println(dType)

}

// будем использовать пакет reflect, который реализует отражение во время выполнения, позволяя программе
// манипулировать объектами произвольных типов.
// Обычно используется для получения значения со статическим типом interface{} и извлечения информации
// о его динамическом типе путем вызова TypeOf, который возвращает тип.
// пример реализации видел здесь https://golang-blog.blogspot.com/2020/05/detect-type-of-object.html
// а описание пакета - здесь https://go.dev/pkg/reflect/?m=old
