package main

import (
	"dev11/internal/server/handlers"
	"dev11/internal/server/middleware"
	"log"
	"net/http"
)

func main() {

	eh := handlers.NewEventHandler()

	http.HandleFunc("/create_event", middleware.Logging(eh.HandlerCreateEvent))
	http.HandleFunc("/update_event", middleware.Logging(eh.HandlerUpdateEvent))
	http.HandleFunc("/delete_event", middleware.Logging(eh.HandlerDeleteEvent))

	http.HandleFunc("/events_for_day", middleware.Logging(eh.HandlerEventsForDay))
	http.HandleFunc("/events_for_week", middleware.Logging(eh.HandlerEventsForWeek))
	http.HandleFunc("/events_for_year", middleware.Logging(eh.HandlerEventsForYear))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
