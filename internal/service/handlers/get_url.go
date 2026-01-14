package handlers

import (
	"net/http"
	"github.com/alwayswannafeed/url_shortener/internal/data/pg"
	"github.com/alwayswannafeed/url_shortener/internal/config"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetURL(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := chi.URLParam(r, "hash")

		urlEntry, err := pg.NewURLsQ(cfg.DB()).Get(hash)
        
		if err != nil {
			Log(r).WithError(err).Error("failed to get url")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		if urlEntry == nil {
			ape.RenderErr(w, problems.NotFound())
			return
		}

		http.Redirect(w, r, urlEntry.OriginalURL, http.StatusFound)
	}
}