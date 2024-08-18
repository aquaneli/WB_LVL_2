package middleware

import (
	"log"
	"net/http"
)

// Logging логирует все входящие запросы
func Logging(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request:[%s] | method request:[%s]", r.URL.String(), r.Method)
		handler.ServeHTTP(w, r)
	}
}
