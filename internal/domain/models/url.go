package models

import "time"

type URL struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	OriginalURL string    `json:"original_url" bson:"original_url"`
	ShortCode   string    `json:"short_code" bson:"short_code"`
	AccessCount int       `json:"access_count" bson:"access_count"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
