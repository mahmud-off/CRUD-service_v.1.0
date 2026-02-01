package crud_server

import (
	"log"
	"net/http"
)

func (s *CRUDServer) LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s | %s \n", r.Method, r.URL, r.Host)
		next(w, r)
	}
}
