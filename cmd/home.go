package cmd

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg/context"
	"github.com/urfave/cli"
)

func HomeCommand() cli.Command {
	return cli.Command{
		Name:  "home",
		Usage: "Get aliases home path",
		Action: func(c *cli.Context) error {
			return HomeAction(c)
		},
	}
}

func HomeAction(c *cli.Context) error {
	// context
	ctx := context.NewContext("", "", "")

	// output
	fmt.Print(ctx.GetHomePath())

	return nil
}
