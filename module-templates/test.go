//	Name: test Module
//	Usage: Used for testing and debugging gorgon
//	Author: dru1d

package modules

import (
//		"os"
		"fmt"
		"errors"

		"github.com/fatih/color"
		"github.com/urfave/cli"
)

var (
		testPrint = cli.Command { 
				Name:		 "test",
				Usage:		 "tests main function and other stuff",
				UsageText: 	 "gorgon test [ARGUMENTS]",
				Action:		 testFunction,
		}
)

func testFunction(c *cli.Context) error {
	if c.Args().First() == "" {
		return errors.New("must supply at least one argument to this command")
	}
	for _, a := range c.Args() {
		fmt.Sprintf(
		"Test Command Output: %s\n",
		color.HiGreenString(a),
		)
	}
	return nil
}