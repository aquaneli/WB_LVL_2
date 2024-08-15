package main

import (
	"dev11/config"
	"dev11/internal/server/handlers"
	"dev11/internal/server/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.ReadConfig("../config/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	eh := handlers.NewEventHandler()

	http.HandleFunc("/create_event", middleware.Logging(eh.HandlerCreateEvent))
	http.HandleFunc("/update_event", middleware.Logging(eh.HandlerUpdateEvent))
	http.HandleFunc("/delete_event", middleware.Logging(eh.HandlerDeleteEvent))

	http.HandleFunc("/events_for_day", middleware.Logging(eh.HandlerEventsForDay))
	http.HandleFunc("/events_for_week", middleware.Logging(eh.HandlerEventsForWeek))
	http.HandleFunc("/events_for_year", middleware.Logging(eh.HandlerEventsForYear))

	address := fmt.Sprintf("%s:%s", cfg.Server.IP, cfg.Server.Port)

	err = http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}

}
