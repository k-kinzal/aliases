package cmd

import (
	"fmt"
	"strings"

	"github.com/k-kinzal/aliases/pkg/aliases"
	"github.com/k-kinzal/aliases/pkg/docker"

	"github.com/k-kinzal/aliases/pkg/aliases/config"
	"github.com/k-kinzal/aliases/pkg/aliases/script"

	"github.com/k-kinzal/aliases/pkg/posix"

	"github.com/urfave/cli"
)

type genContext struct {
	aliases.Context
	cli cli.Context
}

func (ctx *genContext) ExportPath() string {
	p := ctx.cli.String("export-path")
	if p == "" {
		p = ctx.Context.ExportPath()
	}
	return p
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

	if err := ctx.MakeBinaryDir(); err != nil {
		return err
	}

	client, err := docker.NewClient()
	if err != nil {
		return err
	}

	conf, err := config.LoadConfig(ctx.ConfPath())
	if err != nil {
		return err
	}

	for _, binary := range conf.Binaries(ctx.BinaryPath()) {
		if err := docker.Download(binary.Path, binary.Image, binary.Tag); err != nil {
			return err
		}
	}

	if ctx.isExport() {
		for _, opt := range conf.Slice() {
			cmd := script.NewScript(ctx, client, opt)
			if err != nil {
				return err
			}
			_, err := cmd.Write(ctx)
			if err != nil {
				return err
			}
		}
		fmt.Println(posix.PathExport(ctx.ExportPath()))
	} else {
		aliases := make([]posix.Cmd, 0)
		for _, opt := range conf.Slice() {
			cmd := script.NewScript(ctx, client, opt)
			if err != nil {
				return err
			}
			_, err := cmd.Write(ctx)
			if err != nil {
				return err
			}
			aliases = append(aliases, *posix.Alias(cmd.FileName(), strings.Replace(cmd.String(), "$ALIASES_EXPORT_PATH", ctx.ExportPath(), -1)))
		}
		for _, alias := range aliases {
			fmt.Println(alias.String())
		}
	}

	return nil
}
