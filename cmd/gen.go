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
			cli.StringFlag{
				Name: "binary-path",
				Usage: "the directory to put binaries. works only when --binary is specified",
			},
		},
		Action:  func(c *cli.Context) error {
			return GenAction(c)
		},
	}
}

func GenAction(c *cli.Context) error {
	// context
	ctx := aliases.NewContext(c.GlobalString("home"), c.GlobalString("config"), c.String("binary-path"))

	// configuration
	conf, err := aliases.LoadConfFile(ctx)
	if err != nil {
		return err
	}

	// output aliases
	if c.Bool("binary") {
		export.Path(conf, ctx)
	} else {
		export.Aliases(conf, ctx)
	}

	return nil
}