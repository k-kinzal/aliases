package main

import (
	"fmt"
	"os"

	"github.com/k-kinzal/aliases/cmd"
	"github.com/k-kinzal/aliases/pkg/version"
	"github.com/urfave/cli"
)

func main() {
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help",
		Usage: "show help",
	}

	app := cli.NewApp()
	app.Name = "aliases"
	app.Usage = "Generate alias for command on container"

	app.Version = version.GetVersion()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration file",
		},
		cli.StringFlag{
			Name:  "home",
			Usage: "Home directory for aliases",
		},
	}
	app.Commands = []cli.Command{
		cmd.GenCommand(),
		cmd.RunCommand(),
		cmd.HomeCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}
