package service

import (
	"net"
	"net/http"

	"github.com/alwayswannafeed/url_shortener/internal/config"
	"github.com/alwayswannafeed/url_shortener/internal/data"
	"github.com/alwayswannafeed/url_shortener/internal/data/pg"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	urlsQ    data.URLQ
}

func (s *service) Run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		urlsQ: pg.NewURLsQ(cfg.DB()),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).Run(); err != nil {
		panic(err)
	}
}