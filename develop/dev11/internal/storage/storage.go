package storage

import (
	"crypto/rand"
	"fmt"
	"net/http"
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

func (s *Storage) Add(UserIdKey, DateVal string) (string, int) {
	id, dateParse, status := BadRequest(UserIdKey, DateVal)
	if status != http.StatusOK {
		return "", status
	}

	//если пользовтеля не существует
	if _, ok := (*s).Items[id]; !ok {
		(*s).Items[id] = make(map[string]events)
	}

	//если такой даты-ключа не существует
	if _, ok := s.Items[id][DateVal]; !ok {
		s.Items[id][DateVal] = events{UniqCodeEvents: make(map[string]struct{}), Date: dateParse}
	}

	uuid, err := getUniqCode()
	if err != nil {
		return "", http.StatusInternalServerError
	}

	(*s).Items[id][DateVal].UniqCodeEvents[uuid] = struct{}{}

	return uuid, http.StatusOK
}

func (s *Storage) UpDate(UserIdKey, DateValReplace, CodeEvent string) int {
	id, dateParse, status := BadRequest(UserIdKey, DateValReplace)
	if status != http.StatusOK {
		return status
	}

	if _, ok := (*s).Items[id]; !ok {
		return http.StatusServiceUnavailable
	}

	var ok bool
	for key := range (*s).Items[id] {
		_, ok = (*s).Items[id][key].UniqCodeEvents[CodeEvent]
		if ok {
			if _, okReplace := (*s).Items[id][DateValReplace]; !okReplace {
				(*s).Items[id][DateValReplace] = events{UniqCodeEvents: make(map[string]struct{}), Date: dateParse}
			}
			(*s).Items[id][DateValReplace].UniqCodeEvents[CodeEvent] = struct{}{}
			delete((*s).Items[id][key].UniqCodeEvents, CodeEvent)
			break
		}
	}

	if !ok {
		return http.StatusServiceUnavailable
	}

	return http.StatusOK
}

func (s *Storage) Remove(UserIdKey, DateVal, CodeEvent string) int {
	id, _, status := BadRequest(UserIdKey, DateVal)
	if status != http.StatusOK {
		return status
	}

	if _, ok := (*s).Items[id]; !ok {
		return http.StatusServiceUnavailable
	}

	if _, ok := s.Items[id][DateVal]; !ok {
		return http.StatusServiceUnavailable
	}

	if _, ok := s.Items[id][DateVal].UniqCodeEvents[CodeEvent]; !ok {
		return http.StatusServiceUnavailable
	}

	delete(s.Items[id][DateVal].UniqCodeEvents, CodeEvent)

	if len(s.Items[id][DateVal].UniqCodeEvents) == 0 {
		delete(s.Items[id], DateVal)
	}

	return http.StatusOK
}

func (s *Storage) GetEventsForDay(UserIdKey string, DateVal string) (*events, int) {
	id, _, status := BadRequest(UserIdKey, DateVal)
	if status != http.StatusOK {
		return nil, status
	}

	if _, ok := (*s).Items[id]; !ok {
		return nil, http.StatusServiceUnavailable
	}

	if _, ok := (*s).Items[id][DateVal]; !ok {
		return nil, http.StatusServiceUnavailable
	}

	events := (*s).Items[id][DateVal]

	return &events, http.StatusOK
}

func (s *Storage) GetEventsForWeek(UserIdKey string, DateVal string) (*[]events, int) {
	id, startDate, status := BadRequest(UserIdKey, DateVal)
	if status != http.StatusOK {
		return nil, status
	}
	if _, ok := (*s).Items[id]; !ok {
		return nil, http.StatusServiceUnavailable
	}

	eventsForWeek := []events{}

	for _, event := range s.Items[id] {
		if (event.Date.Equal(startDate) || event.Date.After(startDate)) &&
			(event.Date.Before(startDate.AddDate(0, 0, 6)) || event.Date.Equal(startDate.AddDate(0, 0, 6))) {
			eventsForWeek = append(eventsForWeek, event)
		}
	}

	if len(eventsForWeek) == 0 {
		return nil, http.StatusServiceUnavailable
	}

	return &eventsForWeek, http.StatusOK
}

func (s *Storage) GetEventsForYear(UserIdKey string, DateVal string) (*[]events, int) {
	id, startDate, status := BadRequest(UserIdKey, DateVal)
	if status != http.StatusOK {
		return nil, status
	}

	if _, ok := (*s).Items[id]; !ok {
		return nil, http.StatusServiceUnavailable
	}

	eventsForYear := []events{}
	for _, event := range s.Items[id] {
		if (event.Date.Equal(startDate) || event.Date.After(startDate)) &&
			(event.Date.Before(startDate.AddDate(1, 0, 0)) || event.Date.Equal(startDate.AddDate(1, 0, 0))) {
			eventsForYear = append(eventsForYear, event)
		}
	}

	if len(eventsForYear) == 0 {
		return nil, http.StatusServiceUnavailable
	}

	return &eventsForYear, http.StatusOK
}

// обработка ошибки 400
func BadRequest(UserIdKey, DateVal string) (int, time.Time, int) {
	id, err := strconv.Atoi(UserIdKey)
	if err != nil {
		return 0, time.Time{}, http.StatusBadRequest
	}

	if id <= 0 {
		return 0, time.Time{}, http.StatusBadRequest
	}

	dateParse, err := time.Parse("2006-01-02", DateVal)
	if err != nil {
		return 0, time.Time{}, http.StatusBadRequest
	}

	return id, dateParse, http.StatusOK
}

// генерация уникального id
func getUniqCode() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid, nil
}
