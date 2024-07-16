package pattern

import (
	"fmt"
)

/*
	Применимость:	Команда - это поведенческий паттерн проектирования, который превращает
	запросы в объекты, позволяя передавать их как аргументы при вызове методов.

	Плюсы: 	1. Убирает прямую зависимость между объектами.
			2. Позволяет реализовать простую отмену и повтор операций используя историю.
			3. Реализует принцип открытости/закрытости.

	Минусы:
			1. Усложняет код программы за чет дополнительных классов.

	Примеры использования паттерна на практике:
			1. Можно использовать в графическом меню, где разные кнопки вполняют одни и те же действия.
			2. Можно отправить объекты команд по сети для выполнения на другой машине, например действие игрока в компьютерной игре.
*/

/* Command интерфейс */
type Command interface {
	execute()
}

/* AddCommand структура команды добавления данных */
type AddCommand struct {
	db DataBase
}

/* execute выполняет добавление данных в базу */
func (a *AddCommand) execute() {
	a.db.AddData()
}

/* DelCommand структура команды удаления данных */
type DelCommand struct {
	db DataBase
}

/* execute выполняет удаление данных из базы */
func (d *DelCommand) execute() {
	d.db.DelData()
}

/*  Отправитель */
type Invoker struct {
	h HistoryCommand
}

/* Выполняеися команда добавления в базу данных и сохраняется в истории */
func (i *Invoker) PushData(command Command) {
	command.execute()
	i.h.pushHistory(command)
}

/* Выполняеися команда удаления из базы данных и сохраняется в истории */
func (i *Invoker) PopData(command Command) {
	command.execute()
	i.h.pushHistory(command)
}

/* Отменить действие */
func (i *Invoker) undo() {
	typeHistory := i.h.popHistory()
	switch v := typeHistory.(type) {
	case *AddCommand:
		v.db.DelData()
	case *DelCommand:
		v.db.AddData()
	}
}

/*  Получатель */
type DataBase struct{}

/* Добавление в базу данных на стороне получателя */
func (DataBase) AddData() {
	fmt.Println("Данные добавлены")
}

/* Удаление из базы данных на стороне получателя */
func (DataBase) DelData() {
	fmt.Println("Данные удалены")
}

/* HistoryCommand структура для управления историей команд */
type HistoryCommand struct {
	stack []Command
}

/* Сохранить в историю */
func (h *HistoryCommand) pushHistory(c Command) {
	h.stack = append(h.stack, c)
}

/* Удалить элемент из истории */
func (h *HistoryCommand) popHistory() Command {
	if len(h.stack) == 0 {
		return nil
	}
	cmd := h.stack[len(h.stack)-1]
	h.stack = h.stack[:len(h.stack)-1]
	return cmd
}

func CommandConstruct() {
	i := Invoker{}
	db := DataBase{}
	a := AddCommand{db}
	d := DelCommand{db}
	i.PushData(&a)
	i.PopData(&d)
	i.undo()
}
