package main

import (
	"github.com/k-kinzal/aliases/cmd"
	"github.com/k-kinzal/aliases/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	cmds []cli.Command
)

func init() {
	logrus.SetOutput(os.Stdout)

	cmds = append(cmds, cmd.GenCommand())
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
	}
	app.Commands = cmds

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}