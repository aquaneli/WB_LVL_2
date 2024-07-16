package pattern

import (
	"fmt"
)

/*
	Применимость:	Фабричный метод - это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов
	в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

	Плюсы: 	1. Избавляет от привязки к конкретным классам продуктов.
			2. Выделяет код производства продуктов в одно место, упрощая поддержку кода.
			3. Упрощает добавление новых продуктов.
			4. Реализует принцип открытости/закрытости.

	Минусы:
			1. Сильно расширяет параллельные иерархии классов, потому что для каждого класса
			продуктов нужно создать свой подкласс создателя.

	Примеры использования паттерна на практике:
			1. Позволяет добавить новые товары в онлайн магазине.
			2. Можно использовать для расширения библиотеки или фреймворка.
*/

type TransportProduct interface {
	move()
	upgrade()
}

type CarProduct struct{}

func (CarProduct) move() {
	fmt.Println("Автомобиль едет по дороге")
}

func (CarProduct) upgrade() {
	fmt.Println("Автомобиль был улучшен")
}

type ShipProduct struct{}

func (ShipProduct) move() {
	fmt.Println("Корабль плывет в море")
}

func (ShipProduct) upgrade() {
	fmt.Println("Корбаль был улучшен")
}

type TransportCreator interface {
	CreateTransport() TransportProduct
}

type CarCreator struct{}

func (CarCreator) CreateTransport() TransportProduct {
	return &CarProduct{}
}

type ShipCreator struct{}

func (ShipCreator) CreateTransport() TransportProduct {
	return &ShipProduct{}
}

func FactoryMethodConstruct() {
	// Использование CarCreator для создания CarProduct
	var carCreator TransportCreator = &CarCreator{}
	car := carCreator.CreateTransport()
	car.move()
	car.upgrade()

	// Использование ShipCreator для создания ShipProduct
	var shipCreator TransportCreator = &ShipCreator{}
	ship := shipCreator.CreateTransport()
	ship.move()
	ship.upgrade()
}
