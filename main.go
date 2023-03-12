package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/serialt/sugar"
)

func init() {
	flag.BoolVar(&appVersion, "v", false, "Display build and version messages")
	flag.StringVar(&ConfigFile, "c", "config.yaml", "Config file")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("使用说明")
		flag.PrintDefaults()
	}
	flag.ErrHelp = fmt.Errorf("\n\nSome errors have occurred, check and try again !!! ")
	flag.Parse()
	sugar.SetLog("info", "")
	InitConfig()

}
func main() {
	if appVersion {
		fmt.Printf("APPVersion: %v  BuildTime: %v  GitCommit: %v\n",
			APPVersion,
			BuildTime,
			GitCommit)
		return
	}
	service()
}
