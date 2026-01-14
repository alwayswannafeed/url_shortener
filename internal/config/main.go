package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type AppConfig struct {
	BaseURL string `fig:"base_url"`
}

type Config interface {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer
	Copuser() types.Copuser
	AppConfig() *AppConfig
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer
	getter  kv.Getter
	copuser types.Copuser
	appConfig *AppConfig
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Databaser:  pgdb.NewDatabaser(getter),
		Listenerer: comfig.NewListenerer(getter),
		copuser:    copus.NewCopuser(getter),
	}
}

func (c *config) Copuser() types.Copuser {
	return c.copuser
}

func (c *config) AppConfig() *AppConfig {
	if c.appConfig != nil {
		return c.appConfig
	}

	cfg := AppConfig{
		BaseURL: "http://localhost:8000",
	}

	raw, err := c.getter.GetStringMap("app")
	if err != nil {
		c.appConfig = &cfg
		return c.appConfig
	}

	err = figure.Out(&cfg).From(raw).Please()
	if err != nil {
		panic(err)
	}

	c.appConfig = &cfg
	return c.appConfig
}