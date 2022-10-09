package conf

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

var appConf *viper.Viper

func init() {
	// 初始化 viper 库
	appConf = viper.New()

	// 指定 yaml 文件类型
	appConf.SetConfigType("yaml")
}

func ReadConfig(path string) {
	if len(path) == 0 {
		panic(errors.New(fmt.Sprintf("can not read config path:%v", path)))
	}

	appConf.SetConfigFile(path)
	if err := appConf.ReadInConfig(); err != nil {
		panic(err)
	}

}
