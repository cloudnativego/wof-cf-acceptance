package main

import (
	"fmt"
	"os"

	"github.com/cloudnativego/wof-cf-acceptance/command"
	"github.com/urfave/cli"
)

//GlobalFlags -
var GlobalFlags = []cli.Flag{}

//Commands -
var Commands = []cli.Command{
	{
		Name:     "deploy",
		Category: "DEPLOYMENT ACTIONS",
		Usage:    "Deploy World of Fluxcraft to Cloud Foundry",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config, c",
				Usage: "Load configuration from `FILEPATH`",
			},
		},
		Action: command.CmdDeploy,
	},
	{
		Name:     "destroy",
		Category: "DEPLOYMENT ACTIONS",
		Usage:    "Remove World of Fluxcraft from Cloud Foundry",
		Action:   command.CmdDestroy,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config, c",
				Usage: "Load configuration from `FILEPATH`",
			},
		},
	},
	{
		Name:     "acceptance-tests",
		Category: "TEST ACTIONS",
		Usage:    "Run suite of acceptance tests for World of Fluxcraft",
		Action:   command.CmdAcceptanceTests,
		Flags:    []cli.Flag{},
	},
}

//CommandNotFound -
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
