package core

import (
	"tsop/internal/pkg/conf"
	"tsop/internal/pkg/logger"
)

const (
	DefaultConfigPath = "./config.yaml"
)

type App struct {
	log  logger.Logger
	conf *conf.Config
}

// NewApp create app
func NewApp() *App {
	config, err := conf.ParseConf(DefaultConfigPath)
	if err != nil {
		panic(err)
	}
	// initial logger
	l, sync, err := logger.NewZap(config)
	if err != nil {
		panic(err)
	}
	defer sync()

	return &App{
		log:  l,
		conf: config,
	}
}

func (a *App) Start() {
	a.log.Info("info")
	a.log.Infof("infof:%s", "Sans")
	a.log.Debug("debug")
	a.log.Debugf("debug:%s", "Sans")
}
