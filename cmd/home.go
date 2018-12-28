package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func HomeCommand() cli.Command {
	return cli.Command{
		Name:                   "home",
		Usage:                  "Get aliases home path",
		Action:                 HomeAction,
		SkipArgReorder:         true,
		UseShortOptionHandling: true,
	}
}

func HomeAction(c *cli.Context) error {
	fmt.Print(c.GlobalString("home"))

	return nil
}
