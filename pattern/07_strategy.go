package pattern

import (
	"fmt"
	"reflect"
)

/*
	Применимость:	Стратегия - это поведенческий паттер проектирования, который определяет схожие алгоритмы
	и помещает каждый из них в отдельный класс.

	Плюсы: 	1. Изолирует код и данные алгоритмов от остальных классов.
			2. Уход от наследования к делегированию.
			у. Реализует принцип открытости/закрытости.

	Минусы:
			1. Усложняет программу за счет дополнительных классов.
			2. Требуется знать в чем разница каждой стратегии.

	Примеры использования паттерна на практике:
			1. Навигатор где строится оптимальный маршрут для авто или пешехода.
			2. Когда требуется избавиться от множественного условного оператора, где
			каждое условие представляет вариацию алгоритма.
*/

/* Интерфейс для каждой вариации алгоритма */
type LabyrinthStrategy interface {
	execute()
}

/* Конкретная стратегия */
type BinThreeStrategy struct{}

func (b BinThreeStrategy) execute() {
	fmt.Printf("Лабиринт создан с помощью алгоритма %s\n", reflect.TypeOf(b))
}

/* Конкретная стратегия */
type SidewinderStrategy struct{}

func (s SidewinderStrategy) execute() {
	fmt.Printf("Лабиринт создан с помощью алгоритма %s\n", reflect.TypeOf(s))
}

/* Конкретная стратегия */
type EllersStrategy struct{}

func (e EllersStrategy) execute() {
	fmt.Printf("Лабиринт создан с помощью алгоритма %s\n", reflect.TypeOf(e))
}

/* Контекст в котором хранится конкретная стратегия */
type Context struct {
	strategy LabyrinthStrategy
}

/* Установка конкретной стратегии в контекст*/
func (c *Context) SetStrategy(strategy LabyrinthStrategy) {
	c.strategy = strategy
}

/* Выполнить алгоритм конкретной стратегии */
func (c *Context) DoSomething() {
	c.strategy.execute()
}

func StrategyConstruct() {
	three := BinThreeStrategy{}
	sidewinder := SidewinderStrategy{}
	eller := EllersStrategy{}

	context := Context{}
	context.SetStrategy(three)
	context.DoSomething()

	context.SetStrategy(sidewinder)
	context.DoSomething()

	context.SetStrategy(eller)
	context.DoSomething()
}
