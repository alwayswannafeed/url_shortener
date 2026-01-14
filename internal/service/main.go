package service

import (
	"net"
	"net/http"
	"github.com/alwayswannafeed/url_shortener/internal/config"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
)

type service struct {
	log      *logan.Entry
	copus    types.Copuser
	listener net.Listener
	cfg      config.Config
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copuser(),
		listener: cfg.Listener(),
		cfg:      cfg, 
	}
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()
	return http.Serve(s.listener, r)
}