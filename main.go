package main

import (
	"log"
	"net/http"
	"shortener/db"
	"shortener/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализация баз данных
	db.InitPostgres()
	db.InitRedis()

	// Инициализация роутера
	router := mux.NewRouter()

	// Роуты
	router.HandleFunc("/shorten", handlers.ShortenHandler).Methods("POST")
	router.HandleFunc("/{id}", handlers.RedirectHandler).Methods("GET")
	router.HandleFunc("/delete/{id}", handlers.DeleteHandler).Methods("DELETE")
	router.HandleFunc("/stats/{id}", handlers.StatsHandler).Methods("GET")

	log.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
