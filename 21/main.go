// Реализовать паттерн «адаптер» на любом примере.
package main

import "fmt"

// для примера возьмём мой ноут и мак моего друга, которые мы смогли соединить USB кабелем только с помощью адаптера

type user struct { // структура, описывающая пользователя
}

func (c *user) insertUSBIntoComputer(computer computer) { // функция пользователя присоединения usb к обычному компьютеру
	fmt.Println("Пользователь присоединил usb кабель")
	computer.insertIntoUSBPort()
}

type computer interface { // интерфейс подключения к обычному USB
	insertIntoUSBPort()
}

type mac struct{} // структура мака с USB Type-C

func (m *mac) insertIntoUSBCPort() { // функция присоединения к нему
	fmt.Println("Успешное присоединение к USB Type-C")
}

type windows struct{} // структура обычного компа с обычным USB

func (w *windows) insertIntoUSBPort() { // функция присоединения к нему
	fmt.Println("Успешное присоединение к USB")
}

type macAdapter struct { // адаптер
	macbook *mac
}

func (m *macAdapter) insertIntoUSBPort() {
	fmt.Println("Меняем интерфейсы подключения (работает адаптер)")
	m.macbook.insertIntoUSBCPort()
}

func main() {
	user := &user{}
	MSI := &windows{}
	macbook := &mac{}

	user.insertUSBIntoComputer(MSI)

	macAdapter := &macAdapter{
		macbook: macbook,
	}

	user.insertUSBIntoComputer(macAdapter)
}

// в данном случае это структура, которая придерживается того же интерфейса, который ожидает пользователь (обычный USB-порт)
// а запрос от клиента к переводит в дргую структуру (USB Type-C)
