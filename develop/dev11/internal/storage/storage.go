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

func (s *Storage) Add(UserIdKey, DateVal string) error {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return errors.New("некорректный id, введите целое число")
	}

	dateParse, err := time.Parse("2006-01-02", DateVal)
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

func (s *Storage) Update(UserIdKey, DateValReplace, CodeEvent string) error {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return errors.New("некорректный id, введите целое число")
	}

	dateParse, err := time.Parse("2006-01-02", DateValReplace)
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
		return errors.New("такого эвента не существует")
	}

	return nil
}

func (s *Storage) Remove(UserIdKey, DateVal, CodeEvent string) error {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return errors.New("некорректный id, введите целое число")
	}

	_, err = time.Parse("2006-01-02", DateVal)
	if err != nil {
		return errors.New("некорректная дата, введите в формате yyyy-mm-dd")
	}

	if _, ok := (*s).Items[id]; !ok {
		return errors.New("пользователя не существует")
	}

	if _, ok := s.Items[id][DateVal]; !ok {
		return errors.New("на эту дату эвентов нет")
	}

	if _, ok := s.Items[id][DateVal].uniqCodeEvents[CodeEvent]; !ok {
		return errors.New("такого эвента не существет")
	}

	delete(s.Items[id][DateVal].uniqCodeEvents, CodeEvent)

	return nil
}

func (s *Storage) GetEventsForDay(UserIdKey string, DateVal string) (*events, error) {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return nil, errors.New("некорректный id, введите целое число")
	}

	_, err = time.Parse("2006-01-02", DateVal)
	if err != nil {
		return nil, errors.New("некорректная дата, введите в формате yyyy-mm-dd")
	}

	if _, ok := (*s).Items[id]; !ok {
		return nil, errors.New("такого пользователя не существует")
	}

	if _, ok := (*s).Items[id][DateVal]; !ok {
		return nil, errors.New("на этот день эвентов нет")
	}

	events := (*s).Items[id][DateVal]

	return &events, nil
}

func (s *Storage) GetEventsForWeek(UserIdKey string, DateVal string) (*[]events, error) {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return nil, errors.New("некорректный id, введите целое число")
	}

	startDate, err := time.Parse("2006-01-02", DateVal)
	if err != nil {
		return nil, errors.New("некорректная дата, введите в формате yyyy-mm-dd")
	}

	if _, ok := (*s).Items[id]; !ok {
		return nil, errors.New("такого пользователя не существует")
	}

	eventsForWeek := []events{}
	for dateStr, event := range s.Items[id] {
		date, err := time.Parse("2006-01-02", dateStr)
		if err == nil && date.After(startDate) && date.Before(startDate.AddDate(0, 0, 7)) {
			eventsForWeek = append(eventsForWeek, event)
		}
	}

	if len(eventsForWeek) == 0 {
		return nil, errors.New("на этой неделе событий нет")
	}

	return &eventsForWeek, nil
}

func (s *Storage) GetEventsForYear(UserIdKey string, DateVal string) (*[]events, error) {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return nil, errors.New("некорректный id, введите целое число")
	}

	startDate, err := time.Parse("2006-01-02", DateVal)
	if err != nil {
		return nil, errors.New("некорректная дата, введите в формате yyyy-mm-dd")
	}

	if _, ok := (*s).Items[id]; !ok {
		return nil, errors.New("такого пользователя не существует")
	}

	eventsForYear := []events{}
	for dateStr, event := range s.Items[id] {
		// Парсим дату и проверяем год
		date, err := time.Parse("2006-01-02", dateStr)
		if err == nil && date.Year() == startDate.Year() {
			eventsForYear = append(eventsForYear, event)
		}
	}

	// Если нет событий за этот год
	if len(eventsForYear) == 0 {
		return nil, errors.New("за этот год нет событий")
	}

	return &eventsForYear, nil
}
