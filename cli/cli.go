package cli

import (
	"github.com/lcaballero/g-ed/params"
	cmd "gopkg.in/urfave/cli.v2"
)

// CommandProcessor type are actions that take values from the command line
// and then perform some processing.
type CommandProcessor func(params.ValContext) error

// Action properly handles treating a CommandProcessor as a cli Action.
func (cp CommandProcessor) Action() func(*cmd.Context) error {
	return func(c *cmd.Context) error {
		return cp(c)
	}
}

// Processes hold members of CommandProcessors to execute as actions based on
// the sub-command as parsed from the cli.
type Processes struct {
	Serve CommandProcessor
}

// New takes a set of Processes and returns an command Application to run a
// distinct Processor based on command-line arguments.
func New(pr Processes) *cmd.App {
	return &cmd.App{
		Name:    "g-ed",
		Version: "1.0.0",
		Commands: []*cmd.Command{
			{
				Name:   "serve",
				Flags:  flags(),
				Action: pr.Serve.Action(),
			},
		},
	}
}

// flags creates a set of flags for sub-commands.
func flags() []cmd.Flag {
	return []cmd.Flag{
		&cmd.StringFlag{
			Name:  "root",
			Value: "_dest",
			Usage: "Root directory where assets can be found.",
		},
	}
}
