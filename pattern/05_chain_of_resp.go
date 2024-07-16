package pattern

import (
	"fmt"
)

/*
	Применимость:	Цепочка вызовов - это поведенческий паттерн проектирования, который
	позволяет передавать запросы последовательно по цепочке обработчиков.

	Плюсы: 	1. Уменьшает зависимость между клиентом и обработчиком.
			2. Реализует принцип единственной обязанности.
			3. Реализует принцип открытости/закрытости.

	Минусы:
			1. Усложняет код программы за счет дополнительных классов.
			2. Запрос может быть никем не обработан.

	Примеры использования паттерна на практике:
			1. Обработчик сообщений, например проверять спам это или нет. Через цепочку сделать проверку сообщения и если не проходит проверку, то поместить в спам.
			2. Авторизация на сайте. Логин и пароль прогонятьчерез цепочку вызовов, и если все обработчики пройдены успешно, то дать доступ.
*/

/* Интерфейс обработчика */
type Handler interface {
	Handle(code int) error
	SetNext(Handler)
}

/* Базовый обработчик от которого будут строиться все остальные */
type BaseHandler struct {
	Next Handler
}

/* Установка обработчика */
func (base *BaseHandler) SetNext(h Handler) {
	base.Next = h
}

/* Метод обработки входных данных */
func (base *BaseHandler) Handle(code int) error {
	if base.Next != nil {
		return base.Next.Handle(code)
	}
	return nil
}

/* Конкретный обработчик Авторизации */
type AuthorizationHandler struct {
	BaseHandler
}

func (a *AuthorizationHandler) Handle(code int) error {
	if code == 1 {
		fmt.Println("Авторизация пройдена")
		return a.BaseHandler.Handle(code)
	} else if code == 0 {
		fmt.Println("Авторизация не пройдена")
		return nil
	}
	fmt.Println("Невозможно обработать")
	return a.BaseHandler.Handle(code)

}

/* Конкретный обработчик Аутентификации */
type AuthenticationHandler struct {
	BaseHandler
}

func (a *AuthenticationHandler) Handle(code int) error {
	if code == 1 {
		fmt.Println("Аутентификация пройдена")
		return a.BaseHandler.Handle(code)
	} else if code == 0 {
		fmt.Println("Аутентификация не пройдена")
		return nil
	}
	fmt.Println("Невозможно обработать")
	return a.BaseHandler.Handle(code)
}

/* Конкретный обработчик Идентификации */
type IdentificationHandler struct {
	BaseHandler
}

func (i *IdentificationHandler) Handle(code int) error {
	if code == 1 {
		fmt.Println("Идентификация пройдена")
		return i.BaseHandler.Handle(code)
	} else if code == 0 {
		fmt.Println("Идентификация не пройдена")
		return nil
	}
	fmt.Println("Невозможно обработать")
	return i.BaseHandler.Handle(code)
}

func CORConstruct() {
	base := BaseHandler{}
	a := AuthorizationHandler{}
	b := AuthenticationHandler{}
	c := IdentificationHandler{}
	base.SetNext(&a)
	a.SetNext(&b)
	b.SetNext(&c)
	for _, val := range []int{0, 1, 2} {
		base.Handle(val)
	}
}
