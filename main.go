package main

import (
	"github.com/cwahbong/jg/commands"
	"os"
)

func main() {
	switch os.Args[1] {
	case "server":
		commands.JgServer(os.Args[2:])
	case "help":
		fallthrough
	default:
		commands.JgHelp(os.Args[2:])
	}
}
