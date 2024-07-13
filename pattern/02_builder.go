package pattern

import "fmt"

/*
	Применимость:	Строитель - это порождающий паттерн, который помогает создавать новые объекты пошагово избавляя клиента
	от создания усложненного конструирования объекта.

	Плюсы:	1. Позволяет создавать объект пошагово.
			2. Позволяет использовать один и тот же код.
			3. Изолирует сложный код сборки продукта от его основной бизнес-логики.

	Минусы: 1. Усложняет код.
			2. Клиент привязан к конкретным классам строителя.

	Примеры использования паттерна на практике:
			1. Может быть использован в играх, где необходимо создать какого либо персонажа, его поведение может быть одинаково, но характеристики отличаться.
			2. Можно использовать в конструкторе сайтов, где элементы можно конструировать по разному , но концепция везде одинаковая.
*/

type Builder interface {
	SetWheels(int)
	SetDoors(int)
	SetEngine(string)
}

/* Каждый конкретный строитель реализует интерфейс по своему */

/* Это строитель настоящего автомобиля, а у каждого автомобиля свои характеристики */
type CarBuilder struct{}

func (CarBuilder) SetWheels(wheels int) {
	fmt.Printf("Установить %d колеса", wheels)
}

func (CarBuilder) SetDoors(doors int) {
	fmt.Printf("Установить %d двери", doors)
}

func (CarBuilder) SetEngine(engine string) {
	fmt.Printf("Установить %s двигатель", engine)
}

/* Это строитель чертежа автомобиля, а у каждого чертежа свои характеристики */

type CarDrawingBuilder struct{}

func (CarDrawingBuilder) SetWheels(wheels int) {
	fmt.Printf("Нарисовать %d колеса", wheels)
}

func (CarDrawingBuilder) SetDoors(doors int) {
	fmt.Printf("Нарисовать %d двери", doors)
}

func (CarDrawingBuilder) SetEngine(engine string) {
	fmt.Printf("Нарисовать %s двигатель", engine)
}

/*
	Используя поведение каждого строителя мы можем сконструировать нужный нам автомобиль
	Я просто создал конкретных строителей, которые выполняют свою узкую задачу - создать, нарисовать, могу даже например сделать 3D модель,
	но создав строителя который будет распечатывать на 3D принтере детали
*/

type Director struct{}

func (Director) ConstructPassengerCar(builder Builder) {
	builder.SetDoors(5)
	builder.SetEngine("Обычный")
	builder.SetWheels(4)
}

func BuilderConstruct() {
	builder := CarBuilder{}
	d := Director{}
	d.ConstructPassengerCar(builder)

	builderDraw := CarDrawingBuilder{}
	d.ConstructPassengerCar(builderDraw)

}
