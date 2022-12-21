package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/serialt/sugar"
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

func InitConfig() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Get project root path failed: %v", err)
	}

	if len(RootPath) == 0 {
		RootPath = filepath.Dir(exePath)
	}
	// fmt.Printf("Project root path: %s\n", RootPath)
	if len(ConfigFile) == 0 {
		ConfigFile = filepath.Join(RootPath, DefaultConfigFile)
	}

	err = sugar.LoadConfig(ConfigFile, &Config)
	if err != nil {
		Config = new(MyConfig)
	}

	// logfile path
	if len(Config.Log.LogFile) != 0 && !filepath.IsAbs(Config.Log.LogFile) {
		Config.Log.LogFile = filepath.Join(RootPath, Config.Log.LogFile)
	}
	// fmt.Printf("Logfile path %s\n", Config.Log.LogFile)

	Sugar = sugar.NewSugarLogger(Config.Log.LogLevel, Config.Log.LogFile, "", false)
	Sugar.Debugf("Project root path: %s", RootPath)
	Sugar.Debugf("Logfile path %s", Config.Log.LogFile)

}
