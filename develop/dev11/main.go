package main

import (
	storage "dev11/internal/storage"
	"encoding/json"
	"log"
	"net/http"
)

type EventHandlers struct {
	s storage.Storage
}

type RequestErr struct {
	Code string `json:"error"`
}

func main() {

	eh := EventHandlers{s: *storage.NewStorge()}

	http.HandleFunc("/create_event", eh.HandlerCreateEvent)
	http.HandleFunc("/update_event", eh.HandlerUpdateEvent)
	http.HandleFunc("/delete_event", eh.HandlerDeleteEvent)

	http.HandleFunc("/events_for_day", eh.HandlerEventsForDay)
	http.HandleFunc("/events_for_week", eh.HandlerEventsForWeek)
	http.HandleFunc("/events_for_year", eh.HandlerEventsForYear)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func (eh *EventHandlers) HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resMarshal, err := json.Marshal(RequestErr{Code: "HTTP 503"})
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(resMarshal))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		id := r.FormValue("user_id")
		date := r.FormValue("date")
		uuid, err := eh.s.Add(id, date)
		if err != nil {
			log.Fatal(err)
		}

		err = ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: []string{uuid}, Status: "event added"}})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (eh *EventHandlers) HandlerUpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resMarshal, err := json.Marshal(RequestErr{Code: "HTTP 503"})
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(resMarshal))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		id := r.FormValue("user_id")
		date := r.FormValue("date")
		uuid := r.FormValue("uuid")

		err = eh.s.UpDate(id, date, uuid)
		if err != nil {
			log.Fatal(err)
		}

		err = ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: []string{uuid}, Status: "event date updated"}})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (eh *EventHandlers) HandlerDeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resMarshal, err := json.Marshal(RequestErr{Code: "HTTP 503"})
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(resMarshal))
	} else {
		r.ParseForm()
		id := r.FormValue("user_id")
		date := r.FormValue("date")
		uuid := r.FormValue("uuid")
		err := eh.s.Remove(id, date, uuid)
		if err != nil {
			log.Fatal(err)
		}

		err = ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: []string{uuid}, Status: "event deleted"}})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (eh *EventHandlers) HandlerEventsForDay(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("user_id")
	date := r.FormValue("date")

	e, err := eh.s.GetEventsForDay(id, date)
	if err != nil {
		log.Fatal(err)
	}

	size := len(e.UniqCodeEvents)
	uuid := make([]string, 0, size)

	for k := range e.UniqCodeEvents {
		uuid = append(uuid, k)
	}
	err = ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: uuid, Status: "successfully"}})
	if err != nil {
		log.Fatal(err)
	}
}

func (eh *EventHandlers) HandlerEventsForWeek(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("user_id")
	date := r.FormValue("date")

	e, err := eh.s.GetEventsForWeek(id, date)
	if err != nil {
		log.Fatal(err)
	}

	uuid := []string{}
	for _, val := range *e {
		for key := range val.UniqCodeEvents {
			uuid = append(uuid, key)
		}
	}
	err = ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: uuid, Status: "successfully"}})
	if err != nil {
		log.Fatal(err)
	}
}

func (eh *EventHandlers) HandlerEventsForYear(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("user_id")
	date := r.FormValue("date")

	e, err := eh.s.GetEventsForYear(id, date)
	if err != nil {
		log.Fatal(err)
	}

	uuid := []string{}
	for _, val := range *e {
		for key := range val.UniqCodeEvents {
			uuid = append(uuid, key)
		}
	}
	err = ObjectSerialization(w, ResponseResult{Info: Information{UserId: id, Date: date, EventId: uuid, Status: "successfully"}})
	if err != nil {
		log.Fatal(err)
	}
}
