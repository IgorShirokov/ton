package app

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

func InitApp() error {
	if err := InitConfig(); err != nil {
		return err
	}

	if err := InitLogger(); err != nil {
		return err
	}

	client := liteclient.NewConnectionPool()
	if err := client.AddConnectionsFromConfig(context.Background(), CFG.MainnetConfig); err != nil {
		return err
	}

	api := ton.NewAPIClient(client)

	wall, err := wallet.FromSeed(api, CFG.Wallet.Seed, wallet.V4R2)
	if err != nil {
		return err
	}
	logrus.Info(wall.Address())

	return nil
}
