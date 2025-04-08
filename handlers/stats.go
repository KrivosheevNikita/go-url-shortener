package handlers

import (
	"encoding/json"
	"net/http"

	"shortener/db"

	"shortener/models"

	"github.com/gorilla/mux"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	shortID := mux.Vars(r)["id"]

	row := db.GetDB().QueryRow("SELECT long_url, expiry, usage_count FROM short_links WHERE id = $1", shortID)
	var info models.URLInfo
	if err := row.Scan(&info.Url, &info.Expiry, &info.UsageCount); err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
