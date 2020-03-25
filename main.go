package main

import (
	"github.com/imjuanleonard/palu-covid/cmd"
	"github.com/imjuanleonard/palu-covid/cmd/app"
	"github.com/imjuanleonard/palu-covid/config"
	"github.com/imjuanleonard/palu-covid/pkg/logger"
)

var (
	version = "0.0.1"
)

func main() {
	config.Load()

	if err := app.Init(); err != nil {
		logger.Errorf("error initializing the app: %+v", err)
		return
	}
	defer app.Close()

	cli := cmd.NewCLI()
	cli.Version = version
	if err := cli.Execute(); err != nil {
		logger.Errorf("error on executing the CLI: %+v", err)
	}
}
