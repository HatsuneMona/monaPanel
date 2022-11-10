package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"monaPanel/config"
)

var (
	Config      config.MonaPanelConfig
	Log         *zap.Logger
	ConfigViper *viper.Viper
	DB          *gorm.DB
	Redis       *redis.Client
)
