package cmd

import (
	"github.com/k-kinzal/aliases/pkg"
	"github.com/k-kinzal/aliases/pkg/export"
	"github.com/urfave/cli"
)

func GenCommand() cli.Command {
	return cli.Command {
		Name:    "gen",
		Usage:   "Generate aliases",
		Flags: []cli.Flag {
			cli.BoolFlag{
				Name: "binary",
				Usage: "",
			},
		},
		Action:  func(c *cli.Context) error {
			return GenAction(c)
		},
	}
}

func GenAction(c *cli.Context) error {
	// context
	ctx := aliases.NewContext(c.GlobalString("home"), c.GlobalString("config"))
	// configuration
	conf, err := aliases.LoadConfFile(*ctx)
	if err != nil {
		return err
	}
	// generate commands
	cmds := aliases.GenerateCommands(*conf, *ctx)

	// output aliases
	export.WriteFiles(cmds, *conf, *ctx)
	if c.Bool("binary") {
		export.Path(*conf, *ctx)
	} else {
		export.Aliases(cmds)
	}

	return nil
}