package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Инициализация PostgreSQL
func InitPostgres() {
	pgUrl := os.Getenv("POSTGRES_ADDR")
	if pgUrl == "" {
		log.Fatal("Missing POSTGRES_ADDR")
	}

	var err error
	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal("PostgreSQL connection error: ", err)
	} else {
		log.Println("Connected to PostgreSQL")
	}
}

// Возвращает подключение к базе данных
func GetDB() *sql.DB {
	return db
}
