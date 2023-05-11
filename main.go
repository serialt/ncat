package main

import (
	"flag"
	"fmt"

	"github.com/serialt/lancet/cryptor"
	"github.com/serialt/sugar/v3"
	"golang.org/x/exp/slog"
)

func init() {
	flag.BoolVar(&appVersion, "v", false, "Display build and version messages")
	flag.StringVar(&ConfigFile, "c", "config.yaml", "Config file")
	flag.StringVar(&AesData, "d", "", "Plaintext for encryption")
	flag.Parse()

	err := sugar.LoadConfig(ConfigFile, &config)
	if err != nil {
		config = new(Config)
	}
	slog.SetDefault(sugar.New())
	config.DecryptConfig()

}
func main() {
	if appVersion {
		fmt.Printf("APPVersion: %v  BuildTime: %v  GitCommit: %v\n",
			APPVersion,
			BuildTime,
			GitCommit)
		return
	}
	if len(AesData) > 0 {
		fmt.Printf("Encrypted string: %v\n", cryptor.AesCbcEncryptBase64(AesData, AesKey))
		fmt.Printf("Plaintext : %v\n", cryptor.AesCbcDecryptBase64(cryptor.AesCbcEncryptBase64(AesData, AesKey), AesKey))
		return
	}
	service()
}
