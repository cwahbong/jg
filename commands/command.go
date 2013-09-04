package commands

import (
	"errors"
	"fmt"
)

type Command func(args []string) error

func CommandFunc(cmd string) Command {
	switch cmd {
	case "server":
		return jgServer
	case "help":
		return jgHelp
	}
	return badCommand(cmd)
}

func Empty() {
	jgHelp(nil)
}

func badCommand(cmd string) Command {
	return func(args []string) error {
		return errors.New(fmt.Sprintf(`Command "%s" does not exists, see "jg help" for more information.`, cmd))
	}
}
