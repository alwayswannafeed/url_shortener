package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
    "github.com/alwayswannafeed/url_shortener/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
        ape.CtxMiddleware(
            handlers.CtxLog(s.log),
        ),
	)

	r.Route("/urls", func(r chi.Router) {
		r.Post("/", handlers.CreateURL(s.cfg))
	})

	r.Get("/{hash}", handlers.GetURL(s.cfg))

	return r
}