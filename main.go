package main

import (
	"os"
	"os/exec"
	"os/user"
	"path"
	"syscall"

	"github.com/k-kinzal/aliases/pkg/util"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"

	"github.com/k-kinzal/aliases/pkg/aliases/context"

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
		// logger setting
		logger.SetOutput(os.Stderr)
		if ctx.GlobalBool("verbose") {
			logger.SetLogLevel(logger.DebugLevel)
		} else {
			logger.SetLogLevel(logger.WarnLevel)
		}
		// home directory setting
		homePath := ctx.GlobalString("home")
		if homePath != "" && !util.IsFilePath(homePath) {
			return util.FlagError("home", homePath, "invalid path or path denied permission")
		}
		if homePath == "" {
			usr, _ := user.Current()
			homePath = path.Join(usr.HomeDir, ".aliases")
		}
		if err := context.ChangeHomePath(homePath); err != nil {
			return err
		}
		if err := ctx.GlobalSet("home", homePath); err != nil {
			return err
		}
		// configuration file setting
		confPath := ctx.GlobalString("config")
		if confPath != "" && !util.IsFilePath(confPath) {
			return util.FlagError("config, c", confPath, "invalid path or path denied permission")
		}
		if confPath == "" {
			cwd, _ := os.Getwd()
			confPath = path.Join(cwd, "aliases.yaml")
			if _, err := os.Stat(confPath); os.IsNotExist(err) {
				confPath = path.Join(homePath, "aliases.yaml")
			}
		}
		if err := context.ChangeConfPath(confPath); err != nil {
			return err
		}
		if err := ctx.GlobalSet("config", confPath); err != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		switch e := err.(type) {
		case *util.InvalidFlagError:
			logger.Fatal(e)
		case *yaml.UnmarshalError:
			logger.Fatal(e)
		case *exec.ExitError:
			if status, ok := e.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			}
		default:
			logger.Fatalf("aliases: %s", e)
		}
		os.Exit(1)
	}
}
