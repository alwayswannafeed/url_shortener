package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"
	"github.com/alwayswannafeed/url_shortener/internal/data/pg"
	"github.com/alwayswannafeed/url_shortener/internal/config"
	"github.com/alwayswannafeed/url_shortener/internal/data"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

type CreateURLRequest struct {
	OriginalURL string `json:"original_url"`
}

type URLResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func CreateURL(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request CreateURLRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		hasher := sha256.New()
		hasher.Write([]byte(request.OriginalURL))
		fullHash := hex.EncodeToString(hasher.Sum(nil))
		shortHash := fullHash[:8]

		newURL := data.URL{
			Hash:        shortHash,
			OriginalURL: request.OriginalURL,
			CreatedAt:   time.Now(),
		}

		err := pg.NewURLsQ(cfg.DB()).Insert(newURL)
		if err != nil {
			Log(r).WithError(err).Error("failed to save url")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		response := URLResponse{
			ShortURL:    cfg.AppConfig().BaseURL + shortHash,
			OriginalURL: newURL.OriginalURL,
		}

		ape.Render(w, response)
	}
}