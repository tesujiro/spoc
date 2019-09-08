package main

import (
	"flag"
	"os"

	"github.com/tesujiro/spoc/command"
	"github.com/tesujiro/spoc/global"
)

func main() {
	if len(os.Args) < 2 {
		command.Usage()
		os.Exit(1)
	}

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.BoolVar(&global.FlagOnlyIDs, "id", false, "displays only IDs")
	f.BoolVar(&global.FlagRawJson, "json", false, "displays raw json")
	f.Parse(os.Args[1:])
	os.Args = f.Args()

	spoc := NewSpoc()

	cmd := os.Args[0]
	args := os.Args[1:]

	spoc.Run(cmd, args)
}
