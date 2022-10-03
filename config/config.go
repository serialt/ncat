package config

import (
	"fmt"
	"io/ioutil"

	"github.com/mitchellh/go-homedir"
	"github.com/serialt/cli/pkg"
	"github.com/serialt/sugar"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// 存放一些公共的变量
var (
	// Listen   = ":9879"
	// Host     = ""
	// Username = ""
	// Password = ""

	Config MyConfig
	// Logger *zap.Logger
	Sugar *zap.SugaredLogger
)

type Service struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

type Log struct {
	LogLevel string `yaml:"logLevel"` // 日志级别，支持debug,info,warn,error,panic
	LogFile  string `yaml:"logFile"`  // 日志文件存放路径,如果为空，则输出到控制台
}

type MyConfig struct {
	Log     Log     `json:"log" yaml:"log"`
	Service Service `json:"service" yaml:"service"`
}

func LoadConfig(filepath string) {
	filepath, err := homedir.Expand(filepath)
	if err != nil {
		fmt.Printf("get config file failed: %v\n", err)
	}
	if pkg.Exists(filepath) {
		config, _ := ioutil.ReadFile(filepath)
		err = yaml.Unmarshal(config, &Config)
		if err != nil {
			fmt.Printf("Unmarshal to struct, err: %v", err)
		}
	} else {
		Config.Log.LogLevel = "info"
	}
	// fmt.Printf("LoadConfig: %v\n", Config)
	Sugar = sugar.NewSugarLogger(Config.Log.LogLevel, Config.Log.LogFile, "", false)
}
