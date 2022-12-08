package version

import (
	"io"
	"github.com/mellowarex/gonc/cmd/commands"
)

const (
	VERSION = "1.0.0"
)

var CmdVersion = &commands.Command{
	UsageLine: "version",
	Use:				"version",
	Short:		 "Print the current gonc version",
	Long:      "Prints the current gonc version.\n",
	Run:			versionCmd,
}

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdVersion)
}

// versionCmd prints gonc current version
func versionCmd(cmd *commands.Command, args []string) int {
	io.WriteString(cmd.StdOut(), "gonc version "+VERSION + "\n")
	return 0
}
