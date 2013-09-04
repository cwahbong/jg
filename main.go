package main

import (
	"fmt"
	"github.com/cwahbong/jg/commands"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		commands.Empty()
	} else {
		err := commands.CommandFunc(os.Args[1])(os.Args[2:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		}
	}
}
