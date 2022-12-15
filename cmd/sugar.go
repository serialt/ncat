package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/serialt/cli/app"
	"github.com/serialt/cli/pkg"
)

var (
	APPVersion = "v0.2"
	BuildTime  = "20060102"
	GitCommit  = "ccccccccccccccc"
)

func init() {
	// 初始化app信息
	rootCmd.PersistentFlags().StringVarP(&pkg.ConfigFile, "config", "c", pkg.Env("CONFIG", pkg.ConfigFile), "config file path")
	rootCmd.PersistentFlags().StringVarP(&pkg.RootPath, "dir", "d", pkg.Env("WORKSPACE", pkg.RootPath), "workspace dir")
	cobra.OnInitialize(pkg.InitConfig)
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "cli ",
	Short: "cli toolkit",
	Long:  `cli toolkit`,
	Run: func(cmd *cobra.Command, args []string) {
		app.RunServer()
	},
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("APPVersion: %v  BuildTime: %v  GitCommit: %v\n",
			APPVersion,
			BuildTime,
			GitCommit)
	},
}
