package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/serialt/sugar"
)

var (
	RootPath = ""
	// 版本信息
	appVersion bool // 控制是否显示版本
	APPVersion = "v0.0.2"
	BuildTime  = "2006-01-02 15:04:05"
	GitCommit  = "xxxxxxxxxxx"
	ConfigFile = "config.yaml"
	Config     *MyConfig
)

type Service struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

type MyConfig struct {
	Service Service `json:"service" yaml:"service"`
}

func InitConfig() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Get project root path failed: %v", err)
	}

	if len(RootPath) == 0 {
		RootPath = filepath.Dir(exePath)
	}

	ConfigFile = filepath.Join(RootPath, ConfigFile)

	err = sugar.LoadConfig(ConfigFile, &Config)
	if err != nil {
		Config = new(MyConfig)
	}
}
