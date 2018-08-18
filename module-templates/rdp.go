//	Name: RDP Module
//	Usage: Used for bruteforcing RDP services
//  Credits: 
//	Author: dru1d

package main

import (
//	"log"
//	"fmt"
	"errors"

//	"github.com/stacktitan/smb/smb"
//	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var (
rdpSubcommands = []cli.Command {
		{
			Name:   "pth",
			Usage:  "pth against hosts",
			Action: rdpPth,
			Flags: []cli.Flag {
				cli.StringFlag {
					Name:  "hash",
					Usage: "insert hash for pth",
				},
				cli.StringFlag {
					Name:  "host",
					Usage: "insert host for pth",
				},
				cli.StringFlag {
					Name:  "hostfile, H",
					Usage: "Insert host files for pth",
				},
				cli.StringFlag {
					Name:  "user, u",
					Usage: "User for pth",
				},
				cli.StringFlag {
					Name:  "userfile, U",
					Usage: "User files for pth",
				},
			},
		},
		{
			Name:   "combo",
			Usage:  "uses combo file to guess credential",
			Action: rdpCombo,
			Flags: []cli.Flag {
				cli.StringFlag {
					Name:  "file, f",
					Usage: "run combo file attack",
				},
				cli.StringFlag { 
					Name:  "host",
					Usage: "insert for combo attack",
				},
				cli.StringFlag {
					Name:  "hostfile, H",
					Usage: "insert host file for combo attack",
				},
			},
		},
		{
			Name:	"plaintext",
			Usage:	"uses plaintext credentials to bruteforce",
			Action:	rdpPlaintext,
			Flags:	[]cli.Flag {
				cli.StringFlag {
					Name:  "host",
					Usage: "host for bruteforce",
					},
				cli.StringFlag {
					Name:  "hostfile, H",
					Usage: "Insert host files for bruteforce",
					},
				cli.StringFlag {
					Name:	"user, u",
					Usage:	"User for bruteforce",
					},
				cli.StringFlag {
					Name:	"userfile, U",
					Usage:	"User file for bruteforce",
					},
				cli.StringFlag {
					Name:	"password, p",
					Usage:	"Password for bruteforce",
					},
				cli.StringFlag {
					Name:	"passfile, P",
					Usage:	"Password file for bruteforce",
					},
				},
			},
		}

rdpBrute = cli.Command { 
			Name:		 "rdp",
			Usage:		 "bruteforces RDP services",
			Subcommands: rdpSubcommands,
		}
)

func rdpPth(c *cli.Context) error {
	if c.Args().First() == "" {
	return errors.New("must supply at least one argument to this command")
	}
	return moduleNotImplemented(c)
}

func rdpCombo(c *cli.Context) error {
	if c.Args().First() == "" {
	return errors.New("must supply at least one argument to this command")
	}
	return moduleNotImplemented(c)
}

func rdpPlaintext(c *cli.Context) error {
	if c.Args().First() == "" {
		return errors.New("must supply at least one argument to this command")
	}
	return moduleNotImplemented(c)
}