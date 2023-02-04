package main

import (
	"os"

	"github.com/serialt/sugar"
)

func service() {
	sugar.Debug("debug msg")
	sugar.Info("info msg")
	sugar.Error("error msg")
}

func EnvGet(envName string, defaultValue string) (data string) {
	data = os.Getenv(envName)
	if len(data) == 0 {
		data = defaultValue
		return
	}
	return
}
