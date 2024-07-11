package main

import "fmt"

/*
	Применимость:	Фасад - это структурный паттерн проектирования, который
	предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.

	Плюсы: 	1. Упрощает взаимодействие пользователя с компонентами системы и изолирует систему.
			2. Уменьшает зависимость клиента и сложной системой.

	Минусы:
			1. Есть риск стать, что фасад будет зависимым от всех компонентов программы.

	Примеры использования паттерна на практике:
			1. Взаимодействие программы с БД. Нам предоставляется готовый
			пакет в котором реализованы методы и нам не нужно самим писать логику и погружаться во все тонкости.
			2. Если мы подключим фреймворк и нам нужно получить общий результат из нескольких методов, то мы можем сами
			реалиовать фасад и прописать 1 функцию которая внутри себя сделает какие то вычисления, но по итогу мы будем вызывать только 1 метод всегда.
*/

type Facade struct {
	r robot
	b box
}

func (f *Facade) MoveFullBox() {
	f.b.putItem()
	f.r.lowerHand()
	f.r.closeHand()
	f.r.raiseHand()
	f.r.openHand()
	f.b.removeItem()
}

type robot struct{}

func (r *robot) openHand()  { fmt.Println("Robot hand opened") }
func (r *robot) closeHand() { fmt.Println("Robot hand closed") }
func (r *robot) raiseHand() { fmt.Println("Robot hand raised") }
func (r *robot) lowerHand() { fmt.Println("Robot hand lowered") }

type box struct{}

func (b *box) putItem()    { fmt.Println("Item put in the box") }
func (b *box) removeItem() { fmt.Println("Item removed from the box") }

func main() {
	facade := Facade{}
	facade.MoveFullBox()
}
