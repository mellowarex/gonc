package main

import (
	"flag"
	"os"

	"github.com/mellocraft/gonc/cmd"
	"github.com/mellocraft/gonc/cmd/commands"
	"github.com/mellocraft/gonc/config"
	"github.com/mellocraft/gonc/utils"
)

func main() {
	flag.Usage = cmd.Usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		cmd.Usage()
		os.Exit(0)
		return
	}

	if args[0] == "help" {
		cmd.Help(args[1:])
		return
	}

	for _, c := range commands.AvailableCommands {
		if c.Name() == args[0] && c.Run != nil {
			c.Flag.Usage = func() { c.Usage() }
			if c.CustomFlags {
				args = args[1:]
			} else {
				c.Flag.Parse(args[1:])
				args = c.Flag.Args()
			}

			if c.PreRun != nil {
				c.PreRun(c, args)
			}

			// load configurations
			config.LoadConfig()
			os.Exit(c.Run(c, args))
			return
		}
	}

	// end reached therefore unknown command entered
	utils.UnknownSubCommand(args[0])
}