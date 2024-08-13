package storage

import (
	"errors"
	"log"
	"strconv"
	"time"
)

type Storage struct {
	Items map[string]map[string]time.Time
}

func NewStorge() *Storage {
	return &Storage{Items: make(map[string]map[string]time.Time)}
}

func (m *Storage) AddEvent(UserIdKey, DateVal string) {
	date, err := checkingIdAndData(UserIdKey, DateVal, "2010-10-10")
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := (*m).Items[UserIdKey]; !ok {
		(*m).Items[UserIdKey] = make(map[string]time.Time)
		(*m).Items[UserIdKey][DateVal] = date
		return
	}

	if _, ok := (*m).Items[UserIdKey][DateVal]; !ok {
		(*m).Items[UserIdKey][DateVal] = date
		return
	} else {
		log.Fatal(errors.New("эта дата уже занята"))
	}
}

func (m *Storage) UpdateEvent(UserIdKey, DateVal, DataReplace string) {
	_, err := checkingIdAndData(UserIdKey, DateVal, "2010-10-10")
	if err != nil {
		log.Fatal(err)
	}

	newDate, err := time.Parse("2010-10-10", DataReplace)
	if err != nil {
		log.Fatal(errors.New("некорректная дата, введите в формате yyyy-mm-dd"))
	}

	if _, ok := (*m).Items[UserIdKey]; !ok {
		log.Fatal(errors.New("такого пользователя не существует"))
	}

	if _, ok := (*m).Items[UserIdKey][DateVal]; !ok {
		log.Fatal(errors.New("нельзя заменить эту дату, потому что она не добавлена"))
	}

	if _, ok := (*m).Items[UserIdKey][DataReplace]; ok {
		log.Fatal(errors.New("нельзя заменить текущую дату на эту, потому что она уже добавлена"))
	}

	delete((*m).Items[UserIdKey], DateVal)
	(*m).Items[UserIdKey][DataReplace] = newDate
}

func (m *Storage) DeleteEvent(UserIdKey, DateVal string) {
	_, err := checkingIdAndData(UserIdKey, DateVal, "2010-10-10")
	if err != nil {
		log.Fatal(err)
	}
	if _, ok := (*m).Items[UserIdKey]; !ok {
		log.Fatal(errors.New("такого пользователя не существует"))
	}

	if _, ok := (*m).Items[UserIdKey][DateVal]; !ok {
		log.Fatal(errors.New("нельзя удалить эту дату, потому что она не добавлена"))
	}
	delete((*m).Items[UserIdKey], DateVal)
}

func checkingIdAndData(UserIdKey, DateVal, Template string) (time.Time, error) {
	if _, err := strconv.Atoi(UserIdKey); err != nil {
		return time.Time{}, errors.New("некорректный id, введите целое число")
	}

	t, err := time.Parse(Template, DateVal)
	if err != nil {
		return time.Time{}, errors.New("некорректная дата, введите в формате yyyy-mm-dd")
	}

	return t, nil
}

func (m *Storage) GetEventsForDay(UserIdKey string, Day string) {
	_, err := checkingIdAndData(UserIdKey, Day, "10")
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := (*m).Items[UserIdKey]; !ok {
		log.Fatal(errors.New("такого пользователя не существует"))
	}

}
