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
	UniqCodeEvents map[string]struct{}
	Date           time.Time
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

func (s *Storage) Add(UserIdKey, DateVal string) (string, error) {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return "", errors.New("некорректный id, введите целое число")
	}

	DateParse, err := time.Parse("2006-01-02", DateVal)
	if err != nil {
		return "", errors.New("некорректная дата, введите в формате yyyy-mm-dd")
	}

	//если пользовтеля не существует
	if _, ok := (*s).Items[id]; !ok {
		(*s).Items[id] = make(map[string]events)
	}

	//если такой даты-ключа не существует
	if _, ok := s.Items[id][DateVal]; !ok {
		s.Items[id][DateVal] = events{UniqCodeEvents: make(map[string]struct{}), Date: DateParse}
	}
	uuid := getUniqCode()
	(*s).Items[id][DateVal].UniqCodeEvents[uuid] = struct{}{}

	return uuid, nil

}

func (s *Storage) UpDate(UserIdKey, DateValReplace, CodeEvent string) error {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return errors.New("некорректный id, введите целое число")
	}

	DateParse, err := time.Parse("2006-01-02", DateValReplace)
	if err != nil {
		return errors.New("некорректная дата, введите в формате yyyy-mm-dd")
	}

	if _, ok := (*s).Items[id]; !ok {
		return errors.New("пользователя не существует")
	}

	var ok bool
	for key := range (*s).Items[id] {
		_, ok = (*s).Items[id][key].UniqCodeEvents[CodeEvent]
		if ok {
			if _, okReplace := (*s).Items[id][DateValReplace]; !okReplace {
				(*s).Items[id][DateValReplace] = events{UniqCodeEvents: make(map[string]struct{}), Date: DateParse}
			}
			(*s).Items[id][DateValReplace].UniqCodeEvents[CodeEvent] = struct{}{}
			delete((*s).Items[id][key].UniqCodeEvents, CodeEvent)
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

	if _, ok := s.Items[id][DateVal].UniqCodeEvents[CodeEvent]; !ok {
		return errors.New("такого эвента не существет")
	}

	delete(s.Items[id][DateVal].UniqCodeEvents, CodeEvent)

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

	for _, event := range s.Items[id] {

		if (event.Date.Equal(startDate) || event.Date.After(startDate)) &&
			(event.Date.Before(startDate.AddDate(0, 0, 6)) || event.Date.Equal(startDate.AddDate(0, 0, 6))) {
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
	for _, event := range s.Items[id] {

		if (event.Date.Equal(startDate) || event.Date.After(startDate)) && (event.Date.Before(startDate.AddDate(1, 0, 0)) || event.Date.Equal(startDate.AddDate(1, 0, 0))) {
			eventsForYear = append(eventsForYear, event)
		}

	}

	if len(eventsForYear) == 0 {
		return nil, errors.New("за этот год нет событий")
	}

	return &eventsForYear, nil
}
