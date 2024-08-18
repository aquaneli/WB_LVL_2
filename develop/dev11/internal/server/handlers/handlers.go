package handlers

import (
	"dev11/internal/storage"
	"net/http"
	"strconv"
)

// EventHandlers структура в которой хранятся все данные для обработки их handlers
type EventHandlers struct {
	s storage.Storage
}

// NewEventHandler cоздает новый обработчик с возможностью сохранять данные
func NewEventHandler() *EventHandlers {
	return &EventHandlers{*storage.NewStorge()}
}

// HandlerCreateEvent создает новый эвент и сохраняет в памяти
func (eh *EventHandlers) HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ObjectSerialization(w, http.StatusBadRequest, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, http.StatusInternalServerError, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
	}

	id := r.FormValue("user_id")
	date := r.FormValue("date")

	uuid, status := eh.s.Add(id, date)
	if status != http.StatusOK {
		ObjectSerialization(w, status, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	ObjectSerialization(w, http.StatusOK, ResponseResult{Info: Information{UserID: id, Date: date, EventID: []string{uuid}, Status: "event added"}})
}

// HandlerUpdateEvent меняет дату на указанную определнного эвента при указании его uuid
func (eh *EventHandlers) HandlerUpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ObjectSerialization(w, http.StatusBadRequest, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, http.StatusInternalServerError, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
		return
	}

	id := r.FormValue("user_id")
	date := r.FormValue("date")
	uuid := r.FormValue("uuid")

	status := eh.s.UpDate(id, date, uuid)
	if status != http.StatusOK {
		ObjectSerialization(w, status, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	ObjectSerialization(w, http.StatusOK, ResponseResult{Info: Information{UserID: id, Date: date, EventID: []string{uuid}, Status: "event date updated"}})
}

// HandlerDeleteEvent удаляет эвент по uuid
func (eh *EventHandlers) HandlerDeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ObjectSerialization(w, http.StatusBadRequest, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, http.StatusInternalServerError, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
		return
	}

	id := r.FormValue("user_id")
	date := r.FormValue("date")
	uuid := r.FormValue("uuid")

	status := eh.s.Remove(id, date, uuid)
	if status != http.StatusOK {
		ObjectSerialization(w, status, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	ObjectSerialization(w, http.StatusOK, ResponseResult{Info: Information{UserID: id, Date: date, EventID: []string{uuid}, Status: "event deleted"}})
}

// HandlerEventsForDay выводит все эвенты в определнный день определнного user
func (eh *EventHandlers) HandlerEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ObjectSerialization(w, http.StatusBadRequest, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, http.StatusInternalServerError, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
		return
	}

	id := r.FormValue("user_id")
	date := r.FormValue("date")

	e, status := eh.s.GetEventsForDay(id, date)
	if status != http.StatusOK {
		ObjectSerialization(w, status, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	size := len(e.UniqCodeEvents)
	uuid := make([]string, 0, size)

	for k := range e.UniqCodeEvents {
		uuid = append(uuid, k)
	}

	ObjectSerialization(w, http.StatusOK, ResponseResult{Info: Information{UserID: id, Date: date, EventID: uuid, Status: "successfully"}})
}

// HandlerEventsForWeek выводит все эвенты за неделю с текущего дня определнного user
func (eh *EventHandlers) HandlerEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ObjectSerialization(w, http.StatusBadRequest, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, http.StatusInternalServerError, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
		return
	}
	id := r.FormValue("user_id")
	date := r.FormValue("date")

	e, status := eh.s.GetEventsForWeek(id, date)
	if status != http.StatusOK {
		ObjectSerialization(w, status, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	uuid := []string{}
	for _, val := range *e {
		for key := range val.UniqCodeEvents {
			uuid = append(uuid, key)
		}
	}

	ObjectSerialization(w, http.StatusOK, ResponseResult{Info: Information{UserID: id, Date: date, EventID: uuid, Status: "successfully"}})
}

// HandlerEventsForYear выводит все эвенты за год с текущего дня определнного user
func (eh *EventHandlers) HandlerEventsForYear(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ObjectSerialization(w, http.StatusBadRequest, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}
	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, http.StatusInternalServerError, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
		return
	}
	id := r.FormValue("user_id")
	date := r.FormValue("date")

	e, status := eh.s.GetEventsForYear(id, date)
	if status != http.StatusOK {
		ObjectSerialization(w, status, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	uuid := []string{}
	for _, val := range *e {
		for key := range val.UniqCodeEvents {
			uuid = append(uuid, key)
		}
	}

	ObjectSerialization(w, http.StatusOK, ResponseResult{Info: Information{UserID: id, Date: date, EventID: uuid, Status: "successfully"}})
}
