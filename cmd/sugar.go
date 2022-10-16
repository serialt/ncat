package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/serialt/cli/config"
	"github.com/serialt/cli/pkg"
	"github.com/serialt/cli/simau"
	"github.com/serialt/sugar"
)

var (
	APPVersion = "v0.2"
	BuildTime  = "20060102"
	GitCommit  = "ccccccccccccccc"
)

var rootCmd = &cobra.Command{
	Use:   "cli ",
	Short: "cli toolkit",
	Long:  `cli toolkit`,
	Run:   RunServer,
}

func initConfig() {
	err := sugar.LoadConfig(simau.ConfigFile, &simau.Config)
	if err != nil {
		simau.Config = new(config.MyConfig)
	}
	simau.Sugar = sugar.NewSugarLogger(simau.Config.Log.LogLevel, simau.Config.Log.LogFile, "", false)

}

func init() {
	// 初始化app信息
	rootCmd.PersistentFlags().StringVarP(&simau.ConfigFile, "config", "c", pkg.Env("CONFIG", simau.ConfigFile), "config file path")
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
}

func RunServer(cmd *cobra.Command, args []string) {
	simau.Sugar.Infof("config: %v", simau.Config)
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
