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

type GenContext struct {
	*context.Context

	export bool
}

func NewGenContext(c *cli.Context) *GenContext {
	ctx := context.New(
		c.GlobalString("home"),
		c.GlobalString("config"),
		c.String("export-path"),
	)

	return &GenContext{
		Context: ctx,
		export:  c.Bool("export"),
	}
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
	ctx := NewGenContext(c)

	exec, err := executor.New(*ctx.Context)
	if err != nil {
		return err
	}

	commands, err := exec.Commands(*ctx.Context)
	if err != nil {
		return err
	}

	if err := export.Script(*ctx.Context, commands); err != nil {
		return err
	}

	if ctx.export {
		exp := posix.PathExport(ctx.GetExportPath(), false)
		fmt.Println(posix.String(*exp))
	} else {
		for path, cmd := range commands {
			alias := posix.Alias(pathes.Base(path), posix.String(cmd))

			fmt.Println(posix.String(*alias))
		}
	}

	return nil
}
