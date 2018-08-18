package main

import (
		"fmt"
		"os"

		"github.com/fatih/color"
		"github.com/sirupsen/logrus"
		"github.com/urfave/cli"
		"github.com/dru1d-foofus/gorgon"
//		"github.com/dru1d-foofus/gorgon/modules"
		"github.com/dru1d-foofus/gorgon/logger/standard"
)

var (
		//defaultCompileOptions = 
		cliLogger				= standard.NewStandardLogger(nil, "gorgon", "cli", false, false)
		displayBefore 			= true
		debugOutput				= false
)

func init() {
	cli.HelpFlag = cli.BoolFlag{Name: "help, h"}
	cli.VersionFlag = cli.BoolFlag{Name: "version"}
	cli.VersionPrinter = func(c *cli.Context) {
			fmt.Fprintf(c.App.Writer, "%s\n", gorgon.Version)
	}
}

func main() {
	app := cli.NewApp()
	app.Writer = color.Output
	app.ErrWriter = color.Output

	cli.AppHelpTemplate = fmt.Sprintf("%s\n%s", standard.ASCIILogo(), cli.AppHelpTemplate)
	app.Name = "gorgon"
	app.Usage = "Cross platform brute forcing tool"
	app.Description = "Fast and accurate bruteforce tool based on jmk's Medusa."

	app.Flags = []cli.Flag {
			cli.BoolFlag {
					Name:		"debug, d",
					Usage:		"enables debug output",
					Destination: &debugOutput,
			},
	}

	app.Version = gorgon.Version
	app.Authors = []cli.Author {
			cli.Author {
					Name: "Tyler Booth",
					Email: "dru1d@foofus.net",
			},
	}

	app.Copyright = "Copyright (C) 2018 Tyler Booth; inspired by Joe 'jmk' Mondloch"
	app.Commands = []cli.Command {
			sshBrute,
			smbBrute,
	}

	app.Before = func(c *cli.Context) error {
		if debugOutput {
				cliLogger.Logger.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
			cliLogger.Fatalf("Error Encountered: %v", err)
	}
}

func moduleNotImplemented(c *cli.Context) error {
		return fmt.Errorf("%s module not implemented", c.Command.FullName())
}
