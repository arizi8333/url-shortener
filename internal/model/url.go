package model

import "time"

type URL struct {
	ID          int
	OriginalURL string
	ShortCode   string
	Clicks      int
	CreatedAt   time.Time
	ExpiredAt   *time.Time
}
