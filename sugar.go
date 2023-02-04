package main

import (
	"log"
	"os"
)

func service() {
	log.Println("msg")
}

func EnvGet(envName string, defaultValue string) (data string) {
	data = os.Getenv(envName)
	if len(data) == 0 {
		data = defaultValue
		return
	}
	return
}
