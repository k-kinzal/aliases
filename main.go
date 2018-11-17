package main

import (
	"github.com/k-kinzal/aliases/cmd"
	"github.com/k-kinzal/aliases/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"log"
	"os"
)

func init() {
	logrus.SetOutput(os.Stdout)
}

func main() {
	app := cli.NewApp()
	app.Name = "aliases"
	app.Usage = "Generate alias for command on container"

	app.Version = version.GetVersion()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "config, c",
			Usage: "Load configuration from `FILE`",
		},
		cli.StringFlag{
			Name: "home",
			Usage: "Home directory for aliases",
		},
	}
	app.Commands = []cli.Command{
		cmd.GenCommand(),
		cmd.HomeCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}