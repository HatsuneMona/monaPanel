package conf

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var GlobalConf *viper.Viper

//go:embed simpleConfig.yaml
var simpleConfig []byte

func init() {
	// 初始化 viper 库
	GlobalConf = viper.New()

	// 指定 yaml 文件类型
	GlobalConf.SetConfigType("yaml")
}

func ReadConfig(filePath string) {

	GlobalConf.SetConfigFile(filePath)
	if err := GlobalConf.ReadInConfig(); err != nil {
		panic(err)
	}

}

func WriteDefaultConfig(filePath string) error {

	path, _ := filepath.Split(filePath)
	err := os.MkdirAll(path, 0775)
	if err != nil {
		return errors.New(fmt.Sprintf("can not mkdir dir:%v err:%v", path, err))
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0775)
	if err != nil {
		return errors.New(fmt.Sprintf("can not open config err:%v", err))
	}
	defer file.Close()

	_, err = file.Write(simpleConfig)
	if err != nil {
		return errors.New(fmt.Sprintf("file white faild! err:%v", err))
	}

	return nil
}
