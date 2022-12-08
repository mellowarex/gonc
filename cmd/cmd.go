package cmd

import (
	"io"
	"os"
	"github.com/mellowarex/gonc/cmd/commands"
	_ "github.com/mellowarex/gonc/cmd/commands/new"
	_ "github.com/mellowarex/gonc/cmd/commands/version"
	_ "github.com/mellowarex/gonc/cmd/commands/run"
	"github.com/mellowarex/gonc/utils"
)

var usage =`
Usage:
	gonc <command> [arguments]

Available commands:`

var usageEnd = `

Use "gonc help <command>" for more information about a command.

`

// Usage print gonc general invo
// cmd usage
// print available commands
func Usage() {
	for _, cmd := range commands.AvailableCommands {
		usage += `
	` + cmd.UsageLine + `			` + cmd.Short
	}

	// append help cmd usage
	usage += usageEnd

	io.WriteString(os.Stdout, usage)
}

func Help(args []string) {
	if len(args) == 0 {
		Usage()
		return
	}

	if len(args) != 1 {
		utils.PrintErrorAndExit("Too many arguments")
		return
	}

	for _, cmd := range commands.AvailableCommands {
		if cmd.Name() == args[0] {
			cmd.FullUsage()
			return
		}
	}

	utils.PrintErrorAndExit("Unknown help topic")
}