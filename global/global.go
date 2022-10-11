package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"monaPanel/config"
)

var (
	Config      config.MonaPanelConfig
	Log         *zap.Logger
	ConfigViper *viper.Viper
)
