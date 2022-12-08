package commands

import (
	"flag"
	"io"
	"os"
	"fmt"
)

type Command struct {
	// Run runs the command
	// args arguments: subcommands
	Run  					func(cmd *Command, args []string) int

	// PreRun performs an operation before running the command
	PreRun				func(cmd *Command, args []string)

	// UsageLine is the one-line Usage msg
	// first word in the line is taken to be cmd name
	UsageLine			string

	// Use
	Use   				string

	// Options cmd arguments
	Args 					[]string

	// Short short description shown
	Short  				string

	// Long is long msg
	Long  				string

	// Flag is a set of flags specific to this cmd
	Flag   				flag.FlagSet

	// CustomFlags indicates that the cmd will do
	// its own flag parsing
	CustomFlags		bool

	// output out writer if set in SetOutput(w)
	output				*io.Writer
}

var AvailableCommands = []*Command{}
var cmdUsage = `Use gonc help %s for more information.`

// Name returns name of command: usage line
func (this *Command) Name() string {
	return this.UsageLine
}

// StdOut returns out writer of current cmd
// if cmd.output is nil, os.Stdout is used
func (this *Command) StdOut() io.Writer {
	if this.output != nil {
		return *this.output
	}

	return os.Stdout
}

func (this *Command) Write(msg string) {
	io.WriteString(this.StdOut(), msg)
}

// Usage puts out the Usage for the command
func (this *Command) Usage() {
	usageMsg := fmt.Sprint(cmdUsage, this.UsageLine)
	io.WriteString(this.StdOut(), usageMsg)
	os.Exit(0)
}

// Full puts out full command line usage
// usage
// arguments
// description
func (this *Command) FullUsage() {
	usageMsg :=`
Usage:
	gonc ` + this.Use +`
	
`
	optionMsg := ""
	if len(this.Args) > 0 {
		optionMsg += "OPTIONS"
	}
	for _, args := range this.Args{
		optionMsg +=`
	` + args
	}
	if len(optionMsg) > 0 {
		optionMsg += "\n"
	}

	descriptionMsg := `
Description
	` + this.Long

	usage := usageMsg + optionMsg + descriptionMsg
	io.WriteString(this.StdOut(), usage)
	os.Exit(0)
}