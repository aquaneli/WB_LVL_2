package handlers

import (
	"dev11/internal/storage"
	"net/http"
	"strconv"
)

type EventHandlers struct {
	s storage.Storage
}

func NewEventHandler() *EventHandlers {
	return &EventHandlers{*storage.NewStorge()}
}

func (eh *EventHandlers) HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
	}

	id := r.FormValue("user_id")
	date := r.FormValue("date")

	uuid, status := eh.s.Add(id, date)
	if status != http.StatusOK {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: []string{uuid}, Status: "event added"}})
}

func (eh *EventHandlers) HandlerUpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
		return
	}

	id := r.FormValue("user_id")
	date := r.FormValue("date")
	uuid := r.FormValue("uuid")

	status := eh.s.UpDate(id, date, uuid)
	if status != http.StatusOK {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: []string{uuid}, Status: "event date updated"}})

}

func (eh *EventHandlers) HandlerDeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
		return
	}

	id := r.FormValue("user_id")
	date := r.FormValue("date")
	uuid := r.FormValue("uuid")

	status := eh.s.Remove(id, date, uuid)
	if status != http.StatusOK {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: []string{uuid}, Status: "event deleted"}})
}

func (eh *EventHandlers) HandlerEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
		return
	}

	id := r.FormValue("user_id")
	date := r.FormValue("date")

	e, status := eh.s.GetEventsForDay(id, date)
	if status != http.StatusOK {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	size := len(e.UniqCodeEvents)
	uuid := make([]string, 0, size)

	for k := range e.UniqCodeEvents {
		uuid = append(uuid, k)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: uuid, Status: "successfully"}})
}

func (eh *EventHandlers) HandlerEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}

	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
		return
	}
	id := r.FormValue("user_id")
	date := r.FormValue("date")

	e, status := eh.s.GetEventsForWeek(id, date)
	if status != http.StatusOK {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	uuid := []string{}
	for _, val := range *e {
		for key := range val.UniqCodeEvents {
			uuid = append(uuid, key)
		}
	}
	ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: uuid, Status: "successfully"}})
}

func (eh *EventHandlers) HandlerEventsForYear(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusBadRequest)})
		return
	}
	err := r.ParseForm()
	if err != nil {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(http.StatusInternalServerError)})
		return
	}
	id := r.FormValue("user_id")
	date := r.FormValue("date")

	e, status := eh.s.GetEventsForYear(id, date)
	if status != http.StatusOK {
		ObjectSerialization(w, Err{Result: "HTTP " + strconv.Itoa(status)})
		return
	}

	uuid := []string{}
	for _, val := range *e {
		for key := range val.UniqCodeEvents {
			uuid = append(uuid, key)
		}
	}
	ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: uuid, Status: "successfully"}})
}
