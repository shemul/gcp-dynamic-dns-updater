package main

import (
	"github.com/shemul/gcp-dynamic-dns-updater/cmd"
	"github.com/urfave/cli"
)

var (
	app *cli.App
)

func init() {
	app = cli.NewApp()
}

func main() {
	cmd.Run(app)
}
