package pkg

import (
	"go.uber.org/zap"
)

// 存放一些公共的变量
var (
	// Listen   = ":9879"
	// Host     = ""
	// Username = ""
	// Password = ""
	RootPath   = ""
	ConfigFile = "config.yaml" // 默认配置文件路径
	Config     *MyConfig
	// Logger *zap.Logger
	Sugar *zap.SugaredLogger
)
