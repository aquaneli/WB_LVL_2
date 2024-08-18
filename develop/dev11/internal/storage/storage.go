package storage

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Storage cтурктура в которой хранятся информация обо всех эвентах и всех пользователях
type Storage struct {
	Items map[int]map[string]Events
}

// Events хранит информацию об эвентах за день
type Events struct {
	UniqCodeEvents map[string]struct{}
	Date           time.Time
}

// NewStorge cоздает новое хранилище для информации об эвентах
func NewStorge() *Storage {
	return &Storage{Items: make(map[int]map[string]Events)}
}

// Add добавляет данные в структуру
func (s *Storage) Add(UserIDKey, DateVal string) (string, int) {
	id, dateParse, status := BadRequest(UserIDKey, DateVal)
	if status != http.StatusOK {
		return "", status
	}

	//если пользовтеля не существует
	if _, ok := (*s).Items[id]; !ok {
		(*s).Items[id] = make(map[string]Events)
	}

	//если такой даты-ключа не существует
	if _, ok := s.Items[id][DateVal]; !ok {
		s.Items[id][DateVal] = Events{UniqCodeEvents: make(map[string]struct{}), Date: dateParse}
	}

	uuid, err := getUniqCode()
	if err != nil {
		return "", http.StatusInternalServerError
	}

	(*s).Items[id][DateVal].UniqCodeEvents[uuid] = struct{}{}

	return uuid, http.StatusOK
}

// UpDate обновляет дату евента по uuid
func (s *Storage) UpDate(UserIDKey, DateValReplace, CodeEvent string) int {
	id, dateParse, status := BadRequest(UserIDKey, DateValReplace)
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
				(*s).Items[id][DateValReplace] = Events{UniqCodeEvents: make(map[string]struct{}), Date: dateParse}
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

// Remove удаляет эвент
func (s *Storage) Remove(UserIDKey, DateVal, CodeEvent string) int {
	id, _, status := BadRequest(UserIDKey, DateVal)
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

// GetEventsForDay возвращает все эвенты за день определенного пользователя
func (s *Storage) GetEventsForDay(UserIDKey string, DateVal string) (*Events, int) {
	id, _, status := BadRequest(UserIDKey, DateVal)
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

// GetEventsForWeek возвращает все эвенты за неделю определенного пользователя
func (s *Storage) GetEventsForWeek(UserIDKey string, DateVal string) (*[]Events, int) {
	id, startDate, status := BadRequest(UserIDKey, DateVal)
	if status != http.StatusOK {
		return nil, status
	}
	if _, ok := (*s).Items[id]; !ok {
		return nil, http.StatusServiceUnavailable
	}

	eventsForWeek := []Events{}

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

// GetEventsForYear возвращает все эвенты за год определенного пользователя
func (s *Storage) GetEventsForYear(UserIDKey string, DateVal string) (*[]Events, int) {
	id, startDate, status := BadRequest(UserIDKey, DateVal)
	if status != http.StatusOK {
		return nil, status
	}

	if _, ok := (*s).Items[id]; !ok {
		return nil, http.StatusServiceUnavailable
	}

	eventsForYear := []Events{}
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

// BadRequest обрабатывает ошибки 400
func BadRequest(UserIDKey, DateVal string) (int, time.Time, int) {
	id, err := strconv.Atoi(UserIDKey)
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
