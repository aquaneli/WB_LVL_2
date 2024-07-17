package pattern

import "fmt"

/*
	Применимость:	Состояние - это поведенческий паттер проектирования, который позволяет менять
	поведение в зависимости от своего состояния.

	Плюсы: 	1. Избавляет от множества больших условнх операторов.
			2. Упрощает код контекста.

	Минусы:
			1. Может усложнить код, если состояний мало и они редко меняются.

	Примеры использования паттерна на практике:
			1. Публикация файла на сайте, когда он может находиться в состоянии модерации или в состоянии публикации.
			2. В играх если ехать на машине или самолете, клавиши одни и те же, но в контексте транспорта будет разное поведение.
*/

type ATMState interface {
	InsertCard()
	EjectCard()
	EnterPin(pin int)
	RequestCash(amount int)
}

type ATMContext struct {
	state ATMState
	cash  int
	pin   int
}

func (atm *ATMContext) SetState(state ATMState) {
	atm.state = state
}

func (atm *ATMContext) InsertCard() {
	atm.state.InsertCard()
}

func (atm *ATMContext) EjectCard() {
	atm.state.EjectCard()
}

func (atm *ATMContext) EnterPin(pin int) {
	atm.state.EnterPin(pin)
}

func (atm *ATMContext) RequestCash(amount int) {
	atm.state.RequestCash(amount)
}

// InsertCardState: состояние когда карта вставлена
type InsertCardState struct {
	context *ATMContext
}

func (state *InsertCardState) InsertCard() {
	fmt.Println("Карта уже вставлена")
}

func (state *InsertCardState) EjectCard() {
	fmt.Println("Карта извлечена")
	state.context.SetState(&NoCardState{context: state.context})
}

func (state *InsertCardState) EnterPin(pin int) {
	if pin == state.context.pin {
		fmt.Println("PIN верный")
		state.context.SetState(&HasPinState{context: state.context})
	} else {
		fmt.Println("PIN неверный")
		state.context.EjectCard()
	}
}

func (state *InsertCardState) RequestCash(amount int) {
	fmt.Println("Сначала введите PIN")
}

// NoCardState: состояние когда карта не вставлена
type NoCardState struct {
	context *ATMContext
}

func (state *NoCardState) InsertCard() {
	fmt.Println("Карта вставлена")
	state.context.SetState(&InsertCardState{context: state.context})
}

func (state *NoCardState) EjectCard() {
	fmt.Println("Карта не вставлена")
}

func (state *NoCardState) EnterPin(pin int) {
	fmt.Println("Сначала вставьте карту")
}

func (state *NoCardState) RequestCash(amount int) {
	fmt.Println("Сначала вставьте карту")
}

// HasPinState: состояние когда введён правильный PIN
type HasPinState struct {
	context *ATMContext
}

func (state *HasPinState) InsertCard() {
	fmt.Println("Карта уже вставлена")
}

func (state *HasPinState) EjectCard() {
	fmt.Println("Карта извлечена")
	state.context.SetState(&NoCardState{context: state.context})
}

func (state *HasPinState) EnterPin(pin int) {
	fmt.Println("PIN уже введен")
}

func (state *HasPinState) RequestCash(amount int) {
	if amount <= state.context.cash {
		fmt.Printf("Выдано %d\n", amount)
		state.context.cash -= amount
		state.context.EjectCard()
		if state.context.cash <= 0 {
			state.context.SetState(&NoCashState{context: state.context})
		}
	} else {
		fmt.Println("Недостаточно средств")
		state.context.EjectCard()
	}
}

// NoCashState: состояние когда нет денег в банкомате
type NoCashState struct {
	context *ATMContext
}

func (state *NoCashState) InsertCard() {
	fmt.Println("В банкомате нет денег")
}

func (state *NoCashState) EjectCard() {
	fmt.Println("В банкомате нет денег")
}

func (state *NoCashState) EnterPin(pin int) {
	fmt.Println("В банкомате нет денег")
}

func (state *NoCashState) RequestCash(amount int) {
	fmt.Println("В банкомате нет денег")
}

func StateConstruct() {
	atm := &ATMContext{
		cash: 1000,
		pin:  1234,
	}
	noCardState := &NoCardState{context: atm}
	atm.SetState(noCardState)

	atm.InsertCard()
	atm.EnterPin(1234)
	atm.RequestCash(500)
	atm.InsertCard()
	atm.EnterPin(1234)
	atm.RequestCash(600)
	atm.InsertCard()
}
