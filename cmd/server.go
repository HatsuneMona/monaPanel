package cmd

import (
	"github.com/spf13/cobra"
)

var server = &cobra.Command{
	Use:   "server [start] [flags]",
	Short: "A panel to control your servers.",
	// Long:  `A panel to control your servers.`,
	// Run: serverRun,
}

func init() {
	rootCmd.AddCommand(server)
	server.AddCommand(&cobra.Command{
		Run: serverStart,
	})
	server.AddCommand(&cobra.Command{
		Run: createSimpleConfig,
	})

	server.Flags().Int8P("port", "p", 36080, "server listen port")
	server.Flags().StringP("configPath", "conf", "~/.monaPanel/config.yaml", "set the path to mona server config file. Default path is `~/.monaPanel/serverConf.yaml`")
}

func serverStart(cmd *cobra.Command, args []string) {

}

func createSimpleConfig(cmd *cobra.Command, args []string) {

}
