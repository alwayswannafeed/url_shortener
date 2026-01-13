package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
	)

	r.Route("/urls", func(r chi.Router) {
		r.Post("/", s.createURL)
	})

	r.Get("/{hash}", s.getURL)

	return r
}