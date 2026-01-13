package data

import "time"

type URL struct {
	Hash        string    `db:"hash"`
	OriginalURL string    `db:"original_url"`
	CreatedAt   time.Time `db:"created_at"`
}

type URLQ interface {
	New() URLQ
	Insert(url URL) error
	Get(hash string) (*URL, error)
}