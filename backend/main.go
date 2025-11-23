package main

import (
	"backend/internal/db"
	"backend/internal/models"
	"backend/internal/router"
	"log"
	"net/http"
)

func main() {
	if err := db.InitGORM(); err != nil {
		log.Fatal("Database initialization failed: %w", err)
	}

	db.AutoMigrate(
		&models.Task{},
	)

	defer db.CloseDB()

	r := router.NewRouter()
	http.ListenAndServe(":8080", r)
}
