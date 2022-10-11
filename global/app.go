package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"monaPanel/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.MonaPanelConfig
	Log         *zap.Logger
}

// var App = new(Application)
