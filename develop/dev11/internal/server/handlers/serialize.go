package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseResult struct {
	Info Information `json:"result"`
}

type Information struct {
	UserId  string   `json:"user_id"`
	Date    string   `json:"date"`
	EventId []string `json:"event_uuid"`
	Status  string   `json:"status"`
}

type Err struct {
	Result string `json:"error"`
}

func ObjectSerialization(w http.ResponseWriter, status int, msg interface{}) {
	res, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(res)
	if err != nil {
		log.Fatalln(err)
	}
}
