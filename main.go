package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// 版本信息
	appVersion bool // 控制是否显示版本
	APPVersion = "v0.0.2"
	BuildTime  = "2006-01-02 15:04:05"
	GitCommit  = "xxxxxxxxxxx"
)

func init() {
	flag.BoolVar(&appVersion, "v", false, "Display build and version msg")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("使用说明")
		flag.PrintDefaults()
	}
	flag.ErrHelp = fmt.Errorf("\n\nSome errors have occurred, check and try again !!! ")
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

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
