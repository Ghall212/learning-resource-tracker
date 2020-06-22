package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Run the application as an HTTP REST server
func Run(db *sql.DB) {
	r := mux.NewRouter()

	mapRoutes(r, db)

	port := getPort()

	log.Printf("Starting server on port %s\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	return port
}
