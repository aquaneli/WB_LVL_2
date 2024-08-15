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

func ObjectSerialization(w http.ResponseWriter, msg interface{}) {
	res, err := json.Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	_, err = w.Write(res)
	if err != nil {
		log.Fatalln(err)
	}
}
