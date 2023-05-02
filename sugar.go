package main

import (
	"os"

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
