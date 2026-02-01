package crud_server

import "github.com/mahmud-off/crudUsers/internal/services/database"

type CRUDServer struct {
	DB database.DB
}

func NewCRUDServer(db database.DB) *CRUDServer {
	var newServer = CRUDServer{DB: db}
	return &newServer
}
