package main

import (
	"ton/internal/app"

	"github.com/sirupsen/logrus"
)

// https://ton-blockchain.github.io/global.config.json

func main() {
	if err := run(); err != nil {
		panic(err)
	}

}

func run() error {
	if err := app.InitApp(); err != nil {
		return err
	}

	logrus.Info("something info")

	return nil
}
