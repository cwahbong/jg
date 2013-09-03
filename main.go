package main

import (
	"github.com/cwahbong/jg/commands"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		commands.Empty()
	} else {
		commands.CommandFunc(os.Args[1])(os.Args[2:])
	}
}
