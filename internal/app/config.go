package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/xssnick/tonutils-go/liteclient"
)

type appConfig struct {
	MainnetConfig *liteclient.GlobalConfig
	Logger
	Wallet
}

type Logger struct {
	LogLevel string
}

type Wallet struct {
	Seed []string
}

var CFG = &appConfig{}

func InitConfig() error {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Config failed to load")
		return err
	}

	CFG.Logger.LogLevel = os.Getenv("LOG_LVL")

	return nil
}
