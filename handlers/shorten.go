package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"shortener/db"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		URL      string `json:"url"`    // Исходный URL
		Expiry   string `json:"expiry"` // Срок действия
		CustomID string `json:"custom_id"`
	}
	type Response struct {
		ShortURL string `json:"short_url"` // Короткая ссылка
	}

	var req Request
	// Парсинг тела запроса
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || !strings.HasPrefix(req.URL, "http") {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Парсинг даты истечения срока действия
	var expiry time.Time
	if req.Expiry != "" {
		var err error
		expiry, err = time.Parse(time.RFC3339, req.Expiry)
		if err != nil {
			http.Error(w, "Invalid expiry format", http.StatusBadRequest)
			return
		}
	}

	// Если не указан ID, то генерируем случайный
	shortID := req.CustomID
	if shortID == "" {
		shortID = generateID()
	}

	// Проверка существования ID в бд
	var exists int
	err := db.GetDB().QueryRow("SELECT COUNT(*) FROM short_links WHERE id = $1", shortID).Scan(&exists)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	if exists > 0 {
		http.Error(w, "ID already exists", http.StatusConflict)
		return
	}

	// Сохраняем ссылку
	if err := saveLinkToStorage(shortID, req.URL, expiry); err != nil {
		http.Error(w, "Failed to save link", http.StatusInternalServerError)
		return
	}

	// Отправляем в ответе короткий URL
	json.NewEncoder(w).Encode(Response{ShortURL: "http://localhost:8080/" + shortID})
}

// Генерация случайного ID
func generateID() string {
	symbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	s := make([]byte, 5)
	for i := range s {
		s[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(s)
}

// Сохраняет ссылку
func saveLinkToStorage(shortID, Url string, expiry time.Time) error {
	ctx := db.GetContext()
	_, err := db.GetDB().Exec("INSERT INTO short_links (id, long_url, expiry) VALUES ($1, $2, $3)", shortID, Url, expiry)
	if err != nil {
		return err
	}

	// Вычисляем TTL (если указан срок действия)
	ttl := time.Until(expiry)
	if ttl > 0 {
		db.GetRedisClient().Set(ctx, shortID, Url, ttl) // Сохраняем в Redis с TTL
	} else {
		db.GetRedisClient().Set(ctx, shortID, Url, 0) // Без TTL
	}

	return nil
}
