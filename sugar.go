package main

import (
	"os"

	"github.com/serialt/lancet/cryptor"
	"golang.org/x/exp/slog"
)

func service() {
	slog.Debug("debug msg")
	slog.Info("info msg")
	slog.Error("error msg")
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
