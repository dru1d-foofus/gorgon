//	Name: SMB Module
//	Usage: Used for bruteforcing SMB services
//  Credits: 
//	Author: dru1d

package main

import (
	"log"
	"fmt"
	"errors"
	"strings"
	"github.com/stacktitan/smb/smb"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"github.com/dru1d-foofus/gorgon/helpers/files"
)

var (
smbSubcommands = []cli.Command{
		{
			Name:   "combo",
			Usage:  "uses combo file to guess credential",
			Action: smbCombo,
			Flags: []cli.Flag {
				cli.StringFlag {
					Name:  "file, f",
					Usage: "combo file in tab delimited format: ex. HOST[TAB]USER[TAB]PASS format",
				},
				cli.IntFlag {
					Name:  "port",
					Usage: "port for combo attack",
					Value: 445,
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
			Action:	smbPlaintext,
			Flags:	[]cli.Flag {
				cli.StringFlag {
					Name:  "host",
					Usage: "host for bruteforce",
					},
				cli.StringFlag {
					Name:  "hostfile, H",
					Usage: "Insert host files for bruteforce",
					},
				cli.IntFlag {
					Name:	"port",
					Usage:	"Port for bruteforce",
					Value:	445,
					},
				cli.StringFlag {
					Name:	"domain",
					Usage:	"Domain for user",
					Value:	"",
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

smbBrute = cli.Command {
			Name:		 "smb",
			Usage:		 "bruteforces SMB services",
			Subcommands: smbSubcommands,
		}
)

func smbAuth(username string, password string, host string, domain string, port int, timer int) *resp {
	respond := &resp{}
	clientConf := smb.Options{
				Host:		host,
				Port:		port,
				User:		username,
				Domain:		domain,
				Workstation: "",
				Password:	password,
	}
	debug := false
	session, err := smb.NewSession(clientConf, debug)
	if err != nil {
			log.Fatalln("[!]", err)
	}
	defer session.Close()
//	end := time.Now()
//	d := end.Sub(inittime)
//	duration := d.Seconds()
	if session.IsAuthenticated {
			fmt.Printf("\nUser: %s Password: %s [%s]", username, password, color.GreenString("SUCCESS"))
	} else {
			fmt.Printf("\nUser: %s Password: %s [%s]", username, password, color.RedString("FAILED"))
	}
	respond.Error = err
	return respond
}

func smbPth(c *cli.Context) error {
	if c.Args().First() == "" {
	return errors.New("must supply at least one argument to this command")
	}
	return moduleNotImplemented(c)
}

func smbCombo(c *cli.Context) error {
// combo file logic
	if c.String("file") == "" {
	return errors.New("must supply a combo file to this command")
	}
	if c.String("file") != "" {
		combos, err := files.ReadLines(c.String("file"))
		if err != nil {
			log.Fatalf("readLine: %s", err)
		}
		for combo := range combos {
			combosplice := strings.Split(combos[combo], "\t")
			host,user,password,domain := combosplice[0],combosplice[1],combosplice[2],combosplice[3]
			resp := smbAuth(user,password,host,domain,c.Int("port"),c.Int("timeout"))
			resp.mu.Lock()
			if resp.Error == nil {
				resp.mu.Unlock()
			}
		}
	}
	return nil
}

func smbPlaintext(c *cli.Context) error {
		if c.String("userfile") != "" {
		users, err := files.ReadLines(c.String("userfile"))
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		for user := range users {
// userfile + passfile
			if c.String("passfile") != "" {
				passwords, err := files.ReadLines(c.String("passfile"))
				if err != nil {
					log.Fatalf("readLines: %s", err)
				}
				for password := range passwords {
					resp := smbAuth(users[user],passwords[password],c.String("host"), c.String("domain"), c.Int("port"),c.Int("timeout"))
		        	resp.mu.Lock()
	            		if resp.Error == nil {
			        		resp.mu.Unlock()
					}
				}
// userfile + password
			} else {
				resp := smbAuth(users[user],c.String("password"), c.String("host"), c.String("doman"),c.Int("port"),c.Int("timeout"))
				resp.mu.Lock()
					if resp.Error == nil {
						resp.mu.Unlock()
					}
			}
		}
	} else {
// username + passfile
		if c.String("passfile") != "" {
			passwords, err := files.ReadLines(c.String("passfile"))
			if err != nil {
				log.Fatalf("readLine: %s", err)
			}
			for password := range passwords {
				resp := smbAuth(c.String("user"),passwords[password],c.String("host"),c.String("domain"),c.Int("port"),c.Int("timeout"))
				resp.mu.Lock()
					if resp.Error == nil {
						resp.mu.Unlock()
					}
			}
		} else {
// username + password
			resp := smbAuth(c.String("user"),c.String("password"),c.String("host"),c.String("domain"),c.Int("port"),c.Int("timeout"))
			resp.mu.Lock()
				if resp.Error == nil {
					resp.mu.Unlock()
				}
		}
	}
	return nil
}