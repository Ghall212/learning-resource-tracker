package main

import (
	"os"

	"github.com/jacobtie/learning-resource-tracker/pkg/server"
	"github.com/jacobtie/learning-resource-tracker/pkg/storage"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	dbService := &storage.MySQLService{ConnectionString: connectionString}
	db, err := dbService.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	server.Run(db)
}
