package handlers

import (
	"net/http"

	"shortener/db"

	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := db.GetContext()
	shortID := mux.Vars(r)["id"]

	// Удаляем из Postgres
	res, err := db.GetDB().Exec("DELETE FROM short_links WHERE id = $1", shortID)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	// Проверка, была ли строка удалена
	affected, _ := res.RowsAffected()
	if affected == 0 {
		// Если нет, то возвращаем 404
		http.NotFound(w, r)
		return
	}

	// Удаляем из Redis
	db.GetRedisClient().Del(ctx, shortID)
	w.WriteHeader(http.StatusNoContent)
}
