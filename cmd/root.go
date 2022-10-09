/*
Copyright © 2022 HatsuneMona <songmiao39@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "monaPanel",
	Short: "A panel to control your servers.",
	Long:  `A panel to control your servers.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: rootCmdRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.monaPanel.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().StringP("logLevel", "lf", "Info", "logLevel")
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	fmt.Printf("helloWorld!\n")
}
