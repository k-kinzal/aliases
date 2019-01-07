package main

import (
	"os"

	"github.com/k-kinzal/aliases/pkg/aliases"

	"github.com/k-kinzal/aliases/pkg/logger"

	"github.com/k-kinzal/aliases/cmd"
	"github.com/k-kinzal/aliases/pkg/version"
	"github.com/urfave/cli"
)

func main() {
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help",
		Usage: "show help",
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}

	app := cli.NewApp()
	app.Name = "aliases"
	app.Usage = "Generate alias for command on the container"

	app.Version = version.GetVersion()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "load configuration file",
		},
		cli.StringFlag{
			Name:   "home",
			Usage:  "home directory for aliases",
			EnvVar: "ALIASES_HOME",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "enable verbose output",
		},
	}
	app.Commands = []cli.Command{
		cmd.GenCommand(),
		cmd.RunCommand(),
		cmd.HomeCommand(),
	}
	app.Before = func(ctx *cli.Context) error {
		logger.SetOutput(os.Stderr)
		if ctx.GlobalBool("verbose") {
			logger.SetLogLevel(logger.DebugLevel)
		} else {
			logger.SetLogLevel(logger.WarnLevel)
		}

		homePath := ctx.GlobalString("home")
		c, err := aliases.NewContext(homePath, "")
		if err != nil {
			return err
		}
		if err := c.MakeHomeDir(); err != nil {
			return err
		}
		if err := ctx.GlobalSet("home", c.HomePath()); err != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
}
