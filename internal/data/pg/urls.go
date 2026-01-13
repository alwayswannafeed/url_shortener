package pg

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alwayswannafeed/url_shortener/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const urlsTable = "urls"

type urlsQ struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func NewURLsQ(db *pgdb.DB) data.URLQ {
	return &urlsQ{
		db:  db,
		sql: squirrel.Select(fmt.Sprintf("%s.*", urlsTable)).From(urlsTable),
	}
}

func (q *urlsQ) New() data.URLQ {
	return NewURLsQ(q.db.Clone())
}

func (q *urlsQ) Insert(url data.URL) error {
	stmt := squirrel.Insert(urlsTable).SetMap(map[string]interface{}{
		"hash":         url.Hash,
		"original_url": url.OriginalURL,
		"created_at":   url.CreatedAt,
	})
	return q.db.Exec(stmt)
}

func (q *urlsQ) Get(hash string) (*data.URL, error) {
	var result data.URL
	stmt := q.sql.Where(squirrel.Eq{"hash": hash})
	err := q.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &result, err
}