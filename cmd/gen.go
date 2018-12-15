package cmd

import (
	"github.com/k-kinzal/aliases/pkg"
	"github.com/k-kinzal/aliases/pkg/context"
	"github.com/k-kinzal/aliases/pkg/export"
	"github.com/urfave/cli"
)

type GenContext struct {
	*context.Context

	binary bool
}

func NewGenContext(c *cli.Context) *GenContext {
	ctx := context.NewContext(
		c.GlobalString("home"),
		c.GlobalString("config"),
		c.String("binary-path"),
	)

	return &GenContext{
		Context: ctx,
		binary: c.Bool("binary"),
	}
}

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
	ctx := NewGenContext(c)

	// configuration
	conf, err := aliases.LoadConfFile(ctx.Context)
	if err != nil {
		return err
	}

	// output aliases
	if ctx.binary {
		export.Path(ctx.Context, conf)
	} else {
		export.Aliases(ctx.Context, conf)
	}

	return nil
}