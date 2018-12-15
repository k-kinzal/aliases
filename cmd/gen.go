package cmd

import (
	"github.com/k-kinzal/aliases/pkg/conf"
	"github.com/k-kinzal/aliases/pkg/context"
	"github.com/k-kinzal/aliases/pkg/export"
	"github.com/urfave/cli"
)

type GenContext struct {
	*context.Context

	export bool
}

func NewGenContext(c *cli.Context) *GenContext {
	ctx := context.NewContext(
		c.GlobalString("home"),
		c.GlobalString("config"),
		c.String("export-path"),
	)

	return &GenContext{
		Context: ctx,
		export: c.Bool("export"),
	}
}

func GenCommand() cli.Command {
	return cli.Command {
		Name:    "gen",
		Usage:   "Generate aliases",
		Flags: []cli.Flag {
			cli.BoolFlag{
				Name: "export",
				Usage: "If you pass true, you will return export instead of aliase",
			},
			cli.StringFlag{
				Name: "export-path",
				Usage: "The directory to put binaries",
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
	cf, err := conf.LoadConfFile(ctx.Context)
	if err != nil {
		return err
	}

	// output aliases
	if ctx.export {
		export.Path(ctx.Context, cf)
	} else {
		export.Aliases(ctx.Context, cf)
	}

	return nil
}