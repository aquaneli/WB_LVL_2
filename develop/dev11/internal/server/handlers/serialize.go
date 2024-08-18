package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseResult хранит данные о всех эвентах
type ResponseResult struct {
	Info Information `json:"result"`
}

// Information структура в которой хранится информация об эвентах за определенный день и определенного пользователя
type Information struct {
	UserID  string   `json:"user_id"`
	Date    string   `json:"date"`
	EventID []string `json:"event_uuid"`
	Status  string   `json:"status"`
}

// Err структура в которой хранится информация об ошибках
type Err struct {
	Result string `json:"error"`
}

// ObjectSerialization сериализует данные из структуры в байты и отправляет клиенту
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
