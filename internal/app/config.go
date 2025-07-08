package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

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

	jsonConfig, err := os.Open("mainnet-config.json")
	if err != nil {
		return err
	}

	if err := json.NewDecoder(jsonConfig).Decode(&CFG.MainnetConfig); err != nil {
		return err
	}
	defer jsonConfig.Close()

	CFG.Wallet.Seed = strings.Split(os.Getenv("SEED"), " ")

	return nil
}
