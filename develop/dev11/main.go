package main

import (
	"net/http"
	"time"
)

type Events struct {
	UserId int       `json:"user_id"`
	Date   time.Time `json:"date"`
}

func HandlerEventsForDay(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "GET"{

	// }
	w.Write([]byte("Good job"))
}

func HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "POST"
}

func main() {

	http.HandleFunc("/create_event", HandlerCreateEvent)

	http.HandleFunc("/events_for_day", HandlerEventsForDay)

	http.ListenAndServe(":8080", nil)
}
