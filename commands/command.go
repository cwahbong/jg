package commands

import (
	"fmt"
	"os"
)

type Command func(args []string)

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
	return func(args []string) {
		fmt.Fprintf(os.Stderr, "Command \"%s\" does not exists, see \"jg help\" for more information.\n", cmd)
	}
}
