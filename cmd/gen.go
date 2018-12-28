package cmd

import (
	"fmt"
	pathes "path"

	"github.com/k-kinzal/aliases/pkg/posix"

	"github.com/k-kinzal/aliases/pkg/context"
	"github.com/k-kinzal/aliases/pkg/executor"
	"github.com/k-kinzal/aliases/pkg/export"
	"github.com/urfave/cli"
)

type genContext struct {
	context.Context
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

func NewGenContext(c *cli.Context) (*genContext, error) {
	ctx, err := context.New(
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
				Usage: "If you pass true, you will return export instead of aliase",
			},
			cli.StringFlag{
				Name:  "export-path",
				Usage: "The directory to export scripts",
			},
		},
		Action:                 GenAction,
		SkipArgReorder:         true,
		UseShortOptionHandling: true,
	}
}

func GenAction(c *cli.Context) error {
	ctx, err := NewGenContext(c)
	if err != nil {
		return err
	}

	exec, err := executor.New(ctx)
	if err != nil {
		return err
	}

	commands, err := exec.Commands(ctx)
	if err != nil {
		return err
	}

	if err := export.Script(ctx, commands); err != nil {
		return err
	}

	if ctx.isExport() {
		exp := posix.PathExport(ctx.ExportPath(), false)
		fmt.Println(posix.String(*exp))
	} else {
		for path, cmd := range commands {
			alias := posix.Alias(pathes.Base(path), posix.String(cmd))

			fmt.Println(posix.String(*alias))
		}
	}

	return nil
}
