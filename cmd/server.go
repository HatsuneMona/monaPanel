package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"monaPanel/conf"
)

var configPath string
var simpleConfigPath string

var serverRoot = &cobra.Command{
	Use:   "server",
	Short: "A panel to watch&control your servers.",
	// Long:  `A panel to control your servers.`,
	// Run: serverRun,
}
var start = &cobra.Command{
	Use:   "start",
	Short: "start the monaPanel server.",
	Run:   serverStart,
}

var simpleConfig = &cobra.Command{
	Use:   "createSimpleConfig -c configPath",
	Short: "generate a simple config file.",
	Run:   createSimpleConfig,
}

func init() {
	rootCmd.AddCommand(serverRoot)

	serverRoot.AddCommand(start)
	serverRoot.AddCommand(simpleConfig)

	home, err := homedir.Dir()
	if err != nil {
		// TODO log
		panic(err)
	}

	start.Flags().StringVarP(&configPath, "configPath", "c", home+"/.monaPanel/config.yaml", "set the path to mona server config file.")
	simpleConfig.Flags().StringVarP(&simpleConfigPath, "configPath", "c", home+"/.monaPanel/config-simple.yaml", "set the path to mona server simple config file.")
}

// serverStart 启动后端服务
func serverStart(cmd *cobra.Command, args []string) {
	conf.ReadConfig(configPath)
	test := conf.GlobalConf.GetString("server.mysql.schema")
	fmt.Printf("server.mysql.schema = %v", test)
}

// createSimpleConfig 生成参考配置文件
func createSimpleConfig(cmd *cobra.Command, args []string) {
	fmt.Printf("configPath is: %v", simpleConfigPath)
	if err := conf.WriteDefaultConfig(simpleConfigPath); err != nil {
		panic(err)
	}
}
