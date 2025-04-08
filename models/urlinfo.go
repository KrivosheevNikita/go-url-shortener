package models

import "time"

// Структура для хранения информации о URL
type URLInfo struct {
	Url        string    `json:"url"`
	Expiry     time.Time `json:"expiry"`
	UsageCount int       `json:"usage_count"`
}
