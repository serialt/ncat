package config

import (
	"go.uber.org/zap"
)

// 存放一些公共的变量
var (
	// Listen   = ":9879"
	// Host     = ""
	// Username = ""
	// Password = ""

	Config *MyConfig
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
