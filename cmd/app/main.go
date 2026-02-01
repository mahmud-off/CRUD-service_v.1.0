package main

import (
	"log"
	"net/http"

	"github.com/mahmud-off/crudUsers/internal/server/server"
	"github.com/mahmud-off/crudUsers/internal/services/database"
)

func main() {

	db, err := database.NewDB("postgres", "host=127.0.0.1 port=5432 user= dbname= sslmode=disable password=")
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
