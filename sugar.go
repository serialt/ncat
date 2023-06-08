package main

import (
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/serialt/lancet/cryptor"
	"golang.org/x/exp/slog"
)

func service() {
	slog.Debug("debug msg")
	slog.Info("info msg")
	slog.Error("error msg")

	// 定时任务
	c := cron.New()
	c.AddFunc("@every 2s", func() {
		slog.Info("cron job", "output", time.Second)
	})
	c.Start()
}

func EnvGet(envName string, defaultValue string) (data string) {
	data = os.Getenv(envName)
	if len(data) == 0 {
		data = defaultValue
		return
	}
	return
}

func (c *Config) DecryptConfig() {
	if c.Encrypt {
		c.Token = cryptor.AesCbcDecryptBase64(c.Token, AesKey)
		slog.Debug(c.Token)
	}
}
