package cmd

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg"
	"github.com/urfave/cli"
)

func GenCommand() cli.Command {
	return cli.Command {
		Name:    "generate",
		Aliases: []string{"gen"},
		Usage:   "Generate aliases",
		Action:  func(c *cli.Context) error {
			return GenAction(c)
		},
	}
}

func GenAction(c *cli.Context) error {
	// context
	context, err := aliases.NewContext(c.GlobalString("config"))
	if err != nil {
		return err
	}
	// configuration
	conf, err := aliases.LoadConfFile(context.ConfPath)
	if err != nil {
		return err
	}
	// generate commands
	cmds := aliases.GenerateCommands(conf, context)

	// output aliases
	for _, cmd := range cmds {
		fmt.Printf(cmd.ToString())
	}

	return nil
}