package handlers

import (
	"log"
	"net/http"
	"time"

	"shortener/db"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortID := mux.Vars(r)["id"]
	ctx := db.GetContext()
	// Пытаемся получить из Redis
	Url, err := db.GetRedisClient().Get(ctx, shortID).Result()

	// Если в Redis нет данных, или произошла ошибка, то ищем в базе данных
	if err != nil || err == redis.Nil {
		row := db.GetDB().QueryRow("SELECT long_url FROM short_links WHERE id = $1 AND (expiry IS NULL OR expiry > NOW())", shortID)
		if err := row.Scan(&Url); err != nil {
			http.NotFound(w, r)
			return
		}

		// Если данные найдены, то кэшируем их в Redis
		if err := db.GetRedisClient().Set(ctx, shortID, Url, time.Hour*24).Err(); err != nil {
			log.Println("Redis cache update failed:", err)
		}
	}

	// Обновляем счётчик переходов
	_, err = db.GetDB().Exec("UPDATE short_links SET usage_count = usage_count + 1 WHERE id = $1", shortID)
	if err != nil {
		http.Error(w, "Failed to update usage count", http.StatusInternalServerError)
		return
	}

	// Перенаправление на оригинальный URL
	http.Redirect(w, r, Url, http.StatusFound)
}
