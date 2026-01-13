package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alwayswannafeed/url_shortener/internal/data"
	"github.com/go-chi/chi"
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

func (s *service) createURL(w http.ResponseWriter, r *http.Request) {
	var request CreateURLRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	hash := time.Now().Format("150405")
	
	newURL := data.URL{
		Hash:        hash,
		OriginalURL: 	request.OriginalURL,
		CreatedAt:   time.Now(),
	}

	err := s.urlsQ.Insert(newURL)
	if err != nil {
		s.log.WithError(err).Error("failed to save url")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := URLResponse{
		ShortURL:    "http://localhost/" + hash,
		OriginalURL: newURL.OriginalURL,
	}

	ape.Render(w, response)
}

func (s *service) getURL(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hash")

	url, err := s.urlsQ.Get(hash)
	if err != nil {
		s.log.WithError(err).Error("failed to get url")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if url == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, URLResponse{
		ShortURL:    "http://localhost:8000/" + url.Hash,
		OriginalURL: url.OriginalURL,
	})
}