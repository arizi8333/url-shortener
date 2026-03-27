package repository

import (
	"database/sql"
	"url-shortener/internal/model"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

// Save inserts a new URL record into the database and updates the URL struct with the generated ID and CreatedAt timestamp.
func (r *URLRepository) Save(url *model.URL) error {
	query := `
	INSERT INTO urls (original_url, short_code, clicks, expired_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
	`

	return r.db.QueryRow(
		query,
		url.OriginalURL,
		url.ShortCode,
		url.Clicks,
		url.ExpiredAt,
	).Scan(&url.ID, &url.CreatedAt)
}

// FindByCode retrieves a URL record from the database based on the short code.
func (r *URLRepository) FindByCode(code string) (*model.URL, error) {
	query := `
	SELECT id, original_url, short_code, clicks, created_at, expired_at
	FROM urls
	WHERE short_code=$1
	`

	var url model.URL

	err := r.db.QueryRow(query, code).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortCode,
		&url.Clicks,
		&url.CreatedAt,
		&url.ExpiredAt,
	)

	return &url, err
}

// IncrementClicks increments the click count for a given short code.
func (r *URLRepository) IncrementClicks(code string) {
	query := `UPDATE urls SET clicks = clicks + 1 WHERE short_code=$1`
	r.db.Exec(query, code)
}

// LogClick logs a click event for a given short code, including user agent and IP address.
func (r *URLRepository) LogClick(code, userAgent, ip string) {
	query := `
	INSERT INTO url_clicks (short_code, user_agent, ip_address)
	VALUES ($1, $2, $3)
	`
	r.db.Exec(query, code, userAgent, ip)
}

func (r *URLRepository) GetStats(code string) (int, error) {
	query := `SELECT COUNT(*) FROM url_clicks WHERE short_code=$1`

	var count int
	err := r.db.QueryRow(query, code).Scan(&count)

	return count, err
}
