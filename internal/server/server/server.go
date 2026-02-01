package server

import (
	"net/http"

	crud_server "github.com/mahmud-off/crudUsers/internal/server/crud-server"
	"github.com/mahmud-off/crudUsers/internal/services/database"
)

type Server interface {
	UserHandler(w http.ResponseWriter, r *http.Request)
	LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc
}

func NewServer(db database.DB) Server {
	return crud_server.NewCRUDServer(db)
}
