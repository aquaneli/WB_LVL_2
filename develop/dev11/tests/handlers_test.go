package tests

import (
	"dev11/internal/server/handlers"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func createOrAddEvents(EventHandler *handlers.EventHandlers, createURL string, method string) *httptest.ResponseRecorder {
	reqCreate, err := http.NewRequest(method, createURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	reqCreate.ParseForm()
	rrCreate := httptest.NewRecorder()
	EventHandler.HandlerCreateEvent(rrCreate, reqCreate)
	return rrCreate
}

func TestHandlerCreateEvent(t *testing.T) {
	eventHandler := handlers.NewEventHandler()
	createURL := "http://localhost:8080/create_event?user_id=1&date=2010-05-15"
	rrCreate := createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}
}

func TestHandlerUpdateEvent(t *testing.T) {
	eventHandler := handlers.NewEventHandler()
	createURL := "http://localhost:8080/create_event?user_id=1&date=2010-05-15"
	rrCreate := createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	data, err := io.ReadAll(rrCreate.Body)
	if err != nil {
		t.Fatal(err)
	}
	h := handlers.ResponseResult{Info: handlers.Information{}}

	err = json.Unmarshal(data, &h)
	if err != nil {
		log.Fatal(err)
	}

	rrUpdate := httptest.NewRecorder()

	updateURL := "http://localhost:8080/update_event?user_id=1&date=2010-05-12&uuid=" + h.Info.EventID[0]
	reqUpdate, err := http.NewRequest("POST", updateURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	reqUpdate.ParseForm()

	eventHandler.HandlerUpdateEvent(rrUpdate, reqUpdate)
	if !reflect.DeepEqual(rrUpdate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %q but got %q", http.StatusOK, rrUpdate.Result().StatusCode)
	}

}

func TestHandlerDeleteEvent(t *testing.T) {
	eventHandler := handlers.NewEventHandler()
	createURL := "http://localhost:8080/create_event?user_id=1&date=2010-05-15"
	rrCreate := createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	data, err := io.ReadAll(rrCreate.Body)
	if err != nil {
		t.Fatal(err)
	}
	h := handlers.ResponseResult{Info: handlers.Information{}}

	err = json.Unmarshal(data, &h)
	if err != nil {
		log.Fatal(err)
	}

	rrDelete := httptest.NewRecorder()

	deleteURL := "http://localhost:8080/update_event?user_id=1&date=2010-05-12&uuid=" + h.Info.EventID[0]
	reqDelete, err := http.NewRequest("POST", deleteURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	reqDelete.ParseForm()

	eventHandler.HandlerCreateEvent(rrDelete, reqDelete)
	if !reflect.DeepEqual(rrDelete.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %q but got %q", http.StatusOK, rrDelete.Result().StatusCode)
	}

}

func TestHandlerEventsForDay(t *testing.T) {

	eventHandler := handlers.NewEventHandler()
	createURL := "http://localhost:8080/create_event?user_id=1&date=2010-05-15"
	rrCreate := createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	createURL = "http://localhost:8080/create_event?user_id=1&date=2010-05-15"
	rrCreate = createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	createURL = "http://localhost:8080/create_event?user_id=1&date=2010-05-15"
	rrCreate = createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	rrGetEvents := httptest.NewRecorder()

	getEventsURL := "http://localhost:8080/events_for_day?user_id=1&date=2010-05-15"
	req, err := http.NewRequest("GET", getEventsURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.ParseForm()

	eventHandler.HandlerEventsForDay(rrGetEvents, req)

	if !reflect.DeepEqual(rrGetEvents.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %q but got %q", http.StatusOK, rrGetEvents.Result().StatusCode)
	}

}

func TestHandlerEventsForWeek(t *testing.T) {
	eventHandler := handlers.NewEventHandler()
	createURL := "http://localhost:8080/create_event?user_id=1&date=2010-05-15"
	rrCreate := createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	createURL = "http://localhost:8080/create_event?user_id=1&date=2010-05-23"
	rrCreate = createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	createURL = "http://localhost:8080/create_event?user_id=1&date=2010-05-20"
	rrCreate = createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	rrGetEvents := httptest.NewRecorder()

	getEventsURL := "http://localhost:8080/events_for_week?user_id=1&date=2010-05-15"
	req, err := http.NewRequest("GET", getEventsURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.ParseForm()

	eventHandler.HandlerEventsForWeek(rrGetEvents, req)

	if !reflect.DeepEqual(rrGetEvents.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %q but got %q", http.StatusOK, rrGetEvents.Result().StatusCode)
	}

	data, err := io.ReadAll(rrGetEvents.Body)
	if err != nil {
		t.Fatal(err)
	}
	h := handlers.ResponseResult{Info: handlers.Information{}}

	err = json.Unmarshal(data, &h)
	if err != nil {
		log.Fatal(err)
	}

	if len(h.Info.EventID) != 2 {
		t.Errorf("expected len %q but got %q", 2, len(h.Info.EventID))
	}

}

func TestHandlerEventsForYear(t *testing.T) {
	eventHandler := handlers.NewEventHandler()
	createURL := "http://localhost:8080/create_event?user_id=1&date=2010-05-15"
	rrCreate := createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	createURL = "http://localhost:8080/create_event?user_id=1&date=2011-03-23"
	rrCreate = createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	createURL = "http://localhost:8080/create_event?user_id=1&date=2011-02-20"
	rrCreate = createOrAddEvents(eventHandler, createURL, "POST")

	if !reflect.DeepEqual(rrCreate.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %d but got %d", http.StatusOK, rrCreate.Result().StatusCode)
	}

	rrGetEvents := httptest.NewRecorder()

	getEventsURL := "http://localhost:8080/events_for_year?user_id=1&date=2010-05-15"
	req, err := http.NewRequest("GET", getEventsURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.ParseForm()

	eventHandler.HandlerEventsForYear(rrGetEvents, req)

	if !reflect.DeepEqual(rrGetEvents.Result().StatusCode, http.StatusOK) {
		t.Errorf("expected status %q but got %q", http.StatusOK, rrGetEvents.Result().StatusCode)
	}

	data, err := io.ReadAll(rrGetEvents.Body)
	if err != nil {
		t.Fatal(err)
	}
	h := handlers.ResponseResult{Info: handlers.Information{}}

	err = json.Unmarshal(data, &h)
	if err != nil {
		log.Fatal(err)
	}

	if len(h.Info.EventID) != 3 {
		t.Errorf("expected len %d but got %d", 3, len(h.Info.EventID))
	}

}
