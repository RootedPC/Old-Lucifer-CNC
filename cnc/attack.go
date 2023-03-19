package main

import (
	"errors"
	"fmt"

	"github.com/mattn/go-shellwords"
)

type cmdslookup struct {
	cmdsID          uint8
	cmdsFlags       []uint8
	cmdsDescription string
}

var cmdInformation map[string]cmdslookup = map[string]cmdslookup{}

func checkcommand(str string, admin bool) error {
	args, _ := shellwords.Parse(str)

	if len(args) < 0 {
		return errors.New("\033[0mNo Command Entered.\033[0m")
	} else if len(args) == 0 {
		return errors.New("\033[0mNo Command Entered.\033[0m")
	} else {
		var cmdexists bool
		_, cmdexists = cmdInformation[args[0]]
		if !cmdexists {
			return errors.New(fmt.Sprintf("\033[0mCommand Not Found.\033[0m"))
		}
		args = args[1:]
	}
	return nil
}
