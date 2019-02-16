package cmd

import (
	"fmt"
	"path"

	"github.com/k-kinzal/aliases/pkg/util"

	"github.com/k-kinzal/aliases/pkg/aliases/context"
	"github.com/k-kinzal/aliases/pkg/types"

	"github.com/k-kinzal/aliases/pkg/docker"

	"github.com/k-kinzal/aliases/pkg/aliases/config"
	"github.com/k-kinzal/aliases/pkg/aliases/script"

	"github.com/k-kinzal/aliases/pkg/posix"

	"github.com/urfave/cli"
)

// GenCommand returns `aliases gen` command.
//
// `aliases gen` aliases the command defined in aliases.yaml, or export the path of the command.
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

// GenAction is the action of `aliases gen`.
func GenAction(c *cli.Context) error {
	isExport := c.Bool("export")
	exportPath := c.String("export-path")
	if exportPath != "" && !util.IsFilePath(exportPath) {
		return util.FlagError("export-path", exportPath, "invalid path or path denied permission")
	}
	if exportPath == "" {
		exportPath = path.Join(context.HomePath(), types.MD5(context.ConfPath()))
	}
	if err := context.ChangeExportPath(exportPath); err != nil {
		return err
	}

	client, err := docker.NewClient()
	if err != nil {
		return err
	}

	conf, err := config.LoadConfig(context.ConfPath())
	if err != nil {
		return err
	}

	if isExport {
		for _, opt := range conf.Slice() {
			cmd := script.NewScript(client, opt)
			if err != nil {
				return err
			}
			_, err := cmd.Write()
			if err != nil {
				return err
			}
		}
		fmt.Println(posix.PathExport(context.ExportPath()))
	} else {
		aliases := make([]posix.Cmd, 0)
		for _, opt := range conf.Slice() {
			cmd := script.NewScript(client, opt)
			if err != nil {
				return err
			}
			p, err := cmd.Write()
			if err != nil {
				return err
			}
			aliases = append(aliases, *posix.Alias(cmd.FileName(), p))
		}
		for _, alias := range aliases {
			fmt.Println(alias.String())
		}
	}

	return nil
}
