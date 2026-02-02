package main

import (
	"log"
	"net/http"

	cfg "github.com/mahmud-off/crudUsers/internal/config/config"
	"github.com/mahmud-off/crudUsers/internal/server/server"
	"github.com/mahmud-off/crudUsers/internal/services/database"
)

func main() {

	cfg, err := cfg.NewConfig()

	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := database.NewDB(cfg.GetDB(), cfg.GetDBSource())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := server.NewServer(db)

	http.HandleFunc("/user", server.LoggingMiddleware(server.UserHandler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
