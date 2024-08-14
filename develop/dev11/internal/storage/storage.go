package storage

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Storage struct {
	Items map[int]map[string]events
}

type events struct {
	uniqCodeEvents map[string]struct{}
	date           time.Time
}

func NewStorge() *Storage {
	return &Storage{Items: make(map[int]map[string]events)}
}

func (s *Storage) AddEvent(UserIdKey, DateVal string) error {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return errors.New("некорректный id, введите целое число")
	}

	dateParse, err := time.Parse("2010-10-25", DateVal)
	if err != nil {
		return errors.New("некорректная дата, введите в формате yyyy-mm-dd")
	}

	//если пользовтеля не существует
	if _, ok := (*s).Items[id]; !ok {
		(*s).Items[id] = make(map[string]events)
	}

	//если такой даты-ключа не существует
	if _, ok := s.Items[id][DateVal]; !ok {
		s.Items[id][DateVal] = events{uniqCodeEvents: make(map[string]struct{}), date: dateParse}
	}

	(*s).Items[id][DateVal].uniqCodeEvents[getUniqCode()] = struct{}{}

	return nil

}

func getUniqCode() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}

func (s *Storage) UpdateEvent(UserIdKey, DateValReplace, CodeEvent string) error {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return errors.New("некорректный id, введите целое число")
	}

	dateParse, err := time.Parse("2010-10-25", DateValReplace)
	if err != nil {
		return errors.New("некорректная дата, введите в формате yyyy-mm-dd")
	}

	if _, ok := (*s).Items[id]; !ok {
		return errors.New("пользователя не существует")
	}

	var ok bool
	for key := range (*s).Items[id] {
		_, ok = (*s).Items[id][key].uniqCodeEvents[CodeEvent]
		if ok {
			if _, okReplace := (*s).Items[id][DateValReplace]; !okReplace {
				(*s).Items[id][DateValReplace] = events{uniqCodeEvents: make(map[string]struct{}), date: dateParse}
			}
			(*s).Items[id][DateValReplace].uniqCodeEvents[CodeEvent] = struct{}{}
			delete((*s).Items[id][key].uniqCodeEvents, CodeEvent)
			break
		}
	}

	if !ok {
		return errors.New("такого евента не существует")
	}

	return nil
}

// func (m *Storage) DeleteEvent(UserIdKey, DateVal string) {
// 	_, err := checkingIdAndData(UserIdKey, DateVal, "2010-10-10")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if _, ok := (*m).Items[UserIdKey]; !ok {
// 		log.Fatal(errors.New("такого пользователя не существует"))
// 	}

// 	if _, ok := (*m).Items[UserIdKey][DateVal]; !ok {
// 		log.Fatal(errors.New("нельзя удалить эту дату, потому что она не добавлена"))
// 	}
// 	delete((*m).Items[UserIdKey], DateVal)
// }

// func (m *Storage) GetEventsForDay(UserIdKey string, Day string) {
// 	_, err := checkingIdAndData(UserIdKey, Day, "10")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, ok := (*m).Items[UserIdKey]; !ok {
// 		log.Fatal(errors.New("такого пользователя не существует"))
// 	}

// }
