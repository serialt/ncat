package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/serialt/cli/config"
	"github.com/serialt/cli/pkg"
)

var (
	APPVersion = "v0.2"
	BuildTime  = "20060102"
	GitCommit  = "ccccccccccccccc"

	configFile = "config.yaml"
)

var rootCmd = &cobra.Command{
	Use:   "cli ",
	Short: "cli toolkit",
	Long:  `cli toolkit`,
	Run:   RunServer,
}

func initConfig() {
	config.LoadConfig(configFile)

}

func init() {
	// 初始化app信息
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", pkg.Env("CONFIG", configFile), "config file path")
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
}

func RunServer(cmd *cobra.Command, args []string) {
	config.Sugar.Infof("config: %v", config.Config)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

// command: version
// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version of Gins",
	Long:  "print the version of Gins",
	Run:   DisplayVersion,
}

func DisplayVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("APPVersion: %v  BuildTime: %v  GitCommit: %v\n",
		APPVersion,
		BuildTime,
		GitCommit)
}
