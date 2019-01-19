package cmd

import (
	"fmt"
	pathes "path"

	"github.com/k-kinzal/aliases/pkg/aliases"
	"github.com/k-kinzal/aliases/pkg/posix"

	"github.com/k-kinzal/aliases/pkg/export"
	"github.com/urfave/cli"
)

type genContext struct {
	aliases.Context
	cli cli.Context
}

func (ctx *genContext) ExportPath() string {
	path := ctx.cli.String("export-path")
	if path == "" {
		path = ctx.Context.ExportPath()
	}
	return path
}

func (ctx *genContext) isExport() bool {
	return ctx.cli.Bool("export")
}

func newGenContext(c *cli.Context) (*genContext, error) {
	ctx, err := aliases.NewContext(
		c.GlobalString("home"),
		c.GlobalString("config"),
	)
	if err != nil {
		return nil, err
	}

	return &genContext{ctx, *c}, nil
}

func GenCommand() cli.Command {
	return cli.Command{
		Name:  "gen",
		Usage: "Generate aliases",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "export",
				Usage: "if you pass true, you will return export instead of aliase",
			},
			cli.StringFlag{
				Name:  "export-path",
				Usage: "the directory to export scripts",
			},
		},
		Action:                 GenAction,
		SkipArgReorder:         true,
		UseShortOptionHandling: true,
	}
}

func GenAction(c *cli.Context) error {
	ctx, err := newGenContext(c)
	if err != nil {
		return err
	}

	if err := ctx.MakeExportDir(); err != nil {
		return err
	}

	ledger, err := aliases.NewLedgerFromConfig(ctx.ConfPath())
	if err != nil {
		return err
	}

	if ctx.isExport() {
		for _, schema := range ledger.Schemas() {
			cmd, err := aliases.NewCommand(ctx, schema)
			if err != nil {
				return err
			}
			if err := export.Script(pathes.Join(ctx.ExportPath(), schema.FileName), *cmd); err != nil {
				return err
			}
		}
		exp := posix.PathExport(ctx.ExportPath(), false)
		fmt.Println(exp.String())
	} else {
		for _, schema := range ledger.Schemas() {
			for _, dependency := range schema.Dependencies {
				if dependency.IsSchema() {
					for _, s := range dependency.Schemas() {
						cmd, err := aliases.NewCommand(ctx, s)
						if err != nil {
							return err
						}
						if err := export.Script(pathes.Join(ctx.ExportPath(), schema.FileName), *cmd); err != nil {
							return err
						}
					}
				} else {
					s, err := ledger.LookUp(dependency.String())
					if err != nil {
						return err
					}
					cmd, err := aliases.NewCommand(ctx, *s)
					if err != nil {
						return err
					}
					if err := export.Script(pathes.Join(ctx.ExportPath(), schema.FileName), *cmd); err != nil {
						return err
					}
				}
			}
			cmd, err := aliases.NewCommand(ctx, schema)
			if err != nil {
				return err
			}

			alias := posix.Alias(schema.FileName, cmd.String())

			fmt.Println(alias.String())
		}
	}

	return nil
}
