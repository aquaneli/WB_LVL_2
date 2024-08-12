package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Events struct {
	UserId int       `json:"user_id"`
	Date   time.Time `json:"date"`
}

func main() {

	http.HandleFunc("/create_event", HandlerCreateEvent)

	http.HandleFunc("/events_for_day", HandlerEventsForDay)

	http.ListenAndServe(":8080", nil)
}

func HandlerEventsForDay(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// fmt.Println(r.URL.Scheme, "|", r.URL.String())
	// fmt.Println(r.URL.Query(), r.Header)
}

type RequestErr struct {
	Code string `json:"error"`
}

func HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resMarshal, err := json.Marshal(RequestErr{Code: "HTTP 503"})
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(resMarshal))
	} else {
		//парсинг GET запросов
		// for key, val := range r.URL.Query() {
		// 	fmt.Println(key, val)
		// }

		//содержимое тела запроса в виде строки
		// ns := bufio.NewScanner(r.Body)
		// defer r.Body.Close()
		// for ns.Scan() {
		// 	fmt.Println(ns.Text())
		// }
		// fmt.Println("finish")

		//парсинг POST запроса
		r.ParseForm()
		for key, val := range r.Form {
			for _, value := range val {
				fmt.Fprintf(w, "Key: %s, Value: %s\n", key, value)

			}
		}
	}

}
