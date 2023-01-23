//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type human struct {
	name string
	age  int
}

// методы типа  human
func (h human) run() {
	fmt.Println(h.name + " is running")
}
func (h human) walk() {
	fmt.Println(h.name + " is walking")
}
func (h human) fly() {
	fmt.Println(h.name + " is flying")
}

type action struct {
	human // тип human встроен в action
} // теперь все методы и поля структуры human доступны через action

func main() {

	h := human{
		name: "Nick",
		age:  21,
	}

	i := action{
		human: h,
	}
	// обращаемся к методам структуры human через структуру action, т.к. они встроены
	action.run(i)
	action.walk(i)
	action.fly(i)

}

// https://golangify.com/composition-and-forwarding здесь подсмотрел кое что
