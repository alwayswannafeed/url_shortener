-- +migrate Up
CREATE TABLE IF NOT EXISTS urls (
    hash TEXT PRIMARY KEY,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF NOT EXISTS urls;