package pattern

import (
	"fmt"
	"reflect"
)

/*
	Применимость:	Посетитель - это поведенческий паттерн проектирования, который позволяет создавать новые операции, не меняя объектов или
	незначительно дополняя классы, над которыми эти операции могут выполняться.

	Плюсы: 	1. Упрощает добавление операций на структурой объектов.
			2. Объединяет операцци в одном классе.

	Минусы:
			1. Может нарушать инкапсуляцию.
			2. Нет смысла если иерархия компонентов постоянно меняется.

	Примеры использования паттерна на практике:
			1. Когда нужно выполнить операцию над всеми элементами дерева.
			2. Когда не хочется засорять классы лишними операциями, но требуется расшириться.
*/

/* Интерфейс посетителя , в котором был расширен функционал для каждого класса */
type Visitor interface {
	ConverterEuroToRub(Euro)
	ConverterDollarToRub(Dollar)
	ConverterYuanToRub(Yuan)
}

/* Конкретный посетитель это структура в которой реализован интерфейс посетителя длякаждого класса */
type ConcreteVisitor struct{}

func (ConcreteVisitor) ConverterEuroToRub(e Euro) {
	fmt.Printf("Converting %s to Rub\n", reflect.TypeOf(e))
}

func (ConcreteVisitor) ConverterDollarToRub(d Dollar) {
	fmt.Printf("Converting %s to Rub\n", reflect.TypeOf(d))
}

func (ConcreteVisitor) ConverterYuanToRub(y Yuan) {
	fmt.Printf("Converting %s to Rub\n", reflect.TypeOf(y))
}

/* Интерфейс валют в котором реализован метод для посетителя */
type Currency interface {
	Accept(Visitor)
}

type Euro struct{}

func (e Euro) Accept(v Visitor) {
	v.ConverterEuroToRub(e)
}

type Dollar struct{}

func (d Dollar) Accept(v Visitor) {
	v.ConverterDollarToRub(d)
}

type Yuan struct{}

func (y Yuan) Accept(v Visitor) {
	v.ConverterYuanToRub(y)
}

func VisitorConstruct() {
	c := []Currency{Euro{}, Dollar{}, Yuan{}}
	v := ConcreteVisitor{}
	for _, val := range c {
		val.Accept(v)
	}
}
