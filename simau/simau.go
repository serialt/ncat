package simau

import (
	"github.com/serialt/cli/config"
	"go.uber.org/zap"
)

// 存放一些公共的变量
var (
	// Listen   = ":9879"
	// Host     = ""
	// Username = ""
	// Password = ""
	ConfigFile = "config.yaml" // 默认配置文件路径
	Config     *config.MyConfig
	// Logger *zap.Logger
	Sugar *zap.SugaredLogger
)
