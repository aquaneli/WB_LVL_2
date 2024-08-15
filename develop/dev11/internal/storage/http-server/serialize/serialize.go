package serialaze

import (
	"encoding/json"
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

func ObjectSerialization(w http.ResponseWriter, msg ResponseResult) error {
	res, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return nil
}
